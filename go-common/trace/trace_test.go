package trace

import (
	"context"
	"net"
	"net/http"
	"strings"
	"testing"
)

func TestSanitizeHeader(t *testing.T) {
	cases := []struct {
		name string
		raw  string
		want string
	}{
		{"empty", "", ""},
		{"valid", "abcDEF123-_", "abcDEF123-_"},
		{"too_short", "short", ""},
		{"exactly_8", "12345678", "12345678"},
		{"exactly_64", strings.Repeat("a", 64), strings.Repeat("a", 64)},
		{"too_long", strings.Repeat("a", 65), ""},
		{"newline_injection", "abc\nX-Injected: 1", ""},
		{"cr_injection", "abc\r", ""},
		{"space", "abc def", ""},
		{"slash", "abc/def", ""},
		{"unicode", "abcé", ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := SanitizeHeader("X-Test", c.raw); got != c.want {
				t.Fatalf("SanitizeHeader(%q) = %q, want %q", c.raw, got, c.want)
			}
		})
	}
}

func TestNewID(t *testing.T) {
	for i := 0; i < 100; i++ {
		id := NewID()
		if !validID.MatchString(id) {
			t.Fatalf("NewID returned invalid id %q", id)
		}
	}
}

func TestIDValidationFlow(t *testing.T) {
	// missing request id -> generated; missing trace -> equals request id; missing session -> empty.
	reqID := SanitizeHeader(HeaderRequestID, "")
	if reqID != "" {
		t.Fatalf("expected empty request id, got %q", reqID)
	}
	reqID = NewID()
	traceID := SanitizeHeader(HeaderTraceID, "")
	if traceID != "" {
		t.Fatalf("expected empty trace id, got %q", traceID)
	}
	traceID = reqID // gateway fallback
	sessID := SanitizeHeader(HeaderSessionID, "bogus!")
	if sessID != "" {
		t.Fatalf("expected empty session id for invalid value, got %q", sessID)
	}
	ids := IDs{SessionID: sessID, TraceID: traceID, RequestID: reqID}
	if ids.TraceID != ids.RequestID {
		t.Fatalf("trace id should fall back to request id")
	}
}

func TestApplyToRequest(t *testing.T) {
	m := map[string]string{}
	ApplyToRequest(m, IDs{RequestID: "r1", TraceID: "t1", SessionID: "s1"})
	if m[HeaderRequestID] != "r1" || m[HeaderTraceID] != "t1" || m[HeaderSessionID] != "s1" {
		t.Fatalf("ApplyToRequest did not set all headers: %v", m)
	}
	m2 := map[string]string{}
	ApplyToRequest(m2, IDs{}) // all empty -> nothing set
	if len(m2) != 0 {
		t.Fatalf("ApplyToRequest set headers for empty ids: %v", m2)
	}
}

func TestApplyToHeader(t *testing.T) {
	h := http.Header{}
	ApplyToHeader(h, IDs{RequestID: "r1", TraceID: "t1"})
	if h.Get(HeaderRequestID) != "r1" || h.Get(HeaderTraceID) != "t1" {
		t.Fatalf("ApplyToHeader did not set headers: %v", h)
	}
}

func TestContextRoundTrip(t *testing.T) {
	ctx := WithState(WithIDs(context.Background(), IDs{RequestID: "r1"}), &State{Attempt: 1})
	if FromContext(ctx).RequestID != "r1" {
		t.Fatal("FromContext did not return stored ids")
	}
	if FromState(ctx) == nil || FromState(ctx).Attempt != 1 {
		t.Fatal("FromState did not return stored state")
	}
	if FromContext(nil).RequestID != "" {
		t.Fatal("FromContext(nil) should be zero value")
	}
}

func TestClassifyError(t *testing.T) {
	cases := []struct {
		msg      string
		outcome  string
		errtype  string
	}{
		{"dial tcp 127.0.0.1:8002: i/o timeout", OutcomeUpstreamTimeout, "timeout"},
		{"context deadline exceeded", OutcomeUpstreamTimeout, "timeout"},
		{"dial tcp 127.0.0.1:8002: connect: connection refused", OutcomeUpstreamUnreachable, "connection_error"},
		{"unexpected EOF", OutcomeUpstreamUnreachable, "connection_error"},
		{"some other error", "", "transport_error"},
	}
	for _, c := range cases {
		o, e := ClassifyError(&fakeErr{c.msg})
		if o != c.outcome || e != c.errtype {
			t.Fatalf("ClassifyError(%q) = (%q,%q), want (%q,%q)", c.msg, o, e, c.outcome, c.errtype)
		}
	}
}

type fakeErr struct{ s string }

func (e *fakeErr) Error() string { return e.s }

func TestStripInternalHeaders_DirectFacing(t *testing.T) {
	r := &http.Request{
		Header: http.Header{
			"X-B3-Traceid":     {"xyz"},
			"X-Internal-Foo":   {"bar"},
			"X-User-Id":        {"u1"},
			"X-Auth-Token":     {"secret"},
			"X-Forwarded-For":  {"1.2.3.4"},
			"X-Forwarded-Host": {"evil.com"},
			"Traceparent":      {"abc"},
			"True-Client-Ip":   {"9.9.9.9"},
			"Clientid":         {"889999"}, // must be preserved (routing)
		},
		RemoteAddr: "203.0.113.5:55555",
		Host:       "gw.example.com",
	}
	StripInternalHeaders(r, nil)
	for _, k := range []string{"X-B3-Traceid", "X-Internal-Foo", "X-User-Id", "X-Auth-Token", "Traceparent", "True-Client-Ip", "X-Forwarded-Host"} {
		if r.Header.Get(k) != "" {
			t.Fatalf("header %s should have been stripped, got %q", k, r.Header.Get(k))
		}
	}
	if r.Header.Get("Clientid") != "889999" {
		t.Fatalf("clientId must be preserved, got %q", r.Header.Get("Clientid"))
	}
	if got := r.Header.Get("X-Forwarded-For"); got != "203.0.113.5" {
		t.Fatalf("direct-facing XFF should be real peer, got %q", got)
	}
	if got := r.Header.Get("X-Forwarded-Proto"); got != "http" {
		t.Fatalf("direct-facing proto should be http, got %q", got)
	}
	if got := r.Header.Get("X-Forwarded-Host"); got != "gw.example.com" {
		t.Fatalf("direct-facing host should be r.Host, got %q", got)
	}
}

func TestStripInternalHeaders_TrustedProxy(t *testing.T) {
	_, cidr, _ := net.ParseCIDR("10.0.0.0/8")
	r := &http.Request{
		Header: http.Header{
			"X-Forwarded-For":  {"203.0.113.5"},
			"X-Forwarded-Host": {"real.example.com"},
		},
		RemoteAddr: "10.0.0.2:4444",
		Host:       "gw.example.com",
	}
	StripInternalHeaders(r, []net.IPNet{*cidr})
	if got := r.Header.Get("X-Forwarded-For"); got != "203.0.113.5, 10.0.0.2" {
		t.Fatalf("trusted proxy should append hop, got %q", got)
	}
	if got := r.Header.Get("X-Forwarded-Host"); got != "real.example.com" {
		t.Fatalf("trusted proxy host should be preserved, got %q", got)
	}
}

func TestStandardError(t *testing.T) {
	body := StandardError("upstream_unreachable", "connection refused", "req-1")
	if body.Error.Code != "upstream_unreachable" || body.Error.RequestID != "req-1" {
		t.Fatalf("unexpected error body %+v", body)
	}
}
