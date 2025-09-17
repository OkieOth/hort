package protobuf_test

import (
	"os"
	"os/exec"
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

	_ = os.Remove(outputFile)

	var stringBuilder strings.Builder
	err = c.GenerateProtoFile(&parsedSchema, string(templateBytes), packageName, goPackageName, &stringBuilder)
	require.Nil(t, err)
	s := stringBuilder.String()
	file, err := os.Create(outputFile)
	require.Nil(t, err)
	defer file.Close()
	file.WriteString(s)
	require.FileExists(t, outputFile, "couldn't find generated protobuf file")

	typesFile := "../../../examples/crud/protobuf/service.pb.go"
	serviceFile := "../../../examples/crud/protobuf/service_grpc.pb.go"
	_ = os.Remove(typesFile)
	_ = os.Remove(serviceFile)
	cmd := exec.Command("protoc",
		"--proto_path=../../..",
		"--go_out=../../..",
		"--go_opt=paths=source_relative",
		"--go-grpc_out=../../..",
		"--go-grpc_opt=paths=source_relative",
		"examples/crud/protobuf/service.proto")
	out, err := cmd.CombinedOutput()
	require.Nil(t, err, "protoc error: %s", string(out))
	require.FileExists(t, typesFile, "couldn't find generated grpc types file")
	require.FileExists(t, serviceFile, "couldn't find generated grpc service file")
}
