module github.com/okieoth/hort/examples/jsonschemaparser_examples

go 1.25.0

require github.com/okieoth/hort/pkg/jsonschemaparser v0.0.0

replace github.com/okieoth/hort/pkg/jsonschemaparser v0.0.0 => ../../pkg/jsonschemaparser

require (
	github.com/okieoth/goptional/v3 v3.1.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
