package user

type User struct {
	Id          string `bson:"_id,omitempty"`
	FirstName   string `bson:"firstName,omitempty"`
	LastName    string `bson:"lastName,omitempty"`
	Age         int    `bson:"age,omitempty"`
	PhoneNumber string `bson:"phoneNumber,omitempty"`

	//PhoneNumber string `query:"phone" json:"mobile"`
}
