.PHONY: test

test-jsonschemaparser:
	go test -C pkg/jsonschemaparser ./...

test-codegen:
	go test -C pkg/codegen ./...

test-hort.core:
	CGO_ENABLED=1 go test -C pkg/hort.core ./...

test: test-jsonschemaparser test-codegen \
  test-hort.core

all: test
