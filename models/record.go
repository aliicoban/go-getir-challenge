package models

import "go.mongodb.org/mongo-driver/bson"

type RecordRequest struct {
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
	MinCount  float64 `json:"minCount"`
	MaxCount  float64 `json:"maxCount"`
}

type RecordResponse struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Records []bson.M `json:"records"`
}
