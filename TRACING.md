# Request tracing & structured logging

End-to-end request tracing and structured JSON logging for the API gateway, built
on the shared `github.com/msys/go-common/trace` package (see `go-common/`).

## What changed

**Shared library (`go-common/`, module `github.com/msys/go-common`)**
- `go-common/trace/trace.go` — context helpers (`WithIDs`/`FromContext`,
  `WithState`/`FromState`), `SanitizeHeader` (8–64 chars of `[A-Za-z0-9_-]`, never
  logs the raw value), `NewID` (UUID v4), `Logger(ctx, service)`,
  `ApplyToHeader`/`ApplyToRequest` (forward the three IDs), `StandardError`,
  `Recover` (panic → standard 500, stack only logged), `StripInternalHeaders`
  (drops internal-only + normalizes `X-Forwarded-*`), `ClassifyError`, and
  `AccessLog` (single line per request). Versioning tag `v0.1.0`.

**Gateway**
- `middleware/trace.go` — `http.Handler` middleware wrapping `beego.BeeApp.Handlers`.
  Validates/sanitizes the client-supplied IDs (untrusted), generates missing ones,
  strips internal-only headers, injects `X-Request-Id`/`X-Trace-Id` on **every**
  response exactly once, recovers panics into the standard 500 envelope, and emits
  exactly one access-log line per request (including upstream fields + outcome).
- `conf/api/client.go` — `SendRequest(ctx)` now forwards the three validated IDs to
  the upstream and records `upstream_service`/`upstream_status`/`upstream_duration`
  + `outcome`/`error_type` into the per-request `State`.
- `controllers/functions/api_requests.go` — every function takes `ctx
  context.Context`, forwards the IDs, and uses the context-aware logger. All
  request-/response-body debug dumps (PII: account numbers, amounts, mobile
  numbers, tokens) were removed.
- `controllers/service_requests.go` — threads `ctx` to the upstream functions,
  migrates request-path logging to the context-aware logger, and attaches
  `request_id` to gateway-built error bodies (additive field).
- `structs/responses/responses.go` — added `request_id` (omitempty) to the gateway
  error response structs.
- `gatewaylog/slogadapter.go` — a Beego `logs` adapter that writes JSON via `slog`,
  so existing `logs.*` calls keep working with a single log format on stdout.
- `main.go` — configures `slog` from `GATEWAY_LOG_LEVEL`/`GATEWAY_LOG_FORMAT`,
  routes Beego logs through the adapter, disables Beego's built-in recover/access
  log (the middleware owns those), and wraps the handler.
- `go.mod` requires `github.com/msys/go-common v0.1.0` with a local `replace`;
  `go.work` ties the two modules together for local dev.

## Configuration (env vars)
- `GATEWAY_LOG_LEVEL` — debug|info|warn|error (default `info`)
- `GATEWAY_LOG_FORMAT` — json|text (default `json`)
- `TRUSTED_PROXIES` — comma-separated CIDRs of proxies allowed to set
  `X-Forwarded-*`; **empty = directly internet-facing** (default; drops + sets
  XFF/Host/Proto/X-Real-IP from the real connection)
- `LB_REQUEST_ID_HEADER` — if your LB sends its own request id under a separate
  header, name it here; it is logged as `lb_request_id` and is **never** used as
  the request id
- `UPSTREAM_ERROR_MODE` — `legacy` (default, current behavior) | `http` (future:
  502/504 + standard error shape). **Not enabled.**

## Behavior notes / confirmations
- Client-supplied IDs are untrusted: validated to 8–64 `[A-Za-z0-9_-]`. Invalid or
  missing `X-Request-Id` → generated UUID; `X-Trace-Id` missing/invalid → set to the
  request id; `X-Session-Id` missing/invalid → empty. IDs are **not** used for auth,
  routing, or DB keys.
- Gateway currently does **not** retry upstream calls, so the "same request id on
  retry" rule is implemented but currently a no-op; non-idempotent retries are flagged
  if added later.
- Legacy behavior preserved: upstream failures still return HTTP 200 with the
  upstream's body shape; `request_id` is added to gateway-built error bodies and is
  always present on the `X-Request-Id` response header.
- Internal-only headers stripped: `Forwarded`, `X-Forwarded-Port`,
  `X-Forwarded-Prefix`, `traceparent`, `tracestate`, `X-B3-*`, `True-Client-IP`,
  `CF-Connecting-IP`, `X-Internal-*`, `X-User-*`, `X-Auth-*`, and the three trace
  headers (re-set from validated values on the upstream call).

