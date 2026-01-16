.PHONY: build run test lint

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

ARGS ?=

run: build
	./bin/hexlet-path-size $(ARGS)

test:
	go test -v ./...

lint:
	golangci-lint run ./...
