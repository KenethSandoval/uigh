NAME=uigh
WRKDIR=$(shell pwd)
VERSION=$(shell git describe --tags || echo n/a)

all: fmt build

fmt:
	go fmt ./...

clean:
	rm -rf $(WRKDIR)/bin/

build:clean
	go build -o bin/$(NAME) .
