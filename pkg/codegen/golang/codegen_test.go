package golang_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	c "github.com/okieoth/hort/pkg/codegen/golang"
	p "github.com/okieoth/hort/pkg/jsonschemaparser"
	"github.com/okieoth/hort/pkg/jsonschemaparser/types"
)

type GenerateFunc func(parsedSchema *types.ParsedSchema, templateStr, packageName string, outputWriter io.Writer) error

func TestGenerateHortDaoTypes(t *testing.T) {
	fileToUse := "../../jsonschemaparser/_resources/tests/test_schema.json"
	outputFile := "../../../examples/crud/types/hort_types_dao.go"
	templateFile := "../resources/templates/go_dao_types_hort.tmpl"
	generateTypes(t, fileToUse, outputFile, templateFile, c.GenerateDaoTypes)
}

func TestGenerateHortTypes(t *testing.T) {
	fileToUse := "../../jsonschemaparser/_resources/tests/test_schema.json"
	outputFile := "../../../examples/crud/types/hort_types.go"
	templateFile := "../resources/templates/go_types_hort.tmpl"
	generateTypes(t, fileToUse, outputFile, templateFile, c.GenerateTypes)
}

func TestGenerateTypes(t *testing.T) {
	fileToUse := "../../jsonschemaparser/_resources/tests/test_schema.json"
	outputFile := "../../../examples/temp/types.go"
	templateFile := "../resources/templates/go_types.tmpl"
	generateTypes(t, fileToUse, outputFile, templateFile, c.GenerateTypes)
}

func generateTypes(t *testing.T, fileToUse, outputFile, templateFile string, f GenerateFunc) {
	bytes, err := os.ReadFile(fileToUse)
	require.Nil(t, err)
	parsedSchema, err := p.ParseBytes(bytes)
	require.Nil(t, err)
	templateBytes, err := os.ReadFile(templateFile)
	require.Nil(t, err)
	var stringBuilder strings.Builder
	err = f(&parsedSchema, string(templateBytes), "types", &stringBuilder)
	require.Nil(t, err)
	s := stringBuilder.String()
	file, err := os.Create(outputFile)
	require.Nil(t, err)
	defer file.Close()
	file.WriteString(s)
}
