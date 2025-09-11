package types

import (
	o "github.com/okieoth/goptional/v3"
)

type Ordered[T any] struct {
	Name  string
	Value T
}

// This type only covers the main parsed types, that most likely need to be handled
// in the first place.
type ParsedSchema struct {
	ComplexTypes  []Ordered[ComplexType]
	ArrayTypes    []Ordered[ArrayType]
	MapTypes      []Ordered[MapType]
	IntEnums      []Ordered[IntEnumType]
	StringEnums   []Ordered[StringEnumType]
	IntegerTypes  []Ordered[IntegerType]
	NumberTypes   []Ordered[NumberType]
	StringTypes   []Ordered[StringType]
	UUIDTypes     []Ordered[UUIDType]
	DateTypes     []Ordered[DateType]
	DateTimeTypes []Ordered[DateTimeType]
	TimeTypes     []Ordered[TimeType]
	DurationTypes []Ordered[DurationType]
	BoolTypes     []Ordered[BoolType]
	BinaryTypes   []Ordered[BinaryType]
	ObjectTypes   []Ordered[ObjectType]
}

func NewParsedSchema() ParsedSchema {
	return ParsedSchema{
		ComplexTypes:  make([]Ordered[ComplexType], 0),
		ArrayTypes:    make([]Ordered[ArrayType], 0),
		MapTypes:      make([]Ordered[MapType], 0),
		IntEnums:      make([]Ordered[IntEnumType], 0),
		StringEnums:   make([]Ordered[StringEnumType], 0),
		IntegerTypes:  make([]Ordered[IntegerType], 0),
		NumberTypes:   make([]Ordered[NumberType], 0),
		StringTypes:   make([]Ordered[StringType], 0),
		UUIDTypes:     make([]Ordered[UUIDType], 0),
		DateTypes:     make([]Ordered[DateType], 0),
		DateTimeTypes: make([]Ordered[DateTimeType], 0),
		TimeTypes:     make([]Ordered[TimeType], 0),
		DurationTypes: make([]Ordered[DurationType], 0),
		BoolTypes:     make([]Ordered[BoolType], 0),
		BinaryTypes:   make([]Ordered[BinaryType], 0),
		ObjectTypes:   make([]Ordered[ObjectType], 0),
	}
}

func (p *ParsedSchema) GetIntEnumByName(name string) (Ordered[IntEnumType], bool) {
	for _, ie := range p.IntEnums {
		if ie.Name == name {
			return ie, true
		}
	}
	return Ordered[IntEnumType]{}, false
}

func (p *ParsedSchema) GetStringEnumByName(name string) (Ordered[StringEnumType], bool) {
	for _, ie := range p.StringEnums {
		if ie.Name == name {
			return ie, true
		}
	}
	return Ordered[StringEnumType]{}, false
}

func (p *ParsedSchema) GetComplexTypeByName(name string) (Ordered[ComplexType], bool) {
	for _, ie := range p.ComplexTypes {
		if ie.Name == name {
			return ie, true
		}
	}
	return Ordered[ComplexType]{}, false
}

func (p *ParsedSchema) GetArrayTypeByName(name string) (Ordered[ArrayType], bool) {
	for _, ie := range p.ArrayTypes {
		if ie.Name == name {
			return ie, true
		}
	}
	return Ordered[ArrayType]{}, false
}

func (p *ParsedSchema) GetMapTypeByName(name string) (Ordered[MapType], bool) {
	for _, ie := range p.MapTypes {
		if ie.Name == name {
			return ie, true
		}
	}
	return Ordered[MapType]{}, false
}

func (p *ParsedSchema) GetIntegerTypeByName(name string) (Ordered[IntegerType], bool) {
	for _, ie := range p.IntegerTypes {
		if ie.Name == name {
			return ie, true
		}
	}
	return Ordered[IntegerType]{}, false
}

func (p *ParsedSchema) GetNumberTypeByName(name string) (Ordered[NumberType], bool) {
	for _, ie := range p.NumberTypes {
		if ie.Name == name {
			return ie, true
		}
	}
	return Ordered[NumberType]{}, false
}

func (p *ParsedSchema) GetStringTypeByName(name string) (Ordered[StringType], bool) {
	for _, ie := range p.StringTypes {
		if ie.Name == name {
			return ie, true
		}
	}
	return Ordered[StringType]{}, false
}

