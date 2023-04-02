SHELL=/bin/bash
IMAGE_TAG := $(shell git rev-parse HEAD)

.PHONY: all
all: deps lint test build

.PHONY: ci
ci: lint test

.PHONY: build
build:
	go build -mod=vendor -a -o ./artifacts/svc .

.PHONY: deps
deps:
	go mod download

.PHONY: test
test:
	go test -count=1 ./...

.PHONY: cover
cover:
	go test -count=1 -coverprofile cover.out ./... && go tool cover -html=cover.out

.PHONY: dockerise
dockerise:
	docker build -t "elegantcreationism/goclient-polygon:${IMAGE_TAG}" .

.PHONY: lint
lint: gen
	golangci-lint run

