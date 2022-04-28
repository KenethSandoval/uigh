NAME=uigh
WRKDIR=$(pwd)
VERSION=$(shell git describe --tags || echo n/a)

all: fmt build

fmt:
	go fmt ./...

clean:
	rm -rf $(WRKDIR)/build/

build:clean
	go build -o bin/$(NAME) .
