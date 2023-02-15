package routing

import (
	"github.com/labstack/echo/v4"
	"toplearn-api/controller"
)

func SetRouting(e *echo.Echo) error {
	e.POST("/login", controller.LoginUser)

	g := e.Group("users")

	g.GET("/get", controller.GetListOfUser)
	g.POST("/createNewUser", controller.CreateNewUser)
	g.GET("/get/:id/avatars", controller.GetUserAvatar)

	return nil
}
