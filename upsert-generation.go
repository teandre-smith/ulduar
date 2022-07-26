package ulduar

import (
	"fmt"
	"reflect"
	"strings"
)

/*
	Generates Upsert Sql statement based off struct. UpsertOptions.Collection must have `datapi` struct tags.
	noUpdateList is a list of column names that are not included in the update clause. tableName parameter is
	optional, but if not provided, table name will be generated by snake casing struct name.
*/
func GenerateUpsertSqlWithStruct(options *UpsertOptions) (string, error) {
	t := reflect.TypeOf(options.Collection) // Check if not nil
	if t == nil {
		return "", fmt.Errorf("error: UpsertOptions.Collection parameter cannot be nil")
	}

	if options.Target == nil {
		return "", fmt.Errorf("error: UpsertOptions.Target cannot be nil")
	}

	table := getTableName(options.TableName, t.Name())

	v := reflect.ValueOf(options.Collection)

	columnNames := []string{}
	columnValues := []string{}

	for i := 0; i < t.NumField(); i++ {
		columnName := t.Field(i).Tag.Get("datapi") // Get struct tag name 'datapi' value as that is the column name
		if columnName == "" {
			return "", fmt.Errorf("error: struct field must have `datapi` struct tags")
		} else if Contains(options.SkipColumnList, columnName) {
			continue
		}

		fieldType := v.Field(i).Type().String()
		fieldValue := v.Field(i).Interface()

		data, err := checkType(fieldValue, fieldType)
		if err != nil {
			return "", err
		}

		columnNames = append(columnNames, columnName)
		columnValues = append(columnValues, data)
	}

	updateList := []string{}
	for i := 0; i < len(columnNames); i++ {
		if columnNames[i] == *options.Target || Contains(options.NoUpdateList, columnNames[i]) {
			continue
		}

		columnAndValue := fmt.Sprintf("%s = %s", columnNames[i], columnValues[i])
		updateList = append(updateList, columnAndValue)
	}

	columns := strings.Join(columnNames, ", ")
	values := strings.Join(columnValues, ", ")
	updates := strings.Join(updateList, ", ")

	statement := fmt.Sprintf("INSERT INTO %s ( %s ) VALUES ( %s ) ON CONFLICT ( %s ) DO UPDATE SET %s;", table, columns, values, *options.Target, updates)

	return statement, nil
}
