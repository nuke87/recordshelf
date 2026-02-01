# Implementation Log

## 2026-02-01
- Established the base Go module and entrypoint (`go.mod`, `cmd/recordshelf/main.go`).
- Created package boundaries with doc stubs in `internal/db`, `internal/domain`, `internal/repo`, `internal/api`, and `internal/ui`.
- Added an architecture overview in `docs/architecture.md` (module boundaries, data flow, migrations).
- Implemented the initial SQLite schema and migration runner:
  - `internal/db/migrations/0001_init.sql` (albums, tracks, indexes, schema migrations)
  - `internal/db/migrate.go` (embedded migrations, tracking table)
  - `internal/db/db.go` (open, PRAGMA foreign_keys, migrate)
- Added core domain models and format enum in `internal/domain/*`.
- Added repository interface and `ErrNotFound` in `internal/repo/*`.
- Tests: `go test ./...`
