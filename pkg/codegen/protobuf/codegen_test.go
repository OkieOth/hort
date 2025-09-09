package protobuf_test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	c "github.com/okieoth/hort/pkg/codegen/protobuf"
	p "github.com/okieoth/hort/pkg/jsonschemaparser"
)

func TestGenerateProtoFile(t *testing.T) {
	fileToUse := "../../jsonschemaparser/_resources/tests/test_schema.json"
	outputFile := "../../../examples/crud/protobuf/service.proto"
	bytes, err := os.ReadFile(fileToUse)
	require.Nil(t, err)
	parsedSchema, err := p.ParseBytes(bytes)
	require.Nil(t, err)
	templateBytes, err := os.ReadFile("../resources/templates/protobuf_hort.tmpl")
	require.Nil(t, err)
	packageName := "person"
	goPackageName := "github.com/okieoth/hort/examples/crud/protobuf"
	var stringBuilder strings.Builder
	err = c.GenerateProtoFile(&parsedSchema, string(templateBytes), packageName, goPackageName, &stringBuilder)
	require.Nil(t, err)
	s := stringBuilder.String()
	file, err := os.Create(outputFile)
	require.Nil(t, err)
	defer file.Close()
	file.WriteString(s)

	// _ = os.Remove(dbFile)
	// cmd := exec.Command("sqlite3", dbFile, fmt.Sprintf(".read %s", outputFile))
	// out, err := cmd.CombinedOutput()
	// require.Nil(t, err, "sqlite3 error: %s", string(out))

	// // verify that tables were created (example: check sqlite_master)
	// cmd = exec.Command("sqlite3", dbFile, "SELECT name FROM sqlite_master WHERE type='table' ORDER BY name;")
	// out, err = cmd.CombinedOutput()
	// require.Nil(t, err, "failed to query sqlite_master: %s", string(out))
	// tables := strings.TrimSpace(string(out))
	// require.Equal(t, "Person\nPersonContact\nPersonContactAddress\nPersonName\nPerson_roles", tables, "wrong tables created")
}
