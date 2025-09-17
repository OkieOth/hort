package ddl

import (
	"html/template"
	"io"

	"github.com/okieoth/hort/pkg/codegen/helper"
	ct "github.com/okieoth/hort/pkg/codegen/types"
	"github.com/okieoth/hort/pkg/jsonschemaparser/types"
)

type TemplateInput struct {
	ComplexTypes []ct.ComplexTypeDef
}

func getColType(t any) string {
	if _, ok := t.(types.ComplexType); ok {
		return "INTEGER"
	} else if _, ok := t.(types.IntEnumType); ok {
		return "INTEGER"
	} else if _, ok := t.(types.StringEnumType); ok {
		return "TEXT"
	} else if _, ok := t.(types.IntegerType); ok {
		return "INTEGER"
	} else if _, ok := t.(types.NumberType); ok {
		return "REAL"
	} else if _, ok := t.(types.StringType); ok {
		return "TEXT"
	} else if _, ok := t.(types.UUIDType); ok {
		return "INTEGER"
	} else if _, ok := t.(types.DateType); ok {
		return "INTEGER"
	} else if _, ok := t.(types.DateTimeType); ok {
		return "INTEGER"
	} else if _, ok := t.(types.TimeType); ok {
		return "TEXT"
	} else if _, ok := t.(types.DurationType); ok {
		return "TEXT"
	} else if _, ok := t.(types.BoolType); ok {
		return "INTEGER"
	} else if _, ok := t.(types.BinaryType); ok {
		return "BLOB"
	} else {
		return "???"
	}
}

func GenerateCreateTables(parsedSchema *types.ParsedSchema, templateStr string, outputWriter io.Writer) error {
	complexTypes := helper.CreateTemplateInput(parsedSchema)
	tmpl := template.Must(template.New("Tables").Funcs(
		template.FuncMap{
			"isNeitherMapNorArray": helper.IsNeitherMapNorArray,
			"getColType":           getColType,
			"refPrefixIfNeeded":    helper.RefPrefixIfNeeded,
			"getContainerName":     helper.GetContainerName,
		}).Parse(templateStr))
	templateInput := TemplateInput{
		ComplexTypes: complexTypes,
	}
	tmpl.Execute(outputWriter, templateInput)
	return nil
}
