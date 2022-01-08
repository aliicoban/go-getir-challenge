package utils

import (
	"errors"
	"time"

	CONSTANTS "github.com/alicobanserver/constants"
	"github.com/alicobanserver/models"
)

/*
 - Check Record Request Parameters
*/

func CheckParams(record models.RecordRequest) error {

	_, err := time.Parse("2006-01-02", record.StartDate)
	_, err2 := time.Parse("2006-01-02", record.EndDate)

	if err != nil {
		return errors.New(CONSTANTS.START_DATE_ERR)
	} else if err2 != nil {
		return errors.New(CONSTANTS.END_DATE_ERR)
	} else if record.MaxCount < 0 {
		return errors.New(CONSTANTS.MAX_COUNT_ERR)
	} else if record.MinCount < 0 {
		return errors.New(CONSTANTS.MIN_COUNT_ERR)
	} else if record.MinCount > record.MaxCount {
		return errors.New(CONSTANTS.MIN_MAX_ERR)
	}
	return nil
}
