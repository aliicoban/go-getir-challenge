package services

import (
	"encoding/json"
	"errors"
	"github.com/alicobanserver/models"
	"io/ioutil"
	"log"
	"os"
)

type Store struct {
	Db models.Database
}

func DB(db models.Database) Store {
	store := Store{Db: db}
	ReadFileFromDB(&db)
	return store
}

/*
 - Set Service Method.
 - Save key and value to Store
*/
func (store Store) Set(key string, value string) error {

	if key == "" || value == "" {
		return errors.New("")
	}
	store.Db[key] = value
	SaveToDB(store)
	return nil
}

/*
 - Get Service Method.
 - Get value by key from the Store
*/
func (store Store) Get(key string) (interface{}, error) {
	if key == "" {
		return "", errors.New("")
	}
	if store.Db[key] == "" {
		return nil, errors.New("")
	}
	jsonData := make(map[string]string)
	jsonData["key"] = key
	jsonData["value"] = store.Db[key]
	return jsonData, nil
}

/*
 - Save value by key to Store
*/
func SaveToDB(store Store) {
	dbJson, _ := json.Marshal(store.Db)
	ioutil.WriteFile("temp/data.json", dbJson, os.ModeAppend)
}

/*
 - ReadFileFromDB Method
 - ReadFile from "emp/data.json"
*/
func ReadFileFromDB(db *models.Database) {
	bytes, err := ioutil.ReadFile("temp/data.json")
	if err != nil {
		log.Panic("Err:", err.Error())
	}
	json.Unmarshal(bytes, &db)
}
