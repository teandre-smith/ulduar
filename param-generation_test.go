package ulduar

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/rdsdata/types"
	"github.com/stretchr/testify/assert"
)

func TestGenerateSqlParameters(t *testing.T) {
	assert := assert.New(t)

	jst, _ := time.LoadLocation("Asia/Tokyo")

	testCollection := TestStruct{
		SomeString: "hi",
		SomeInt:    0,
		SomeTime:   time.Date(2016, 1, 1, 0, 0, 0, 0, jst),
		SomeFloat:  0,
	}

	params := &SqlParamOptions{
		Collection:     testCollection,
		SkipColumnList: []string{},
	}

	sqlParams, err := GenerateSqlParameters(params)

	names := []string{"someString", "someInt", "someTime", "someFloat"}

	expectedResults := []types.SqlParameter{
		{
			Name: &names[0],
			Value: &types.FieldMemberStringValue{
				Value: "hi",
			},
		},
		{
			Name: &names[1],
			Value: &types.FieldMemberLongValue{
				Value: 0,
			},
		},
		{
			Name: &names[2],
			Value: &types.FieldMemberStringValue{
				Value: "2016-01-01T00:00:00+09:00",
			},
		},
		{
			Name: &names[3],
			Value: &types.FieldMemberDoubleValue{
				Value: 0,
			},
		},
	}

	assert.Equal(expectedResults, sqlParams)
	assert.NoError(err)

}
