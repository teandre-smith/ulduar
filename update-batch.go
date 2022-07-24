package ulduar

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rdsdata"
	"github.com/aws/aws-sdk-go-v2/service/rdsdata/types"
)

/*
	This function batch updates the provided records.

	Each UpdateBatch.Options.Collection's structure can be different as the dml template is generated for each individual record.

	If updating data with the same dml sql template, please use
	BatchUpdateWithSameDmlTemplate.
*/
func (da *DataApi) BatchUpdateRecords(u *UpdateBatch) error {

	sqlParams := [][]types.SqlParameter{}
	batchArray := []rdsdata.BatchExecuteStatementInput{}

	if u.BatchAmount == nil || *u.BatchAmount <= 0 {
		defAmount := 100
		u.BatchAmount = &defAmount
	}

	for index, record := range u.Options {
		sql, err := GenerateDmlUpdateSql(record)
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

		if len(sqlParams)%*u.BatchAmount == 0 || index+1 == len(u.Options) {

			batchArray = append(batchArray, rdsdata.BatchExecuteStatementInput{
				ResourceArn:   da.ResourceArn,
				SecretArn:     da.SecretArn,
				Sql:           &sql,
				Database:      da.DbName,
				ParameterSets: sqlParams,
			})

			sqlParams = [][]types.SqlParameter{}
		}
	}

	for _, batchRecord := range batchArray {
		_, err := da.RDSClient.BatchExecuteStatement(context.TODO(), &batchRecord)
		if err != nil {
			return err
		}
	}

	return nil
}

/*
	This function batch updates the provided records.

	Each UpdateBatch.Options.Collection's structure is assumed to be the exact same.

	If updating data with the varying structure, please use
	BatchUpdateRecords.
*/
func (da *DataApi) BatchUpdateWithSameDmlTemplate(u *UpdateBatch) error {
	sqlParams := [][]types.SqlParameter{}
	batchArray := []rdsdata.BatchExecuteStatementInput{}

	sql, err := GenerateDmlUpdateSql(u.Options[0])
	if err != nil {
		return err
	}

	if u.BatchAmount == nil || *u.BatchAmount <= 0 {
		defAmount := 100
		u.BatchAmount = &defAmount
	}

	for index, record := range u.Options {

		sqlParam, err := GenerateSqlParameters(&SqlParamOptions{
			Collection:     record.Collection,
			SkipColumnList: record.SkipColumnList,
		})
		if err != nil {
			return err
		}

		sqlParams = append(sqlParams, sqlParam)

		if len(sqlParams)%*u.BatchAmount == 0 || index+1 == len(u.Options) {

			batchArray = append(batchArray, rdsdata.BatchExecuteStatementInput{
				ResourceArn:   da.ResourceArn,
				SecretArn:     da.SecretArn,
				Sql:           &sql,
				Database:      da.DbName,
				ParameterSets: sqlParams,
			})

			sqlParams = [][]types.SqlParameter{}
		}
	}

	for _, batchRecord := range batchArray {
		_, err := da.RDSClient.BatchExecuteStatement(context.TODO(), &batchRecord)
		if err != nil {
			return err
		}
	}

	return nil
}
