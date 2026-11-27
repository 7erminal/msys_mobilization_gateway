// Package trace provides request-tracing primitives shared by the API gateway and
// the downstream services (customer management, authentication, system, payment,
// transaction, ...). It is intentionally dependency-light: only the standard
// library plus google/uuid.
//
// It exposes:
//   - context helpers: WithIDs / FromContext, WithState / FromState
//   - ID validation: SanitizeHeader, NewID
//   - a structured logger: Logger, Configure
//   - outbound helpers that forward the three trace headers: ApplyToHeader, ApplyToRequest
//   - the standard error envelope: StandardError, Recover
//   - request hygiene: StripInternalHeaders, ClassifyError
//   - the single access-log line: AccessLog
package trace

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"regexp"
	"runtime/debug"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Header names used to carry trace context between the app, the gateway, and upstreams.
const (
	HeaderSessionID = "X-Session-Id"
	HeaderTraceID   = "X-Trace-Id"
	HeaderRequestID = "X-Request-Id"
)

// IDs carries the three client-supplied (validated) trace identifiers.
// They are treated as UNTRUSTED: callers must sanitize before use and must
// never use them for auth, authorization, routing, or as database keys.
type IDs struct {
	SessionID string
	TraceID   string
	RequestID string
}

// State is a per-request, mutable bag for metrics gathered while handling a
// request. It is stored in the request context and populated by the outbound
// HTTP helper so the access log can report upstream timing/result.
type State struct {
	MatchedRoute      string
	UpstreamService   string
	UpstreamStatus    int
	UpstreamDuration  time.Duration
	Outcome           string // upstream_timeout | upstream_unreachable | upstream_5xx
	ErrorType         string
	Attempt           int
	LBRequestID       string
	Bytes             int
	Panic             bool
}

// Outcome values reported in the access log.
const (
	OutcomeUpstreamTimeout     = "upstream_timeout"
	OutcomeUpstreamUnreachable = "upstream_unreachable"
	OutcomeUpstream5xx         = "upstream_5xx"
)

type contextKey string

// Exported context keys so callers can store/retrieve IDs and State.
const (
	IDsKey   contextKey = "msys/trace:ids"
	StateKey contextKey = "msys/trace:state"
)

// WithIDs returns a copy of ctx carrying the validated IDs.
func WithIDs(ctx context.Context, ids IDs) context.Context {
	return context.WithValue(ctx, IDsKey, ids)
}

// FromContext returns the IDs stored in ctx, or the zero value if absent.
func FromContext(ctx context.Context) IDs {
	if ctx == nil {
		return IDs{}
	}
	if ids, ok := ctx.Value(IDsKey).(IDs); ok {
		return ids
	}
	return IDs{}
}

// WithState stores the per-request State in ctx.
func WithState(ctx context.Context, st *State) context.Context {
	return context.WithValue(ctx, StateKey, st)
}

// FromState returns the per-request State, or nil if absent.
func FromState(ctx context.Context) *State {
	if ctx == nil {
		return nil
	}
	if st, ok := ctx.Value(StateKey).(*State); ok {
		return st
	}
	return nil
}

// validID matches the allowed header value: 8-64 chars of [A-Za-z0-9_-].
var validID = regexp.MustCompile(`^[A-Za-z0-9_-]{8,64}$`)

// NewID returns a new random UUID v4.
func NewID() string { return uuid.NewString() }

// IsValidID reports whether s is a valid trace header value (8-64 chars of [A-Za-z0-9_-]).
func IsValidID(s string) bool { return validID.MatchString(s) }

// ConfigureCapture installs a slog JSON handler writing to w. It is intended for
// tests that need to assert on structured log output.
func ConfigureCapture(w io.Writer) { slog.SetDefault(slog.New(slog.NewJSONHandler(w, nil))) }

// SanitizeHeader validates an inbound header value.
// Returns "" when the value is missing, longer than 64 chars, or contains any
// disallowed character (including control chars / newlines). The raw value is
// NEVER returned and MUST NOT be logged.
func SanitizeHeader(name, raw string) string {
	if raw == "" {
		return ""
	}
	if len(raw) > 64 {
		return ""
	}
	if !validID.MatchString(raw) {
		return ""
	}
	return raw
}

