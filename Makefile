GO_BIN = go

BIN_DIR = bin
TMP_DIR = tmp

PROJECT_PACKAGE = github.com/kevinmarquesp/go-postr
PROJECT_MAIN = cmd/application
PROJECT_BIN = gopostr-application

run: build
	./$(BIN_DIR)/$(PROJECT_BIN)

build:
	$(GO_BIN) build -o $(BIN_DIR)/$(PROJECT_BIN) $(PROJECT_PACKAGE)/$(PROJECT_MAIN)

clean:
	rm -rf $(BIN_DIR) $(TMP_DIR)

.PHONY: run build clean

MIGRATIONS_TARGET = ./db/sqlite/migrations
DATABASE_URL = ./tmp/application.db
DATABASE_PROVIDER = sqlite3

dependency-goose:
	$(GO_BIN) install github.com/pressly/goose/v3/cmd/goose@latest

goose-add:
	@mkdir -vp $(MIGRATIONS_TARGET)
	@read -rp "(read) File name: " file; \
		GOOSE_DBSTRING=$(DATABASE_URL) goose -dir=$(MIGRATIONS_TARGET) $(DATABASE_PROVIDER) create "$$file" sql

goose-up:
	GOOSE_DBSTRING=$(DATABASE_URL) goose -dir=$(MIGRATIONS_TARGET) $(DATABASE_PROVIDER) up

goose-reset:
	GOOSE_DBSTRING=$(DATABASE_URL) goose -dir=$(MIGRATIONS_TARGET) $(DATABASE_PROVIDER) reset

.PHONY: migration-add migration-up migration-reset
