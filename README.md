## Go Template

A minimal, production-minded Golang HTTP service template. It uses chi for routing, Zap for structured logging, Viper for configuration, and includes sensible middleware (logging, recovery, and rate limiting). The project is organized with a clean separation of concerns across controller/router, handler, service, and models.

### Features
- **HTTP server**: `chi` router with configurable server options
- **Middleware**: request logging, panic recovery, and sliding-window rate limiting (per IP and per user)
- **Config**: Viper-backed config with env and optional file overrides
- **Logging**: Uber Zap with configurable levels
- **Graceful shutdown**: listens for SIGINT and closes the server cleanly
- **Example endpoint**: `GET /v1/hello?name=YourName`

### Tech Stack
- Router: `github.com/go-chi/chi/v5`
- Rate limiting: `github.com/go-chi/httprate`
- Config: `github.com/spf13/viper`
- Logging: `go.uber.org/zap`

### Project Layout
```
cmd/
  gotemp/
    main.go           # App entrypoint (bootstraps gotemp)
config/
  config.go           # Viper-backed config loader
internal/
  controller/httprouter/
    router.go         # Route and middleware wiring
    middleware/
      logger.go       # Request logging middleware
      recovery.go     # Panic recovery middleware
      ratelimiter.go  # Rate limiter helpers (IP/user)
  handler/
    hello.go          # Transport layer: decode/encode and HTTP handlers
    errors.go         # Common transport errors
  models/
    hello.go          # Request/response DTOs
  service/
    service.go        # Service interfaces
    hello/hello.go    # Example business logic implementation
pkg/
  httpserver/         # Thin wrapper over net/http with options
    server.go
    options.go
  logger/
    logger.go         # Zap logger factory
gotemp.go             # Core app wiring (config, logger, server, shutdown)
options.go            # App-level option helpers
```

### Requirements
- Go 1.23+ (module sets toolchain `go1.24.9`)

### Getting Started
1) Clone and enter the repository
```bash
git clone <your-repo-url>
cd go-template
```

2) Run the service
```bash
go run ./cmd/gotemp
```

The server starts on the configured port (default `:80`).

### Configuration
Configuration is loaded via Viper in this order:
1. Defaults (hardcoded)
2. `config.(yaml|yml|json)` if present in repo root
3. Environment variables (override both)

Available settings (env keys shown):
- `APP_NAME` (default: `go-microservice`)
- `PORT` (default: `80`)
- `ENV` (default: `development`)
- `LOG_LEVEL` (default: `info`) — one of: `debug`, `info`, `warn`, `error`
- `REQUEST_LIM_MIN` (default: `100`) — requests per minute

Optional `config.yaml` example:
```yaml
APP_NAME: go-template
PORT: "8080"
ENV: development
LOG_LEVEL: debug
REQUEST_LIM_MIN: 120
```

You can also override via environment variables:
```bash
export PORT=8080 LOG_LEVEL=debug REQUEST_LIM_MIN=120
go run ./cmd/gotemp
```

### HTTP API
- `GET /v1/hello?name=YourName`
  - Response 200:
    ```json
    { "message": "Hello YourName" }
    ```

Rate limiting:
- Global middleware applies IP-based limit: `REQUEST_LIM_MIN` per minute
- Inside `/v1`, an additional limiter demonstrates per-user limiting using a `userID` value from the request context

### Logging
Zap logger is initialized with the configured `LOG_LEVEL`. Example fields include method, path, and duration from the logging middleware.

### Graceful Shutdown
The app listens for `SIGINT` and closes the HTTP server gracefully before exiting.

### Extending the Template
- Add new services under `internal/service/<domain>/`
- Define request/response DTOs in `internal/models/`
- Add HTTP handlers in `internal/handler/`
- Register routes in `internal/controller/httprouter/router.go`

Server options (timeouts, port) can be tuned via `pkg/httpserver/options.go` and passed to `httpserver.New(...)`.

### Run Examples
```bash
# Default (port 80) — may require sudo depending on your OS
go run ./cmd/gotemp

# With custom env
PORT=8080 LOG_LEVEL=debug REQUEST_LIM_MIN=60 go run ./cmd/gotemp

# Call the example endpoint
curl "http://localhost:8080/v1/hello?name=Harsh"
```

### License
MIT (or update to your preferred license)


