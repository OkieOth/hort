module github.com/okieoth/hort/examples/crud

go 1.25.0

replace github.com/okieoth/hort/pkg/jsonschemaparser v0.0.0 => ../../pkg/jsonschemaparser

replace github.com/okieoth/hort/pkg/codegen v0.0.0 => ../../pkg/codegen

require (
	github.com/mattn/go-sqlite3 v1.14.32
	github.com/okieoth/hort/pkg/jsonschemaparser v0.0.0
	google.golang.org/grpc v1.75.0
	google.golang.org/protobuf v1.36.8
)

require (
	github.com/okieoth/goptional/v3 v3.1.0 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250707201910-8d1bb00bc6a7 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
