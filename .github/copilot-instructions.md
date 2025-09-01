# Portfolio GOTTH - AI Coding Assistant Instructions

## Architecture Overview

This is a **Go portfolio website** using **Templ** for type-safe HTML templating and **TailwindCSS** for styling. The architecture follows a clean separation:

- **`cmd/web/main.go`**: Entry point that loads `data.json` into memory at startup
- **`internals/`**: Core business logic (handlers, models, middleware, utils)
- **`web/view/`**: Templ templates organized by layout/pages/components
- **`static/`**: Generated CSS and client-side assets

## Critical Workflow Patterns

### Development Commands
```bash
make run          # Generate templates + CSS, then run server
make tmux-run     # Start parallel watchers (Templ + Tailwind) in tmux
make generate     # Generate .templ files + compile CSS
```

### Template Generation Flow
1. **Templ files** (`.templ`) → **Go code** via `go generate ./...`
2. **Input CSS** (`web/css/input.css`) → **Static CSS** (`static/css/app.css`) via TailwindCSS
3. **Both must complete** before server starts

## Project-Specific Conventions

### Data-Driven Architecture
- **Single source of truth**: `data.json` contains all portfolio content
- **Startup data loading**: JSON unmarshaled once in `init()`, enriched with computed fields
- **Computed fields**: `AllSkills` and `ProjectTechnologies` generated from base data

### Templ Component Patterns
```go
// Handler pattern - inject data via closure
func IndexHandeler(data models.DataModel) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        home := home.HomePage(data)
        home.Render(context.Background(), w)
    }
}

// Template composition pattern
templ HomePage(data models.DataModel) {
    @layout.BaseLayout("Home", homeNav()) {
        @WelcomeSection(data.Name)
        @AboutSection(data)
    }
}
```

### CSS + Alpine.js Integration
- **TailwindCSS** processes from `web/view/` directory (templates as source)
- **Alpine.js** for client-side interactivity (mobile nav, project filtering)
- **CSS custom properties**: `--primary-color` for theming
- **Utility classes**: `.hover-glow`, `.project-card` for consistent animations

### Directory Structure Rules
- **`web/view/generate.go`**: Contains `//go:generate templ generate ./...` directive
- **`internals/`**: Never import from `web/` - one-way dependency flow
- **Generated files**: `*_templ.go` files are auto-generated, edit `.templ` sources

### Utility Patterns
- **Technology aggregation**: `utils.GetAllTechnologies()` deduplicates across skills + projects
- **JSON tags**: Use `json:"-"` for computed fields not in source data
- **Error handling**: Render errors to HTTP response, log details server-side

## Key Integration Points

### Static Assets
- **File server**: `/static/` prefix stripped and served from `static/` directory
- **Asset references**: Use `static/css/app.css` (not `/static/...`) in templates
- **Font Awesome**: Loaded via npm for icon components

### Component System
- **Icon component**: `@components.Icon("name", "classes")` for SVG icons
- **Project cards**: Uniform styling with hover effects and tech tag filtering
- **Nav components**: Composition pattern with mobile-responsive Alpine.js state

When editing this codebase:
1. **Always run `make generate`** after modifying `.templ` files or CSS
2. **Update `data.json`** for content changes (not hardcoded in templates)
3. **Follow the handler → template → component** data flow pattern
4. **Use Templ's type safety** - pass strongly-typed model structs to templates
