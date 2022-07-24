package ulduar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBatchInsertRecords(t *testing.T) {
	assert := assert.New(t)

	mockClient := &MockClient{}

	api := DataApi{
		RDSClient:   mockClient,
		ResourceArn: new(string),
		SecretArn:   new(string),
		DbName:      new(string),
	}

	table := "test"
	batchAmount := 200

	collection := TestStruct{
		SomeString: "",
		SomeInt:    0,
		SomeTime:   time.Time{},
		SomeFloat:  0,
	}

	param := &InsertBatch{
		Options: []*InsertOptions{{
			Collection:     collection,
			SkipColumnList: []string{},
			TableName:      &table,
		}},
		BatchAmount: &batchAmount,
	}

	err := api.BatchInsertRecords(param)

	assert.NoError(err)
}

func TestBatchInsertWithSameDmlTemplate(t *testing.T) {
	assert := assert.New(t)

	mockClient := &MockClient{}

	api := DataApi{
		RDSClient:   mockClient,
		ResourceArn: new(string),
		SecretArn:   new(string),
		DbName:      new(string),
	}

	table := "test"
	batchAmount := 200

	collection := TestStruct{
		SomeString: "",
		SomeInt:    0,
		SomeTime:   time.Time{},
		SomeFloat:  0,
	}

	param := &InsertBatch{
		Options: []*InsertOptions{{
			Collection:     collection,
			SkipColumnList: []string{},
			TableName:      &table,
		}},
		BatchAmount: &batchAmount,
	}

	err := api.BatchInsertWithSameDmlTemplate(param)

	assert.NoError(err)
}
