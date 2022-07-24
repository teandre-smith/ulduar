package ulduar

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rdsdata"
)

/*
	Inserts a record using datapi.
*/
func (da *DataApi) InsertRecord(i *Insert) error {

	sql, err := GenerateInsertSqlWithStruct(i.Options)
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
	Inserts a record by generating dml sql template.
*/
func (da *DataApi) InsertDmlRecord(i *Insert) error {
	sql, err := GenerateDmlInsertSql(i.Options)
	if err != nil {
		return err
	}

	sqlParams, err := GenerateSqlParameters(&SqlParamOptions{
		Collection:     i.Options.Collection,
		SkipColumnList: i.Options.SkipColumnList,
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
