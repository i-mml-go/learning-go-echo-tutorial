package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"toplearn-api/database"
	"toplearn-api/model/user"
)

type UserRepository interface {
	GetUserList() ([]user.User, error)
	GetUserById(id string) (user.User, error)
	GetUserByUserNameAndPassword(username, password string) (user.User, error)
	InsertUser(user user.User) (string, error)
	UpdateUserById(user user.User) error
	DeleteUserById(id string) error
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

func (userRepository userRepository) GetUserById(id string) (user.User, error) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user.User{}, err
	}

	userCollection := userRepository.db.GetUserCollection()

	var userObject user.User

	err = userCollection.FindOne(context.TODO(), bson.D{{"_id", objectId}}).Decode(&userObject)

	if err != nil {
		return user.User{}, err
	}

	return userObject, nil

}

func (userRepository userRepository) GetUserByUserNameAndPassword(username, password string) (user.User, error) {
	userCollection := userRepository.db.GetUserCollection()

	var userObject user.User
	err := userCollection.FindOne(context.TODO(), bson.D{{"userName", username}, {"password", password}}).Decode(&userObject)

	if err != nil {
		return user.User{}, err
	}

	return userObject, nil
}

func (userRepository userRepository) InsertUser(user user.User) (string, error) {
	userCollection := userRepository.db.GetUserCollection()
	result, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return "", err
	}

	objectId := result.InsertedID.(primitive.ObjectID).Hex()

	return objectId, err
}

func (userRepository userRepository) UpdateUserById(user user.User) error {

	objectId, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return err
	}
	// because we don't want the id in body of update , it's not editable
	user.Id = ""

	userCollection := userRepository.db.GetUserCollection()

	_, err = userCollection.UpdateByID(context.TODO(), bson.D{{"_id", objectId}}, bson.D{{"&set", user}})

	if err != nil {
		return err
	}

	return err
}

func (userRepository userRepository) DeleteUserById(id string) error {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	userCollection := userRepository.db.GetUserCollection()

	_, err = userCollection.DeleteOne(context.TODO(), bson.D{{"_id", objectId}})

	if err != nil {
		return err
	}

	return err
}