## Sample log output (one access line per request)

Success:
```json
{"time":"2026-11-27T12:00:00Z","level":"INFO","msg":"access","service":"gateway","session_id":"ses-VALID12","trace_id":"trc-VALID12","request_id":"req-VALID12","method":"POST","path":"/v2/api/credit-account","matched_route":"/v2/api/credit-account","upstream_service":"http://127.0.0.1:8003","upstream_status":200,"upstream_duration_ms":42,"duration_ms":45,"bytes":312,"remote_ip":"203.0.113.9","user_agent":"msys-flutter/1.4.0","attempt":1,"status":200,"outcome":"","error_type":"","lb_request_id":"","panic":false}
```

Upstream 5xx (gateway still returns 200 to the app, body carries the upstream
envelope):
```json
{"time":"2026-11-27T12:00:01Z","level":"ERROR","msg":"access","service":"gateway","session_id":"","trace_id":"trc-abc12345","request_id":"req-abc12345","method":"GET","path":"/v2/api/name-inquiry/233","matched_route":"/v2/api/name-inquiry/233","upstream_service":"http://127.0.0.1:8003","upstream_status":503,"upstream_duration_ms":30,"duration_ms":33,"bytes":0,"remote_ip":"203.0.113.9","user_agent":"msys-flutter/1.4.0","attempt":1,"status":200,"outcome":"upstream_5xx","error_type":"http_5xx","lb_request_id":"","panic":false}
```

Panic (stack only in the `panic recovered` line, never in the client response):
```json
{"time":"2026-11-27T12:00:02Z","level":"ERROR","msg":"panic recovered","service":"gateway","session_id":"","trace_id":"trc-abc12345","request_id":"req-abc12345","panic":"boom","stack":"goroutine 1 [...]"}
{"time":"2026-11-27T12:00:02Z","level":"ERROR","msg":"access","service":"gateway","session_id":"","trace_id":"trc-abc12345","request_id":"req-abc12345","method":"GET","path":"/v2/api/x","matched_route":"/v2/api/x","upstream_service":"","upstream_status":0,"upstream_duration_ms":0,"duration_ms":2,"bytes":79,"remote_ip":"203.0.113.9","user_agent":"msys-flutter/1.4.0","attempt":1,"status":500,"outcome":"","error_type":"","lb_request_id":"","panic":true}
```

## Upstream error shapes found (today)
The gateway **passes through** the raw upstream JSON. Each client backend returns an
envelope of the form `{"data":{"StatusCode":int,"StatusDesc":string,"Result":...,
"Client":string}}` (see `structs/responses/responses.go`). On error the upstream
puts a non-200 `StatusCode` in the **body** while the gateway returns HTTP **200**
(`c.ServeJSON()` defaults to 200). Some older functions return a bare
`err.Error()` string on transport failure (connection refused / timeout).

⚠️ **Do not normalize the upstream envelope without sign-off**: the Flutter app
reads `StatusCode`/`StatusDesc` from the JSON **body**, not from the HTTP status.
Changing the shape would break the current app. `UPSTREAM_ERROR_MODE=http` is the
gated path for a future 502/504 + standard-shape transition.

## How to build & test
```bash
export GOPRIVATE=github.com/msys/go-common
cd go-common && go mod tidy   # needs network to fetch github.com/google/uuid
cd ..
go mod tidy
go build ./...
go test ./...
```
`go.work` already links `./` and `./go-common` for local development.

## Open questions (see chat)
1. Is an LB / TLS-terminating proxy in front? `TRUSTED_PROXIES` defaults to
   directly-internet-facing.
2. What header does the LB use for its own request id (`LB_REQUEST_ID_HEADER`), and
   does it rewrite `X-Request-Id` (which would clobber the app's value)?
3. `go.mod` says `go 1.22` but `.github/workflows/deploy.yml` pins `go-version:
   '1.20'` — bump to 1.22 (slog needs 1.21+).
4. File logging was removed (stdout JSON only) — want a file handler too?
5. Approve enabling `UPSTREAM_ERROR_MODE=http` (502/504 + standard shape)?
6. `go.sum` for `github.com/google/uuid` must be generated via `go mod tidy` in a
   networked environment (could not be generated in this sandbox).
