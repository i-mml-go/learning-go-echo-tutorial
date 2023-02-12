package user

type User struct {
	Id          string `bson:"_id"`
	FirstName   string
	LastName    string
	Age         int
	PhoneNumber string

	//PhoneNumber string `query:"phone" json:"mobile"`
}
