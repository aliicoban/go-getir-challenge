package controllers

import (
	"encoding/json"
	"net/http"

	CONSTANTS "github.com/alicobanserver/constants"
	"github.com/alicobanserver/models"
	"github.com/alicobanserver/services"
)

var db = services.DB(models.Database{})

/*
 - Set Endpoint
 - Set Key and Value
*/
func Set(w http.ResponseWriter, r *http.Request) {

	var memory = models.InMemory{}

	err := json.NewDecoder(r.Body).Decode(&memory)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Message{CONSTANTS.SET_DATA_ERR})
		return
	}

	key := memory.Key
	value := memory.Value

	err = db.Set(key, value)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Message{CONSTANTS.SET_DATA_ERR})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(memory)
}

/*
 - Get Endpoint(GET)
 - Get Value by key
*/
func Get(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Message{CONSTANTS.MISS_URL_QUERY})
		return
	}

	key := keys[0]
	getValue, err := db.Get(key)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Message{CONSTANTS.GET_DATA_ERR})
		return
	}

	if getValue == "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.Message{CONSTANTS.NO_DATA_FOUND})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getValue)

}
