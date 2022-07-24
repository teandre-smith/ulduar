package ulduar

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/rdsdata"
)

/*
	Upserts a record using datapi.
*/
func (da *DataApi) UpsertRecord(u *Upsert) error {

	sql, err := GenerateUpsertSqlWithStruct(u.Options)
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
	Upserts a record using generated dml sql template.
*/
func (da *DataApi) UpsertDmlRecord(u *Upsert) error {

	if u.Options.Target == nil {
		return fmt.Errorf("target is a required struct field")
	}

	sql, err := GenerateDmlUpsertSql(u.Options)
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
