include .env

debug:
	@go run ./... --source-directory $(DFSYNC_FROM)

help:
	@go run ./... --help

version:
	@go run ./... --version

test:
	@go test -v ./...

run:
	@docker container run --name dfsync --interactive --tty --rm --volume $(shell pwd):/home/golang/dfsync/ -e "TERM=xterm-256color" -e LOG_LEVEL=$(LOG_LEVEL) dfsync bash

build:
	@docker image build --tag dfsync .

