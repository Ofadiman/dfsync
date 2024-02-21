FROM golang

# Ensure linux permissions are correct when mounting files.
RUN groupadd --gid 1000 dfsync
RUN useradd --create-home --home-dir /home/golang/ --shell /bin/bash --gid 1000 --uid 1000 golang

USER golang

WORKDIR /home/golang/dfsync/

COPY go.mod go.sum ./

RUN go mod download
