include .env

debug:
	@go run ./... --source-directory $(DFSYNC_FROM) --target-directory $(DFSYNC_TO)

help:
	@go run ./... --help

version:
	@go run ./... --version

test:
	@go test -v ./...
