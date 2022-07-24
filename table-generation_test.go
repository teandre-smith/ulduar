package ulduar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateCreateTableSqlWithStruct(t *testing.T) {
	assert := assert.New(t)

	jst, _ := time.LoadLocation("Asia/Tokyo")

	collection := TestStruct{
		SomeString: "something",
		SomeInt:    0,
		SomeTime:   time.Date(2016, 1, 1, 0, 0, 0, 0, jst),
		SomeFloat:  0,
	}

	param := &TableOptions{
		Collection: collection,
		TableName:  new(string),
	}

	sql, err := GenerateCreateTableSqlWithStruct(param)

	expectedResults := "CREATE TABLE IF NOT EXISTS test_struct ( id SERIAL UNIQUE, someString text, someInt bigint, someTime timestamp with time zone, someFloat numeric, CONSTRAINT test_struct_pkey PRIMARY KEY (id) );"

	assert.Equal(len([]byte(expectedResults)), len([]byte(sql)))
	assert.NoError(err)
}

func TestGenerateDropTableSql(t *testing.T) {
	assert := assert.New(t)

	data1, _ := GenerateDropTableSql([]string{"test"}...)
	date2, err := GenerateDropTableSql([]string{"test1, test2"}...)

	expectedResults1 := "DROP TABLE test"
	expectedResults2 := "DROP TABLE test1, test2"

	assert.Equal(expectedResults1, data1)
	assert.Equal(expectedResults2, date2)
	assert.NoError(err)
}
