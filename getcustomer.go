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


func GetCustomer(w http.ResponseWriter, r *http.Request) {
	// set header.
   w.Header().Set("Content-Type", "application/json")
   defer func() {
	if r := recover(); r != nil {
		helper.Insufficient("not found", w)
	}
  }()
	

	var customer models.Customer
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
		panic(err)
	}

	json.NewEncoder(w).Encode(customer)
}