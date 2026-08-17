package types

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string        `bson:"firstName" json:"firstName"`
	LastName  string        `bson:"lastName" json:"lastName"`
}
