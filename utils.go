package ulduar

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/rdsdata/types"
)

func Contains[T comparable](arrayValue []T, value T) bool {
	for _, item := range arrayValue {
		if item == value {
			return true
		}
	}

	return false
}

func getTableName(providedTable *string, structName string) string {
	if providedTable == nil || *providedTable == "" {
		return snakeCaseName(structName)
	}

	return *providedTable
}

func snakeCaseName(name string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(name, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func checkType(value any, item string) (string, error) {
	switch {
	case strings.Contains(item, "int"):
		return fmt.Sprintf("%d", value), nil

	case strings.Contains(item, "float"):
		return fmt.Sprintf("%f", value), nil

	case item == "bool":
		return fmt.Sprintf("%t", value), nil

	case item == "string":
		return fmt.Sprintf("'%s'", value), nil

	case strings.Contains(item, "time"):
		var t time.Time = value.(time.Time)
		return fmt.Sprintf("'%s'::timestamp", t.Format(time.RFC3339)), nil
	}

	return "", fmt.Errorf("error: %s did not match any of the types", item)
}

/*
	DataAPI sends back data as types.Field and those items will be strings in
	most cases so to accomodate that you will have to Check and assert type.
	Example of this is below:
		someFloat := CheckQueryFieldType(datapoint).(string)
		num, _ := strconv.ParseFloat(someFloat, 64)
*/
func CheckQueryFieldType(item types.Field) any {
	var value any
	// type switches can be used to check the union value
	switch v := item.(type) {
	case *types.FieldMemberArrayValue:
		value = v.Value // Value is types.ArrayValue

	case *types.FieldMemberBlobValue:
		value = v.Value // Value is []byte

	case *types.FieldMemberBooleanValue:
		value = v.Value // Value is bool

	case *types.FieldMemberDoubleValue:
		value = v.Value // Value is float64

	case *types.FieldMemberIsNull:
		value = v.Value // Value is bool

	case *types.FieldMemberLongValue:
		value = v.Value // Value is int64

	case *types.FieldMemberStringValue:
		value = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}

	return value
}
