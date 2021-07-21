package main

import (
	"encoding/json"
	"net/http"
	"restapi/helper"
	"restapi/models"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func Transfer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var balance models.BalanceNew
	var sendcustomer models.Customer
	var reccustomer models.Customer

	decodeerror := json.NewDecoder(r.Body).Decode(&balance)
	if balance.Balance==0{
		helper.Insufficient("transfer more than 0",w)
		return
	}


	if decodeerror != nil {
		helper.Insufficient("server error", w)
		return
    }
	
	sendid,senderiderr :=  primitive.ObjectIDFromHex(r.Header.Get("userid"))
	if senderiderr != nil {
		helper.Insufficient("server error", w)
		return
    }
	recid,reciderr :=  primitive.ObjectIDFromHex(balance.ReceiverID)
	if reciderr != nil {
		helper.Insufficient("server error", w)
		return
    }
	// Create filter
	sendfilter := bson.M{"_id":sendid}
	recfilter := bson.M{"_id":recid}

	senderr := collection.FindOne(r.Context(), sendfilter).Decode(&sendcustomer)

	if senderr != nil {
		helper.GetError(senderr, w)
		return
	}

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		// Important: You must pass sessCtx as the Context parameter to the operations for them to be executed in the
		// transaction.
		recerr := collection.FindOne(sessCtx, recfilter).Decode(&reccustomer)
		if recerr != nil {
			return nil,recerr
		}

		sendupdate := bson.D{
			{"$set", bson.D{
				{"name",sendcustomer.Name},
				{"email",sendcustomer.Email},
				{"balance",sendcustomer.Balance-balance.Balance},
			}},
		}
		erro := collection.FindOneAndUpdate(sessCtx, sendfilter, sendupdate).Decode(&sendcustomer)
	    if erro != nil {
			return nil,erro
		}

		recupdate := bson.D{
			{"$set", bson.D{
				{"name",reccustomer.Name},
				{"email",reccustomer.Email},
				{"balance",reccustomer.Balance+balance.Balance},
			}},
		}
		recerro := collection.FindOneAndUpdate(sessCtx, recfilter, recupdate).Decode(&reccustomer)
		if recerro != nil {
			return nil, recerro
		}

		insertbalance:=models.BalanceNew{Sender:sendcustomer.Name,SenderID:balance.SenderID,Receiver:reccustomer.Name,ReceiverID:balance.ReceiverID,Balance:balance.Balance,Time:time.Now()  }
	    result, er := transaction.InsertOne(sessCtx, insertbalance)
		if er != nil {		
			return nil, er
		}

		update := bson.D{{"$push", bson.D{{"ids", result.InsertedID}}}}
		recerro = collection.FindOneAndUpdate(sessCtx, recfilter, update).Decode(&reccustomer)
		if recerro != nil {	
			return nil, recerro
		}

		update = bson.D{{"$push", bson.D{{"ids", result.InsertedID}}}}
		senderro := collection.FindOneAndUpdate(sessCtx, sendfilter, update).Decode(&sendcustomer)
		if senderro != nil {
			return nil, senderro
		   }

		return result, nil
	}

	if sendcustomer.Balance>balance.Balance {
        client, err := helper.GetClient()
		session, err := client.StartSession()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			var response = helper.ErrorResponse{
				ErrorMessage: "cannot proceed try after sometime",
				StatusCode:   http.StatusInternalServerError,
			}
			message, _ := json.Marshal(response)
			w.Write(message)
			return
		}
		defer session.EndSession(r.Context())
		result, err := session.WithTransaction(r.Context(), callback)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			var response = helper.ErrorResponse{
				ErrorMessage: "cannot proceed try after sometime",
				StatusCode:   http.StatusInternalServerError,
			}
			message, _ := json.Marshal(response)
			w.Write(message)
			return
		}
		json.NewEncoder(w).Encode(result)

	}else{
		helper.Insufficient("not enough balance", w)
	}
}