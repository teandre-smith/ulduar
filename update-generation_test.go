package ulduar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateUpdateSqlWithStruct(t *testing.T) {
	assert := assert.New(t)

	jst, _ := time.LoadLocation("Asia/Tokyo")

	collection := TestStruct{
		SomeString: "something",
		SomeInt:    0,
		SomeTime:   time.Date(2016, 1, 1, 0, 0, 0, 0, jst),
		SomeFloat:  0,
	}

	condition := "id = 1"

	param := &UpdateOptions{
		Collection:     collection,
		Condition:      &condition,
		SkipColumnList: []string{},
		TableName:      new(string),
	}

	sql, err := GenerateUpdateSqlWithStruct(param)

	expectedResults := "UPDATE test_struct SET someString = 'something', someInt = 0, someTime = '2016-01-01T00:00:00+09:00'::timestamp, someFloat = 0.000000 WHERE id = 1;"

	assert.Equal(len([]byte(expectedResults)), len([]byte(sql)))
	assert.NoError(err)
}
