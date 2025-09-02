.PHONY: test

test:
	go test -C pkg/jsonschemaparser ./...
	go test -C pkg/codegen ./...

all: test
