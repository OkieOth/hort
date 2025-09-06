package helper

import (
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
