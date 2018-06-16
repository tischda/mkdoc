# ---------------------------------------------------------------------------
# Makefile for mkdoc
# ---------------------------------------------------------------------------

PROJECT_DIR=$(notdir $(shell pwd))

BUILD_TAG=$(shell git describe --tags)
LDFLAGS=all=-ldflags "-X main.version=${BUILD_TAG} -s -w"

all: get build

build:
	go build ${LDFLAGS}

get:
	govendor sync

test: clean vet
	govendor test -v -cover

cover:
	govendor test -coverprofile=coverage.out +local
	go tool cover -html=coverage.out

fmt:
	govendor fmt +local

vet:
	govendor vet -v +local

install:
	go install ${LDFLAGS}

dist: clean build
	upx -9 ${PROJECT_DIR}.exe

clean:
	go clean
