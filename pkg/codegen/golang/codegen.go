package golang

import (
	"fmt"
	"io"

	"github.com/okieoth/hort/pkg/jsonschemaparser/types"
)

func GenerateTypes(parsedSchema *types.ParseResult, template, packageName string, outputWriter io.Writer) error {
	// TODO
	fmt.Println("hello from codegen/golang")
	return nil
}
