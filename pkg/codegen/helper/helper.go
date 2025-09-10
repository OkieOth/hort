package helper

import (
	"slices"
	"strings"

	ct "github.com/okieoth/hort/pkg/codegen/types"
	"github.com/okieoth/hort/pkg/jsonschemaparser/types"
)

func CreateTemplateInput(parsedSchema *types.ParsedSchema) []ct.ComplexTypeDef {
	ret := make([]ct.ComplexTypeDef, 0)
	for _, complexType := range parsedSchema.ComplexTypes {
		ret = append(ret, ct.NewComplexTypeDef(complexType))
	}
	return ret
}

func GetContainerName(t any) string {
	if x, ok := t.(types.ComplexType); ok {
		return x.Name
	} else {
		return "value"
	}
}

func UpperFirstCase(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}

func UpperCase(s string) string {
	return strings.ToUpper(s)
}

func LowerCase(s string) string {
	return strings.ToLower(s)
}

func LowerFirstCase(s string) string {
	return strings.ToLower(s[0:1]) + s[1:]
}

func GetUpperContainerName(t any) string {
	if x, ok := t.(types.ComplexType); ok {
		return UpperFirstCase(x.Name)
	} else {
		return "Value"
	}
}

func IsNeitherMapNorArray(t any) bool {
	if _, ok := t.(types.ArrayType); ok {
		return false
	} else if _, ok := t.(types.MapType); ok {
		return false
	} else {
		return true
	}
}

func RefPrefixIfNeeded(t any, prefix string) string {
	if _, ok := t.(types.ComplexType); ok {
		return prefix
	} else if _, ok := t.(types.ArrayType); ok {
		return prefix
	} else if _, ok := t.(types.MapType); ok {
		return prefix
	} else {
		return ""
	}
}

func HasGoTimeAttribs(parsedSchema *types.ParsedSchema) bool {
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

func hasAttribWithType[T any](parsedSchema *types.ParsedSchema) bool {
	for _, ct := range parsedSchema.ComplexTypes {
		for _, p := range ct.Properties {
			if _, ok := p.ValueType.(T); ok {
				return true
			}
		}
	}
	return false
}

func HasTimeAttribs(parsedSchema *types.ParsedSchema) bool {
	return hasAttribWithType[types.TimeType](parsedSchema)
}

func HasTimestampAttribs(parsedSchema *types.ParsedSchema) bool {
	return hasAttribWithType[types.DateTimeType](parsedSchema)
}

func HasDateAttribs(parsedSchema *types.ParsedSchema) bool {
	return hasAttribWithType[types.DateType](parsedSchema)
}

func HasPureObjectAttribs(parsedSchema *types.ParsedSchema) bool {
	return hasAttribWithType[types.ObjectType](parsedSchema)
}

func TypeHasTag(t types.ComplexType, tag string) bool {
	return slices.Contains(t.Tags, tag)
}

func Add(v1, v2 int) int {
	return v1 + v2
}
