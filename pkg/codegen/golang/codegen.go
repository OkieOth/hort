package golang

import (
	"fmt"
	"io"
	"slices"
	"text/template"

	"github.com/okieoth/hort/pkg/codegen/helper"
	ct "github.com/okieoth/hort/pkg/codegen/types"
	"github.com/okieoth/hort/pkg/jsonschemaparser/types"
)

type TypesTemplateInput struct {
	Schema      *types.ParsedSchema
	PackageName string
}

type HortTypesTemplateInput struct {
	ComplexTypes []ct.ComplexTypeDef
	PackageName  string
}

func getGolangType(propName string, t any) string {
	if o, ok := t.(types.ComplexType); ok {
		return o.Name
	} else if o, ok := t.(types.ArrayType); ok {
		return fmt.Sprintf("[]%s", getGolangType(propName, o.ValueType))
	} else if o, ok := t.(types.MapType); ok {
		return fmt.Sprintf("map[string]%s", getGolangType(propName, o.ValueType))
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

func getDaoType(t any) string {
	if _, ok := t.(types.ComplexType); ok {
		return "int64"
	} else if _, ok := t.(types.IntEnumType); ok {
		return "int32"
	} else if _, ok := t.(types.StringEnumType); ok {
		return "string"
	} else if o, ok := t.(types.IntegerType); ok {
		if format, isSet := o.Format.Get(); isSet {
			switch format {
			case "int32":
				return "int32"
			case "int64":
				return "int64"
			default:
				return "int64"
			}
		} else {
			return "int64"
		}
	} else if o, ok := t.(types.NumberType); ok {
		if format, isSet := o.Format.Get(); isSet {
			switch format {
			case "float":
				return "float32"
			case "double":
				return "float64"
			default:
				return "float64"
			}
		} else {
			return "float64"
		}
	} else if _, ok := t.(types.StringType); ok {
		return "string"
	} else if _, ok := t.(types.UUIDType); ok {
		return "string // UUID"
	} else if _, ok := t.(types.DateType); ok {
		return "int64 // Date"
	} else if _, ok := t.(types.DateTimeType); ok {
		return "int64 // DateTime"
	} else if _, ok := t.(types.TimeType); ok {
		return "string // Time"
	} else if _, ok := t.(types.DurationType); ok {
		return "string // Duration"
	} else if _, ok := t.(types.BoolType); ok {
		return "int32 // bool"
	} else if _, ok := t.(types.BinaryType); ok {
		return "string // Binary"
	} else {
		return "??? unsupported type"
	}
}

func typeHasTag(t types.ComplexType, tag string) bool {
	return slices.Contains(t.Tags, tag)
}

func getDaoNeededProperties(t types.ComplexType) []types.Property {
	ret := make([]types.Property, 0)
	for _, p := range t.Properties {
		if _, isArray := p.ValueType.(types.ArrayType); isArray {
			continue
		}
		if _, isMap := p.ValueType.(types.MapType); isMap {
			continue
		}
		ret = append(ret, p)
	}
	return ret
}

func getMainTypes(parsedSchema *types.ParsedSchema) []types.ComplexType {
	ret := make([]types.ComplexType, 0)
	for _, ct := range parsedSchema.ComplexTypes {
		if slices.Contains(ct.Tags, "hort.main-type") {
			ret = append(ret, ct)
		}
	}
	return ret
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
			"getGolangType":      getGolangType,
			"isTimeImportNeeded": isTimeImportNeeded,
			"typeHasTag":         typeHasTag,
			"upperFirstCase":     helper.UpperFirstCase,
		}).Parse(templateStr))
	templateInput := TypesTemplateInput{
		Schema:      parsedSchema,
		PackageName: packageName,
	}
	tmpl.Execute(outputWriter, templateInput)
	return nil
}

func GenerateDaoTypes(parsedSchema *types.ParsedSchema, templateStr, packageName string, outputWriter io.Writer) error {
	complexTypes := helper.CreateTemplateInput(parsedSchema)
	tmpl := template.Must(template.New("Tables").Funcs(
		template.FuncMap{
			"isNeitherMapNorArray":  helper.IsNeitherMapNorArray,
			"getDaoType":            getDaoType,
			"refPrefixIfNeeded":     helper.RefPrefixIfNeeded,
			"getUpperContainerName": helper.GetUpperContainerName,
			"upperFirstCase":        helper.UpperFirstCase,
		}).Parse(templateStr))
	templateInput := HortTypesTemplateInput{
		ComplexTypes: complexTypes,
		PackageName:  packageName,
	}
	tmpl.Execute(outputWriter, templateInput)

	return nil
}
