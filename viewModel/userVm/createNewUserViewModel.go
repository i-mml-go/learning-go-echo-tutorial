package userVm

type CreateNewUserViewModel struct {
	FirstName       string
	LastName        string `validate:"required"`
	Email           string `validate:"required"`
	UserName        string `validate:"required"`
	Password        string `validate:"required"`
	CreatorUserName string
}
