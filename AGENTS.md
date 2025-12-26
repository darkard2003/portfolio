# AGENTS.md

## Build/Testing Commands

### Run All Tests
```bash
make test
```

### Run Single Test
```bash
go test -run TestNamePattern
```

### Lint Code
```bash
make lint
```

### Build Project
```bash
make build
```

### Generate Templates
```bash
make generate
```

### Development Server
```bash
make run
```

## Code Style Guidelines

### Go Conventions
- Use `gofmt` for formatting
- Prefer snake_case for variables
- Use `json:"-"` for unexported fields
- Always return errors explicitly

### Templ Templates
- Use type-safe HTML templates
- Follow component composition pattern
- Use `@layout.BaseLayout("Page")` for layouts

### TailwindCSS
- Use utility-first approach
- Prioritize semantic class names
- Use `--primary-color` for theming

### Naming
- Use PascalCase for types
- Use snake_case for variables/functions
- Use `project-card` for components

### Error Handling
- Return errors via HTTP status codes
- Log detailed errors to stderr
- Use `models.DataModel` as primary data type

## Project Rules

### Data Architecture
- All content comes from `data.json`
- No hardcoded values in templates
- Use `utils.GetAllTechnologies()` for tech lists

### File Structure
- Never import from `web/` directory
- Generated files (e.g., `_templ.go`) are read-only
- Edit `.templ` files for template changes

### CI/CD
- `make generate` must succeed before CI passes
- All templates must compile successfully
- Use `go generate ./...` for template generation

## Integration Rules

### Static Assets
- Use `static/css/app.css` for CSS
- Use `static/scripts/alpine.min.js` for Alpine.js
- Never reference `/static/` directly in templates

### Component System
- Use `@components.Icon("name")` for icons
- Use `@components.Nav()` for navigation
- Use `@components.PostCard()` for blog posts

### Dependency Management
- Go modules are used for dependencies
- Node.js packages are in `package.json`
- Use `npm install` for frontend dependencies

## Special Instructions
- Always run `make generate` after template changes
- Update `data.json` for content changes
- Use `go mod tidy` to manage dependencies
- Use `make tmux-run` for parallel development

### Cursor Rules
<!-- Add Cursor rules here if found -->

### Copilot Instructions
<!-- Add Copilot instructions here if found -->