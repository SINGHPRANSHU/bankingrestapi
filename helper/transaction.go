package helper

import (
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
)

// ConnectDB : This is helper function to connect mongoDB
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func ConnectTransactionDb() *mongo.Collection {
	// Set client options

	// Connect to MongoDB
	client, err := GetClient()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to transfer collection MongoDB!")

	collection := client.Database("go_rest_api").Collection("transfers")
	
	
	return collection
}

