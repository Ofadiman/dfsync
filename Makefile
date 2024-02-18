include .env

debug:
	@go run ./... --from $(DFSYNC_FROM) --to $(DFSYNC_TO)

help:
	@go run ./... --help

version:
	@go run ./... --version

test:
	@go test -v ./...
