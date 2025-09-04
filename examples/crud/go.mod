module github.com/okieoth/hort/examples/crud

go 1.25.0

replace github.com/okieoth/hort/pkg/jsonschemaparser v0.0.0 => ../../pkg/jsonschemaparser

replace github.com/okieoth/hort/pkg/codegen v0.0.0 => ../../pkg/codegen

require (
	github.com/mattn/go-sqlite3 v1.14.32
	github.com/okieoth/hort/pkg/jsonschemaparser v0.0.0
)

require (
	github.com/okieoth/goptional/v3 v3.1.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
