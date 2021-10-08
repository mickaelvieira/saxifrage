OS          := $(shell uname -s)
SHELL       := /bin/bash
GOFMT       := gofmt -s -w -l
STATICCHECK := staticcheck
GOSEC       := gosec --quiet
APP_VERSION := $(shell cat .github/.version)
GO_LDFLAGS  := -ldflags "-s -w -X main.version=$(APP_VERSION)"

.PHONY: build docker-build buildah-build binaries run test fmt clean lint help

build:	## Build the binary for the current platform
	CGO_ENABLED=0 go build $(GO_LDFLAGS)

docker-build:	## Build application in a docker image
	sudo docker build --rm --build-arg APP_VERSION=$(APP_VERSION) -t saxifrage:latest .

buildah-build: ## Build application in a buildah image
	sudo buildah build -f Dockerfile --build-arg TARGETOS=linux --build-arg TARGETARCH=amd64 --rm --layers --build-arg APP_VERSION=$(APP_VERSION) -t saxifrage:latest .

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
	$(STATICCHECK) ./...
	$(GOSEC) ./...

help:	## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
