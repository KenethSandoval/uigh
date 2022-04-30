NAME=uigh
WRKDIR=$(shell pwd)
VERSION=$(shell git describe --tags || echo n/a)

all: fmt build

fmt:
	go fmt ./...

clean:
	@echo "****** CLEANING ******"
	@echo
	rm -rf $(WRKDIR)/bin/

build:clean
	@echo "****** BUILDING BINARY ******"
	@echo
	go build -o bin/$(NAME) .
