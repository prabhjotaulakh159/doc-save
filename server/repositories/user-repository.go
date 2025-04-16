package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/prabhjotaulakh159/doc-save/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"errors"
)

type UserRepository interface {
	CheckIfUserExists(username string) (bool, error)
	CreateNewUser(username string, password string) error
	GetUserByUsername(username string) (*models.UserModel, error)
}

type CrudUserRepository struct {
	Collection *mongo.Collection
}

func (c *CrudUserRepository) CheckIfUserExists(username string) (bool, error) {
	var result models.UserModel
	err := c.Collection.FindOne(context.TODO(), bson.D{{"username", username}}, nil).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil		
		} 
		return false, err
	}
	return true, nil 
}

func (c *CrudUserRepository) CreateNewUser(username string, password string) error {
	_, err := c.Collection.InsertOne(context.TODO(), &models.UserModel{Username: username, Password: password})
	if err != nil {
		return err
	}
	return nil
}

func (c *CrudUserRepository) GetUserByUsername(username string) (*models.UserModel, error) {
	var user *models.UserModel
	err := c.Collection.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil		
		}	
		return nil, err
	}
	return user, nil
}