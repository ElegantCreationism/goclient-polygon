SHELL=/bin/bash
IMAGE_TAG := $(shell git rev-parse HEAD)

.PHONY: all
all: deps lint test build

.PHONY: ci
ci: lint test

docker_repository = "elegantcreationism"

.PHONY: deps
deps:
	go mod download

.PHONY: test
test:
	go test -count=1 -coverprofile cover.out ./...

.PHONY: cover
cover: test
	go tool cover -html=cover.out

.PHONY: build
build:
	docker build -t "${docker_repository}/goclient-polygon:${IMAGE_TAG}" .

.PHONY: run
run: build
	docker run elegantcreationism/goclient-polygon:${IMAGE_TAG}

.PHONY: push
push: build
	docker tag elegantcreationism/goclient-polygon:${IMAGE_TAG} elegantcreationism/goclient-polygon:latest
	docker push elegantcreationism/goclient-polygon:latest

.PHONY: lint
lint: gen
	golangci-lint run

