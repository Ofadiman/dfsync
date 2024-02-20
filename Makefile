include .env

debug:
	@go run ./... --source-directory $(DFSYNC_FROM) --target-directory $(DFSYNC_TO)

help:
	@go run ./... --help

version:
	@go run ./... --version

test:
	@go test -v ./...

# `ctrl+p` followed by `ctrl+q` allows to detach session from running container without stopping it.
container_spawn:
	@docker run --name dfsync --interactive --tty --detach --rm --workdir /home/go/dfsync/ --volume $(shell pwd):/home/go/dfsync/ golang bash

container_attach:
	@docker container attach dfsync
 
container_stop:
	@docker container stop dfsync
