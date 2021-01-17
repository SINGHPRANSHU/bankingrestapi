package main


import (
	"context"
	"encoding/json"
	"net/http"
	"restapi/helper"
	"restapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"		
)


func UpdateBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	



	var balance models.BalanceNew
	var sendcustomer models.Customer
	var reccustomer models.Customer

/////////////////////////////////////////////////////////////////////////////////////////////////////
	// Read update model from body request
	decodeerror := json.NewDecoder(r.Body).Decode(&balance)
	if balance.Balance==0{
		helper.Insufficient("transfer more than 0",w)
		return
	}


	if decodeerror != nil {
		helper.Insufficient("server error", w)
		return
    }
	
	sendid,senderiderr :=  primitive.ObjectIDFromHex(balance.SenderID)
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

	senderr := collection.FindOne(context.TODO(), sendfilter).Decode(&sendcustomer)

	if senderr != nil {
		helper.GetError(senderr, w)
		return
	}
//////////////////////////////////////////////////////////////////////////////////////////////////////////////
	if(sendcustomer.Balance>balance.Balance){


		recerr := collection.FindOne(context.TODO(), recfilter).Decode(&reccustomer)

		if recerr != nil {
			helper.GetError(recerr, w)
			return
		}
		
		
		
		// prepare update model.
	
			sendupdate := bson.D{
				{"$set", bson.D{
					{"name",sendcustomer.Name},
					{"email",sendcustomer.Email},
					{"balance",sendcustomer.Balance-balance.Balance},
				}},
			}
			erro := collection.FindOneAndUpdate(context.TODO(), sendfilter, sendupdate).Decode(&sendcustomer)
		
			   if erro != nil {
				
				 helper.GetError(erro, w)
				 return
				}
				
		 recupdate := bson.D{
				{"$set", bson.D{
					{"name",reccustomer.Name},
					{"email",reccustomer.Email},
					{"balance",reccustomer.Balance+balance.Balance},
				}},
			}
			recerro := collection.FindOneAndUpdate(context.TODO(), recfilter, recupdate).Decode(&reccustomer)
		
			   if recerro != nil {
				
				 helper.GetError(recerro, w)
				 return
				}
	
		
		
		insertbalance:=models.BalanceNew{Sender:sendcustomer.Name,SenderID:balance.SenderID,Receiver:reccustomer.Name,ReceiverID:balance.ReceiverID,Balance:balance.Balance,Time:time.Now()  }
	
			
		
		
		result, er := transaction.InsertOne(context.TODO(), insertbalance)
		   
				if er != nil {
				
					helper.GetError(er, w)
					return
				}
	
		

	c1 := make(chan int)
	c2 := make(chan int)
 			
    go	func (){
		update := bson.D{{"$push", bson.D{{"ids", result.InsertedID}}}}
		recerro := collection.FindOneAndUpdate(context.TODO(), recfilter, update).Decode(&reccustomer)
		if recerro != nil {	
			//rollback func
			helper.GetError(recerro, w)
			return
		   }
        c1<-1
	}()
	
	
	go func (){
		update := bson.D{{"$push", bson.D{{"ids", result.InsertedID}}}}
		senderro := collection.FindOneAndUpdate(context.TODO(), sendfilter, update).Decode(&sendcustomer)
		if senderro != nil {
			//rollback func
			helper.GetError(recerro, w)
			return
		   }
		   c2<-1
	}()
  

	//for synchroization
	<-c1
	<-c2
	


	json.NewEncoder(w).Encode(result)
		
	}else{
		helper.Insufficient("not enough balance", w)
	}

	
}





