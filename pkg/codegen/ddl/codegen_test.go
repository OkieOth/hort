package ddl_test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	c "github.com/okieoth/hort/pkg/codegen/ddl"
	p "github.com/okieoth/hort/pkg/jsonschemaparser"
)

func TestGenerateCreateTables(t *testing.T) {
	fileToUse := "../../jsonschemaparser/_resources/tests/test_schema.json"
	outputFile := "../../../examples/temp/testgeneratedtypes.sql"
	bytes, err := os.ReadFile(fileToUse)
	require.Nil(t, err)
	parsedSchema, err := p.ParseBytes(bytes)
	require.Nil(t, err)
	templateBytes, err := os.ReadFile("../resources/templates/create_tables.tmpl")
	require.Nil(t, err)
	var stringBuilder strings.Builder
	err = c.GenerateCreateTables(&parsedSchema, string(templateBytes), &stringBuilder)
	require.Nil(t, err)
	s := stringBuilder.String()
	file, err := os.Create(outputFile)
	require.Nil(t, err)
	defer file.Close()
	file.WriteString(s)
}