// ApplyToHeader sets the three trace headers on an http.Header (only when non-empty).
func ApplyToHeader(h http.Header, ids IDs) {
	if h == nil {
		return
	}
	if ids.RequestID != "" {
		h.Set(HeaderRequestID, ids.RequestID)
	}
	if ids.TraceID != "" {
		h.Set(HeaderTraceID, ids.TraceID)
	}
	if ids.SessionID != "" {
		h.Set(HeaderSessionID, ids.SessionID)
	}
}

// ApplyToRequest sets the three trace headers on a map[string]string. This is the
// shape used by the gateway's httplib Request.HeaderField, but it works for any
// outbound call that takes a header map.
func ApplyToRequest(h map[string]string, ids IDs) {
	if h == nil {
		return
	}
	if ids.RequestID != "" {
		h[HeaderRequestID] = ids.RequestID
	}
	if ids.TraceID != "" {
		h[HeaderTraceID] = ids.TraceID
	}
	if ids.SessionID != "" {
		h[HeaderSessionID] = ids.SessionID
	}
}

// Configure installs a slog default logger. format is "json" or "text" (default
// json); level is one of debug|info|warn|error (default info). Every line uses
// the configured handler; structured fields are added per-log via Logger/AccessLog.
func Configure(level, format string) {
	var l slog.Level
	switch strings.ToLower(level) {
	case "debug":
		l = slog.LevelDebug
	case "warn", "warning":
		l = slog.LevelWarn
	case "error":
		l = slog.LevelError
	default:
		l = slog.LevelInfo
	}
	opts := &slog.HandlerOptions{Level: l}
	var h slog.Handler
	if strings.ToLower(format) == "text" {
		h = slog.NewTextHandler(os.Stdout, opts)
	} else {
		h = slog.NewJSONHandler(os.Stdout, opts)
	}
	slog.SetDefault(slog.New(h))
}

// Logger returns a slog.Logger annotated with service and trace identifiers.
func Logger(ctx context.Context, service string) *slog.Logger {
	ids := FromContext(ctx)
	return slog.With(
		slog.String("service", service),
		slog.String("session_id", ids.SessionID),
		slog.String("trace_id", ids.TraceID),
		slog.String("request_id", ids.RequestID),
	)
}

// ErrorBody is the standard gateway error envelope.
type ErrorBody struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail is a single error entry in the standard envelope.
type ErrorDetail struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

// StandardError builds the standard error envelope.
func StandardError(code, message, requestID string) ErrorBody {
	return ErrorBody{Error: ErrorDetail{Code: code, Message: message, RequestID: requestID}}
}

// Recover catches a panic in the calling goroutine, logs it (with stack) using the
// provided logger, and writes the standard 500 envelope to w. The stack is never
// sent to the client. It returns true if a panic was recovered.
func Recover(ctx context.Context, w http.ResponseWriter, requestID string, log *slog.Logger) (recovered bool) {
	if r := recover(); r != nil {
		if log == nil {
			log = slog.Default()
		}
		log.Log(ctx, slog.LevelError, "panic recovered",
			slog.Any("panic", r),
			slog.String("stack", string(debug.Stack())),
		)
		if w.Header().Get("Content-Type") == "" {
			w.Header().Set("Content-Type", "application/json")
		}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(StandardError("internal", "internal server error", requestID))
		return true
	}
	return false
}

// ClassifyError maps a transport/HTTP error to an access-log outcome and error
// type. It returns ("", "") when err is nil.
func ClassifyError(err error) (outcome, errorType string) {
	if err == nil {
		return "", ""
	}
	msg := err.Error()
	switch {
	case strings.Contains(msg, "context deadline exceeded"),
		strings.Contains(msg, "Client.Timeout"),
		strings.Contains(msg, "i/o timeout"),
		strings.Contains(msg, "Timeout exceeded"):
		return OutcomeUpstreamTimeout, "timeout"
	case strings.Contains(msg, "connection refused"),
		strings.Contains(msg, "no such host"),
		strings.Contains(msg, "network is unreachable"),
		strings.Contains(msg, "cannot assign requested address"),
		strings.Contains(msg, "EOF"):
		return OutcomeUpstreamUnreachable, "connection_error"
	default:
		return "", "transport_error"
	}
}

