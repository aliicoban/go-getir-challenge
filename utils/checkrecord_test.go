package utils

import (
	//"fmt"
	"testing"

	CONSTANTS "github.com/alicobanserver/constants"
	"github.com/alicobanserver/models"
	"github.com/stretchr/testify/assert"
)

func TestStartDate(t *testing.T) {
	var sDTest = models.RecordRequest{
		StartDate: "2016",
		EndDate:   "2018-02-02",
		MinCount:  2700,
		MaxCount:  3000,
	}
	err := CheckParams(sDTest)
	if err != nil {
		assert.EqualValues(t, err.Error(), CONSTANTS.START_DATE_ERR)
	}
}

func TestEndDate(t *testing.T) {
	var sDTest = models.RecordRequest{
		StartDate: "2016-01-26",
		EndDate:   "2012-10-01T09:45:00.000+02:00",
		MinCount:  2700,
		MaxCount:  3000,
	}
	err := CheckParams(sDTest)
	if err != nil {
		assert.EqualValues(t, err.Error(), CONSTANTS.END_DATE_ERR)
	}
}

func TestIsMinCountNegative(t *testing.T) {
	var sDTest = models.RecordRequest{
		StartDate: "2016-01-26",
		EndDate:   "2018-02-02",
		MinCount:  -1,
		MaxCount:  3000,
	}
	err := CheckParams(sDTest)
	if err != nil {
		assert.EqualValues(t, err.Error(), CONSTANTS.MIN_COUNT_ERR)
	}
}

func TestIsMaxCountNegative(t *testing.T) {
	var sDTest = models.RecordRequest{
		StartDate: "2016-01-26",
		EndDate:   "2018-02-02",
		MinCount:  2700,
		MaxCount:  -1,
	}
	err := CheckParams(sDTest)
	if err != nil {
		assert.EqualValues(t, err.Error(), CONSTANTS.MAX_COUNT_ERR)
	}
}

func TestCompareMinMax(t *testing.T) {
	var sDTest = models.RecordRequest{
		StartDate: "2016-01-26",
		EndDate:   "2018-02-02",
		MinCount:  4700,
		MaxCount:  3000,
	}
	err := CheckParams(sDTest)
	if err != nil {
		assert.EqualValues(t, err.Error(), CONSTANTS.MIN_MAX_ERR)
	}
}
func TestRecordReqSucess(t *testing.T) {
	var sDTest = models.RecordRequest{
		StartDate: "2016-01-26",
		EndDate:   "2018-02-02",
		MinCount:  2700,
		MaxCount:  3000,
	}
	err := CheckParams(sDTest)
	if err != nil {
		assert.NotNil(t, err)
	}
	assert.Nil(t, err)
}
