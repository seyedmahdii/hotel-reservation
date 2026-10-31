package db

import (
	"context"

	"github.com/seyedmahdii/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const bookingsColl = "bookings"

type BookingStore interface {
	InsertBooking(context.Context, *types.Booking) (*types.Booking, error)
}

type MongoBookingStore struct {
	client *mongo.Client
	coll   *mongo.Collection

	HotelStore
}

func NewMongoBookingStore(client *mongo.Client) *MongoBookingStore {
	return &MongoBookingStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(bookingsColl),
	}
}

func (s *MongoBookingStore) InsertBooking(ctx context.Context, booking *types.Booking) (*types.Booking, error) {
	res, err := s.coll.InsertOne(ctx, booking)
	if err != nil {
		return nil, err
	}

	booking.ID = res.InsertedID.(bson.ObjectID)
	return booking, nil
}
