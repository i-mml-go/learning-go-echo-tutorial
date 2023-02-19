package service

import (
	"time"
	"toplearn-api/model/user"
	"toplearn-api/repository"
	"toplearn-api/viewModel/userVm"
)

type UserService interface {
	GetUserList() ([]user.User, error)
	CreateNewUser(userInput userVm.CreateNewUserViewModel) (string, error)
	GetUserByUserNameAndPassword(loginViewModel userVm.LoginUserViewModel) (user.User, error)
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

func (userService) GetUserByUserNameAndPassword(loginViewModel userVm.LoginUserViewModel) (user.User, error) {
	userRepository := repository.NewUserRepository()
	user, err := userRepository.GetUserByUserNameAndPassword(loginViewModel.UserName, loginViewModel.Password)

	return user, err
}

func (userService) CreateNewUser(userInput userVm.CreateNewUserViewModel) (string, error) {

	userEntity := user.User{
		FirstName:     userInput.FirstName,
		LastName:      userInput.LastName,
		Email:         userInput.Email,
		UserName:      userInput.UserName,
		Password:      userInput.Password,
		RegisterData:  time.Now(),
		CreatorUserId: userInput.CreatorUserName,
	}

	userRepository := repository.NewUserRepository()
	userId, err := userRepository.InsertUser(userEntity)

	return userId, err
}
