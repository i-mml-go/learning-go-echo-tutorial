package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"log"
	"toplearn-api/Utility"
	"toplearn-api/config"
	"toplearn-api/routing"
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
	// must be written without ()
	server.Pre(RootLevel)
	server.Use(AfterRooter)

	// start server
	server.Start(":" + config.AppConfig.Server.Port)

}
