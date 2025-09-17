# Requirements

* protobuf - `sudo apt  install protobuf-compiler`
* golang codegen - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
# gRPC workflow

1. create a proto file
2. Generate go service code
```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    examples/crud/protobuf/service.proto
```
