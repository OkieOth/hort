package golang_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	c "github.com/okieoth/hort/pkg/codegen/golang"
	p "github.com/okieoth/hort/pkg/jsonschemaparser"
)

func TestGenerateTypes(t *testing.T) {
	fileToUse := "../../pkg/jsonschemaparser/_resources/tests/test_schema.json"
	bytes, err := os.ReadFile(fileToUse)
	require.Nil(t, err)
	parsedSchema, err := p.ParseBytes(bytes)
	if err != nil {
		fmt.Println("error:", err)
		panic("error while parsing test file")
	}
	templateBytes, err := os.ReadFile("../../pkg/codegen/resources/templates/go_types.tmpl")
	if err != nil {
		fmt.Println("error:", err)
		panic("error while reading template file")
	}
	err = c.GenerateTypes(&parsedSchema, string(templateBytes), "dummy", os.Stdout)

}
