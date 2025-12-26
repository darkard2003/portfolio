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
- Use `gofmt` for formatting: `gofmt -w ./...` to reformat code
- Variable example: `var projectID string` (snake_case)
- Unexported field example: `type User struct { id int json:"-" }`
- Error handling example: 
  ```go
  if err := validateInput(); err != nil {
    log.Fatal(err)
  }
  ```

### TailwindCSS
- Utility-first example: `class="p-4 bg-[var:--primary-color] text-[var:--secondary-color]"`
- Semantic class example: `.project-card { ... }` not `.card`
- Theming example: 
  ```html
  <style>
  :root {
    --primary-color: #3b82f6;
  }
  </style>
  ```

### Naming
- Component pattern: `project-card` (snake_case) vs `ProjectCard` (PascalCase)
- Service example: `func FetchPosts() ([]Post, error)`
- Type example: `type Portfolio struct { ... }`

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