package golang_test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	c "github.com/okieoth/hort/pkg/codegen/golang"
	p "github.com/okieoth/hort/pkg/jsonschemaparser"
)

func TestGenerateHortTypes(t *testing.T) {
	fileToUse := "../../jsonschemaparser/_resources/tests/test_schema.json"
	outputFile := "../../../examples/crud/types/types.go"
	templateFile := "../resources/templates/go_types_hort.tmpl"
	generateTypes(t, fileToUse, outputFile, templateFile)
}

func TestGenerateTypes(t *testing.T) {
	fileToUse := "../../jsonschemaparser/_resources/tests/test_schema.json"
	outputFile := "../../../examples/temp/types.go"
	templateFile := "../resources/templates/go_types.tmpl"
	generateTypes(t, fileToUse, outputFile, templateFile)
}

func generateTypes(t *testing.T, fileToUse, outputFile, templateFile string) {
	bytes, err := os.ReadFile(fileToUse)
	require.Nil(t, err)
	parsedSchema, err := p.ParseBytes(bytes)
	require.Nil(t, err)
	templateBytes, err := os.ReadFile(templateFile)
	require.Nil(t, err)
	var stringBuilder strings.Builder
	err = c.GenerateTypes(&parsedSchema, string(templateBytes), "types", &stringBuilder)
	require.Nil(t, err)
	s := stringBuilder.String()
	file, err := os.Create(outputFile)
	require.Nil(t, err)
	defer file.Close()
	file.WriteString(s)
}
