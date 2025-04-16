package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Username string
	Password string
} 