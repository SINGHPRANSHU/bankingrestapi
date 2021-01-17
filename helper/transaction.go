package helper

import (
	"context"
	"fmt"
	"log"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB : This is helper function to connect mongoDB
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func ConnectTransactionDb() *mongo.Collection {
	
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://"+os.Getenv("username")+":"+os.Getenv("password")+"@cluster0.dgzyl.mongodb.net/"+os.Getenv("database")+"?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to transfer collection MongoDB!")

	collection := client.Database("go_rest_api").Collection("transfers")
	
	
	return collection
}
