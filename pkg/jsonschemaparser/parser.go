package jsonschemaparser

import (
	"encoding/json"
	"fmt"
	"reflect"
	"slices"
	"strings"
	"time"
	"unicode"

	o "github.com/okieoth/goptional/v3"
	ordered "github.com/okieoth/hort/pkg/jsonschemaparser/orderedreader"
	"github.com/okieoth/hort/pkg/jsonschemaparser/types"
)

const NO_NAME = ""

func ignoreIfEmptyStr(s string) bool {
	return s != ""
}

func ParseBytes(input []byte) (types.ParsedSchema, error) {
	var parsedSchema ordered.OrderedValue
	extractedTypes := types.NewParsedSchema()

	if err := json.Unmarshal(input, &parsedSchema); err != nil {
		return extractedTypes, fmt.Errorf("error while unmarshalling schema: %v", err)
	}

	var definitions ordered.OrderedObject
	var found bool
	definitions, found = parsedSchema.ObjectValueForKey("definitions")
	if !found {
		// respects json schema draft 2020-12
		definitions, found = parsedSchema.ObjectValueForKey("$defs")
	}
	if found {
		err := parseTypesFromDefinition(definitions, &extractedTypes)
		if err != nil {
			return extractedTypes, fmt.Errorf("error while parsing types in the definitions section: %v", err)
		}
	}

	// to level object
	err := parseTopLevelType(&parsedSchema, &extractedTypes)
	if err != nil {
		return extractedTypes, fmt.Errorf("error while parsing main type: %v", err)
	}

	// resolve so far skipped references, that were later extract in the parsing and
	// temporary replaced with DummyTypes
	if err := resolveDummyTypesForComplexTypes(&extractedTypes); err != nil {
		return extractedTypes, fmt.Errorf("error while resolving dummy types for complex types: %v", err)
	}
	if err := resolveDummyTypesForMapTypes(&extractedTypes); err != nil {
		return extractedTypes, fmt.Errorf("error while resolving dummy types for map types: %v", err)
	}
	if err := resolveDummyTypesForArrayTypes(&extractedTypes); err != nil {
		return extractedTypes, fmt.Errorf("error while resolving dummy types for array types: %v", err)
	}
	return extractedTypes, nil
}

func getTypeByNameFromMap[T any](nameToFind string, mapToCheck []types.Ordered[T]) (any, bool) {
	for _, t := range mapToCheck {
		if t.Name == nameToFind {
			return t, true
		}
	}
	return types.DummyType{}, false
}

func getTypeByName(extractedTypes *types.ParsedSchema, typeName string) (any, error) {
	if t, found := getTypeByNameFromMap(typeName, extractedTypes.StringEnums); found {
		return t, nil
	}
	if t, found := getTypeByNameFromMap(typeName, extractedTypes.ComplexTypes); found {
		return t, nil
	}
	if t, found := getTypeByNameFromMap(typeName, extractedTypes.IntEnums); found {
		return t, nil
	}
	if t, found := getTypeByNameFromMap(typeName, extractedTypes.ArrayTypes); found {
		return t, nil
	}
	if t, found := getTypeByNameFromMap(typeName, extractedTypes.MapTypes); found {
		return t, nil
	}
	if t, found := getTypeByNameFromMap(typeName, extractedTypes.IntegerTypes); found {
		return t, nil
	}
	if t, found := getTypeByNameFromMap(typeName, extractedTypes.NumberTypes); found {
		return t, nil
	}
	if t, found := getTypeByNameFromMap(typeName, extractedTypes.StringTypes); found {
		return t, nil
	}
	if t, found := getTypeByNameFromMap(typeName, extractedTypes.UUIDTypes); found {
		return t, nil
	}
	if t, found := getTypeByNameFromMap(typeName, extractedTypes.DateTypes); found {
		return t, nil
	}
	if t, found := getTypeByNameFromMap(typeName, extractedTypes.DateTimeTypes); found {
		return t, nil
	}
	if t, found := getTypeByNameFromMap(typeName, extractedTypes.TimeTypes); found {
		return t, nil
	}
	if t, found := getTypeByNameFromMap(typeName, extractedTypes.DurationTypes); found {
		return t, nil
	}
	if t, found := getTypeByNameFromMap(typeName, extractedTypes.BoolTypes); found {
		return t, nil
	}
	if t, found := getTypeByNameFromMap(typeName, extractedTypes.BinaryTypes); found {
		return t, nil
	}
	if t, found := getTypeByNameFromMap(typeName, extractedTypes.ObjectTypes); found {
		return t, nil
	}
	return types.DummyType{}, fmt.Errorf("couldn't find type with name: %s", typeName)
}

