APP = github.com/kevinmarquesp/go-postr
SERVER = cmd/server

DB = tmp/database.sqlite3
PROVIDER = sqlite3
MIGRATIONS = db/sqlite3/migrations

.PHONY: run
run:
	@go run $(APP)/$(SERVER)

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
