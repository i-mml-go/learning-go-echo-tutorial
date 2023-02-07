package routing

import (
	"github.com/labstack/echo/v4"
	"toplearn-api/controller"
)

func SetRouting(e *echo.Echo) error {

	g := e.Group("users")

	//// add Name to the route
	//g.GET("/get/list", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "this is the list of our users")
	//}).Name = "getUsersList"

	// use a external function , instead inline function
	g.GET("/get/:id/avatars", controller.GetUserAvatar)

	g.GET("/get", controller.GetListOfUser)

	g.POST("/create", controller.CreateUser)

	return nil
}
