package main

import (
	"context"
	"encoding/json"
	"net/http"
	"restapi/helper"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerWithPassword struct{
	ID        primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string                 `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Email     string                 `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Password  string                 `json:"password,omitempty" bson:"password,omitempty" validate:"required,nefield=Email"`
	Balance   int                    `json:"balance,omitempty" bson:"balance,omitempty"`
	IDs       []primitive.ObjectID   `json:"ids,omitempty" bson:"ids,omitempty"` 
}



func Signup(w http.ResponseWriter, r *http.Request) {
  var user CustomerWithPassword
  var alreadyExistUser CustomerWithPassword
	decodeerror := json.NewDecoder(r.Body).Decode(&user)
	if decodeerror != nil {
		helper.Insufficient("request body is not valid", w)
		return
    }

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
        validationErrors := err.(validator.ValidationErrors)
       
        responseBody := map[string]string{"error": validationErrors.Error()}
        json.NewEncoder(w).Encode(responseBody)
        return
    }
	filter := bson.M{"email": user.Email}
    erro := collection.FindOne(context.TODO(),filter).Decode(&alreadyExistUser)
	if erro != nil {
		if erro.Error() == "mongo: no documents in result"{
			hashPassword := helper.CreateHash(user.Password)
			collection.InsertOne(context.TODO(), CustomerWithPassword{Name: user.Name, Email: user.Email,Password: hashPassword,Balance: 100})
		    helper.Success("user created", w)
		}
		return
	}

	helper.Insufficient("user already exist", w)

	
}