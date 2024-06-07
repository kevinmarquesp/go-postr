SERVER_SRC := cmd/server/main.go
SERVER_BIN := bin/server
STATIC_TAILWINDCSS := static/css/tailwind.css
MIGRATION_DIR := db/migrations
ENV := .env

DATABASE := postgres
DATABASE_DNS := postgres://$$POSTGRES_USER:$$POSTGRES_PASSWORD@$$POSTGRES_HOST:$$POSTGRES_PORT/$$POSTGRES_DB?sslmode=disable
GOOSE_VARS := GOOSE_DRIVER=$(DATABASE) GOOSE_DBSTRING=$(DATABASE_DNS)

NPM := pnpm
NPX := pnpm

.PHONY: build
build:
	templ generate
	go build -o $(SERVER_BIN) $(SERVER_SRC)
	$(NPX) tailwindcss build --output $(STATIC_TAILWINDCSS) --minify

.PHONY: run
run: build
	./$(SERVER_BIN)

.PHONY: deps
deps:
	go install github.com/a-h/templ/cmd/templ@v0.2.707       # to render the views
	go install github.com/air-verse/air@v1.52.2              # to allow live reloading
	go install github.com/pressly/goose/v3/cmd/goose@v3.20.0 # to do migrations
	$(NPM) add --save-dev tailwindcss
	which asdf &>/dev/null && asdf reshim
	go mod tidy

.PHONY: dev
dev:
	air -build.bin "$(SERVER_BIN)"

.PHONY: db/connect
db/connect:
	@. ./$(ENV) && psql "$(DATABASE_DNS)"

.PHONY: migrations/create
migrations/create:
	@read -rp "Migration name: " new_migration; \
		$(GOOSE_VARS) goose -dir=${MIGRATION_DIR} create "$$new_migration" sql

.PHONY: migrations/up
migrations/up:
	@. ./$(ENV) && $(GOOSE_VARS) goose -dir=${MIGRATION_DIR} up

.PHONY: migrations/status
migrations/status:
	@. ./$(ENV) && $(GOOSE_VARS) goose -dir=${MIGRATION_DIR} status

.PHONY: migrations/reset
migrations/reset:
	@. ./$(ENV) && $(GOOSE_VARS) goose -dir=${MIGRATION_DIR} reset
