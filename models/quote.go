package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Quote struct {
	ID                      primitive.ObjectID `bson:"_id" example:"63f79148c3a63a346225fd99"`
	UserId                  primitive.ObjectID `bson:"userid" example:"63f78fc8c3a63a346225fd96"`
	GallonsRequested        int32              `json:"gallonsrequested" example:"99999"`
	DeliveryAddress         string             `json:"deliveryaddress" example:"best st"`
	DeliveryDate            primitive.DateTime `bson:"deliverydate" swaggertype:"primitive,string" example:"2023-02-24T00:00:00Z"`
	SuggestedPricePerGallon float64            `json:"suggestedpricepergallon" example:"1.74"`
	TotalAmountDue          float64            `json:"totalamountdue" example:"173998.26"`
}
