package ulduar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateDmlInsertSQL(t *testing.T) {
	assert := assert.New(t)

	testStruct := TestStruct{
		SomeString: "",
		SomeInt:    0,
		SomeTime:   time.Time{},
		SomeFloat:  0,
	}

	tableName := "test"

	options := &InsertOptions{
		Collection:     testStruct,
		SkipColumnList: []string{},
		TableName:      &tableName,
	}

	sql, err := GenerateDmlInsertSql(options)

	expectedResults := `INSERT INTO test ( someString, someInt, someTime, someFloat ) VALUES ( :someString, :someInt, :someTime::timestamp, :someFloat )`

	assert.Equal(expectedResults, sql)
	assert.NoError(err)
}

func TestGenerateDmlUpsertSQL(t *testing.T) {
	assert := assert.New(t)

	testStruct := TestStruct{
		SomeString: "",
		SomeInt:    0,
		SomeTime:   time.Time{},
		SomeFloat:  0,
	}

	target := "someString"

	options := &UpsertOptions{
		Collection:     testStruct,
		Target:         &target,
		SkipColumnList: []string{},
		NoUpdateList:   []string{},
		TableName:      new(string),
	}

	sql, err := GenerateDmlUpsertSql(options)

	expectedResults := `INSERT INTO test_struct ( someString, someInt, someTime, someFloat ) VALUES ( :someString, :someInt, :someTime::timestamp, :someFloat ) ON CONFLICT ( someString ) DO UPDATE SET someInt = :someInt, someTime = :someTime::timestamp, someFloat = :someFloat`

	assert.Equal(expectedResults, sql)
	assert.NoError(err)
}

func TestGenerateDmlUpdateSQL(t *testing.T) {
	assert := assert.New(t)

	testStruct := TestStruct{
		SomeString: "",
		SomeInt:    0,
		SomeTime:   time.Time{},
		SomeFloat:  0,
	}

	condition := "someString = 'chubu'"
	tableName := "test"

	options := &UpdateOptions{
		Collection:     testStruct,
		Condition:      &condition,
		SkipColumnList: []string{"someString"},
		TableName:      &tableName,
	}

	sql, err := GenerateDmlUpdateSql(options)

	expectedResults := "UPDATE test SET someInt = :someInt, someTime = :someTime::timestamp, someFloat = :someFloat WHERE someString = 'chubu'"

	assert.Equal(expectedResults, sql)
	assert.NoError(err)
}
