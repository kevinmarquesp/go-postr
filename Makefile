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
