package types

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Booking struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    bson.ObjectID `bson:"userID,omitempty" json:"userID,omitempty"`
	RoomID    bson.ObjectID `bson:"roomID,omitempty" json:"roomID,omitempty"`
	NumPeople int           `bson:"numPeople,omitempty" json:"numPeople,omitempty"`
	FromDate  time.Time     `bson:"fromDate,omitempty" json:"fromDate,omitempty"`
	TillDate  time.Time     `bson:"tillDate,omitempty" json:"tillDate,omitempty"`
}
