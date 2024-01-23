BIN_NAME := gopostr
VERSION := 0.1.0
ARCH := amd64

SERVER_SRC := cmd/server/main.go
TARGET := bin

build:
	GOARCH=$(ARCH) GOOS=darwin go build -o $(TARGET)/$(BIN_NAME)-$(VERSION)-darwin.app $(SERVER_SRC)
	GOARCH=$(ARCH) GOOS=linux go build -o $(TARGET)/$(BIN_NAME)-$(VERSION)-linux.elf $(SERVER_SRC)
	GOARCH=$(ARCH) GOOS=windows go build -o $(TARGET)/$(BIN_NAME)-$(VERSION)-win.exe $(SERVER_SRC)
.PHONY: build

run:
	go run $(SERVER_SRC)
.PHONY: run
