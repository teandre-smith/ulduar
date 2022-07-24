package ulduar

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rdsdata"
)

/*
	Selects a record using datapi.
*/
func (da *DataApi) QueryRecord(s *Select) (*rdsdata.ExecuteStatementOutput, error) {

	sql, err := GenerateSelectSql(s.Options)
	if err != nil {
		return nil, err
	}

	params := &rdsdata.ExecuteStatementInput{
		ResourceArn: da.ResourceArn,
		SecretArn:   da.SecretArn,
		Sql:         &sql,
		Database:    da.DbName,
	}

	record, err := da.RDSClient.ExecuteStatement(context.TODO(), params)
	if err != nil {
		return nil, err
	}

	return record, nil
}
