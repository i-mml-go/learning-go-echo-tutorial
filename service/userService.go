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

	return userList, err

}
