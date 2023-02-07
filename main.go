package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"toplearn-api/config"
	"toplearn-api/database"
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

	db, err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	} else {
		userCollection := db.GetUserCollection()

		var result bson.M
		err1 := userCollection.FindOne(context.TODO(), bson.D{{"firstName", "Laura"}}).Decode(&result)
		if err1 != nil {
			log.Fatalln(err1)
		}

		fmt.Println(result)
	}

	err = config.GetConfig()
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
