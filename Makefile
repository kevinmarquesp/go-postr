APP = github.com/kevinmarquesp/go-postr
SERVER = cmd/gopostr
SEED = cmd/seed

BIN = bin/server

DB = tmp/database.sqlite3
PROVIDER = sqlite3
MIGRATIONS = db/sqlite3/migrations

.PHONY: run
run: build
	@./$(BIN)

.PHONY: build
build:
	@mkdir -vp bin &>/dev/null
	@go build -o $(BIN) $(APP)/$(SERVER)

.PHONY: deps
deps:
	go install gotest.tools/gotestsum@latest
	go install github.com/air-verse/air@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: seed
seed:
	@go run $(APP)/$(SEED)

.PHONY: test
test:
	@gotestsum --format-hide-empty-pkg --format-icons codicons $(UNIT)

.PHONY: test-watch
test-watch:
	@inotifywait --include '.\.(go|sql)$$' -rm ./ -e modify | \
		while read -r dir action file; do \
			clear; printf "\e[?25l\n"; \
			gotestsum --format-hide-empty-pkg --format-icons codicons; \
		done

.PHONY: migration-create
migration-create:
	@read -rp "Migration name: " file; \
		GOOSE_DBSTRING=$(DB) goose -dir=$(MIGRATIONS) $(PROVIDER) create "$$file" sql

.PHONY: migration-up
migration-up:
	GOOSE_DBSTRING=$(DB) goose -dir=$(MIGRATIONS) $(PROVIDER) up

.PHONY: migration-reset
migration-reset:
	GOOSE_DBSTRING=$(DB) goose -dir=$(MIGRATIONS) $(PROVIDER) reset

.PHONY: migration-status
migration-status:
	GOOSE_DBSTRING=$(DB) goose -dir=$(MIGRATIONS) $(PROVIDER) status
