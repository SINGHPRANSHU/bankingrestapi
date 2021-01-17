package main


import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"restapi/helper"
	"restapi/models"
	"go.mongodb.org/mongo-driver/bson"
		
		
)


func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
    (w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// we created Customer array
	var customers []models.Customer

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var customer models.Customer
		// & character returns the memory address of the following variable.
		err := cur.Decode(&customer) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
			helper.GetError(err, w)
		}

		// add item our array
		customers = append(customers, customer)
	}

	if err := cur.Err(); err != nil {
		helper.GetError(err, w)
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(customers) // encode similar to serialize process.
}