package ulduar

import (
	"fmt"
	"reflect"
	"strings"
)

/*
	Generates Sql statement by using provided UpdateOptions.Collection.

	Struct must have `datapi` struct tags so column names can be identified.
*/
func GenerateUpdateSqlWithStruct(options *UpdateOptions) (string, error) {
	t := reflect.TypeOf(options.Collection) // Check if not nil
	if t == nil {
		return "", fmt.Errorf("error: UpdateOptions.Collection cannot be nil")
	}

	if options.Condition == nil || *options.Condition == "" {
		return "", fmt.Errorf("error: UpdateOptions.Condition cannot be nil or empty")
	}

	table := getTableName(options.TableName, t.Name())

	v := reflect.ValueOf(options.Collection)

	columnMap := map[string]string{}

	for i := 0; i < t.NumField(); i++ {
		columnName := t.Field(i).Tag.Get("datapi") // Get struct tag name 'datapi' value as that is the column name
		if columnName == "" {
			return "", fmt.Errorf("error: struct values must have `datapi` struct tag")
		} else if Contains(options.SkipColumnList, columnName) {
			continue
		}

		fieldType := v.Field(i).Type().String()
		fieldValue := v.Field(i).Interface()

		data, err := checkType(fieldValue, fieldType)
		if err != nil {
			return "", err
		}

		columnMap[columnName] = data
	}

	updateList := []string{}

	for columnName, columnValue := range columnMap {
		columnAndValue := fmt.Sprintf("%s = %s", columnName, columnValue)
		updateList = append(updateList, columnAndValue)
	}

	updates := strings.Join(updateList, ", ")
	statement := fmt.Sprintf("UPDATE %s SET %s WHERE %s;", table, updates, *options.Condition)

	return statement, nil
}
