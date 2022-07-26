package ulduar

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rdsdata"
)

/*
	Executes provided sql statement.
*/
func (da DataApi) CustomQuery(sql string) error {
	params := &rdsdata.ExecuteStatementInput{
		ResourceArn: da.ResourceArn,
		SecretArn:   da.SecretArn,
		Sql:         &sql,
		Database:    da.DbName,
	}

	_, err := da.RDSClient.ExecuteStatement(context.TODO(), params)
	if err != nil {
		return err
	}

	return nil
}