func (p *ParsedSchema) GetUUIDTypeByName(name string) (Ordered[UUIDType], bool) {
	for _, ie := range p.UUIDTypes {
		if ie.Name == name {
			return ie, true
		}
	}
	return Ordered[UUIDType]{}, false
}

func (p *ParsedSchema) GetDateTypeByName(name string) (Ordered[DateType], bool) {
	for _, ie := range p.DateTypes {
		if ie.Name == name {
			return ie, true
		}
	}
	return Ordered[DateType]{}, false
}

func (p *ParsedSchema) GetDateTimeTypeByName(name string) (Ordered[DateTimeType], bool) {
	for _, ie := range p.DateTimeTypes {
		if ie.Name == name {
			return ie, true
		}
	}
	return Ordered[DateTimeType]{}, false
}

func (p *ParsedSchema) GetTimeTypeByName(name string) (Ordered[TimeType], bool) {
	for _, ie := range p.TimeTypes {
		if ie.Name == name {
			return ie, true
		}
	}
	return Ordered[TimeType]{}, false
}

func (p *ParsedSchema) GetDurationTypeByName(name string) (Ordered[DurationType], bool) {
	for _, ie := range p.DurationTypes {
		if ie.Name == name {
			return ie, true
		}
	}
	return Ordered[DurationType]{}, false
}

func (p *ParsedSchema) GetBoolTypeByName(name string) (Ordered[BoolType], bool) {
	for _, ie := range p.BoolTypes {
		if ie.Name == name {
			return ie, true
		}
	}
	return Ordered[BoolType]{}, false
}

func (p *ParsedSchema) GetBinaryTypeByName(name string) (Ordered[BinaryType], bool) {
	for _, ie := range p.BinaryTypes {
		if ie.Name == name {
			return ie, true
		}
	}
	return Ordered[BinaryType]{}, false
}

func (p *ParsedSchema) GetObjectTypeByName(name string) (Ordered[ObjectType], bool) {
	for _, ie := range p.ObjectTypes {
		if ie.Name == name {
			return ie, true
		}
	}
	return Ordered[ObjectType]{}, false
}

func (p *ParsedSchema) AddIntEnum(name string, ie IntEnumType) {
	o := Ordered[IntEnumType]{
		Name:  name,
		Value: ie,
	}
	p.IntEnums = append(p.IntEnums, o)
}

func (p *ParsedSchema) AddStringEnum(name string, ie StringEnumType) {
	o := Ordered[StringEnumType]{
		Name:  name,
		Value: ie,
	}
	p.StringEnums = append(p.StringEnums, o)
}

func (p *ParsedSchema) AddIntegerType(name string, v IntegerType) {
	o := Ordered[IntegerType]{
		Name:  name,
		Value: v,
	}
	p.IntegerTypes = append(p.IntegerTypes, o)
}

func (p *ParsedSchema) AddNumberType(name string, v NumberType) {
	o := Ordered[NumberType]{
		Name:  name,
		Value: v,
	}
	p.NumberTypes = append(p.NumberTypes, o)
}

func (p *ParsedSchema) AddBoolType(name string, v BoolType) {
	o := Ordered[BoolType]{
		Name:  name,
		Value: v,
	}
	p.BoolTypes = append(p.BoolTypes, o)
}

func (p *ParsedSchema) AddDateType(name string, v DateType) {
	o := Ordered[DateType]{
		Name:  name,
		Value: v,
	}
	p.DateTypes = append(p.DateTypes, o)
}

func (p *ParsedSchema) AddDateTimeType(name string, v DateTimeType) {
	o := Ordered[DateTimeType]{
		Name:  name,
		Value: v,
	}
	p.DateTimeTypes = append(p.DateTimeTypes, o)
}

func (p *ParsedSchema) AddTimeType(name string, v TimeType) {
	o := Ordered[TimeType]{
		Name:  name,
		Value: v,
	}
	p.TimeTypes = append(p.TimeTypes, o)
}

func (p *ParsedSchema) AddUUIDType(name string, v UUIDType) {
	o := Ordered[UUIDType]{
		Name:  name,
		Value: v,
	}
	p.UUIDTypes = append(p.UUIDTypes, o)
}

func (p *ParsedSchema) AddDurationType(name string, v DurationType) {
	o := Ordered[DurationType]{
		Name:  name,
		Value: v,
	}
	p.DurationTypes = append(p.DurationTypes, o)
}

