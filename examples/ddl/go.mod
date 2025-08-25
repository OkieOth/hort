module github.com/okieoth/hort/examples/ddl

go 1.25.0

require github.com/okieoth/hort/pkg/jsonschemaparser v0.0.0

require (
	github.com/okieoth/goptional/v3 v3.1.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/okieoth/hort/pkg/jsonschemaparser v0.0.0 => ../../pkg/jsonschemaparser
