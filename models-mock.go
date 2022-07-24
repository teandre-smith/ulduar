package ulduar

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/rdsdata"
)

type MockClient struct{}

func (m MockClient) ExecuteStatement(ctx context.Context, params *rdsdata.ExecuteStatementInput, optFns ...func(*rdsdata.Options)) (*rdsdata.ExecuteStatementOutput, error) {
	return &rdsdata.ExecuteStatementOutput{}, nil
}

func (m MockClient) BatchExecuteStatement(ctx context.Context, params *rdsdata.BatchExecuteStatementInput, optFns ...func(*rdsdata.Options)) (*rdsdata.BatchExecuteStatementOutput, error) {
	return &rdsdata.BatchExecuteStatementOutput{}, nil
}

type TestStruct struct {
	SomeString string    `datapi:"someString"`
	SomeInt    int64     `datapi:"someInt"`
	SomeTime   time.Time `datapi:"someTime"`
	SomeFloat  float64   `datapi:"someFloat"`
}
