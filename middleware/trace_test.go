package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"msys_api_gateway/conf/api"

	gctrace "github.com/msys/go-common/trace"
)

// captureLogs redirects the slog default logger to a buffer and returns it.
func captureLogs(t *testing.T) *bytes.Buffer {
	t.Helper()
	buf := &bytes.Buffer{}
	gctrace.ConfigureCapture(buf)
	return buf
}

// recordingUpstream returns a test server that records the headers it received
// and replies 200 with a tiny JSON body.
func recordingUpstream(t *testing.T, status int) (*httptest.Server, *http.Header, *sync.Mutex) {
	t.Helper()
	var mu sync.Mutex
	var got http.Header
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		got = r.Header.Clone()
		mu.Unlock()
		w.WriteHeader(status)
		_, _ = w.Write([]byte(`{"data":{"StatusCode":` + itoa(status) + `}}`))
	}))
	return srv, &got, &mu
}

func itoa(i int) string {
	if i == 200 {
		return "200"
	}
	return "500"
}

// --- Middleware: response headers always present, exactly once ---

func TestResponseHeadersPresentOnSuccess(t *testing.T) {
	buf := captureLogs(t)
	defer buf.Reset()
	h := Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"ok":true}`))
	}), Config{})

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/v2/api/name-inquiry/233", nil)
	h.ServeHTTP(rec, req)

	if rec.Header().Get(gctrace.HeaderRequestID) == "" {
		t.Fatal("missing X-Request-Id on success response")
	}
	if rec.Header().Get(gctrace.HeaderTraceID) == "" {
		t.Fatal("missing X-Trace-Id on success response")
	}
	// exactly once
	if len(rec.Header().Values(gctrace.HeaderRequestID)) != 1 {
		t.Fatal("X-Request-Id duplicated")
	}
}

func TestResponseHeadersPresentOnPanic(t *testing.T) {
	buf := captureLogs(t)
	defer buf.Reset()
	h := Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("boom")
	}), Config{})

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/x", nil)
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", rec.Code)
	}
	body := rec.Body.String()
	if strings.Contains(body, "boom") || strings.Contains(body, "stack") {
		t.Fatalf("panic detail leaked to client: %s", body)
	}
	if rec.Header().Get(gctrace.HeaderRequestID) == "" {
		t.Fatal("missing X-Request-Id on panic response")
	}
}

func TestResponseHeadersPresentOnUpstreamError(t *testing.T) {
	buf := captureLogs(t)
	defer buf.Reset()
	// inner handler calls a dead upstream (connection refused) then returns 200.
	h := Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ids := gctrace.FromContext(r.Context())
		req := api.NewRequest("http://127.0.0.1:1", "/x", api.POST) // nothing listening
		client := api.Client{Request: req, Type_: "params"}
		_, _ = client.SendRequest(r.Context())
		_ = ids
		w.WriteHeader(200)
	}), Config{})

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/v2/api/x", nil)
	h.ServeHTTP(rec, req)

	if rec.Header().Get(gctrace.HeaderRequestID) == "" {
		t.Fatal("missing X-Request-Id when upstream errored")
	}
}

// --- ID validation / generation ---

func TestValidHeadersEchoedAndForwarded(t *testing.T) {
	buf := captureLogs(t)
	defer buf.Reset()
	srv, got, _ := recordingUpstream(t, 200)
	defer srv.Close()

	h := Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ids := gctrace.FromContext(r.Context())
		req := api.NewRequest(srv.URL, "/x", api.POST)
		client := api.Client{Request: req, Type_: "params"}
		resp, err := client.SendRequest(r.Context())
		if err != nil {
			t.Errorf("SendRequest: %v", err)
		} else {
			resp.Body.Close()
		}
		w.WriteHeader(200)
	}), Config{})

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/v2/api/x", nil)
	req.Header.Set(gctrace.HeaderRequestID, "req-VALID12")
	req.Header.Set(gctrace.HeaderTraceID, "trc-VALID12")
	req.Header.Set(gctrace.HeaderSessionID, "ses-VALID12")
	h.ServeHTTP(rec, req)

	if rec.Header().Get(gctrace.HeaderRequestID) != "req-VALID12" {
		t.Fatalf("valid request id not echoed: %q", rec.Header().Get(gctrace.HeaderRequestID))
	}
	if rec.Header().Get(gctrace.HeaderTraceID) != "trc-VALID12" {
		t.Fatalf("valid trace id not echoed: %q", rec.Header().Get(gctrace.HeaderTraceID))
	}
	// session id is a label only -> NOT echoed on response
	if rec.Header().Get(gctrace.HeaderSessionID) != "" {
		t.Fatal("session id must not be set on response")
	}
	// forwarded upstream unchanged
	if (*got).Get(gctrace.HeaderRequestID) != "req-VALID12" {
		t.Fatalf("upstream did not receive validated request id: %q", (*got).Get(gctrace.HeaderRequestID))
	}
	if (*got).Get(gctrace.HeaderTraceID) != "trc-VALID12" {
		t.Fatal("upstream did not receive validated trace id")
	}
	if (*got).Get(gctrace.HeaderSessionID) != "ses-VALID12" {
		t.Fatal("upstream did not receive validated session id")
	}
}

func TestInvalidAndMissingIDs(t *testing.T) {
	buf := captureLogs(t)
	defer buf.Reset()
	h := Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ids := gctrace.FromContext(r.Context())
		w.WriteHeader(200)
		// ensure generated request id is valid format
		if !gctrace.IsValidID(ids.RequestID) {
			t.Errorf("generated request id invalid: %q", ids.RequestID)
		}
		// trace falls back to request id when missing/invalid
		if ids.TraceID != ids.RequestID {
			t.Errorf("trace id should equal request id, got %q vs %q", ids.TraceID, ids.RequestID)
		}
	}), Config{})

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/x", nil)
	// invalid: too short + disallowed chars + newline injection attempt
	req.Header.Set(gctrace.HeaderRequestID, "ab\nX-Injected:1")
	req.Header.Set(gctrace.HeaderTraceID, "!!bad!!")
	req.Header.Set(gctrace.HeaderSessionID, "way-too-long-session-value-that-exceeds-the-allowed-sixty-four-character-limit-by-a-lot")
	h.ServeHTTP(rec, req)

	if rec.Header().Get(gctrace.HeaderRequestID) == "ab\nX-Injected:1" {
		t.Fatal("invalid raw request id was echoed (injection!)")
	}
	if rec.Header().Get(gctrace.HeaderRequestID) == "" {
		t.Fatal("request id not generated for invalid input")
	}
	// the raw invalid value must never appear in logs
	if strings.Contains(buf.String(), "X-Injected") {
		t.Fatal("raw invalid header value leaked into logs")
	}
}

func TestMissingHeadersGenerateIDs(t *testing.T) {
	cases := []struct {
		name    string
		reqID   string
		traceID string
		sessID  string
	}{
		{"all missing", "", "", ""},
		{"request only", "req-ONLY1234", "", ""},
		{"trace only", "", "trc-ONLY1234", ""},
		{"session only", "", "", "ses-ONLY1234"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			buf := captureLogs(t)
			defer buf.Reset()
			h := Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ids := gctrace.FromContext(r.Context())
				w.WriteHeader(200)
				if ids.RequestID == "" {
					t.Error("request id should be generated")
				}
				if !gctrace.IsValidID(ids.RequestID) {
					t.Errorf("generated request id invalid: %q", ids.RequestID)
				}
			}), Config{})
			rec := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/x", nil)
			if c.reqID != "" {
				req.Header.Set(gctrace.HeaderRequestID, c.reqID)
			}
			if c.traceID != "" {
				req.Header.Set(gctrace.HeaderTraceID, c.traceID)
			}
			if c.sessID != "" {
				req.Header.Set(gctrace.HeaderSessionID, c.sessID)
			}
			h.ServeHTTP(rec, req)
		})
	}
}

// --- Internal-only header stripping ---

func TestInternalHeadersStripped(t *testing.T) {
	buf := captureLogs(t)
	defer buf.Reset()
	h := Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, k := range []string{"X-B3-Traceid", "X-Internal-Foo", "X-User-Id", "X-Auth-Token", "Traceparent", "True-Client-Ip", "X-Forwarded-For"} {
			if r.Header.Get(k) != "" {
				t.Errorf("header %s was not stripped: %q", k, r.Header.Get(k))
			}
		}
		// clientId (routing) must survive
		if r.Header.Get("clientId") != "889999" {
			t.Errorf("clientId stripped: %q", r.Header.Get("clientId"))
		}
		w.WriteHeader(200)
	}), Config{})

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/x", nil)
	req.Header.Set("clientId", "889999")
	req.Header.Set("X-B3-Traceid", "abc")
	req.Header.Set("X-Internal-Foo", "bar")
	req.Header.Set("X-User-Id", "u1")
	req.Header.Set("X-Auth-Token", "secret")
	req.Header.Set("Traceparent", "tp")
	req.Header.Set("True-Client-Ip", "9.9.9.9")
	req.Header.Set("X-Forwarded-For", "1.2.3.4")
	// direct-facing: XFF must be reset to the real peer
	req.RemoteAddr = "203.0.113.9:5555"
	h.ServeHTTP(rec, req)
}

// --- Access log: exactly once, with upstream fields ---

func TestAccessLogOnceWithUpstreamFields(t *testing.T) {
	buf := captureLogs(t)
	defer buf.Reset()
	srv, _, _ := recordingUpstream(t, 200)
	defer srv.Close()

	h := Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := api.NewRequest(srv.URL, "/x", api.POST)
		client := api.Client{Request: req, Type_: "params"}
		resp, err := client.SendRequest(r.Context())
		if err != nil {
			t.Errorf("SendRequest: %v", err)
			w.WriteHeader(502)
			return
		}
		resp.Body.Close()
		w.WriteHeader(200)
	}), Config{})

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/v2/api/name-inquiry/233", nil)
	req.Header.Set(gctrace.HeaderRequestID, "req-VALID12")
	h.ServeHTTP(rec, req)

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	var access int
	var last map[string]interface{}
	for _, l := range lines {
		if l == "" {
			continue
		}
		var m map[string]interface{}
		if err := json.Unmarshal([]byte(l), &m); err != nil {
			continue
		}
		if m["msg"] == "access" {
			access++
			last = m
		}
	}
	if access != 1 {
		t.Fatalf("expected exactly 1 access log line, got %d (buf=%s)", access, buf.String())
	}
	if last["upstream_service"] == nil || last["upstream_service"] == "" {
		t.Fatalf("access log missing upstream_service: %v", last)
	}
	if last["upstream_status"] != float64(200) {
		t.Fatalf("access log missing upstream_status: %v", last)
	}
}

func TestAccessLogUpstream5xxOutcome(t *testing.T) {
	buf := captureLogs(t)
	defer buf.Reset()
	srv, _, _ := recordingUpstream(t, 500)
	defer srv.Close()

	h := Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := api.NewRequest(srv.URL, "/x", api.POST)
		client := api.Client{Request: req, Type_: "params"}
		resp, err := client.SendRequest(r.Context())
		if err != nil {
			t.Errorf("SendRequest: %v", err)
			return
		}
		resp.Body.Close()
		w.WriteHeader(200)
	}), Config{})

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/x", nil)
	h.ServeHTTP(rec, req)

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	for _, l := range lines {
		if l == "" {
			continue
		}
		var m map[string]interface{}
		if err := json.Unmarshal([]byte(l), &m); err != nil {
			continue
		}
		if m["msg"] == "access" {
			if m["outcome"] != gctrace.OutcomeUpstream5xx {
				t.Fatalf("expected outcome %q, got %v", gctrace.OutcomeUpstream5xx, m["outcome"])
			}
			if m["upstream_status"] != float64(500) {
				t.Fatalf("expected upstream_status 500, got %v", m["upstream_status"])
			}
			return
		}
	}
	t.Fatal("no access log line found")
}
