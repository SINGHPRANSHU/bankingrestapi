package helper

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client = nil
var err error = nil
func GetClient() (*mongo.Client, error){

	if client == nil {
		clientOptions := options.Client().ApplyURI(os.Getenv("dbconnectionstring"))
		client, err = mongo.Connect(context.TODO(), clientOptions)
	}

    return client,err
}