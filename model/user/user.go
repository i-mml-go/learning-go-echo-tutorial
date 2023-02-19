package user

import "time"

type User struct {
	Id            string    `bson:"_id,omitempty"`
	FirstName     string    `bson:"firstName,omitempty"`
	LastName      string    `bson:"lastName,omitempty"`
	Email         string    `bson:"email,omitempty"`
	UserName      string    `bson:"userName,omitempty"`
	Password      string    `bson:"password,omitempty"`
	RegisterData  time.Time `bson:"registerDate,omitempty"`
	CreatorUserId string    `bson:"CreatorUserId,omitempty"`

	//PhoneNumber string `query:"phone" json:"mobile"`
}
