package controller

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"toplearn-api/Utility"
	"toplearn-api/model/user"
	"toplearn-api/service"
	"toplearn-api/viewModel/common/security"
	"toplearn-api/viewModel/userVm"
)

func GetUserAvatar(c echo.Context) error {
	idString := c.Param("id")

	return c.String(http.StatusOK, "/users/get/[with param ("+idString+")]/avatars/")
}

// the structure of incoming fields

func GetListOfUser(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	fmt.Println(apiContext.GetUserId())

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
	apiContext := c.(*Utility.ApiContext)

	newUser := new(userVm.CreateNewUserViewModel)

	if err := c.Bind(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	if err := c.Validate(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	creator, err := apiContext.GetUserId()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	newUser.CreatorUserName = creator

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

func LoginUser(c echo.Context) error {
	loginModel := new(userVm.LoginUserViewModel)

	if err := c.Bind(loginModel); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	if err := c.Validate(loginModel); err != nil {
		return c.JSON(http.StatusBadRequest, "Model not Valid")
	}

	userService := service.NewUserService()
	user, err := userService.GetUserByUserNameAndPassword(*loginModel)

	if err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	claims := &security.JwtClaims{
		user.UserName,
		user.Id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString([]byte("our-secret-key-in-golang-project"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	loginResData := struct {
		Token string
	}{
		Token: stringToken,
	}

	return c.JSON(http.StatusOK, loginResData)
}
