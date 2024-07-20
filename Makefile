GO_BIN = go
TARGET = bin

PROJECT_PACKAGE = github.com/kevinmarquesp/go-postr
PROJECT_MAIN = cmd/go_postr
PROJECT_BIN = go_postr
PROJECT_PORT = 8000

GOTEST_BIN = gotestsum
TEST_TARGET = ./...

MIGRATIONS_TARGET = ./db/sqlite3/migrations
DATABASE_URL = ./tmp/database.sqlite3
DATABASE_PROVIDER = sqlite3

run: build
	./$(TARGET)/$(PROJECT_BIN)
.PHONY: run

build:
	$(GO_BIN) build -o $(TARGET)/$(PROJECT_BIN) $(PROJECT_PACKAGE)/$(PROJECT_MAIN)
.PHONY: build

test:
	$(GOTEST_BIN) --format-hide-empty-pkg --format-icons codicons $(TEST_TARGET)
.PHONY: test

test-watch:
	inotifywait --include '.\.(go|sql)$$' -rm ./ -e modify | \
		while read -r dir action file; do \
			clear; printf "\e[?25l\n"; \
			gotestsum --format-hide-empty-pkg --format-icons codicons; \
		done
.PHONY: test-watch

air:
	air -build.bin=$(TARGET)/$(PROJECT_BIN)
.PHONY: air

templ:
	templ generate
.PHONY: templ

templ-watch:
	templ generate -watch -proxy=http://localhost:$(PROJECT_PORT) -open-browser=false
.PHONY: templ-watch

depsget-all: depsget depsget-gose depsget-gotestsum depsget-templ depsget-air
.PHONY: depsget-all

depsget:
	$(GO_BIN) mod download
.PHONY: depsget

depsget-gose:
	$(GO_BIN) install github.com/pressly/goose/v3/cmd/goose@latest
.PHONY: depsget-gotestsum

depsget-gotestsum:
	$(GO_BIN) install gotest.tools/gotestsum@latest
.PHONY: depsget-gotestsum

depsget-templ:
	$(GO_BIN) install github.com/a-h/templ/cmd/templ@latest
.PHONY: depsget-templ

depsget-air:
	$(GO_BIN) install github.com/air-verse/air@latest
.PHONY: depsget-air

migration-add:
	@mkdir -vp $(MIGRATIONS_TARGET)
	@read -rp "(read) File name: " file; \
		GOOSE_DBSTRING=$(DATABASE_URL) goose -dir=$(MIGRATIONS_TARGET) $(DATABASE_PROVIDER) create "$$file" sql
.PHONY: migration-add

migration-up:
	GOOSE_DBSTRING=$(DATABASE_URL) goose -dir=$(MIGRATIONS_TARGET) $(DATABASE_PROVIDER) up
.PHONY: migration-add

migration-reset:
	GOOSE_DBSTRING=$(DATABASE_URL) goose -dir=$(MIGRATIONS_TARGET) $(DATABASE_PROVIDER) reset
.PHONY: migration-add
