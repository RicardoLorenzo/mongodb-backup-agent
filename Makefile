#!/usr/bin/make

SHELL := /bin/bash
PWD = $(shell pwd)

BIN = $(shell basename $(PWD))
GO := /usr/local/go/bin/go

FIND_STD_DEPS = $(GO) list std | sort | uniq
DEPS          = "$(PKG) $(FIND_STD_DEPS)"

.PHONY: %

all: build
build:
	export GOPATH=/Users/ricardo.lorenzo/Development
	export GOBIN=$(GOPATH)/bin
	$(GO) install github.com/RicardoLorenzo/mongodb-backup-agent
fmt:
	$(GO) fmt $(PKG)
clean:
	$(GO) clean -i $(PKG)
clean-all:
	$(GO) clean -i -r $(PKG)
deps:
	#$(GO) install github.com/RicardoLorenzo/mongodb-backup-agent/storage
	#$(GO) install github.com/RicardoLorenzo/mongodb-backup-agent
run: all
	./$(BIN)