// AccessRecord holds the fields emitted on the single per-request access log line.
type AccessRecord struct {
	Service          string
	Method           string
	Path             string
	MatchedRoute     string
	UpstreamService  string
	UpstreamStatus   int
	UpstreamDuration time.Duration
	Duration         time.Duration
	Bytes            int
	RemoteIP         string
	UserAgent        string
	Attempt          int
	Status           int
	Outcome          string
	ErrorType        string
	LBRequestID      string
	Panic            bool
}

// AccessLog emits exactly one structured access line for the request.
func AccessLog(ctx context.Context, rec AccessRecord) {
	ids := FromContext(ctx)
	lvl := slog.LevelInfo
	if rec.Panic || rec.Status >= 500 || rec.Outcome != "" {
		lvl = slog.LevelError
	}
	slog.Log(ctx, lvl, "access",
		slog.String("service", rec.Service),
		slog.String("session_id", ids.SessionID),
		slog.String("trace_id", ids.TraceID),
		slog.String("request_id", ids.RequestID),
		slog.String("method", rec.Method),
		slog.String("path", rec.Path),
		slog.String("matched_route", rec.MatchedRoute),
		slog.String("upstream_service", rec.UpstreamService),
		slog.Int("upstream_status", rec.UpstreamStatus),
		slog.Int64("upstream_duration_ms", rec.UpstreamDuration.Milliseconds()),
		slog.Int64("duration_ms", rec.Duration.Milliseconds()),
		slog.Int("bytes", rec.Bytes),
		slog.String("remote_ip", rec.RemoteIP),
		slog.String("user_agent", rec.UserAgent),
		slog.Int("attempt", rec.Attempt),
		slog.Int("status", rec.Status),
		slog.String("outcome", rec.Outcome),
		slog.String("error_type", rec.ErrorType),
		slog.String("lb_request_id", rec.LBRequestID),
		slog.Bool("panic", rec.Panic),
	)
}

// internalHeader reports whether an inbound header is reserved for the gateway or
// internal callers and must be stripped entirely.
func internalHeader(name string) bool {
	n := strings.ToLower(name)
	switch n {
	case "forwarded", "x-forwarded-port", "x-forwarded-prefix",
		"traceparent", "tracestate", "true-client-ip", "cf-connecting-ip":
		return true
	}
	if strings.HasPrefix(n, "x-b3-") ||
		strings.HasPrefix(n, "x-internal-") ||
		strings.HasPrefix(n, "x-user-") ||
		strings.HasPrefix(n, "x-auth-") {
		return true
	}
	return false
}

// StripInternalHeaders removes headers that only the gateway or internal callers
// may set, and normalizes X-Forwarded-For / -Host / -Proto / X-Real-IP.
//
// trusted lists the IP ranges of proxies allowed to set X-Forwarded-* (an empty
// slice means the gateway is directly internet-facing). The client-supplied trace
// headers are removed here; the validated values are forwarded upstream separately
// by the outbound helper.
func StripInternalHeaders(r *http.Request, trusted []net.IPNet) {
	for k := range r.Header {
		if internalHeader(k) {
			r.Header.Del(k)
		}
	}
	r.Header.Del(HeaderSessionID)
	r.Header.Del(HeaderTraceID)
	r.Header.Del(HeaderRequestID)

	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	remoteIP := net.ParseIP(host)
	trustedPeer := len(trusted) > 0 && remoteIP != nil && isTrusted(remoteIP, trusted)

	if trustedPeer {
		if existing := r.Header.Get("X-Forwarded-For"); existing != "" {
			r.Header.Set("X-Forwarded-For", existing+", "+host)
		} else {
			r.Header.Set("X-Forwarded-For", host)
		}
		r.Header.Set("X-Real-IP", host)
		// X-Forwarded-Proto / X-Forwarded-Host are trusted from the proxy.
	} else {
		r.Header.Del("X-Forwarded-For")
		r.Header.Del("X-Forwarded-Host")
		r.Header.Del("X-Forwarded-Proto")
		r.Header.Del("X-Real-IP")
		r.Header.Set("X-Forwarded-For", host)
		r.Header.Set("X-Forwarded-Proto", scheme(r))
		r.Header.Set("X-Forwarded-Host", r.Host)
		r.Header.Set("X-Real-IP", host)
	}
}

func scheme(r *http.Request) string {
	if r.TLS != nil {
		return "https"
	}
	return "http"
}

func isTrusted(ip net.IP, nets []net.IPNet) bool {
	for _, n := range nets {
		if n.Contains(ip) {
			return true
		}
	}
	return false
}
