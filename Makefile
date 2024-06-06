BIN_NAME := gopostr
VERSION := 0.1.0
ARCH := amd64

SERVER_SRC := cmd/server/main.go
TARGET := bin

run:
	templ generate
	go run $(SERVER_SRC)
.PHONY: run

depget:
	go install github.com/a-h/templ/cmd/templ@v0.2.707
	go mod tidy
.PHONY: depget
