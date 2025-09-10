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
