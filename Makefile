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
	# Development tools
	@go get -v -u github.com/cespare/reflex
	@go get -v -u github.com/mailgun/godebug

	# Dependencies
	@go get -v -u golang.org/x/net/http2
	@go get -v -u github.com/spf13/viper
	@go get -v -u github.com/Sirupsen/logrus
	@go get -v -u github.com/stretchr/testify

	# Code Coverage
	@go get -v -u golang.org/x/tools/cmd/cover
	@go get -v -u github.com/mattn/goveralls
	@go get -v -u github.com/axw/gocov/gocov

run: install
	server

devel: install
	reflex --regex='\.go$$' --start-service -- sh -c 'make run'

PHONY: all test clean install deps run devel
