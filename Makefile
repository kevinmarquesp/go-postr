APP = github.com/kevinmarquesp/go-postr
SERVER = cmd/server

.PHONY: run
run:
	@go run $(APP)/$(SERVER)
