package main

import (
	"fmt"
	"log"
	"toplearn-api/config"
	"toplearn-api/routing"

	"github.com/labstack/echo/v4"
)

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

	// routing
	routing.SetRouting(server)

	// start server
	server.Start(":" + config.AppConfig.Server.Port)

}
