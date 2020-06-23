OS          := $(shell uname -s)
SHELL       := /bin/bash
GOFMT       := gofmt -s -w -l
GOLINT      := golint
GOVET       := go vet
GOSHDW      := go vet -vettool=$$(which shadow)
GOSEC       := gosec --quiet
APP_VERSION := $(shell cat .github/.version)
GO_LDFLAGS  := -ldflags "-s -w -X main.version=$(APP_VERSION)"

.PHONY: build binaries run test fmt clean lint help

build:	## Build the binary for the current platform
	go generate
	CGO_ENABLED=0 go build $(GO_LDFLAGS)

docker-build:	## Build application in a docker image
	docker build  --rm --build-arg APP_VERSION=$(APP_VERSION) -t saxifrage:latest .

binaries:	## Build and zip the binary for Linux and MacOS
	./scripts/create-binaries

run:	## Run the application
	CGO_ENABLED=0 go run sax.go

test:	## Run the tests
	CGO_ENABLED=0 go test ./...

fmt:	## Format source code
	$(GOFMT) *.go
	$(GOFMT) **/*.go

clean:	## Tidy and clean
	go mod tidy
	go clean

lint:	## Run golint, govet, gosec, etc...
	$(GOLINT) ./...
	$(GOVET) ./...
	$(GOSEC) ./...
	$(GOSHDW) ./...

help:	## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
