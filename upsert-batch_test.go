package ulduar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBatchUpsertRecords(t *testing.T) {
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

	target := "someString"

	param := &UpsertBatch{
		Options: []*UpsertOptions{{
			Collection:     collection,
			Target:         &target,
			SkipColumnList: []string{},
			NoUpdateList:   []string{},
			TableName:      new(string),
		}},
		BatchAmount: new(int),
	}

	err := api.BatchUpsertRecords(param)

	assert.NoError(err)
}

func TestBatchUpsertWithSameDmlTemplate(t *testing.T) {
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

	target := "someString"

	param := &UpsertBatch{
		Options: []*UpsertOptions{{
			Collection:     collection,
			Target:         &target,
			SkipColumnList: []string{},
			NoUpdateList:   []string{},
			TableName:      new(string),
		}},
		BatchAmount: new(int),
	}

	err := api.BatchUpsertWithSameDmlTemplate(param)

	assert.NoError(err)
}
