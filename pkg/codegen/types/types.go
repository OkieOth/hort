package types

import "github.com/okieoth/hort/pkg/jsonschemaparser/types"

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