func resolveDummyTypesForComplexTypes(extractedTypes *types.ParsedSchema) error {
	for j, o := range extractedTypes.ComplexTypes {
		complexType := o.Value
		wasChanged := false
		for i, property := range complexType.Properties {
			if dummyType, isDummy := property.ValueType.(types.DummyType); isDummy {
				foundType, err := getTypeByName(extractedTypes, dummyType.Name)
				if err != nil {
					return fmt.Errorf("couldn't replace property dummy type, type: %s, property: %s, dummyType: %s, error: %v ",
						complexType.Name, property.Name, dummyType.Name, err)
				}
				complexType.Properties[i].ValueType = foundType
				wasChanged = true
			}
		}
		if wasChanged {
			extractedTypes.ComplexTypes[j].Value = complexType
		}
	}
	return nil
}

func resolveDummyTypesForMapTypes(extractedTypes *types.ParsedSchema) error {
	for j, o := range extractedTypes.MapTypes {
		mapType := o.Value
		if dummyType, isDummy := mapType.ValueType.(types.DummyType); isDummy {
			foundType, err := getTypeByName(extractedTypes, dummyType.Name)
			if err != nil {
				return fmt.Errorf("couldn't replace value dummy type in map, type: %v, dummyType: %s, error: %v ",
					mapType.Name, dummyType.Name, err)
			}
			mapType.ValueType = foundType
			extractedTypes.MapTypes[j].Value = mapType
		}
	}
	return nil
}

func resolveDummyTypesForArrayTypes(extractedTypes *types.ParsedSchema) error {
	for j, o := range extractedTypes.ArrayTypes {
		arrayType := o.Value
		if dummyType, isDummy := arrayType.ValueType.(types.DummyType); isDummy {
			foundType, err := getTypeByName(extractedTypes, dummyType.Name)
			if err != nil {
				return fmt.Errorf("couldn't replace value dummy type in array, type: %v, dummyType: %s, error: %v",
					arrayType.Name, dummyType.Name, err)
			}
			arrayType.ValueType = foundType
			extractedTypes.ArrayTypes[j].Value = arrayType
		}
	}
	return nil
}

