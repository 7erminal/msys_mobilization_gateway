// Package middleware contains the gateway's HTTP middleware that implements
// end-to-end request tracing and structured access logging using the shared
// github.com/msys/go-common/trace package.
package middleware

import (
	"net"
	"net/http"
	"time"

	gctrace "github.com/msys/go-common/trace"
)

// Config controls the tracing middleware.
type Config struct {
	// ServiceName is emitted on every log line (defaults to "gateway").
	ServiceName string
	// TrustedProxies are the IP ranges allowed to set X-Forwarded-* headers.
	// Empty means the gateway is directly internet-facing.
	TrustedProxies []net.IPNet
	// LBRequestIDHeader, if set, names a header carrying a load-balancer request id.
	// It is recorded as `lb_request_id` and is NEVER used as the request id.
	LBRequestIDHeader string
	// UpstreamErrorMode is "legacy" (default) or "http". Reserved for future use;
	// the gateway currently always returns the upstream status it already used.
	UpstreamErrorMode string
}

// statusRecorder captures the response status code and body byte count so the
// access log can report them. WriteHeader/Write are only invoked once by net/http.
type statusRecorder struct {
	http.ResponseWriter
	status int
	bytes  int
}

func (s *statusRecorder) WriteHeader(code int) {
	if s.status == 0 {
		s.status = code
	}
	s.ResponseWriter.WriteHeader(code)
}

func (s *statusRecorder) Write(b []byte) (int, error) {
	if s.status == 0 {
		s.status = http.StatusOK
	}
	n, err := s.ResponseWriter.Write(b)
	s.bytes += n
	return n, err
}

// Middleware wraps an http.Handler with request tracing:
//   - validates/sanitizes the client-supplied IDs (untrusted), generating missing ones
//   - strips internal-only and LB-controlled headers; normalizes X-Forwarded-*
//   - injects X-Request-Id / X-Trace-Id on EVERY response, exactly once
//   - recovers panics into the standard 500 envelope (stack only logged)
//   - emits exactly one structured access log line per request
func Middleware(next http.Handler, cfg Config) http.Handler {
	if cfg.ServiceName == "" {
		cfg.ServiceName = "gateway"
	}
	if cfg.UpstreamErrorMode == "" {
		cfg.UpstreamErrorMode = "legacy"
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// 1. Validate / sanitize client-supplied IDs (untrusted).
		rawReq := r.Header.Get(gctrace.HeaderRequestID)
		rawTrace := r.Header.Get(gctrace.HeaderTraceID)
		rawSess := r.Header.Get(gctrace.HeaderSessionID)

		// Always prefer a valid app-supplied X-Request-Id; regenerate only if missing/invalid.
		reqID := gctrace.SanitizeHeader(gctrace.HeaderRequestID, rawReq)
		if reqID == "" {
			reqID = gctrace.NewID()
		}
		traceID := gctrace.SanitizeHeader(gctrace.HeaderTraceID, rawTrace)
		if traceID == "" {
			// If X-Trace-Id is missing/invalid, set it to the request ID.
			traceID = reqID
		}
		// If X-Session-Id is missing/invalid, leave empty.
		sessID := gctrace.SanitizeHeader(gctrace.HeaderSessionID, rawSess)

		ids := gctrace.IDs{SessionID: sessID, TraceID: traceID, RequestID: reqID}

		// 2. LB request id: logged separately, never used as the request id.
		var lbID string
		if cfg.LBRequestIDHeader != "" {
			lbID = r.Header.Get(cfg.LBRequestIDHeader)
		}

		// 3. Per-request state for metrics.
		st := &gctrace.State{Attempt: 1, LBRequestID: lbID, MatchedRoute: r.URL.Path}

		ctx := gctrace.WithState(gctrace.WithIDs(r.Context(), ids), st)
		r = r.WithContext(ctx)

		// 4. Strip/normalize internal and forwarded headers.
		gctrace.StripInternalHeaders(r, cfg.TrustedProxies)

		// 5. Inject response headers exactly once, BEFORE the handler writes the body.
		w.Header().Set(gctrace.HeaderRequestID, reqID)
		w.Header().Set(gctrace.HeaderTraceID, traceID)

		rec := &statusRecorder{ResponseWriter: w}
		log := gctrace.Logger(ctx, cfg.ServiceName)

		defer func() {
			// Panic recovery -> standard 500 (stack only logged, never leaked).
			if gctrace.Recover(ctx, rec, reqID, log) {
				st.Panic = true
			}
			status := rec.status
			if status == 0 {
				status = http.StatusOK
			}
			remoteIP := r.RemoteAddr
			if h, _, err := net.SplitHostPort(remoteIP); err == nil {
				remoteIP = h
			}
			// Exactly one access log line, including upstream fields and outcome.
			gctrace.AccessLog(ctx, gctrace.AccessRecord{
				Service:          cfg.ServiceName,
				Method:           r.Method,
				Path:             r.URL.Path,
				MatchedRoute:     st.MatchedRoute,
				UpstreamService:  st.UpstreamService,
				UpstreamStatus:   st.UpstreamStatus,
				UpstreamDuration: st.UpstreamDuration,
				Duration:         time.Since(start),
				Bytes:            rec.bytes,
				RemoteIP:         remoteIP,
				UserAgent:        r.UserAgent(),
				Attempt:          st.Attempt,
				Status:           status,
				Outcome:          st.Outcome,
				ErrorType:        st.ErrorType,
				LBRequestID:      st.LBRequestID,
				Panic:            st.Panic,
			})
		}()

		next.ServeHTTP(rec, r)
	})
}
