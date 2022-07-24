package ulduar

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rdsdata"
)

/*
	Interface for *rdsdata.Client. Makes for easy mocking and testing
	of aws datapi client.
*/
type DataApiInterface interface {
	ExecuteStatement(ctx context.Context, params *rdsdata.ExecuteStatementInput, optFns ...func(*rdsdata.Options)) (*rdsdata.ExecuteStatementOutput, error)
	BatchExecuteStatement(ctx context.Context, params *rdsdata.BatchExecuteStatementInput, optFns ...func(*rdsdata.Options)) (*rdsdata.BatchExecuteStatementOutput, error)
}

type DataApi struct {
	RDSClient   DataApiInterface
	ResourceArn *string
	SecretArn   *string
	DbName      *string
}

type Insert struct {
	Options *InsertOptions
}

type InsertBatch struct {
	Options []*InsertOptions

	// The amount of records in each batch request
	//
	// 	@Default - 100
	BatchAmount *int
}

type Upsert struct {
	Options *UpsertOptions
}

type UpsertBatch struct {
	Options []*UpsertOptions

	// The amount of records in each batch request
	//
	// 	@Default - 100
	BatchAmount *int
}

type Update struct {
	Options *UpdateOptions
}

type UpdateBatch struct {
	Options []*UpdateOptions

	// The amount of records in each batch request
	//
	// 	@Default - 100
	BatchAmount *int
}

type Delete struct {
	Options *DeleteOptions
}

type DeleteBatch struct {
	Options []*DeleteOptions

	// The amount of records in each batch request
	//
	// 	@Default - 100
	BatchAmount *int
}

type Select struct {
	Options *SelectOptions
}

type Table struct {
	Options *TableOptions
}

type InsertOptions struct {

	// Struct of data to generate query from.
	//
	// This field is required.
	Collection interface{}

	// Columns that are not included in INSERT INTO x (...)
	// VALUES (...) portion of statement generation.
	//
	// This field is optional.
	SkipColumnList []string

	// Name of the table.
	// If not provided, table name will be generated based off provided struct
	//
	// This field is optional
	TableName *string
}

type UpsertOptions struct {

	// Struct of data to generate query from.
	//
	// This field is required.
	Collection interface{}

	// Target column for ON CONFLICT (x) DO UPDATE...
	//
	// This field is required.
	Target *string

	// Columns that are not included in INSERT INTO x (...)
	// VALUES (...) portion of statement generation.
	//
	// This field is optional.
	SkipColumnList []string

	// Columns that are not included in DO UPDATE SET...
	// portion of statement generation
	//
	// This field is optional
	NoUpdateList []string

	// Name of the table.
	// If not provided, table name will be generated based off provided struct
	//
	// This field is optional
	TableName *string
}

type UpdateOptions struct {

	// Struct of data to generate query from.
	//
	// This field is required.
	Collection interface{}

	// Condition to match for WHERE clause
	//
	// This field is required
	Condition *string

	// Columns that are not included in UPDATE SET (...)
	// VALUES (...) portion of statement generation.
	//
	// This field is optional.
	SkipColumnList []string

	// Name of the table.
	// If not provided, table name will be generated based off provided struct
	//
	// This field is optional
	TableName *string
}

type DeleteOptions struct {
	// Name of the table.
	// If not provided, table name will be generated based off provided struct
	//
	// This field is required
	TableName *string

	// Condition to match for WHERE clause
	//
	// This field is required
	Condition *string
}

type SqlParamOptions struct {
	// Struct of data to generate query from.
	//
	// This field is required.
	Collection interface{}

	// Columns that are not included in UPDATE SET (...)
	// VALUES (...) portion of statement generation.
	//
	// This field is optional.
	SkipColumnList []string
}

type TableOptions struct {
	// Struct of data to generate query from.
	//
	// This field is required.
	Collection interface{}

	// Name of the table.
	//
	// This field is required
	TableName *string
}

type SelectOptions struct {
	// Struct of data to generate query from.
	//
	// This field is required.
	Collection interface{}

	// Columns that are not included in SELECT x... FROM y
	// portion of statement generation.
	//
	// This field is optional.
	SkipColumnList []string

	// Name of the table.
	// If not provided, table name will be generated based off provided struct
	//
	// This field is required
	TableName *string

	// Name of column to sort by
	//
	// This field is optional
	SortByColumn *string

	// Sort direction. Can either be ASC or DESC
	//
	// This field is optional
	//
	// @Default - DESC
	SortDirection *string

	// Filter for WHERE CLAUSE
	//
	// This field is optional
	//
	Filter *string

	// Max number of records returned
	//
	// This field is optional
	Limit *int
}
