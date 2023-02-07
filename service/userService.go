package service

import "toplearn-api/model/user"

type UserService interface {
	GetUserList() ([]user.User, error)
}

type userService struct {
}

func NewUserService() UserService {
	return userService{}
}

func (userService) GetUserList() ([]user.User, error) {
	return []user.User{
		{
			FirstName:   "Mohammad",
			LastName:    "Javadi",
			Age:         22,
			PhoneNumber: "091212345667",
		},
		{
			FirstName:   "Kazem",
			LastName:    "jodi",
			Age:         51,
			PhoneNumber: "0911454545987",
		},
	}, nil

}
