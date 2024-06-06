NPM := pnpm
NPX := pnpm

SERVER_SRC := cmd/server/main.go
SERVER_BIN := bin/server
STATIC_TAILWINDCSS := static/css/tailwind.css

.PHONY: build
build:
	templ generate
	go build -o $(SERVER_BIN) $(SERVER_SRC)
	$(NPX) tailwindcss build --output $(STATIC_TAILWINDCSS) --minify

.PHONY: run
run: build
	./$(SERVER_BIN)

.PHONY: depget
depget:
	$(NPM) add --save-dev tailwindcss
	go install github.com/a-h/templ/cmd/templ@v0.2.707
	go install github.com/air-verse/air@v1.52.2
	go mod tidy

.PHONY: dev
dev:
	air -build.bin "$(SERVER_BIN)"
