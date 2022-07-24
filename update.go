package ulduar

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rdsdata"
)

/*
	Update a record using datapi.
*/
func (da *DataApi) UpdateRecord(u *Update) error {

	sql, err := GenerateUpdateSqlWithStruct(u.Options)
	if err != nil {
		return err
	}

	params := &rdsdata.ExecuteStatementInput{
		ResourceArn: da.ResourceArn,
		SecretArn:   da.SecretArn,
		Sql:         &sql,
		Database:    da.DbName,
	}

	_, err = da.RDSClient.ExecuteStatement(context.TODO(), params)
	if err != nil {
		return err
	}

	return nil
}

/*
	Updates a record using generated dml sql template.
*/
func (da *DataApi) UpdateDmlRecord(u *Update) error {
	sql, err := GenerateDmlUpdateSql(u.Options)
	if err != nil {
		return err
	}

	sqlParams, err := GenerateSqlParameters(&SqlParamOptions{
		Collection:     u.Options.Collection,
		SkipColumnList: u.Options.SkipColumnList,
	})
	if err != nil {
		return err
	}

	params := &rdsdata.ExecuteStatementInput{
		ResourceArn: da.ResourceArn,
		SecretArn:   da.SecretArn,
		Sql:         &sql,
		Database:    da.DbName,
		Parameters:  sqlParams,
	}

	_, err = da.RDSClient.ExecuteStatement(context.TODO(), params)
	if err != nil {
		return err
	}

	return nil
}
