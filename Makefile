.PHONY: test

test:
	go test -C pkg/jsonschemaparser ./...

all: test
