package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenResponse struct {
	Token            string                          `json:"token"` 
	ID               primitive.ObjectID              `json:"_id,omitempty" bson:"_id,omitempty"` 
	Name             string                          `json:"name,omitempty" bson:"name,omitempty"`
	Email            string                          `json:"email,omitempty" bson:"email,omitempty"`	
}

type User struct{
	ID                  primitive.ObjectID              `json:"_id,omitempty" bson:"_id,omitempty"` 
	Email               string                          `json:"email,omitempty" bson:"email,omitempty"`
	Name                string                          `json:"name,omitempty" bson:"name,omitempty"`
	Password            string                          `json:"password,omitempty" bson:"password,omitempty"`	
}


func Login(w http.ResponseWriter, r *http.Request){
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Detected panic")
			helper.Insufficient("invalid credentials", w)
        }
    }()
	// email, password := r.Body
	var user User
	decodeerror := json.NewDecoder(r.Body).Decode(&user)
	if decodeerror != nil {
		helper.Insufficient("request body is not valid", w)
		return
    }
	filter := bson.M{"email":user.Email, "password": user.Password}
	recerr := collection.FindOne(context.TODO(), filter).Decode(&user)
	if recerr != nil {
		// helper.GetError(recerr, w)
		panic("cannot find user")
		// return
	}
	token, tokenerr := helper.GenerateToken(user.ID)
	if tokenerr != nil {
		helper.GetError(tokenerr, w)
		
	}else {
		tokenres := TokenResponse{Token: token, Email: user.Email,ID: user.ID, Name: user.Name}
		json.NewEncoder(w).Encode(tokenres)
	}
	
	
}