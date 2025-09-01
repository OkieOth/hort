package ddl

import (
	"html/template"
	"io"

	"github.com/okieoth/hort/pkg/jsonschemaparser/types"
)

type ContainerType struct {
	PropName  string
	ValueType any
}

func NewContainerType(name string, valueType any) ContainerType {
	return ContainerType{
		PropName:  name,
		ValueType: valueType,
	}
}

type ComplexTypeDef struct {
	Type       types.ComplexType
	ArrayTypes []ContainerType // array definitions for the properties of that complex type
	MapTypes   []ContainerType // map definitions for the properties of that type
}

func arrayTypesFromComplexType(t types.ComplexType) []ContainerType {
	ret := make([]ContainerType, 0)
	for _, p := range t.Properties {
		if a, ok := p.ValueType.(types.ArrayType); ok {
			ret = append(ret, NewContainerType(p.Name, a.ValueType))
		}
	}
	return ret
}

func mapTypesFromComplexType(t types.ComplexType) []ContainerType {
	ret := make([]ContainerType, 0)
	for _, p := range t.Properties {
		if m, ok := p.ValueType.(types.MapType); ok {
			ret = append(ret, NewContainerType(p.Name, m.ValueType))
		}
	}
	return ret
}

func NewComplexTypeDef(t types.ComplexType) ComplexTypeDef {
	arrayTypes := arrayTypesFromComplexType(t)
	mapTypes := mapTypesFromComplexType(t)
	return ComplexTypeDef{
		Type:       t,
		ArrayTypes: arrayTypes,
		MapTypes:   mapTypes,
	}
}

type TemplateInput struct {
	ComplexTypes []ComplexTypeDef
}

func isNeitherMapNorArray(t any) bool {
	if _, ok := t.(types.ArrayType); ok {
		return false
	} else if _, ok := t.(types.MapType); ok {
		return false
	} else {
		return true
	}
}

func refPrefixIfNeeded(t any) string {
	if _, ok := t.(types.ComplexType); ok {
		return "_"
	} else if _, ok := t.(types.ArrayType); ok {
		return "_"
	} else if _, ok := t.(types.MapType); ok {
		return "_"
	} else {
		return ""
	}
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

func createTemplateInput(parsedSchema *types.ParsedSchema) []ComplexTypeDef {
	ret := make([]ComplexTypeDef, 0)
	for _, complexType := range parsedSchema.ComplexTypes {
		ret = append(ret, NewComplexTypeDef(complexType))
	}
	return ret
}

func getContainerName(t any) string {
	if x, ok := t.(types.ComplexType); ok {
		return x.Name
	} else {
		return "value"
	}
}

func GenerateCreateTables(parsedSchema *types.ParsedSchema, templateStr string, outputWriter io.Writer) error {
	complexTypes := createTemplateInput(parsedSchema)
	tmpl := template.Must(template.New("Tables").Funcs(
		template.FuncMap{
			"isNeitherMapNorArray": isNeitherMapNorArray,
			"getColType":           getColType,
			"refPrefixIfNeeded":    refPrefixIfNeeded,
			"getContainerName":     getContainerName,
		}).Parse(templateStr))
	templateInput := TemplateInput{
		ComplexTypes: complexTypes,
	}
	tmpl.Execute(outputWriter, templateInput)
	return nil
}
