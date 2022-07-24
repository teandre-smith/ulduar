package ulduar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInsertRecord(t *testing.T) {
	assert := assert.New(t)

	mockClient := &MockClient{}

	api := DataApi{
		RDSClient:   mockClient,
		ResourceArn: new(string),
		SecretArn:   new(string),
		DbName:      new(string),
	}

	testItem := TestStruct{
		SomeString: "",
		SomeInt:    0,
		SomeTime:   time.Time{},
		SomeFloat:  0,
	}

	param := &Insert{
		Options: &InsertOptions{
			Collection:     testItem,
			SkipColumnList: []string{},
			TableName:      new(string),
		},
	}

	err := api.InsertRecord(param)

	assert.NoError(err)
}

func TestInsertDmlRecord(t *testing.T) {
	assert := assert.New(t)

	mockClient := &MockClient{}

	api := DataApi{
		RDSClient:   mockClient,
		ResourceArn: new(string),
		SecretArn:   new(string),
		DbName:      new(string),
	}

	testItem := TestStruct{
		SomeString: "",
		SomeInt:    0,
		SomeTime:   time.Time{},
		SomeFloat:  0,
	}

	param := &Insert{
		Options: &InsertOptions{
			Collection:     testItem,
			SkipColumnList: []string{},
			TableName:      new(string),
		},
	}

	err := api.InsertDmlRecord(param)

	assert.NoError(err)
}
