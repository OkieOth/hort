package main

import (
	"fmt"
	"os"

	c "github.com/okieoth/hort/pkg/codegen/golang"
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
	c.Dummy(&parseResult)
}
