# go-common

Shared Go module for the MSYS services. Currently hosts the `trace` package used by
the API gateway and the downstream services (customer management, authentication,
system, payment, transaction, …) for end-to-end request tracing and structured
logging.

Module path: `github.com/msys/go-common`
Trace package: `github.com/msys/go-common/trace`

## Design

Deliberately dependency-light: only the standard library plus `github.com/google/uuid`.

The `trace` package provides:

- Context helpers: `WithIDs` / `FromContext`, `WithState` / `FromState`
- ID validation: `SanitizeHeader` (8–64 chars of `[A-Za-z0-9_-]`, never logs raw), `NewID`
- Structured logger: `Configure` (sets `slog` default), `Logger(ctx, service)`
- Outbound helpers that forward the three trace headers:
  `ApplyToHeader(http.Header, ids)`, `ApplyToRequest(map[string]string, ids)`
- Standard error envelope: `StandardError`, `Recover` (panic recovery, never leaks stack)
- Request hygiene: `StripInternalHeaders`, `ClassifyError`
- Single access-log line: `AccessLog`

## Versioning

Tag releases as `v0.1.0`, `v0.2.0`, … (this initial drop is intended to be `v0.1.0`).

```
git tag v0.1.0
git push origin v0.1.0
```

## GOPRIVATE

The module is private, so tell Go not to use the public proxy / checksum database:

```
export GOPRIVATE=github.com/msys/go-common
# if you pull it from a private GitHub org over SSH:
export GONOSUMCHECK=1            # or configure GONOSUMDB / GOSUMDB=off
git config --global url."git@github.com:".insteadOf "https://github.com/"
```

## Local development (go.work / replace)

Two supported setups:

1. **go.work (recommended for local work across this gateway + go-common):**
   A `go.work` at the repository root already lists both `./` (the gateway) and
   `./go-common`. Just run `go build ./...` / `go test ./...` from the root.

2. **replace directive (for CI / consumers without go.work):**
   The gateway's `go.mod` contains:
   ```
   require github.com/msys/go-common v0.1.0
   replace github.com/msys/go-common => ./go-common
   ```
   For a published release, replace the local path with the version tag:
   ```
   require github.com/msys/go-common v0.1.0
   # (drop the replace, or point it at a fork)
   ```

## Bootstrapping go.sum

`go-common` needs `github.com/google/uuid`. After cloning, generate sums (requires
network to the module proxy / GitHub):

```
cd go-common && go mod tidy
```
