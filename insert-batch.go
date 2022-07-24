package ulduar

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rdsdata"
	"github.com/aws/aws-sdk-go-v2/service/rdsdata/types"
)

/*
	This function batch inserts the provided records.

	Each InsertBatch.Options.Collection's structure can be different as the dml template is generated for each individual record.

	If inserting data with the same dml sql template, please use
	BatchInsertWithSameDmlTemplate.
*/
func (da *DataApi) BatchInsertRecords(i *InsertBatch) error {

	sqlParams := [][]types.SqlParameter{}
	batchInsertArray := []rdsdata.BatchExecuteStatementInput{}

	if i.BatchAmount == nil {
		defAmount := 100
		i.BatchAmount = &defAmount
	}

	for index, record := range i.Options {
		sql, err := GenerateDmlInsertSql(record)
		if err != nil {
			return err
		}

		sqlParam, err := GenerateSqlParameters(&SqlParamOptions{
			Collection:     record.Collection,
			SkipColumnList: record.SkipColumnList,
		})
		if err != nil {
			return err
		}

		sqlParams = append(sqlParams, sqlParam)

		if len(sqlParams)%*i.BatchAmount == 0 || index+1 == len(i.Options) {

			batchInsertArray = append(batchInsertArray, rdsdata.BatchExecuteStatementInput{
				ResourceArn:   da.ResourceArn,
				SecretArn:     da.SecretArn,
				Sql:           &sql,
				Database:      da.DbName,
				ParameterSets: sqlParams,
			})

			sqlParams = [][]types.SqlParameter{}
		}
	}

	for _, batchRecord := range batchInsertArray {
		_, err := da.RDSClient.BatchExecuteStatement(context.TODO(), &batchRecord)
		if err != nil {
			return err
		}
	}

	return nil
}

/*
	This function batch inserts the provided records.

	Each InsertBatch.Options.Collection's structure is assumed to be the exact same.

	If inserting data with the varying structure, please use
	BatchInsertRecords.
*/
func (da *DataApi) BatchInsertWithSameDmlTemplate(i *InsertBatch) error {
	sqlParams := [][]types.SqlParameter{}
	batchInsertArray := []rdsdata.BatchExecuteStatementInput{}

	sql, err := GenerateDmlInsertSql(i.Options[0])
	if err != nil {
		return err
	}

	if i.BatchAmount == nil {
		defAmount := 100
		i.BatchAmount = &defAmount
	}

	for index, record := range i.Options {

		sqlParam, err := GenerateSqlParameters(&SqlParamOptions{
			Collection:     record.Collection,
			SkipColumnList: record.SkipColumnList,
		})
		if err != nil {
			return err
		}

		sqlParams = append(sqlParams, sqlParam)

		if len(sqlParams)%*i.BatchAmount == 0 || index+1 == len(i.Options) {

			batchInsertArray = append(batchInsertArray, rdsdata.BatchExecuteStatementInput{
				ResourceArn:   da.ResourceArn,
				SecretArn:     da.SecretArn,
				Sql:           &sql,
				Database:      da.DbName,
				ParameterSets: sqlParams,
			})

			sqlParams = [][]types.SqlParameter{}
		}
	}

	for _, batchRecord := range batchInsertArray {
		_, err := da.RDSClient.BatchExecuteStatement(context.TODO(), &batchRecord)
		if err != nil {
			return err
		}
	}

	return nil
}
