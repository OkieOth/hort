.PHONY: test

test-jsonschemaparser:
	go test -C pkg/jsonschemaparser ./...

test-codegen:
	go test -C pkg/codegen ./...

test-hort.core:
	CGO_ENABLED=1 go test -C pkg/hort.core ./...

test: test-jsonschemaparser test-codegen \
  test-hort.core


generate-example-hort:
	protoc --go_out=. --go_opt=paths=source_relative \
     --go-grpc_out=. --go-grpc_opt=paths=source_relative examples/crud/protobuf/service.proto

all: test
