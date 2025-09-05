package golang

import (
	"fmt"
	"io"
	"slices"
	"text/template"

	"github.com/okieoth/hort/pkg/jsonschemaparser/types"
)

type TemplateInput struct {
	Schema      *types.ParsedSchema
	PackageName string
}

func getTypeName(propName string, t any) string {
	if o, ok := t.(types.ComplexType); ok {
		return o.Name
	} else if o, ok := t.(types.ArrayType); ok {
		return fmt.Sprintf("[]%s", getTypeName(propName, o.ValueType))
	} else if o, ok := t.(types.MapType); ok {
		return fmt.Sprintf("map[string]%s", getTypeName(propName, o.ValueType))
	} else if o, ok := t.(types.IntEnumType); ok {
		return o.Name
	} else if o, ok := t.(types.StringEnumType); ok {
		return o.Name
	} else if o, ok := t.(types.IntegerType); ok {
		if format, isSet := o.Format.Get(); isSet {
			switch format {
			case "int32":
				return "int32"
			case "int64":
				return "int64"
			default:
				return "int"
			}
		} else {
			return "int"
		}
	} else if o, ok := t.(types.NumberType); ok {
		if format, isSet := o.Format.Get(); isSet {
			switch format {
			case "float":
				return "float32"
			case "double":
				return "float64"
			default:
				return "float"
			}
		} else {
			return "float"
		}
	} else if _, ok := t.(types.StringType); ok {
		return "string"
	} else if _, ok := t.(types.UUIDType); ok {
		return "string // UUID"
	} else if _, ok := t.(types.DateType); ok {
		return "time.Time // Date"
	} else if _, ok := t.(types.DateTimeType); ok {
		return "time.Time // DateTime"
	} else if _, ok := t.(types.TimeType); ok {
		return "string // Time"
	} else if _, ok := t.(types.DurationType); ok {
		return "string // Duration"
	} else if _, ok := t.(types.BoolType); ok {
		return "bool"
	} else if _, ok := t.(types.BinaryType); ok {
		return "string // Binary"
	} else {
		return "??? unsupported type: " + propName
	}
}

func typeHasTag(t types.ComplexType, tag string) bool {
	return slices.Contains(t.Tags, tag)
}

func hasTimeAttribs(parsedSchema *types.ParsedSchema) bool {
	for _, ct := range parsedSchema.ComplexTypes {
		for _, p := range ct.Properties {
			if _, ok := p.ValueType.(types.DateTimeType); ok {
				return true
			}
			if _, ok := p.ValueType.(types.DateType); ok {
				return true
			}
		}
	}
	return false
}

func GenerateTypes(parsedSchema *types.ParsedSchema, templateStr, packageName string, outputWriter io.Writer) error {
	isTimeImportNeeded := func() bool {
		return hasTimeAttribs(parsedSchema)
	}
	tmpl := template.Must(template.New("GolangTypes").Funcs(
		template.FuncMap{
			"getTypeName":        getTypeName,
			"isTimeImportNeeded": isTimeImportNeeded,
			"typeHasTag":         typeHasTag,
		}).Parse(templateStr))
	templateInput := TemplateInput{
		Schema:      parsedSchema,
		PackageName: packageName,
	}
	tmpl.Execute(outputWriter, templateInput)

	return nil
}
