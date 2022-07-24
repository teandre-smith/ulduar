package ulduar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateUpsertSqlWithStruct(t *testing.T) {
	assert := assert.New(t)

	jst, _ := time.LoadLocation("Asia/Tokyo")

	collection := TestStruct{
		SomeString: "something",
		SomeInt:    0,
		SomeTime:   time.Date(2016, 1, 1, 0, 0, 0, 0, jst),
		SomeFloat:  0,
	}

	target := "someTime"

	param := &UpsertOptions{
		Collection:     collection,
		Target:         &target,
		SkipColumnList: []string{},
		NoUpdateList:   []string{},
		TableName:      new(string),
	}

	sql, err := GenerateUpsertSqlWithStruct(param)

	expectedResults := "INSERT INTO test_struct ( someString, someInt, someTime, someFloat ) VALUES ( 'something', 0, '2016-01-01T00:00:00+09:00'::timestamp, 0.000000 ) ON CONFLICT ( someTime ) DO UPDATE SET someString = 'something', someInt = 0, someFloat = 0.000000;"

	assert.Equal(expectedResults, sql)
	assert.NoError(err)
}
