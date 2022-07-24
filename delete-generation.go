package ulduar

import "fmt"

func GenerateDeleteDataSql(options *DeleteOptions) (string, error) {
	if options.TableName == nil || *options.TableName == "" {
		return "", fmt.Errorf("tableName cannot be nil or an empty string")
	}

	if options.Condition == nil || *options.Condition == "" {
		return "", fmt.Errorf("condition cannot be nil or an empty string")
	}

	statement := fmt.Sprintf("DELETE FROM %s WHERE %s", *options.TableName, *options.Condition)

	return statement, nil
}
