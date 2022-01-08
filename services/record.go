package services

import (
	"context"
	"fmt"
	"net/http"
	"time"

	CONSTANTS "github.com/alicobanserver/constants"
	"github.com/alicobanserver/log"
	"github.com/alicobanserver/models"
	"github.com/alicobanserver/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func FetcData(data []bson.M, recordResp models.RecordResponse, record models.RecordRequest) (resp interface{}, err error) {

	recordResp.Code = http.StatusBadRequest
	recordResp.Records = data

	startDate, err := time.Parse("2006-01-02", record.StartDate)
	if err != nil {
		recordResp.Msg = err.Error()
		log.Error(err)
		return recordResp, err
	}

	endDate, err := time.Parse("2006-01-02", record.EndDate)
	if err != nil {
		recordResp.Msg = err.Error()
		log.Error(err)
		return recordResp, err
	}

	filter := []bson.M{
		{
			"$match": bson.M{
				"createdAt": bson.M{
					"$gt": startDate,
					"$lt": endDate,
				},
			},
		},
		{
			"$project": bson.M{
				"_id":        0,
				"key":        1,
				"createdAt":  1,
				"totalCount": bson.M{"$sum": "$counts"},
			},
		},
		{
			"$match": bson.M{
				"totalCount": bson.M{
					"$gt": record.MinCount,
					"$lt": record.MaxCount,
				},
			},
		},
	}

	cursor, err := utils.Recordcollection.Aggregate(context.TODO(), filter)
	if err != nil {
		recordResp.Msg = err.Error()
		log.Error(err)
		return recordResp, err
	}

	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &data); err != nil {
		recordResp.Msg = err.Error()
		log.Error(err)
		return recordResp, err
	}

	if len(data) > 0 {
		recordResp.Code = 0
		recordResp.Msg = CONSTANTS.SUCCESS
		recordResp.Records = data
		return recordResp, nil
	}

	recordResp.Code = http.StatusNoContent
	recordResp.Msg = CONSTANTS.NO_DATA_FOUND
	recordResp.Records = data
	err = fmt.Errorf(CONSTANTS.NO_DATA_FOUND)
	return recordResp, err
}
