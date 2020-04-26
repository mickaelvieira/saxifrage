OS          := $(shell uname -s)
SHELL       := /bin/bash
GOFMT       := gofmt -s -w -l
GOLINT      := golint
GOVET       := go vet
GOSHDW      := go vet -vettool=$$(which shadow)
GOSEC       := gosec --quiet
APP_VERSION := $(shell cat .github/.version)
GO_LDFLAGS  := -ldflags "-s -w -X main.version=$(APP_VERSION)"

run:
	CGO_ENABLED=0 go run sax.go

build:
	CGO_ENABLED=0 go build $(GO_LDFLAGS)

test:
	CGO_ENABLED=0 go test ./...

fmt:
	$(GOFMT) *.go
	$(GOFMT) **/*.go

clean:
	go mod tidy
	go clean

lint:
	$(GOLINT) ./...
	$(GOVET) ./...
	$(GOSEC) ./...
	$(GOSHDW) ./...

