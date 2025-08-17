# Copilot Instructions for Portfolio Project

## Architecture Overview

This is a single-page portfolio website built with **Go + Gin** serving **embedded static assets**. The entire application compiles to a single binary with all assets embedded using `go:embed` directives.

### Key Components
- **`main.go`**: Entry point with embedded file systems (`static/*`, `templates/*`, `data.json`)
- **`models/models.go`**: Type definitions for portfolio data structure
- **`data.json`**: Content source - edit this to update portfolio information
- **`templates/index.html`**: Single HTML template using Go template syntax
- **`static/`**: CSS, JS, and assets served at `/static/` route

## Critical Patterns

### Embedded Assets Pattern
```go
//go:embed static/*
var static embed.FS
```
All static files are embedded at compile time. Changes to `static/` or `templates/` require rebuilding the binary.

### Data-Driven Content
Portfolio content lives in `data.json` and maps to `models.Portfolio` struct. The JSON structure directly drives the template rendering - no database involved.

### Template Data Flow
`data.json` → `models.Portfolio` → Go template rendering in `index.html`. Use Go template syntax like `{{ .Name }}` and `{{ range .Projects }}`.

## Development Workflows

### Local Development
```bash
go run main.go  # Starts server on :8080 or $PORT
```

### Building & Deployment
- **Docker**: Multi-stage build with cross-compilation support via `TARGETOS`/`TARGETARCH` args
- **Fly.io**: Configured in `fly.toml` for Mumbai region (`bom`) deployment
- **Binary**: Single executable with all assets embedded

### Content Updates
1. Edit `data.json` for portfolio content changes
2. Modify `models/models.go` if adding new data fields
3. Update `templates/index.html` for layout/display changes
4. Rebuild binary to apply changes

## Project Conventions

### Performance-First Approach
- **14KB initial payload limit** - optimize for instant loading on slow connections
- All assets served locally from `static/` - no external CDN dependencies
- Minimal JavaScript - prefer CSS-only solutions when possible
- Use HTMX only when true interactivity is required

### File Organization
- Keep all static assets in `static/` (auto-served at `/static/` route)
- Single template approach - avoid template fragmentation
- Models define JSON structure exactly - maintain 1:1 mapping

### Styling & Frontend
- **PicoCSS only** - use local copy in `static/pico.min.css`, never CDN for performance
- Custom styles in `static/styles.css` - keep minimal and semantic
- **Local JavaScript only** - all JS files served from `static/`
- Use **HTMX sparingly** - only add if interactivity truly requires it
- Icons via **Feather icons** loaded locally in JavaScript
- **Performance target**: Keep initial payload under 14KB for instant loading

### Port Configuration
Application respects `PORT` environment variable (defaults to 8080) - critical for Fly.io deployment.

## Integration Points

### External Dependencies
- **Gin framework** for HTTP routing and template rendering
- **PicoCSS** for base styling (local copy only - never CDN)
- **Feather Icons** for UI icons (loaded locally)

### Deployment Targets
- **Fly.io**: Primary deployment platform (see `fly.toml`)
- **Docker**: Cross-platform containerization with Alpine base
- **Local**: Direct Go binary execution

## Common Modifications

### Adding New Portfolio Sections
1. Extend `models.Portfolio` struct with new field
2. Add corresponding data in `data.json`
3. Update `templates/index.html` to render new section

### Styling Changes
Edit `static/styles.css` - the project uses CSS custom properties and follows PicoCSS conventions.

### Static Asset Updates
Place new files in `static/` directory. They'll be accessible at `/static/<filename>` after rebuilding.