func (p *ParsedSchema) AddObjectType(name string, v ObjectType) {
	o := Ordered[ObjectType]{
		Name:  name,
		Value: v,
	}
	p.ObjectTypes = append(p.ObjectTypes, o)
}

func (p *ParsedSchema) AddBinaryType(name string, v BinaryType) {
	o := Ordered[BinaryType]{
		Name:  name,
		Value: v,
	}
	p.BinaryTypes = append(p.BinaryTypes, o)
}

func (p *ParsedSchema) AddStringType(name string, v StringType) {
	o := Ordered[StringType]{
		Name:  name,
		Value: v,
	}
	p.StringTypes = append(p.StringTypes, o)
}

func (p *ParsedSchema) AddArrayType(name string, v ArrayType) {
	o := Ordered[ArrayType]{
		Name:  name,
		Value: v,
	}
	p.ArrayTypes = append(p.ArrayTypes, o)
}

func (p *ParsedSchema) AddMapType(name string, v MapType) {
	o := Ordered[MapType]{
		Name:  name,
		Value: v,
	}
	p.MapTypes = append(p.MapTypes, o)
}

func (p *ParsedSchema) AddComplexType(name string, v ComplexType) {
	o := Ordered[ComplexType]{
		Name:  name,
		Value: v,
	}
	p.ComplexTypes = append(p.ComplexTypes, o)
}

type StringType struct {
	Name      o.Optional[string]
	Default   o.Optional[string]
	Format    o.Optional[string]
	MinLength o.Optional[int]
	MaxLength o.Optional[int]
	Pattern   o.Optional[string]
}

type IntEnumType struct {
	Name    string
	Default o.Optional[int]
	Values  []int
}

type StringEnumType struct {
	Name    string
	Default o.Optional[string]
	Values  []string
}

type UUIDType struct {
	Name    o.Optional[string]
	Default o.Optional[string]
}

type DateType struct {
	Name             o.Optional[string]
	Default          o.Optional[string]
	Minimum          o.Optional[string]
	ExclusiveMinimum o.Optional[string]
	Maximum          o.Optional[string]
	ExclusiveMaximum o.Optional[string]
}

type DateTimeType struct {
	Name             o.Optional[string]
	Default          o.Optional[string]
	Minimum          o.Optional[string]
	ExclusiveMinimum o.Optional[string]
	Maximum          o.Optional[string]
	ExclusiveMaximum o.Optional[string]
}

type TimeType struct {
	Name             o.Optional[string]
	Default          o.Optional[string]
	Minimum          o.Optional[string]
	ExclusiveMinimum o.Optional[string]
	Maximum          o.Optional[string]
	ExclusiveMaximum o.Optional[string]
}

type DurationType struct {
	Name    o.Optional[string]
	Default o.Optional[string]
}

type IntegerType struct {
	Name             o.Optional[string]
	Format           o.Optional[string]
	Default          o.Optional[int]
	MultipleOf       o.Optional[int]
	Minimum          o.Optional[int]
	ExclusiveMinimum o.Optional[int]
	Maximum          o.Optional[int]
	ExclusiveMaximum o.Optional[int]
}

type NumberType struct {
	Name             o.Optional[string]
	Format           o.Optional[string]
	Default          o.Optional[float64]
	Minimum          o.Optional[float64]
	ExclusiveMinimum o.Optional[float64]
	Maximum          o.Optional[float64]
	ExclusiveMaximum o.Optional[float64]
}

type BoolType struct {
	Name    o.Optional[string]
	Default o.Optional[bool]
}

type BinaryType struct {
	Name        o.Optional[string]
	Description o.Optional[string]
}

type ObjectType struct {
	Name        o.Optional[string]
	Description o.Optional[string]
}

type ComplexType struct {
	Source      string
	TopLevel    bool
	Name        string
	Description o.Optional[string]
	Properties  []Property
	Tags        []string
}

type ArrayType struct {
	Source      string
	Name        o.Optional[string]
	MinItems    o.Optional[int]
	MaxItems    o.Optional[int]
	Description o.Optional[string]
	ValueType   any
}

type MapType struct {
	Source      string
	Name        string
	Description o.Optional[string]
	ValueType   any
	TopLevel    bool
}

// This type is only used as place holder while parsing references
type DummyType struct {
	Source string
	Name   string
}

type Property struct {
	Name         string
	ValueType    any
	ForeignKeyTo o.Optional[string]
	Description  o.Optional[string]
	Tags         []string
}
