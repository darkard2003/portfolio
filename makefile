.PHONY: build
build: generate
	go build -o ./bin/main ./cmd/web/main.go
	
.PHONY: generate
generate:
	go generate ./...
	npx tailwindcss -i ./web/css/input.css -o ./static/css/app.css  --minify
.PHONY: run
run: generate
	go run ./cmd/web/main.go

.PHONY: run-bin
run-bin: build
	./bin/main

.PHONY: tailwind-watch
tailwind-watch:
	npx tailwindcss -i ./web/css/input.css -o ./static/css/app.css --watch

.PHONY: templ-watch
templ-watch:
	templ generate --watch

.PHONY: tmux-run
tmux-run:
	tmux split-window -h "make tailwind-watch"
	tmux split-window -v "make templ-watch"
	air

