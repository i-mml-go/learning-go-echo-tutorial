package controller

import (
	"fmt"
	"net/http"
	"toplearn-api/model/user"
	"toplearn-api/service"

	"github.com/labstack/echo/v4"
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

func CreateUser(c echo.Context) error {
	userInput := new(user.User)

	err := c.Bind(userInput)
	if err != nil {
		return err
	}
	fmt.Println(userInput)

	return c.String(http.StatusOK, "what the fuck are you talking about")
}