func ToProperName(input string) string {
	// Split the string into fields by any non-letter or digit character
	var result strings.Builder
	capNext := true

	for _, r := range input {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			capNext = true
			continue
		}
		if capNext {
			result.WriteRune(unicode.ToUpper(r))
			capNext = false
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func parseTopLevelType(parsedSchema *ordered.OrderedValue, alreadyExtractedTypes *types.ParsedSchema) error {
	if !hasToplevelType(parsedSchema) {
		return nil
	}
	var typeName string
	if titleEntry, ok := parsedSchema.StringValueForKey("title"); ok {
		typeName = ToProperName(titleEntry)
	} else {
		currentDate := time.Now().Format("20060102")
		typeName = "UnknownTitle_" + currentDate
	}
	_, err := extractType(typeName, parsedSchema, alreadyExtractedTypes, true)
	return err
}

func parseTypesFromDefinition(definitionsMap ordered.OrderedObject, alreadyExtractedTypes *types.ParsedSchema) error {
	for _, v := range definitionsMap {
		typeName := v.Key
		_, err := extractType(typeName, v.Value, alreadyExtractedTypes, true)
		if err != nil {
			return fmt.Errorf("error while extracting type: %v", err)
		}
	}
	return nil
}

func hasToplevelType(valuesMap *ordered.OrderedValue) bool {
	if _, ok := valuesMap.ValueForKey("enum"); ok {
		// found enum entry
		return true
	} else if _, ok := valuesMap.ValueForKey("$ref"); ok {
		// found ref entry
		return true
	} else if _, ok := valuesMap.ValueForKey("type"); ok {
		// found type entry
		return true
	}
	return false
}

func extractType(name string, valuesMap *ordered.OrderedValue, alreadyExtractedTypes *types.ParsedSchema, topLevel bool) (any, error) {
	if v, ok := valuesMap.ArrayValueForKey("enum"); ok {
		// found enum entry
		return extractEnumType(name, alreadyExtractedTypes, v)
	} else if refStr, ok := valuesMap.StringValueForKey("$ref"); ok {
		return extractRefType(name, alreadyExtractedTypes, refStr)
	} else if typeStr, ok := valuesMap.StringValueForKey("type"); ok {
		return extractNormalType(name, valuesMap, alreadyExtractedTypes, topLevel, typeStr)
	}
	return types.DummyType{}, fmt.Errorf("missing type, ref or enum entry for type: %s", name)
}

func toStringArray(a []any) []string {
	ret := make([]string, 0)
	for _, v := range a {
		if s, ok := v.(string); ok {
			ret = append(ret, s)
		}
	}
	return ret
}

func toIntArray(a []any) []int {
	ret := make([]int, 0)
	for _, v := range a {
		if s, ok := v.(int); ok {
			ret = append(ret, s)
		} else if f, ok := v.(float64); ok {
			ret = append(ret, int(f))
		}
	}
	return ret
}

func extractEnumType(name string, alreadyExtractedTypes *types.ParsedSchema, enumValues any) (any, error) {
	if a, ok := enumValues.([]any); ok {
		if len(a) > 0 {
			if _, isInt := a[0].(int); isInt {
				newType := types.IntEnumType{
					Name:   name,
					Values: toIntArray(a),
				}
				if _, exist := alreadyExtractedTypes.GetIntEnumByName(name); exist {
					return types.IntEnumType{}, fmt.Errorf("int enum with name already exist: %s", name)
				}
				alreadyExtractedTypes.AddIntEnum(name, newType)
				return newType, nil
			} else if _, isStr := a[0].(string); isStr {
				newType := types.StringEnumType{
					Name:   name,
					Values: toStringArray(a),
				}
				if _, exist := alreadyExtractedTypes.GetStringEnumByName(name); exist {
					return types.StringEnumType{}, fmt.Errorf("string enum with name already exist: %s", name)
				}
				alreadyExtractedTypes.AddStringEnum(name, newType)
				return newType, nil
			} else if _, isFloat := a[0].(float64); isFloat {
				// int values are read as numbers by go ... that means float64
				newType := types.IntEnumType{
					Name:   name,
					Values: toIntArray(a),
				}
				if _, exist := alreadyExtractedTypes.GetIntEnumByName(name); exist {
					return types.StringEnumType{}, fmt.Errorf("int enum with name already exist: %s", name)
				}
				alreadyExtractedTypes.AddIntEnum(name, newType)
				return newType, nil
			} else {
				return types.StringEnumType{}, fmt.Errorf("unknown array entry for enum type with name: %s, type: %v", name, reflect.TypeOf(a[0]))
			}
		} else {
			return types.StringEnumType{}, fmt.Errorf("enum array entry has len 0 for enum type with name: %s", name)
		}
	} else {
		return types.StringEnumType{}, fmt.Errorf("no array entry for enum type with name: %s", name)
	}
}

func extractRefType(name string, alreadyExtractedTypes *types.ParsedSchema, refStr string) (any, error) {
	// supported refStr: #/definitions/MY_TYPE
	lastSlash := strings.LastIndex(refStr, "/")
	if lastSlash == -1 {
		return types.DummyType{}, fmt.Errorf("refStr has no '/', seems to have the wrong format: %s, type: %s", refStr, name)
	}
	typeName := refStr[lastSlash+1:]
	if t, exist := alreadyExtractedTypes.GetComplexTypeByName(typeName); exist {
		return t, nil
	}
	if t, exist := alreadyExtractedTypes.GetStringEnumByName(typeName); exist {
		return t, nil
	}
	if t, exist := alreadyExtractedTypes.GetIntEnumByName(typeName); exist {
		return t, nil
	}
	if t, exist := alreadyExtractedTypes.GetArrayTypeByName(typeName); exist {
		return t, nil
	}
	if t, exist := alreadyExtractedTypes.GetMapTypeByName(typeName); exist {
		return t, nil
	}
	if t, exist := alreadyExtractedTypes.GetIntegerTypeByName(typeName); exist {
		return t, nil
	}
	if t, exist := alreadyExtractedTypes.GetNumberTypeByName(typeName); exist {
		return t, nil
	}
	if t, exist := alreadyExtractedTypes.GetStringTypeByName(typeName); exist {
		return t, nil
	}
	if t, exist := alreadyExtractedTypes.GetUUIDTypeByName(typeName); exist {
		return t, nil
	}
	if t, exist := alreadyExtractedTypes.GetDateTypeByName(typeName); exist {
		return t, nil
	}
	if t, exist := alreadyExtractedTypes.GetDateTimeTypeByName(typeName); exist {
		return t, nil
	}
	if t, exist := alreadyExtractedTypes.GetTimeTypeByName(typeName); exist {
		return t, nil
	}
	if t, exist := alreadyExtractedTypes.GetDurationTypeByName(typeName); exist {
		return t, nil
	}
	if t, exist := alreadyExtractedTypes.GetBoolTypeByName(typeName); exist {
		return t, nil
	}
	if t, exist := alreadyExtractedTypes.GetBinaryTypeByName(typeName); exist {
		return t, nil
	}
	if t, exist := alreadyExtractedTypes.GetObjectTypeByName(typeName); exist {
		return t, nil
	}
	return types.DummyType{
		Name: typeName,
	}, nil
}

func extractNormalType(name string, valuesMap *ordered.OrderedValue, alreadyExtractedTypes *types.ParsedSchema,
	topLevel bool, typeStr string) (any, error) {
	switch typeStr {
	case "integer":
		return extractIntegerType(name, valuesMap, alreadyExtractedTypes, topLevel)
	case "number":
		return extractNumberType(name, valuesMap, alreadyExtractedTypes, topLevel)
	case "boolean":
		return extractBooleanType(name, valuesMap, alreadyExtractedTypes, topLevel)
	case "string":
		return extractStringType(name, valuesMap, alreadyExtractedTypes, topLevel)
	case "array":
		return extractArrayType(name, valuesMap, alreadyExtractedTypes, topLevel)
	case "object":
		return extractObjectType(name, valuesMap, alreadyExtractedTypes, topLevel)
	default:
		return types.DummyType{}, fmt.Errorf("unknown type for name: %s, type: %s", name, typeStr)
	}
}

func getOptionalString(key string, valuesMap *ordered.OrderedValue, allowed []string) o.Optional[string] {
	if s, ok := valuesMap.StringValueForKey(key); ok {
		if allowed == nil || slices.Contains(allowed, s) {
			return o.NewOptionalValue(s)
		}
	}
	return o.NewOptional[string]()
}

func getOptionalInt(key string, valuesMap *ordered.OrderedValue, allowed []int) o.Optional[int] {
	if f, ok := valuesMap.FloatValueForKey(key); ok {
		s := int(f)
		if allowed == nil || slices.Contains(allowed, s) {
			return o.NewOptionalValue(s)
		}
	}
	return o.NewOptional[int]()
}

func getOptionalNumber(key string, valuesMap *ordered.OrderedValue, allowed []float64) o.Optional[float64] {
	if f, ok := valuesMap.FloatValueForKey(key); ok {
		if allowed == nil || slices.Contains(allowed, f) {
			return o.NewOptionalValue(f)
		}
	}
	return o.NewOptional[float64]()
}

func getOptionalBool(key string, valuesMap *ordered.OrderedValue) o.Optional[bool] {
	if b, ok := valuesMap.BoolValueForKey(key); ok {
		return o.NewOptionalValue(b)
	}
	return o.NewOptional[bool]()
}

func nameIfTopLevelElseEmpty(name string, topLevel bool) string {
	if !topLevel {
		return NO_NAME
	} else {
		return name
	}
}

func extractIntegerType(name string, valuesMap *ordered.OrderedValue, alreadyExtractedTypes *types.ParsedSchema, topLevel bool) (types.IntegerType, error) {
	intType := types.IntegerType{
		Name:             o.NewOptionalConditional[string](nameIfTopLevelElseEmpty(name, topLevel), ignoreIfEmptyStr),
		Format:           getOptionalString("format", valuesMap, []string{"int32", "int64", "uint32", "uint64"}),
		Default:          getOptionalInt("default", valuesMap, nil),
		MultipleOf:       getOptionalInt("multipleOf", valuesMap, nil),
		Minimum:          getOptionalInt("minimum", valuesMap, nil),
		ExclusiveMinimum: getOptionalInt("exclusiveMinimum", valuesMap, nil),
		Maximum:          getOptionalInt("maximum", valuesMap, nil),
		ExclusiveMaximum: getOptionalInt("exclusiveMaximum", valuesMap, nil),
	}
	if topLevel && name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes.GetIntegerTypeByName(name)
		if exist {
			return intType, fmt.Errorf("can't add int type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes.AddIntegerType(name, intType)
	}
	return intType, nil
}

func extractNumberType(name string, valuesMap *ordered.OrderedValue, alreadyExtractedTypes *types.ParsedSchema, topLevel bool) (types.NumberType, error) {
	numberType := types.NumberType{
		Name:             o.NewOptionalConditional[string](nameIfTopLevelElseEmpty(name, topLevel), ignoreIfEmptyStr),
		Format:           getOptionalString("format", valuesMap, []string{"float32", "float64"}),
		Default:          getOptionalNumber("default", valuesMap, nil),
		Minimum:          getOptionalNumber("minimum", valuesMap, nil),
		ExclusiveMinimum: getOptionalNumber("exclusiveMinimum", valuesMap, nil),
		Maximum:          getOptionalNumber("maximum", valuesMap, nil),
		ExclusiveMaximum: getOptionalNumber("exclusiveMaximum", valuesMap, nil),
	}
	if topLevel && name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes.GetNumberTypeByName(name)
		if exist {
			return numberType, fmt.Errorf("can't add float type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes.AddNumberType(name, numberType)
	}
	return numberType, nil
}
func extractBooleanType(name string, valuesMap *ordered.OrderedValue, alreadyExtractedTypes *types.ParsedSchema, topLevel bool) (types.BoolType, error) {
	boolType := types.BoolType{
		Name:    o.NewOptionalConditional[string](nameIfTopLevelElseEmpty(name, topLevel), ignoreIfEmptyStr),
		Default: getOptionalBool("default", valuesMap),
	}
	if topLevel && name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes.GetBoolTypeByName(name)
		if exist {
			return boolType, fmt.Errorf("can't add bool type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes.AddBoolType(name, boolType)
	}
	return boolType, nil
}
func extractStringType(name string, valuesMap *ordered.OrderedValue, alreadyExtractedTypes *types.ParsedSchema, topLevel bool) (any, error) {
	f := getOptionalString("format", valuesMap, nil)
	if formatValue, isSet := f.Get(); isSet {
		switch formatValue {
		case "date":
			return extractDateType(name, valuesMap, alreadyExtractedTypes, topLevel)
		case "time":
			return extractTimeType(name, valuesMap, alreadyExtractedTypes, topLevel)
		case "date-time":
			return extractDateTimeType(name, valuesMap, alreadyExtractedTypes, topLevel)
		case "uuid":
			return extractUuidType(name, valuesMap, alreadyExtractedTypes, topLevel)
		case "duration":
			return extractDurationType(name, valuesMap, alreadyExtractedTypes, topLevel)
		case "binary":
			return extractBinaryType(name, valuesMap, alreadyExtractedTypes, topLevel)
		}
	}
	return extractPureStringType(name, valuesMap, alreadyExtractedTypes, topLevel, f)
}

func extractDateType(name string, valuesMap *ordered.OrderedValue, alreadyExtractedTypes *types.ParsedSchema, topLevel bool) (types.DateType, error) {
	t := types.DateType{
		Name:             o.NewOptionalConditional[string](nameIfTopLevelElseEmpty(name, topLevel), ignoreIfEmptyStr),
		Default:          getOptionalString("default", valuesMap, nil),
		Minimum:          getOptionalString("minimum", valuesMap, nil),
		ExclusiveMinimum: getOptionalString("exclusiveMinimum", valuesMap, nil),
		Maximum:          getOptionalString("maximum", valuesMap, nil),
		ExclusiveMaximum: getOptionalString("exclusiveMaximum", valuesMap, nil),
	}
	if topLevel && name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes.GetDateTypeByName(name)
		if exist {
			return t, fmt.Errorf("can't add date type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes.AddDateType(name, t)
	}
	return t, nil
}

func extractTimeType(name string, valuesMap *ordered.OrderedValue, alreadyExtractedTypes *types.ParsedSchema, topLevel bool) (types.TimeType, error) {
	t := types.TimeType{
		Name:             o.NewOptionalConditional[string](nameIfTopLevelElseEmpty(name, topLevel), ignoreIfEmptyStr),
		Default:          getOptionalString("default", valuesMap, nil),
		Minimum:          getOptionalString("minimum", valuesMap, nil),
		ExclusiveMinimum: getOptionalString("exclusiveMinimum", valuesMap, nil),
		Maximum:          getOptionalString("maximum", valuesMap, nil),
		ExclusiveMaximum: getOptionalString("exclusiveMaximum", valuesMap, nil),
	}
	if topLevel && name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes.GetTimeTypeByName(name)
		if exist {
			return t, fmt.Errorf("can't add time type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes.AddTimeType(name, t)
	}
	return t, nil
}

func extractDateTimeType(name string, valuesMap *ordered.OrderedValue, alreadyExtractedTypes *types.ParsedSchema, topLevel bool) (types.DateTimeType, error) {
	t := types.DateTimeType{
		Name:             o.NewOptionalConditional[string](nameIfTopLevelElseEmpty(name, topLevel), ignoreIfEmptyStr),
		Default:          getOptionalString("default", valuesMap, nil),
		Minimum:          getOptionalString("minimum", valuesMap, nil),
		ExclusiveMinimum: getOptionalString("exclusiveMinimum", valuesMap, nil),
		Maximum:          getOptionalString("maximum", valuesMap, nil),
		ExclusiveMaximum: getOptionalString("exclusiveMaximum", valuesMap, nil),
	}
	if topLevel && name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes.GetDateTimeTypeByName(name)
		if exist {
			return t, fmt.Errorf("can't add date-time type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes.AddDateTimeType(name, t)
	}
	return t, nil
}

func extractUuidType(name string, valuesMap *ordered.OrderedValue, alreadyExtractedTypes *types.ParsedSchema, topLevel bool) (types.UUIDType, error) {
	t := types.UUIDType{
		Name:    o.NewOptionalConditional[string](nameIfTopLevelElseEmpty(name, topLevel), ignoreIfEmptyStr),
		Default: getOptionalString("default", valuesMap, nil),
	}
	if topLevel && name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes.GetUUIDTypeByName(name)
		if exist {
			return t, fmt.Errorf("can't add uuid type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes.AddUUIDType(name, t)
	}
	return t, nil
}

func extractDurationType(name string, valuesMap *ordered.OrderedValue, alreadyExtractedTypes *types.ParsedSchema, topLevel bool) (types.DurationType, error) {
	t := types.DurationType{
		Name:    o.NewOptionalConditional[string](nameIfTopLevelElseEmpty(name, topLevel), ignoreIfEmptyStr),
		Default: getOptionalString("default", valuesMap, nil),
	}
	if topLevel && name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes.GetDurationTypeByName(name)
		if exist {
			return t, fmt.Errorf("can't add duration type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes.AddDurationType(name, t)
	}
	return t, nil
}

func extractBinaryType(name string, valuesMap *ordered.OrderedValue, alreadyExtractedTypes *types.ParsedSchema, topLevel bool) (types.BinaryType, error) {
	t := types.BinaryType{
		Name: o.NewOptionalConditional[string](nameIfTopLevelElseEmpty(name, topLevel), ignoreIfEmptyStr),
	}
	if topLevel && name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes.GetBinaryTypeByName(name)
		if exist {
			return t, fmt.Errorf("can't add binary type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes.AddBinaryType(name, t)
	}
	return t, nil
}

func extractPureStringType(name string, valuesMap *ordered.OrderedValue, alreadyExtractedTypes *types.ParsedSchema, topLevel bool, formatValue o.Optional[string]) (types.StringType, error) {
	t := types.StringType{
		Name:      o.NewOptionalConditional[string](nameIfTopLevelElseEmpty(name, topLevel), ignoreIfEmptyStr),
		Default:   getOptionalString("default", valuesMap, nil),
		Format:    formatValue,
		MinLength: getOptionalInt("minLength", valuesMap, nil),
		MaxLength: getOptionalInt("minLength", valuesMap, nil),
		Pattern:   getOptionalString("pattern", valuesMap, nil),
	}
	if topLevel && name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes.GetStringTypeByName(name)
		if exist {
			return t, fmt.Errorf("can't add string type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes.AddStringType(name, t)
	}
	return t, nil
}

func getValueType(name, key string, valuesMap *ordered.OrderedValue, alreadyExtractedTypes *types.ParsedSchema) (any, error) {
	if v, ok := valuesMap.ValueForKey(key); ok {
		return extractType(name, v, alreadyExtractedTypes, false)
	} else {
		return types.DummyType{}, fmt.Errorf("couldn't find key to extract the value type")
	}
}

func extractArrayType(name string, valuesMap *ordered.OrderedValue, alreadyExtractedTypes *types.ParsedSchema, topLevel bool) (types.ArrayType, error) {
	itemsTypeName := ToProperName(name + " Items")
	valueType, err := getValueType(itemsTypeName, "items", valuesMap, alreadyExtractedTypes)
	if err != nil {
		return types.ArrayType{}, fmt.Errorf("error while extract value type (name: %s): %v", name, err)
	}
	t := types.ArrayType{
		Name:        o.NewOptionalConditional[string](nameIfTopLevelElseEmpty(name, topLevel), ignoreIfEmptyStr),
		MinItems:    getOptionalInt("minItems", valuesMap, nil),
		MaxItems:    getOptionalInt("maxItems", valuesMap, nil),
		Description: getOptionalString("description", valuesMap, nil),
		ValueType:   valueType,
	}
	if topLevel && name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes.GetArrayTypeByName(name)
		if exist {
			return t, fmt.Errorf("can't add Array type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes.AddArrayType(name, t)
	}
	return t, nil
}

func extractProperties(parentTypeName string, propertiesMap map[string]any, alreadyExtractedTypes *types.ParsedSchema) ([]types.Property, error) {
	ret := make([]types.Property, 0)
	for key, value := range propertiesMap {
		var valuesMap *ordered.OrderedValue
		if m, isMap := value.(map[string]any); !isMap {
			return []types.Property{}, fmt.Errorf("schema value no map for complex type: %s, property: %s",
				parentTypeName, key)
		} else {
			valuesMap = m
		}
		tags := getTags(valuesMap)

		newTypeName := ToProperName(parentTypeName + " " + ToProperName(key))
		valueType, err := extractType(newTypeName, valuesMap, alreadyExtractedTypes, false)
		if err != nil {
			return []types.Property{}, fmt.Errorf("error while building property for complex type: %s, property: %s",
				parentTypeName, key)
		}
		ret = append(ret, types.Property{
			Name:         key,
			ValueType:    valueType,
			ForeignKeyTo: getOptionalString("x-ref", valuesMap, nil),
			Description:  getOptionalString("description", valuesMap, nil),
			Tags:         tags,
		})
	}
	return ret, nil
}

func extractComplexType(name string, propertiesMap map[string]any, description o.Optional[string],
	alreadyExtractedTypes *types.ParsedSchema, topLevel bool) (types.ComplexType, error) {
	properties, err := extractProperties(name, propertiesMap, alreadyExtractedTypes)
	if err != nil {
		return types.ComplexType{}, fmt.Errorf("couldn't extract properties for complex type: %s", name)
	}
	t := types.ComplexType{
		Name:        name,
		TopLevel:    topLevel,
		Description: description,
		Properties:  properties,
	}
	// only the case for toplevel types
	_, exist := alreadyExtractedTypes.GetComplexTypeByName(name)
	if exist {
		return t, fmt.Errorf("can't add Array type, because a type with the same name already exists, name: %s", name)
	}
	alreadyExtractedTypes.AddComplexType(name, t)
	return t, nil
}

func extractMapType(name string, propertiesMap map[string]any, description o.Optional[string],
	alreadyExtractedTypes *types.ParsedSchema, topLevel bool) (types.MapType, error) {
	valueType, err := extractType(name, propertiesMap, alreadyExtractedTypes, false)
	if err != nil {
		return types.MapType{}, fmt.Errorf("error while extract value type for map type (name: %s): %v", name, err)
	}
	t := types.MapType{
		Name:        name,
		Description: description,
		ValueType:   valueType,
		TopLevel:    topLevel,
	}
	if topLevel && name != "" {
		// only the case for toplevel types
		_, exist := alreadyExtractedTypes.GetMapTypeByName(name)
		if exist {
			return t, fmt.Errorf("can't add map type, because a type with the same name already exists, name: %s", name)
		}
		alreadyExtractedTypes.AddMapType(name, t)
	}
	return t, nil
}

func getTags(valuesMap *ordered.OrderedValue) []string {
	if tags, ok := valuesMap.ArrayValueForKey("x-tags"); ok {
		if tagsArray, isArray := tags.([]any); isArray {
			ret := make([]string, 0)
			for _, t := range tagsArray {
				ret = append(ret, fmt.Sprintf("%v", t))
			}
			return ret
		}
	}
	return []string{}
}

func extractObjectType(name string, valuesMap *ordered.OrderedValue, alreadyExtractedTypes *types.ParsedSchema, topLevel bool) (any, error) {
	description := getOptionalString("description", valuesMap, nil)
	if properties, ok := valuesMap.ObjectValueForKey("properties"); ok {
		if m, isMap := properties.(map[string]any); isMap {
			// found normal complex type
			a, err := extractComplexType(name, m, description, alreadyExtractedTypes, topLevel)
			a.Tags = getTags(valuesMap)
			if len(a.Tags) > 0 {
				alreadyExtractedTypes.AddComplexType(a.Name, a)
			}
			return a, err
		} else {
			return types.ObjectType{}, fmt.Errorf("properties content has not map format, type: %s", name)
		}
	} else if additionalProperties, ok := valuesMap.ObjectValueForKey("additionalProperties"); ok {
		if m, isMap := additionalProperties.(map[string]any); isMap {
			// found dictionary/map type
			return extractMapType(name, m, description, alreadyExtractedTypes, topLevel)
		} else {
			return types.ObjectType{}, fmt.Errorf("additionalProperties content has not map format, type: %s", name)
		}
	}
	return types.ObjectType{
		Name: o.NewOptionalConditional(nameIfTopLevelElseEmpty(name, topLevel), ignoreIfEmptyStr),
	}, nil
}
