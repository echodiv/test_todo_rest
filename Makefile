.PHONY: build
build:
		go build -v ./cmd/tasks_srv

.PHONY: test
test:
		go test -v -race ./...

.DEFAULT_GOAL := build