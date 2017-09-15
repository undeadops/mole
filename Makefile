SHELL := /bin/bash

REV := $(shell git rev-parse HEAD)
CHANGES := $(shell test -n "$$(git status --porcelain)" && echo '+CHANGES' || true)

TARGET := mole
VERSION := $(shell cat VERSION)

OS := darwin linux
ARCH := amd64

fmt:
	go fmt ./...

env:
	@go env

build: deps
	go build -v \
	   -ldflags "$(LDFLAGS)" \
	   -o "$(TARGET)" .

build-all: deps
	mkdir -v -p $(CURDIR)/artifacts/$(VERSION)
	gox -verbose \
	    -os "$(OS)" -arch "$(ARCH)" \
	    -ldflags "$(LDFLAGS)" \
	    -output "$(CURDIR)/artifacts/$(VERSION)/{{.OS}}_{{.Arch}}/$(TARGET)" .
	cp -v -f \
	   $(CURDIR)/artifacts/$(VERSION)/$$(go env GOOS)_$$(go env GOARCH)/$(TARGET) .

doc:
	godoc -http=:8080 -index

version:
	@go version
