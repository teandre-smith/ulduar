package ulduar

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rdsdata"
)

/*
	Deletes a record using datapi.
*/
func (da *DataApi) DeleteRecord(d *Delete) error {

	sql, err := GenerateDeleteDataSql(d.Options)
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
	Deletes an array of records using datapi.
*/
func (da *DataApi) DeleteRecords(d *DeleteBatch) error {

	for _, record := range d.Options {
		sql, err := GenerateDeleteDataSql(record)
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
	}

	return nil
}
