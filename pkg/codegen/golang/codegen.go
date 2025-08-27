package golang

import (
	"io"
	"text/template"

	"github.com/okieoth/hort/pkg/jsonschemaparser/types"
)

type TemplateInput struct {
	Schema      *types.ParsedSchema
	PackageName string
}

func GenerateTypes(parsedSchema *types.ParsedSchema, templateStr, packageName string, outputWriter io.Writer) error {
	tmpl := template.Must(template.New("GolangTypes").Funcs(
		template.FuncMap{}).Parse(templateStr))
	templateInput := TemplateInput{
		Schema:      parsedSchema,
		PackageName: packageName,
	}
	tmpl.Execute(outputWriter, templateInput)

	return nil
}
