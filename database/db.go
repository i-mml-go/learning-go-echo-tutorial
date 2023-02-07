package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Db struct {
	client *mongo.Client
}

func Connect() (Db, error) {
	ctx := context.TODO()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return Db{}, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return Db{}, err
	}

	return Db{client: client}, nil

	//userColention := client.Database("toplearn-api-golang").Collection("users")
	//
	//var result bson.M
	//err1 := userColention.FindOne(ctx, bson.D{{"firstName", "Laura"}}).Decode(&result)
	//if err1 != nil {
	//	log.Fatalln(err1)
	//}
	//
	//fmt.Println(result)
}

func (db Db) GetUserCollection() *mongo.Collection {
	userColention := db.client.Database("toplearn-api-golang").Collection("users")

	return userColention

}
