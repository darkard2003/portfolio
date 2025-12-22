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
		templ generate --watch --watch-pattern='(.+\.go$$)|(.+\.templ$$)|(.+_templ\.txt$$)|(.+\.md$$)' --proxy="http://localhost:8080" --cmd="go run -tags dev ./cmd/web/main.go"

.PHONY: templ-watch-release
templ-watch-release:
		templ generate --watch --watch-pattern='(.+\.go$$)|(.+\.templ$$)|(.+_templ\.txt$$)|(.+\.md$$)' --proxy="http://localhost:8080" --cmd="go run  ./cmd/web/main.go"

.PHONY: tmux-run
tmux-run:
	tmux split-window -h "make tailwind-watch"
	tmux split-window -v "make templ-watch"

# Docker commands
.PHONY: docker-build
docker-build:
	docker build -t portfolio-gotth .

.PHONY: docker-run

docker-run:
	docker run -p 8080:8080 portfolio-gotth

.PHONY: docker-build-run
docker-build-run: docker-build docker-run

