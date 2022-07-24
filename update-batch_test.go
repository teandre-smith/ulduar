package ulduar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBatchUpdateRecords(t *testing.T) {
	assert := assert.New(t)

	mockClient := &MockClient{}

	api := DataApi{
		RDSClient:   mockClient,
		ResourceArn: new(string),
		SecretArn:   new(string),
		DbName:      new(string),
	}

	collection := TestStruct{
		SomeString: "",
		SomeInt:    0,
		SomeTime:   time.Time{},
		SomeFloat:  0,
	}

	param := &UpdateBatch{
		Options: []*UpdateOptions{{
			Collection:     collection,
			Condition:      new(string),
			SkipColumnList: []string{},
			TableName:      new(string),
		}},
		BatchAmount: new(int),
	}

	err := api.BatchUpdateRecords(param)

	assert.NoError(err)
}

func TestBatchUpdateWithSameDmlTemplate(t *testing.T) {
	assert := assert.New(t)

	mockClient := &MockClient{}

	api := DataApi{
		RDSClient:   mockClient,
		ResourceArn: new(string),
		SecretArn:   new(string),
		DbName:      new(string),
	}

	collection := TestStruct{
		SomeString: "",
		SomeInt:    0,
		SomeTime:   time.Time{},
		SomeFloat:  0,
	}

	param := &UpdateBatch{
		Options: []*UpdateOptions{{
			Collection:     collection,
			Condition:      new(string),
			SkipColumnList: []string{},
			TableName:      new(string),
		}},
	}

	err := api.BatchUpdateWithSameDmlTemplate(param)

	assert.NoError(err)
}
