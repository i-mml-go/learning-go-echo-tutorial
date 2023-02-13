package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"toplearn-api/model/user"
	"toplearn-api/service"
	"toplearn-api/viewModel/userVm"
)

func GetUserAvatar(c echo.Context) error {
	idString := c.Param("id")

	return c.String(http.StatusOK, "/users/get/[with param ("+idString+")]/avatars/")
}

// the structure of incoming fields

func GetListOfUser(c echo.Context) error {
	userService := service.NewUserService()

	userList, errList := userService.GetUserList()
	if errList != nil {
		println(errList)
	}

	userInput := new(user.User)

	err := c.Bind(userInput)

	if err != nil {
		return err
	}

	fmt.Println(userInput)

	return c.JSON(http.StatusOK, userList)
}

func CreateNewUser(c echo.Context) error {
	newUser := new(userVm.CreateNewUserViewModel)

	if err := c.Bind(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	if err := c.Validate(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userService := service.NewUserService()
	newUserId, err := userService.CreateNewUser(*newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		NewUserId string
	}{
		NewUserId: newUserId,
	}

	return c.JSON(http.StatusOK, userResData)
}
