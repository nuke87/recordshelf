# RecordShelf – Project Instructions (for Codex/Contributors)

You are an expert software engineer and open-source maintainer.

## Project
**RecordShelf** is an open-source desktop application to manage a personal music collection (vinyl records and CDs).

## Core goals
- Clean, maintainable, well-documented code
- Modular architecture
- Cross-platform (Linux & Windows)
- Open-source friendly: clear structure, readable commits, good docs

## Functional requirements
- Store collection data in a local database
- Fetch album metadata (cover images, tracklists, artist, year, label) via public APIs
- Support **vinyl records** and **CDs**
- Allow manual edits/corrections
- **Offline-first** (APIs only for enrichment)

## Proposed stack (current)
- Language: **Go**
- UI: **Fyne** (default choice for simplicity) or **Tauri** (optional alternative)
- Database: **SQLite**
- Metadata APIs: **Discogs** (primary), **MusicBrainz** (fallback)

## Development rules
- Prefer simple, explicit code over clever code
- Follow idiomatic Go conventions
- Always add brief architectural reasoning (short, practical)
- Add tests for core logic where reasonable
- Never assume hidden context; if something is unclear, ask
- Keep dependencies minimal
- Document each implementation task in `documentation/` (brief summary of changes, rationale, and tests)

## Documentation
- Always record work completed in `documentation/` so the project history stays readable without digging through commits.

## Workflow / order of implementation
1. Project structure (packages, internal boundaries)
2. Database schema + migration approach
3. Domain models + repository layer
4. API clients (Discogs + MusicBrainz) with rate limiting & caching
5. Import/enrichment workflow
6. UI screens (list/search/details/edit)
7. Export/backup features (later)

## Git / commits
- Small commits
- Clear messages, e.g.:
  - `db: add initial schema`
  - `api: implement discogs search`
  - `ui: add album details view`
