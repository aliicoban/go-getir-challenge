package utils

import "go.mongodb.org/mongo-driver/mongo"

var Recordcollection *mongo.Collection


func RecordCollection(c *mongo.Database) {
	Recordcollection = c.Collection("records")
}
