package db

import (
	"context"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetMongoClient() (*mongo.Client, error) {
	opts := options.Client().ApplyURI(os.Getenv("doc-save-connection-string"))
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err	
	}
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}
	return client, nil
}