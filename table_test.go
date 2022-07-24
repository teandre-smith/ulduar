package ulduar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTable(t *testing.T) {
	assert := assert.New(t)

	mockClient := &MockClient{}

	api := DataApi{
		RDSClient:   mockClient,
		ResourceArn: new(string),
		SecretArn:   new(string),
		DbName:      new(string),
	}

	testStruct := TestStruct{}

	param := &Table{
		Options: &TableOptions{
			Collection: testStruct,
			TableName:  new(string),
		},
	}

	err := api.CreateTable(param)

	assert.NoError(err)
}

func TestDropTable(t *testing.T) {
	assert := assert.New(t)

	mockClient := &MockClient{}

	api := DataApi{
		RDSClient:   mockClient,
		ResourceArn: new(string),
		SecretArn:   new(string),
		DbName:      new(string),
	}

	testStruct := TestStruct{}

	param := &Table{
		Options: &TableOptions{
			Collection: testStruct,
			TableName:  new(string),
		},
	}

	err := api.DropTable(param)

	assert.NoError(err)
}

func TestDropTables(t *testing.T) {
	assert := assert.New(t)

	mockClient := &MockClient{}

	api := DataApi{
		RDSClient:   mockClient,
		ResourceArn: new(string),
		SecretArn:   new(string),
		DbName:      new(string),
	}

	err := api.DropTables([]string{"test1", "test2"})

	assert.NoError(err)
}
