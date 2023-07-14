package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Quote struct {
	ID                      primitive.ObjectID `bson:"_id"`
	userid                  primitive.ObjectID `bson:"userid"`
	gallonsrequested        int32              `bson:"gallonsrequested"`
	deliveryaddress         string             `bson:"deliveryaddress"`
	deliverydate            primitive.DateTime `bson:"deliverydate"`
	suggestedpricepergallon float64            `bson:"suggestedpricepergallon"`
	totalamountdue          float64            `bson:"totalamountdue"`
}

