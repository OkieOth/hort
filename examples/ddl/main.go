package main

import (
	"fmt"
	"os"

	p "github.com/okieoth/hort/pkg/jsonschemaparser"
)

func main() {
	fileToUse := "../../pkg/jsonschemaparser/_resources/tests/test_schema.json"
	bytes, err := os.ReadFile(fileToUse)
	if err != nil {
		panic(fmt.Sprintf("can't read input file: %v", err))
	}
	parseResult, err := p.ParseBytes(bytes)
	if err != nil {
		fmt.Println("error:", err)
		panic("error while parsing test file")
	}
	for name, complexType := range parseResult.ComplexTypes {
		fmt.Println("TODO", name, complexType)
	}
}
