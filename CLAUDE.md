# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

### Development
```bash
air                          # Hot reload dev server (runs gen-docs before build automatically)
go build -o ./bin/main ./cmd/api  # Manual build
```

### Migrations
```bash
make migration name=<name>   # Create new migration files
make migrate-up              # Apply all pending migrations
make migrate-down            # Rollback one migration (add steps=N for multiple)
make seed                    # Seed the database with test data
```

### Documentation
```bash
make gen-docs                # Regenerate Swagger docs (also runs automatically via air)
```

### Tests
```bash
go test ./...                # Run all tests
go test ./internal/store/... # Run store tests only
```

### Environment
Uses `direnv` with `.envrc`. Required vars: `ADDR`, `DB_ADDR`. The database runs via `docker-compose up`.

## Architecture

### Request Flow
```
HTTP Request
  → Chi middleware (RequestID, RealIP, Logger, Recoverer, Timeout)
  → Route-specific context middleware (pre-loads resources like post/user)
  → Handler (reads context, calls store, writes JSON response)
  → Store method (executes SQL with 5s context timeout)
```

### Key Structure
- `cmd/api/` — HTTP handlers, routing, application struct (`application` in `api.go`)
- `internal/store/` — Data layer: interfaces in `storage.go`, implementations per entity
- `internal/db/` — PostgreSQL connection pool setup
- `internal/env/` — Typed env var helpers (`GetString`, `GetInt` with fallbacks)

### Application Struct (`cmd/api/api.go`)
All handlers are methods on `*application`, which carries `config`, `store.Storage`, and a Zap logger. Config is assembled from env vars in `main.go` and passed to `run()`.

### Storage Layer (`internal/store/storage.go`)
`Storage` is a struct of interfaces (`PostStore`, `UserStore`, etc.), not a single interface. Use `withTx()` for multi-step operations requiring atomicity — it handles rollback automatically.

### Context Middleware Pattern
Resources are pre-fetched in middleware and stored in request context before reaching the handler:
```go
// Middleware stores resource
ctx := context.WithValue(r.Context(), postCtx, post)

// Handler retrieves it
post := getPostFromCtx(r)
```

### Error Handling
Store packages define sentinel errors (`ErrNotFound`, `ErrDuplicateEmail`, etc.). Handlers match with `errors.Is()` and call the appropriate response method (`notFoundError`, `badRequestError`, `conflictResponse`, `internalServerError`).

### Token Security Pattern (used in invitations)
Store the hash, send the plain token:
```go
plainToken := uuid.New().String()
hash := sha256.Sum256([]byte(plainToken))
hashToken := hex.EncodeToString(hash[:])
// save hashToken to DB, send plainToken in email
```
On confirmation: re-hash the received token and look up by hash.

### Optimistic Concurrency Control
Post updates include a `version` column check to detect concurrent modifications. Increment `version` on each update and match it in the `WHERE` clause.

### API Docs
Swagger annotations live in handler files. Run `make gen-docs` to regenerate `docs/`. The `/v1/swagger/*` route serves the UI.
