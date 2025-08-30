# GEMINI.md

## Project Overview

This is a web portfolio project built with Go, HTMX, and Tailwind CSS. The project uses `templ` for server-side rendering of HTML components and `air` for live reloading during development.

**Key Technologies:**

*   **Backend:** Go
*   **Frontend:**
    *   `templ` for Go-based templating
    *   `tailwindcss` for CSS styling
*   **Development:**
    *   `make` for build automation
    *   `air` for live reloading
    *   `tmux` for running multiple processes in a single terminal session

## Building and Running

### Prerequisites

*   Go
*   Node.js and npm
*   `templ`
*   `air`
*   `tmux`

### Commands

*   **`make build`**: Builds the Go application and generates the CSS file.
*   **`make generate`**: Generates Go code from `templ` files and the CSS file from `tailwindcss`.
*   **`make run`**: Runs the application.
*   **`make run-bin`**: Builds and runs the application.
*   **`make tailwind-watch`**: Watches for changes in the `input.css` file and recompiles the `app.css` file.
*   **`make templ-watch`**: Watches for changes in the `templ` files and regenerates the Go code.
*   **`make tmux-run`**: Starts a `tmux` session with `tailwind-watch` and `templ-watch` running in separate panes.

## Development Conventions

*   **Styling:** Use `tailwindcss` utility classes in the `web/css/input.css` file.
*   **Templating:** Create `templ` components in the `web/view/` directory.
*   **Handlers:** Create new handlers in the `internals/handelers/` directory.
*   **Routing:** Add new routes to the `main.go` file.

## Directions for GEMINI

* If i ask a question, just answer the question. Don't try to edit or run anything.
* Only try to edit or run if i ask you to do something.
* Always assume live server is running, so no need to run the generate or build commands
