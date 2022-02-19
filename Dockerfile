FROM golang:1.7.1

COPY . /go/src/github.com/docker/go
WORKDIR /go/src/github.com/docker/go
