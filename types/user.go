package types

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost = 12
)

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID                bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // primitive.ObjectId
	FirstName         string        `bson:"firstName" json:"firstName"`
	LastName          string        `bson:"lastName" json:"lastName"`
	Email             string        `bson:"email" json:"email"`
	EncryptedPassword string        `bson:"encryptedPassword" json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encryptedPassword),
	}, nil
}
