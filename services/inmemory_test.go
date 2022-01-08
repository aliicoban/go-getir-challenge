package services

import (
	"github.com/alicobanserver/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

var Db = models.Database{}
var St = Store{Db: Db}

func TestSetData(t *testing.T) {
	key := "active-tabs"
	value := "getir"
	err := St.Set(key, value)
	if err != nil {
		assert.NotNil(t, err)
		return
	}
	assert.EqualValues(t, Db[key], value)
}

func TestGetData(t *testing.T) {
	key := "active-tabs"
	value, err := St.Get(key)
	if err != nil {
		assert.NotNil(t, err)
		return
	}
	assert.NotNil(t, value)
}
