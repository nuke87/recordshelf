# Architecture overview

## Goals
- Keep the core collection model independent of storage, UI, and external APIs.
- Support offline-first usage; remote APIs only enrich local data.
- Maintain clear module boundaries to keep changes local and testable.

## Module boundaries
- `internal/domain`: Entities and business rules (albums, tracks, formats).
- `internal/repo`: Repository interfaces plus persistence implementations.
- `internal/db`: Database setup, migrations, and connection helpers for SQLite.
- `internal/api`: Metadata enrichment clients (Discogs, MusicBrainz) with caching.
- `internal/ui`: Desktop UI composition and view models.
- `cmd/recordshelf`: Application entrypoint and wiring.

## Data flow (high level)
1. UI requests a list/search/detail view.
2. Repositories fetch from the local database.
3. Enrichment workflow optionally calls API clients to fill gaps.
4. Updates are stored locally and reflected in the UI.

## Migrations
Migrations are embedded SQL files applied in filename order, with applied versions tracked in `schema_migrations`.
