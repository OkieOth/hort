package types

import (
	o "github.com/okieoth/goptional/v3"
)

// This type only covers the main parsed types, that most likely need to be handled
// in the first place.
type ParsedSchema struct {
	ComplexTypes  []ComplexType
	ArrayTypes    []ArrayType
	MapTypes      []MapType
	IntEnums      []IntEnumType
	StringEnums   []StringEnumType
	IntegerTypes  []IntegerType
	NumberTypes   []NumberType
	StringTypes   []StringType
	UUIDTypes     []UUIDType
	DateTypes     []DateType
	DateTimeTypes []DateTimeType
	TimeTypes     []TimeType
	DurationTypes []DurationType
	BoolTypes     []BoolType
	BinaryTypes   []BinaryType
	ObjectTypes   []ObjectType
}

func NewParsedSchema() ParsedSchema {
	return ParsedSchema{
		ComplexTypes:  make([]ComplexType, 0),
		ArrayTypes:    make([]ArrayType, 0),
		MapTypes:      make([]MapType, 0),
		IntEnums:      make([]IntEnumType, 0),
		StringEnums:   make([]StringEnumType, 0),
		IntegerTypes:  make([]IntegerType, 0),
		NumberTypes:   make([]NumberType, 0),
		StringTypes:   make([]StringType, 0),
		UUIDTypes:     make([]UUIDType, 0),
		DateTypes:     make([]DateType, 0),
		DateTimeTypes: make([]DateTimeType, 0),
		TimeTypes:     make([]TimeType, 0),
		DurationTypes: make([]DurationType, 0),
		BoolTypes:     make([]BoolType, 0),
		BinaryTypes:   make([]BinaryType, 0),
		ObjectTypes:   make([]ObjectType, 0),
	}
}

type NamedObject interface {
	GetName() string
}

type StringType struct {
	Name      o.Optional[string]
	Default   o.Optional[string]
	Format    o.Optional[string]
	MinLength o.Optional[int]
	MaxLength o.Optional[int]
	Pattern   o.Optional[string]
}

func (t StringType) GetName() string {
	if s, isSet := t.Name.Get(); isSet {
		return s
	} else {
		return ""
	}
}

type IntEnumType struct {
	Name    string
	Default o.Optional[int]
	Values  []int
}

func (t IntEnumType) GetName() string {
	return t.Name
}

type StringEnumType struct {
	Name    string
	Default o.Optional[string]
	Values  []string
}

func (t StringEnumType) GetName() string {
	return t.Name
}

type UUIDType struct {
	Name    o.Optional[string]
	Default o.Optional[string]
}

func (t UUIDType) GetName() string {
	if s, isSet := t.Name.Get(); isSet {
		return s
	} else {
		return ""
	}
}

type DateType struct {
	Name             o.Optional[string]
	Default          o.Optional[string]
	Minimum          o.Optional[string]
	ExclusiveMinimum o.Optional[string]
	Maximum          o.Optional[string]
	ExclusiveMaximum o.Optional[string]
}

func (t DateType) GetName() string {
	if s, isSet := t.Name.Get(); isSet {
		return s
	} else {
		return ""
	}
}

type DateTimeType struct {
	Name             o.Optional[string]
	Default          o.Optional[string]
	Minimum          o.Optional[string]
	ExclusiveMinimum o.Optional[string]
	Maximum          o.Optional[string]
	ExclusiveMaximum o.Optional[string]
}

func (t DateTimeType) GetName() string {
	if s, isSet := t.Name.Get(); isSet {
		return s
	} else {
		return ""
	}
}

type TimeType struct {
	Name             o.Optional[string]
	Default          o.Optional[string]
	Minimum          o.Optional[string]
	ExclusiveMinimum o.Optional[string]
	Maximum          o.Optional[string]
	ExclusiveMaximum o.Optional[string]
}

func (t TimeType) GetName() string {
	if s, isSet := t.Name.Get(); isSet {
		return s
	} else {
		return ""
	}
}

type DurationType struct {
	Name    o.Optional[string]
	Default o.Optional[string]
}

func (t DurationType) GetName() string {
	if s, isSet := t.Name.Get(); isSet {
		return s
	} else {
		return ""
	}
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

func (t IntegerType) GetName() string {
	if s, isSet := t.Name.Get(); isSet {
		return s
	} else {
		return ""
	}
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

func (t NumberType) GetName() string {
	if s, isSet := t.Name.Get(); isSet {
		return s
	} else {
		return ""
	}
}

type BoolType struct {
	Name    o.Optional[string]
	Default o.Optional[bool]
}

func (t BoolType) GetName() string {
	if s, isSet := t.Name.Get(); isSet {
		return s
	} else {
		return ""
	}
}

type BinaryType struct {
	Name        o.Optional[string]
	Description o.Optional[string]
}

func (t BinaryType) GetName() string {
	if s, isSet := t.Name.Get(); isSet {
		return s
	} else {
		return ""
	}
}

type ObjectType struct {
	Name        o.Optional[string]
	Description o.Optional[string]
}

func (t ObjectType) GetName() string {
	if s, isSet := t.Name.Get(); isSet {
		return s
	} else {
		return ""
	}
}

type ComplexType struct {
	Source      string
	TopLevel    bool
	Name        string
	Description o.Optional[string]
	Properties  []Property
	Tags        []string
}

func (t ComplexType) GetName() string {
	return t.Name
}

type ArrayType struct {
	Source      string
	Name        o.Optional[string]
	MinItems    o.Optional[int]
	MaxItems    o.Optional[int]
	Description o.Optional[string]
	ValueType   any
}

func (t ArrayType) GetName() string {
	if s, isSet := t.Name.Get(); isSet {
		return s
	} else {
		return ""
	}
}

type MapType struct {
	Source      string
	Name        string
	Description o.Optional[string]
	ValueType   any
	TopLevel    bool
}

func (t MapType) GetName() string {
	return t.Name
}

// This type is only used as place holder while parsing references
type DummyType struct {
	Source string
	Name   string
}

func (t DummyType) GetName() string {
	return t.Name
}

type Property struct {
	Name         string
	ValueType    any
	ForeignKeyTo o.Optional[string]
	Description  o.Optional[string]
	Tags         []string
}

func (t Property) GetName() string {
	return t.Name
}
