package orderedreader_test

import (
	"encoding/json"
	"os"
	"testing"

	ordered "github.com/okieoth/hort/pkg/jsonschemaparser/orderedreader"
)

func TestJsonKeepOrder(t *testing.T) {
	inputFile := "../_resources/tests/test_schema.json"
	outputFile := "../../../temp/test_schema_output.json"

	inputBytes, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	var root ordered.OrderedValue
	if err := json.Unmarshal(inputBytes, &root); err != nil {
		panic(err)
	}

	// Write back to output.json
	outputBytes, err := json.MarshalIndent(&root, "", "  ")
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(outputFile, outputBytes, 0644); err != nil {
		panic(err)
	}
}
