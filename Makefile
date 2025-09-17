.PHONY: test

test-jsonschemaparser:
	go test -C pkg/jsonschemaparser ./...

test-codegen:
	go test -C pkg/codegen ./...

test: test-jsonschemaparser \
	test-codegen

generate-example-hort:
	protoc --go_out=. --go_opt=paths=source_relative \
     --go-grpc_out=. --go-grpc_opt=paths=source_relative examples/crud/protobuf/service.proto

all: test
