SERVER_SRC = github.com/kevinmarquesp/go-postr/cmd/server
SERVER_BIN = bin/server

NPM = pnpm
NPX = pnpm

.PHONY: run
run:
	[ -e ./$(SERVER_BIN) ] || make build; \
		./$(SERVER_BIN)

.PHONY: build
build:
	go build -o $(SERVER_BIN) $(SERVER_SRC)

.PHONY: build/production
build/production: deps templ tailwind build

.PHONY: deps
deps:
	$(NPM) install --force
	go mod download
	go install github.com/a-h/templ/cmd/templ@v0.2.707       # To render the views.
	go install github.com/air-verse/air@v1.52.2              # To allow live reloading.
	go install github.com/pressly/goose/v3/cmd/goose@v3.20.0 # To do migrations.

.PHONY: clean
clean:
	rm -vr node_modules bin static/css/* views/**/*_templ.go

.PHONY: air
air:
	air -build.bin "$(SERVER_BIN)"

# ------------------------------------------------------------------------------
# Tailwind related recipes.

STATIC_TAILWINDCSS = static/dist/css/tailwind.css

.PHONY: tailwind
tailwind:
	$(NPX) tailwindcss build --output $(STATIC_TAILWINDCSS) --minify

.PHONY: tailwind/watch
tailwind/watch:
	$(NPX) tailwindcss --watch --output $(STATIC_TAILWINDCSS) --minify

# ------------------------------------------------------------------------------
# Go's Templ related recipes, to build the `.templ` files (deppends on the
# `.env` file, or just its environment variables).

.PHONY: templ
templ:
	templ generate

.PHONY: templ/watch
templ/watch:
	[ -e ./$(DOTENV) ] && . ./$(DOTENV); \
		templ generate -watch -proxy=http://localhost:$${PORT}

# ------------------------------------------------------------------------------
# Postgres related recipes, to connect and run migrations (deppends on the
# `.env` file, or just its environment variables).

SEED_SRC = github.com/kevinmarquesp/go-postr/cmd/seed
MIGRATION_DIR = db/migrations
DOTENV = .env

DATABASE = postgres
DATABASE_DNS = postgres://$$POSTGRES_USER:$$POSTGRES_PASSWORD@$$POSTGRES_HOST:$$POSTGRES_PORT/$$POSTGRES_DB?sslmode=disable
GOOSE_VARS = GOOSE_DRIVER=$(DATABASE) GOOSE_DBSTRING=$(DATABASE_DNS)

.PHONY: postgres
postgres:
	@[ -e ./$(DOTENV) ] && . ./$(DOTENV); \
		psql "$(DATABASE_DNS)"

.PHONY: postgres/seed
postgres/seed:
	go run $(SEED_SRC)

.PHONY: postgres/migrations/create
postgres/migrations/create:
	@read -rp "Migration name: " new_migration; \
		$(GOOSE_VARS) goose -dir=${MIGRATION_DIR} create "$$new_migration" sql

.PHONY: postgres/migrations/up
postgres/migrations/up:
	@[ -e ./$(DOTENV) ] && . ./$(DOTENV); \
		$(GOOSE_VARS) goose -dir=${MIGRATION_DIR} up

.PHONY: postgres/migrations/status
postgres/migrations/status:
	@[ -e ./$(DOTENV) ] && . ./$(DOTENV); \
		$(GOOSE_VARS) goose -dir=${MIGRATION_DIR} status

.PHONY: postgres/migrations/reset
postgres/migrations/reset:
	@[ -e ./$(DOTENV) ] && . ./$(DOTENV); \
		$(GOOSE_VARS) goose -dir=${MIGRATION_DIR} reset
