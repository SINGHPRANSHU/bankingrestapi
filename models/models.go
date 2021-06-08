package models

import (
	    "go.mongodb.org/mongo-driver/bson/primitive"
		"time"
	   )



type Customer struct{
	ID        primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string                 `json:"name,omitempty" bson:"name,omitempty"`
	Email     string                 `json:"email,omitempty" bson:"email,omitempty"`
	Balance   int                    `json:"balance,omitempty" bson:"balance,omitempty"`
	IDs       []primitive.ObjectID   `json:"ids,omitempty" bson:"ids,omitempty"` 
}

type BalanceNew struct{
	ID                  primitive.ObjectID              `json:"_id,omitempty" bson:"_id,omitempty"` 
	Sender              string                          `json:"sender,omitempty" bson:"sender,omitempty"` 
	SenderID            string                          `json:"senderid,omitempty" bson:"senderid,omitempty"`          
	Receiver            string                          `json:"receiver,omitempty" bson:"receiver,omitempty"` 
	ReceiverID          string                          `json:"receiverid,omitempty" bson:"receiverid,omitempty"`          
	Balance             int                             `json:"balance,omitempty" bson:"balance,omitempty"`
	Time                time.Time                       `json:"time,omitempty" bson:"time,omitempty"`
}
