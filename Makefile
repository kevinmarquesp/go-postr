GO_BIN = go
TARGET = bin

PROJECT_PACKAGE = github.com/kevinmarquesp/go-postr
PROJECT_MAIN = cmd/go_postr
PROJECT_BIN = go_postr

GOTEST_BIN = gotestsum
TEST_TARGET = ./...

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

depsget:
	$(GO_BIN) mod download
.PHONY: depsget

depsget-gotestsum:
	$(GO_BIN) install gotest.tools/gotestsum@latest
.PHONY: depsget-gotestsum
