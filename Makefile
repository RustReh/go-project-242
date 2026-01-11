.PHONY: build run test lint

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

run: build
	./bin/hexlet-path-size

test:
	go test ./...

lint:
	golangci-lint run ./...
