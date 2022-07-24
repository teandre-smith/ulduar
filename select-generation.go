package ulduar

import (
	"fmt"
	"reflect"
	"strings"
)

/*
	Generates Select Sql statement.
*/
func GenerateSelectSql(options *SelectOptions) (string, error) {
	t := reflect.TypeOf(options.Collection)
	if t == nil {
		return "", fmt.Errorf("error: SelectOptions.Collection cannot be nil")
	}

	table := getTableName(options.TableName, t.Name())

	columns := ""

	if len(options.SkipColumnList) == 0 {
		columns = "*"
	} else {
		columnNames := []string{}

		for i := 0; i < t.NumField(); i++ {
			columnName := t.Field(i).Tag.Get("datapi")
			if columnName == "" {
				return "", fmt.Errorf("error: struct value %s does not possess `datapi` tags", t.Field(i).Name)
			} else if Contains(options.SkipColumnList, columnName) {
				continue
			}

			columnNames = append(columnNames, columnName)
		}

		columns = strings.Join(columnNames, ", ")
	}

	statement := []string{}
	statement = append(statement, fmt.Sprintf("SELECT %s FROM %s", columns, table))

	if options.Filter != nil {
		statement = append(statement, fmt.Sprintf("WHERE %s", *options.Filter))
	}

	if options.SortByColumn != nil {
		direction := "DESC"
		if options.SortDirection != nil && *options.SortDirection == "ASC" {
			direction = *options.SortDirection
		}
		statement = append(statement, fmt.Sprintf("ORDER BY %s %s", *options.SortByColumn, direction))
	}

	if options.Limit != nil {
		statement = append(statement, fmt.Sprintf("LIMIT %d", *options.Limit))
	}

	return strings.Join(statement, " "), nil

}
