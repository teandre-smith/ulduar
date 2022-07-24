package ulduar

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/rdsdata/types"
)

/* Generates types.SqlParameters for BatchExecuteStatementInput.ParameterSets || ExecuteStatement.Input.Parameters */
func GenerateSqlParameters(options *SqlParamOptions) ([]types.SqlParameter, error) {
	t := reflect.TypeOf(options.Collection) // Check if not nil
	if t == nil {
		return []types.SqlParameter{}, fmt.Errorf("error: SqlParamOptions.Collection cannot be nil")
	}
	params := []types.SqlParameter{}

	v := reflect.ValueOf(options.Collection)

	for i := 0; i < v.NumField(); i++ {
		columnName := t.Field(i).Tag.Get("datapi") // Get struct tag name 'datapi' value as that is the column name
		if columnName == "" {
			return []types.SqlParameter{}, fmt.Errorf("error: struct field must possess `datapi` struct tags")
		} else if Contains(options.SkipColumnList, columnName) {
			continue
		}

		fieldType := v.Field(i).Type().String()
		fieldValue := v.Field(i).Interface()

		sqlParamItem := generateSqlParameter(fieldValue, fieldType, columnName)
		params = append(params, sqlParamItem)
	}

	return params, nil
}

/*
	Helper function for GenerateSqlParameters that returns types.SqlParameters
*/
func generateSqlParameter(value any, valueType string, columnName string) types.SqlParameter {
	param := types.SqlParameter{
		Name: &columnName,
	}

	switch {
	case strings.Contains(valueType, "time"):
		newTime := value.(time.Time)
		param.Value = &types.FieldMemberStringValue{Value: newTime.Format(time.RFC3339)}

	case valueType == "string":
		newString := value.(string)
		param.Value = &types.FieldMemberStringValue{Value: newString}

	case valueType == "bool":
		newBool := value.(bool)
		param.Value = &types.FieldMemberBooleanValue{Value: newBool}

	case strings.Contains(valueType, "int"):
		newInt := value.(int64)
		param.Value = &types.FieldMemberLongValue{Value: newInt}

	case strings.Contains(valueType, "float"):
		newFloat := value.(float64)
		param.Value = &types.FieldMemberDoubleValue{Value: newFloat}
	}

	return param
}
