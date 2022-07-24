package ulduar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteRecord(t *testing.T) {
	assert := assert.New(t)

	mockClient := &MockClient{}

	api := DataApi{
		RDSClient:   mockClient,
		ResourceArn: new(string),
		SecretArn:   new(string),
		DbName:      new(string),
	}

	table := "test"
	condition := "id = 1"

	param := &Delete{
		Options: &DeleteOptions{
			TableName: &table,
			Condition: &condition,
		},
	}

	err := api.DeleteRecord(param)

	assert.NoError(err)
}

func TestDeleteRecords(t *testing.T) {
	assert := assert.New(t)

	mockClient := &MockClient{}

	api := DataApi{
		RDSClient:   mockClient,
		ResourceArn: new(string),
		SecretArn:   new(string),
		DbName:      new(string),
	}

	table := "test"
	condition := "id = 1"

	param := &DeleteBatch{
		Options:     []*DeleteOptions{{TableName: &table, Condition: &condition}},
		BatchAmount: new(int),
	}

	err := api.DeleteRecords(param)

	assert.NoError(err)
}
