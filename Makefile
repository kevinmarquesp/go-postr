GO_BIN = go
TARGET = bin

PROJECT_PACKAGE = github.com/kevinmarquesp/go-postr
PROJECT_MAIN = cmd/go_postr
PROJECT_BIN = go_postr

run: build
	./$(PROJECT_BIN)
.PHONY: run

build:
	$(GO_BIN) build -o $(TARGET)/$(PROJECT_BIN) $(PROJECT_PACKAGE)/$(PROJECT_MAIN)
.PHONY: build
