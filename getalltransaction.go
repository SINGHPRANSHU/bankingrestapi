package main

import (
	"context"
	"encoding/json"
	
	"net/http"
		"restapi/helper"
		"restapi/models"
		"go.mongodb.org/mongo-driver/bson"
		
		
		
)

func Getalltransaction(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	


	var balances []models.BalanceNew

	cur, err := transaction.Find(context.TODO(), bson.M{})

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

	if err := cur.Err(); err != nil {
		helper.GetError(err, w)
		
	}

	json.NewEncoder(w).Encode(balances)
}