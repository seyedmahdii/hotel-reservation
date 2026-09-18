package types

import "go.mongodb.org/mongo-driver/v2/bson"

type Hotel struct {
	ID       bson.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string          `bson:"name" json:"name"`
	Location string          `bson:"location" json:"location"`
	Rooms    []bson.ObjectID `bson:"rooms" json:"room"`
	Rating   int             `bson:"rating" json:"rating"`
}

type RoomType int

const (
	_ RoomType = iota
	SingleRoomType
	DoubleRoomType
	SeaSideRoomType
	DeluxeRoomType
)

type Room struct {
	ID bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	// small, normal, kingsize
	Size    string        `bson:"size" json:"size"`
	SeaSide bool          `bson:"seaSide" json:"seaSide"`
	Price   float64       `bson:"price" json:"price"`
	HotelID bson.ObjectID `bson:"hotelID" json:"hotelID"`
}
