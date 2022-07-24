package ulduar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUpsertRecord(t *testing.T) {
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

	target := "someString"

	param := &Upsert{
		Options: &UpsertOptions{
			Collection:     testItem,
			Target:         &target,
			SkipColumnList: []string{},
			NoUpdateList:   []string{},
			TableName:      new(string),
		},
	}

	err := api.UpsertRecord(param)

	assert.NoError(err)
}

func TestUpsertDmlRecord(t *testing.T) {
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

	target := "someString"

	param := &Upsert{
		Options: &UpsertOptions{
			Collection:     testItem,
			Target:         &target,
			SkipColumnList: []string{},
			NoUpdateList:   []string{},
			TableName:      new(string),
		},
	}

	err := api.UpsertDmlRecord(param)

	assert.NoError(err)
}
