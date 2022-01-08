package controllers

import (
	"encoding/json"
	"net/http"

	LOG "github.com/alicobanserver/log"
	CONSTANTS "github.com/alicobanserver/constants"

	"github.com/alicobanserver/models"
	"github.com/alicobanserver/services"
	"github.com/alicobanserver/utils"
	"go.mongodb.org/mongo-driver/bson"
)

/*
 - Fetch "records" table
 - startDate, endDate, minCount, maxCount parameters 
*/

func Fetch(w http.ResponseWriter, r *http.Request) {

	var data []bson.M
	var recordResp models.RecordResponse
	var record models.RecordRequest

	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		LOG.Error(err)
		return
	}

	isReqValid := utils.CheckParams(record)
	if isReqValid != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.Message{isReqValid.Error()})
		return
	}

	resp, err := services.FetcData(data, recordResp, record)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.Message{CONSTANTS.NO_DATA_FOUND})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
