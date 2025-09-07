module github.com/okieoth/hort/pkg/hort.core

go 1.25.0

require github.com/okieoth/hort/pkg/jsonschemaparser v0.0.0

require (
	github.com/mattn/go-sqlite3 v1.14.32 // indirect
	github.com/stretchr/testify v1.11.1 // indirect
)

replace github.com/okieoth/hort/pkg/jsonschemaparser v0.0.0 => ../jsonschemaparser
