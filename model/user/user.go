package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id          primitive.ObjectID
	FirstName   string
	LastName    string
	Age         int
	PhoneNumber string

	//PhoneNumber string `query:"phone" json:"mobile"`
}
