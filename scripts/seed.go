package main

import (
	"context"
	"log"

	"github.com/seyedmahdii/hotel-reservation/db"
	"github.com/seyedmahdii/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	client     *mongo.Client
	hotelStore *db.MongoHotelStore
	roomStore  *db.MongoRoomStore
	ctx        = context.Background()
)

func seedHotel(name, location string, rating int) {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []bson.ObjectID{},
		Rating:   rating,
	}
	rooms := []types.Room{
		{
			Size:  "normal",
			Price: 99.9,
		},
		{
			Size:  "small",
			Price: 199.9,
		},
		{
			Size:  "kingsize",
			Price: 133.9,
		},
	}

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		_, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	seedHotel("King Mahdi Hotel", "Iran", 5)
	seedHotel("King MP Hotel", "France", 4)
	seedHotel("King Ramos Hotel", "Spain", 3)
}

func init() {
	var err error
	client, err = mongo.Connect(options.Client().
		ApplyURI(db.DBURI))
	if err = client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
}
