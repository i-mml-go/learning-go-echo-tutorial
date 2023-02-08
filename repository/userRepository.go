package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"toplearn-api/database"
	"toplearn-api/model/user"
)

type UserRepository interface {
	GetUserList() ([]user.User, error)
}

type userRepository struct {
	db database.Db
}

func NewUserRepository() UserRepository {
	db, err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	return userRepository{
		db: db,
	}
}

func (userRepository userRepository) GetUserList() ([]user.User, error) {
	userColection := userRepository.db.GetUserCollection()

	cursor, err := userColection.Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}

	var users []user.User
	err = cursor.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}

	return users, nil

}
