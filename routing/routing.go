package routing

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"toplearn-api/controller"
)

func GroupLevel(next echo.HandlerFunc) echo.HandlerFunc {
	// this middleware is for group routing
	fmt.Println("i am groupLevel")
	return next
}

func RouteLevel(next echo.HandlerFunc) echo.HandlerFunc {
	// this middleware is for group routing
	fmt.Println("i am routeLevel")
	return next
}
func RouteLevel2(next echo.HandlerFunc) echo.HandlerFunc {
	// this middleware is for group routing
	fmt.Println("i am routeLevel 2")
	return next
}

func SetRouting(e *echo.Echo) error {

	g := e.Group("users", GroupLevel)

	g.GET("/get", controller.GetListOfUser, RouteLevel, RouteLevel2)
	g.POST("/createNewUser", controller.CreateNewUser)
	g.GET("/get/:id/avatars", controller.GetUserAvatar)

	return nil
}
