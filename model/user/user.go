package user

type User struct {
	FirstName   string `query:"name" json:"firstName"`
	LastName    string `query:"family" json:"lastName"`
	Age         int    `query:"age"`
	PhoneNumber string `query:"phone" json:"mobile"`
}
