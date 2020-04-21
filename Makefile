OS     := $(shell uname -s)
SHELL  := /bin/bash
GOFMT  := gofmt -s -w -l
GOLINT := golint
GOVET  := go vet
GOSHDW := go vet -vettool=$$(which shadow)
GOSEC  := gosec --quiet

run:
	go run sax.go

build:
	go build -o sax

test:
	go test ./...

fmt:
	$(GOFMT) **/*.go

clean:
	go mod tidy
	go clean

lint:
	$(GOLINT) ./...
	$(GOVET) ./...
	$(GOSEC) ./...
	$(GOSHDW) ./...

