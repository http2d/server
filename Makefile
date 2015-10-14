GOFLAGS ?= $(GOFLAGS:)
PACKAGE := github.com/http2d/server

all: install test

install:
ifdef GODEBUG
		godebug build -instrument "${GODEBUG}" ./...
else
	  @go install -x $(GOFLAGS) ./...
endif

test: install
	@go test $(GOFLAGS) ./...

bench: install
	@go test -run=NONE -bench=. $(GOFLAGS) ./...

clean:
	@go clean $(GOFLAGS) -i ./...

deps:
	@go get -v -u github.com/cespare/reflex
	@go get -v -u github.com/mailgun/godebug
	@go get -v -u github.com/stretchr/testify

run: install
	server

devel: install
	reflex --regex='\.go$$' --start-service -- sh -c 'make run'

PHONY: all test clean install deps run devel
