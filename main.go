package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"toplearn-api/Utility"
	"toplearn-api/config"
	"toplearn-api/routing"
	"toplearn-api/viewModel/common/security"
)

func RootLevel(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("i am RootLevel")
	return next
}
func AfterRooter(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("i am after")
	return next
}

func main() {
	// ***** structure of projects *****
	// get config
	// init server
	// routing
	// middleware
	// start server

	// get config
	err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("server port :", config.AppConfig.Server.Port)

	// init server
	server := echo.New()
	server.Validator = &Utility.CustomValidator{Validator: validator.New()}

	// routing
	routing.SetRouting(server)

	// middleware
	server.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			apiContext := &Utility.ApiContext{Context: c}
			return next(apiContext)
		}
	})
	server.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:             []byte("our-secret-key-in-golang-project"),
		Claims:                 &security.JwtClaims{},
		ContinueOnIgnoredError: true,
		ErrorHandlerWithContext: func(err error, c echo.Context) error {
			return nil
		},
	}))

	server.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))

	// start server
	server.Start(":" + config.AppConfig.Server.Port)

}
