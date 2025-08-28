package golang_test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	c "github.com/okieoth/hort/pkg/codegen/golang"
	p "github.com/okieoth/hort/pkg/jsonschemaparser"
)

func TestGenerateTypes(t *testing.T) {
	fileToUse := "../../jsonschemaparser/_resources/tests/test_schema.json"
	outputFile := "../../../examples/temp/testgeneratedtypes.go"
	bytes, err := os.ReadFile(fileToUse)
	require.Nil(t, err)
	parsedSchema, err := p.ParseBytes(bytes)
	require.Nil(t, err)
	templateBytes, err := os.ReadFile("../resources/templates/go_types.tmpl")
	require.Nil(t, err)
	var stringBuilder strings.Builder
	err = c.GenerateTypes(&parsedSchema, string(templateBytes), "dummy", &stringBuilder)
	require.Nil(t, err)
	s := stringBuilder.String()
	file, err := os.Create(outputFile)
	require.Nil(t, err)
	defer file.Close()
	file.WriteString(s)
}
