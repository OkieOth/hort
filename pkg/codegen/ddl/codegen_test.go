package ddl_test

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	c "github.com/okieoth/hort/pkg/codegen/ddl"
	p "github.com/okieoth/hort/pkg/jsonschemaparser"
)

func TestGenerateCreateTables(t *testing.T) {
	fileToUse := "../../jsonschemaparser/_resources/tests/test_schema.json"
	outputFile := "../../../examples/temp/testgeneratedtypes.sql"
	dbFile := "../../../examples/temp/testgeneratedtypes.sqlite"
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

	_ = os.Remove(dbFile)
	cmd := exec.Command("sqlite3", dbFile, fmt.Sprintf(".read %s", outputFile))
	out, err := cmd.CombinedOutput()
	require.Nil(t, err, "sqlite3 error: %s", string(out))

	// verify that tables were created (example: check sqlite_master)
	cmd = exec.Command("sqlite3", dbFile, "SELECT name FROM sqlite_master WHERE type='table' ORDER BY name;")
	out, err = cmd.CombinedOutput()
	require.Nil(t, err, "failed to query sqlite_master: %s", string(out))
	tables := strings.TrimSpace(string(out))
	require.Equal(t, "Person\nPersonContact\nPersonContactAddress\nPersonName\nPerson_roles", tables, "wrong tables created")
}
