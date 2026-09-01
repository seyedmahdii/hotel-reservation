package db

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const hotelsColl = "hotels"

type MongoHotelStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client, dbname string) *MongoHotelStore {
	return &MongoHotelStore{
		client: client,
		coll:   client.Database(dbname).Collection(hotelsColl),
	}
}
