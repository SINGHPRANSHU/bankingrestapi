package main

import (
	"context"
	"encoding/json"
	"net/http"
	"restapi/helper"
	"restapi/models"

	"go.mongodb.org/mongo-driver/bson"

	// "github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type retCustomer struct{
	ID        primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string                 `json:"name,omitempty" bson:"name,omitempty"`
	Email     string                 `json:"email,omitempty" bson:"email,omitempty"`
	Balance   int                    `json:"balance,omitempty" bson:"balance,omitempty"`
	IDs       []models.BalanceNew    `json:"ids,omitempty" bson:"ids,omitempty"` 
}


func GetCustomer(w http.ResponseWriter, r *http.Request) {
	// set header.
   w.Header().Set("Content-Type", "application/json")
   defer func() {
	if r := recover(); r != nil {
		helper.Insufficient("not found", w)
	}
  }()
	

	var customer models.Customer
	var balances []models.BalanceNew
	// we get params with mux.
	// var params = mux.Vars(r)

	// string to primitive.ObjectID
	id,senderiderr := primitive.ObjectIDFromHex(r.Header.Get("userid"))

	if senderiderr != nil {
		helper.Insufficient("not found", w)
		return
    }

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&customer)
	if err != nil {
		helper.GetError(err, w)
		return
	}

	// for _, value := range customer.IDs {
	// 	fmt.Println(value)
	// }

	cur, err := transaction.Find(context.TODO(), bson.M{"_id": bson.M{"$in": customer.IDs}})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var balance models.BalanceNew
		// & character returns the memory address of the following variable.
		err := cur.Decode(&balance) // decode similar to deserialize process.
		if err != nil {
			helper.GetError(err, w)
		}

		// add item our array
		balances = append(balances, balance)

		
	}

	newCustomer := retCustomer{ID: customer.ID,Name: customer.Name,Email: customer.Email, Balance: customer.Balance, IDs: balances}
    
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(newCustomer)
}