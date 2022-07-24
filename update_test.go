package ulduar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUpdateRecord(t *testing.T) {
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

	condition := "id = 1"

	param := &Update{
		Options: &UpdateOptions{
			Collection:     testItem,
			Condition:      &condition,
			SkipColumnList: []string{},
			TableName:      new(string),
		},
	}

	err := api.UpdateRecord(param)

	assert.NoError(err)
}

func TestUpdateDmlRecord(t *testing.T) {
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

	condition := "id = 1"

	param := &Update{
		Options: &UpdateOptions{
			Collection:     testItem,
			Condition:      &condition,
			SkipColumnList: []string{},
			TableName:      new(string),
		},
	}

	err := api.UpdateDmlRecord(param)

	assert.NoError(err)
}
