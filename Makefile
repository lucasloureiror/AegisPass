.DEFAULT_GOAL := build

vet: test
	go test ./...

.PHONY: vet

fmt: test
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



