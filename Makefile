.DEFAULT_GOAL := build

fmt:
	go fmt ./...

.PHONY: fmt

vet: fmt
	go vet ./...

.PHONY: vet

run: vet
	go run cmd/aegis/main.go

.PHONY: build

build: vet
	go build -o ../build/bin/aegis ../cmd/aegis

.PHONY: build

test: 
	go test ./...

.PHONY: test



