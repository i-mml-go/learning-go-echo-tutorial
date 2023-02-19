package routing

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"toplearn-api/controller"
	"toplearn-api/viewModel/common/security"
)

func SetRouting(e *echo.Echo) error {
	e.POST("/login", controller.LoginUser)

	g := e.Group("users")

	g.GET("/get", controller.GetListOfUser)
	g.GET("/get/:id/avatars", controller.GetUserAvatar)

	jwtConfig := middleware.JWTConfig{SigningKey: []byte("our-secret-key-in-golang-project"), Claims: &security.JwtClaims{}}
	g.POST("/createNewUser", controller.CreateNewUser, middleware.JWTWithConfig(jwtConfig))

	return nil
}
