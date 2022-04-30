NAME=uigh
WRKDIR=$(shell pwd)
VERSION=$(shell git describe --tags || echo n/a)

all: fmt build

fmt:
	go fmt ./...

clean:
	@echo "****** CLEANING ******"
	rm -rf $(WRKDIR)/bin/

build:clean
	@echo "****** BUILDING BINARY ******"
	go build -o bin/$(NAME) .
