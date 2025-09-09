package protobuf

import (
	"fmt"
	"io"
	"text/template"

	"github.com/okieoth/hort/pkg/codegen/helper"
	"github.com/okieoth/hort/pkg/jsonschemaparser/types"
)

type ProtoTemplateInput struct {
	Schema        *types.ParsedSchema
	PackageName   string
	GoPackageName string
}

func getProtoType(propName string, t any) string {
	if o, ok := t.(types.ComplexType); ok {
		return o.Name
	} else if o, ok := t.(types.ArrayType); ok {
		return fmt.Sprintf("repeated %s", getProtoType(propName, o.ValueType))
	} else if o, ok := t.(types.MapType); ok {
		return fmt.Sprintf("map[string]%s", getProtoType(propName, o.ValueType))
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
				return "int64"
			}
		} else {
			return "int32"
		}
	} else if o, ok := t.(types.NumberType); ok {
		if format, isSet := o.Format.Get(); isSet {
			switch format {
			case "float":
				return "float"
			case "double":
				return "double"
			default:
				return "float"
			}
		} else {
			return "float"
		}
	} else if _, ok := t.(types.StringType); ok {
		return "string"
	} else if _, ok := t.(types.UUIDType); ok {
		return "string"
	} else if _, ok := t.(types.DateType); ok {
		return "time.Time"
	} else if _, ok := t.(types.DateTimeType); ok {
		return "google.protobuf.Timestamp"
	} else if _, ok := t.(types.TimeType); ok {
		return "google.protobuf.TimeOfDay"
	} else if _, ok := t.(types.DurationType); ok {
		return "string"
	} else if _, ok := t.(types.BoolType); ok {
		return "bool"
	} else if _, ok := t.(types.BinaryType); ok {
		return "string"
	} else if _, ok := t.(types.ObjectType); ok {
		return "google.protobuf.Any"
	} else {
		return "??? unsupported type: " + propName
	}
}

func GenerateProtoFile(parsedSchema *types.ParsedSchema, templateStr, packageName, goPackageName string, outputWriter io.Writer) error {
	containsTime := func() bool {
		return helper.HasTimeAttribs(parsedSchema)
	}
	containsDate := func() bool {
		return helper.HasDateAttribs(parsedSchema)
	}
	containsTimestamp := func() bool {
		return helper.HasTimestampAttribs(parsedSchema)
	}
	containsObject := func() bool {
		return helper.HasPureObjectAttribs(parsedSchema)
	}
	tmpl := template.Must(template.New("GolangTypes").Funcs(
		template.FuncMap{
			"add":               helper.Add,
			"getProtoType":      getProtoType,
			"containsObject":    containsObject,
			"containsDate":      containsDate,
			"containsTime":      containsTime,
			"containsTimestamp": containsTimestamp,
			"typeHasTag":        helper.TypeHasTag,
			"upperFirstCase":    helper.UpperFirstCase,
			"upperCase":         helper.UpperCase,
		}).Parse(templateStr))
	templateInput := ProtoTemplateInput{
		Schema:        parsedSchema,
		PackageName:   packageName,
		GoPackageName: goPackageName,
	}
	tmpl.Execute(outputWriter, templateInput)
	return nil
}
