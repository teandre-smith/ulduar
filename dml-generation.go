package ulduar

import (
	"fmt"
	"reflect"
	"strings"
)

/*
	Generates insert dml sql statement for BatchExecuteStatementInput.SQL || ExecuteStatementInput.SQL using struct values.
		InsertOptions.Collection fields must possess `datapi` struct tags.
*/
func GenerateDmlInsertSql(options *InsertOptions) (string, error) {
	t := reflect.TypeOf(options.Collection) // Check if not nil
	if t == nil {
		return "", fmt.Errorf("error: InsertOptions.Collection cannot be nil")
	}

	table := getTableName(options.TableName, t.Name())

	var columnNames []string
	var dmlNames []string

	v := reflect.ValueOf(options.Collection)

	for i := 0; i < t.NumField(); i++ {
		columnName := t.Field(i).Tag.Get("datapi") // Get struct tag name 'datapi' value as that is the column name
		if columnName == "" {
			return "", fmt.Errorf("error: struct field must possess `datapi` struct tags")
		} else if Contains(options.SkipColumnList, columnName) {
			continue
		}

		fieldType := v.Field(i).Type().String()

		data := checkTypeDml(columnName, fieldType)

		columnNames = append(columnNames, columnName)
		dmlNames = append(dmlNames, data)
	}

	fields := strings.Join(columnNames, ", ")
	dmlValues := strings.Join(dmlNames, ", ")

	statement := fmt.Sprintf(`INSERT INTO %s ( %s ) VALUES ( %s )`, table, fields, dmlValues)

	return statement, nil
}

/*
	Generates upsert dml sql statement for BatchExecuteStatementInput.SQL || ExecuteStatementInput.SQL using struct values.
		UpsertOptions.Collection fields must possess `datapi` struct tags.
*/
func GenerateDmlUpsertSql(options *UpsertOptions) (string, error) {
	t := reflect.TypeOf(options.Collection) // Check if not nil
	if t == nil {
		return "", fmt.Errorf("error: collection parameter cannot be nil")
	}

	if options.Target == nil {
		return "", fmt.Errorf("error: UpsertOptions.Target cannot be nil")
	}

	table := getTableName(options.TableName, t.Name())

	var columnNames []string
	var dmlNames []string

	v := reflect.ValueOf(options.Collection)

	for i := 0; i < t.NumField(); i++ {
		columnName := t.Field(i).Tag.Get("datapi") // Get struct tag name 'datapi' value as that is the column name
		if columnName == "" {
			return "", fmt.Errorf("error: struct field must possess `datapi` struct tag")
		} else if Contains(options.SkipColumnList, columnName) {
			continue
		}

		fieldType := v.Field(i).Type().String()

		data := checkTypeDml(columnName, fieldType)

		columnNames = append(columnNames, columnName)
		dmlNames = append(dmlNames, data)
	}

	updateList := []string{}
	for i := 0; i < len(columnNames); i++ {
		if columnNames[i] == *options.Target || Contains(options.NoUpdateList, columnNames[i]) {
			continue
		}

		columnAndValue := fmt.Sprintf("%s = %s", columnNames[i], dmlNames[i])
		updateList = append(updateList, columnAndValue)
	}

	columns := strings.Join(columnNames, ", ")
	values := strings.Join(dmlNames, ", ")
	updates := strings.Join(updateList, ", ")

	statement := fmt.Sprintf("INSERT INTO %s ( %s ) VALUES ( %s ) ON CONFLICT ( %s ) DO UPDATE SET %s", table, columns, values, *options.Target, updates)

	return statement, nil
}

/*
	Generates update dml sql statement for BatchExecuteStatementInput.SQL || ExecuteStatementInput.SQL using struct values.
		UpdateOptions.Collection fields must possess `datapi` struct tags.
*/
func GenerateDmlUpdateSql(options *UpdateOptions) (string, error) {

	t := reflect.TypeOf(options.Collection) // Check if not nil
	if t == nil {
		return "", fmt.Errorf("error: UpdateOptions.Collection cannot be nil")
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

		data := checkTypeDml(columnName, fieldType)

		columnMap[columnName] = data
	}

	updateList := []string{}

	for columnName, columnValue := range columnMap {
		columnAndValue := fmt.Sprintf("%s = %s", columnName, columnValue)
		updateList = append(updateList, columnAndValue)
	}

	updates := strings.Join(updateList, ", ")
	statement := fmt.Sprintf("UPDATE %s SET %s WHERE %s", table, updates, *options.Condition)

	return statement, nil
}

// Dependent Funcs below

// Checks type and explicitly casts any time.Time types
func checkTypeDml(value string, item string) string {
	switch {
	case strings.Contains(item, "time"):
		return fmt.Sprintf(":%s::timestamp", value)

	default:
		return fmt.Sprintf(":%s", value)
	}
}
