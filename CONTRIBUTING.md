# Contributing to RecordShelf

Thanks for considering contributing!

## Ground rules
- Keep changes small and focused
- Prefer readable, boring code over clever code
- Stick to idiomatic Go
- Minimize dependencies
- Add tests for core logic where practical

## Workflow
1. Fork the repo
2. Create a feature branch:
   - `feature/db-schema`
   - `fix/ui-crash`
3. Make commits with clear messages:
   - `db: add initial schema`
   - `api: handle discogs rate limit`
4. Open a Pull Request with:
   - What changed
   - Why it changed
   - How to test it

## Coding style
- Run `gofmt` on all Go files
- Use meaningful names
- Avoid huge files; split by responsibility
- Handle errors explicitly

## Reporting bugs / feature requests
Open an issue with:
- Steps to reproduce (for bugs)
- Expected vs actual behavior
- Screenshots/logs if relevant
- OS + version info (Linux/Windows)

