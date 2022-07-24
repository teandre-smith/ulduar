package ulduar

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateDeleteDataSql(t *testing.T) {
	assert := assert.New(t)

	table := "test"
	condition := "id = 1"

	params := &DeleteOptions{
		TableName: &table,
		Condition: &condition,
	}
	sql, err := GenerateDeleteDataSql(params)

	expectedResults := fmt.Sprintf("DELETE FROM %s WHERE %s", table, condition)

	assert.Equal(expectedResults, sql)
	assert.NoError(err)
}

func TestGenerateDeleteDataSqlWithEmptyParams(t *testing.T) {
	assert := assert.New(t)

	params := &DeleteOptions{}
	sql, err := GenerateDeleteDataSql(params)

	expectedResults := ""

	assert.Equal(expectedResults, sql)
	assert.Error(err)
}
