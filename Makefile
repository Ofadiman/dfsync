include .env

debug:
	@go run main.go --from $(DFSYNC_FROM) --to $(DFSYNC_TO)

test:
	@go test -v ./...
