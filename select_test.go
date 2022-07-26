package ulduar

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/rdsdata"
	"github.com/stretchr/testify/assert"
)

func TestQueryRecord(t *testing.T) {
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

	param := &Select{
		Options: &SelectOptions{
			Collection:     testItem,
			SkipColumnList: []string{},
			TableName:      new(string),
			SortByColumn:   new(string),
			SortDirection:  new(string),
			Filter:         new(string),
			Limit:          new(int),
		},
	}

	output, err := api.QueryRecord(param)

	expectedResult := &rdsdata.ExecuteStatementOutput{}

	assert.Equal(expectedResult, output)
	assert.NoError(err)
}

func TestQueryRecordJson(t *testing.T) {
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

	param := &Select{
		Options: &SelectOptions{
			Collection:     testItem,
			SkipColumnList: []string{},
			TableName:      new(string),
			SortByColumn:   new(string),
			SortDirection:  new(string),
			Filter:         new(string),
			Limit:          new(int),
		},
	}

	output, err := api.QueryRecordJson(param)

	expectedResult := &rdsdata.ExecuteStatementOutput{}

	assert.Equal(expectedResult, output)
	assert.NoError(err)
}
