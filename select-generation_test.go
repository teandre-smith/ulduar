package ulduar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSelectSql(t *testing.T) {
	assert := assert.New(t)

	jst, _ := time.LoadLocation("Asia/Tokyo")

	collection := TestStruct{
		SomeString: "something",
		SomeInt:    0,
		SomeTime:   time.Date(2016, 1, 1, 0, 0, 0, 0, jst),
		SomeFloat:  0,
	}

	sortByColumn := "someTime"
	sortDirection := "ASC"
	filter := "area = chubu"
	limit := 10

	param := &SelectOptions{
		Collection:     collection,
		SkipColumnList: []string{},
		TableName:      new(string),
		SortByColumn:   &sortByColumn,
		SortDirection:  &sortDirection,
		Filter:         &filter,
		Limit:          &limit,
	}

	sql, err := GenerateSelectSql(param)

	expectedResults := "SELECT * FROM test_struct WHERE area = chubu ORDER BY someTime ASC LIMIT 10"

	assert.Equal(expectedResults, sql)
	assert.NoError(err)
}
