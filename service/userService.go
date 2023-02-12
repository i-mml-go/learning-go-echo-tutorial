package service

import (
	"toplearn-api/model/user"
	"toplearn-api/repository"
)

type UserService interface {
	GetUserList() ([]user.User, error)
}

type userService struct {
}

func NewUserService() UserService {
	return userService{}
}

func (userService) GetUserList() ([]user.User, error) {
	userRepository := repository.NewUserRepository()
	userList, err := userRepository.GetUserList()
	id, err := userRepository.InsertUser(user.User{
		Id:          "",
		FirstName:   "kazem",
		LastName:    "ghiyasi",
		Age:         44,
		PhoneNumber: "09121234152",
	})
	println(id)

	//user, err := userRepository.GetUserById("63df8a55329f4660805f4133")
	//fmt.Println("this is new user log", user)

	return userList, err

}
