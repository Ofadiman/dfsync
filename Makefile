include .env

debug:
	@go run ./... --from $(DFSYNC_FROM) --to $(DFSYNC_TO)

test:
	@go test -v ./...
