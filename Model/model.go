package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Quote struct {
	ID                      primitive.ObjectID `json:"_id"`
	UserId                  primitive.ObjectID `json:"userid"`
	GallonsRequested        int32              `json:"gallonsrequested"`
	DeliveryAddress         string             `json:"deliveryaddress"`
	DeliveryDate            primitive.DateTime `json:"deliverydate"`
	SuggestedPricePerGallon float64            `json:"suggestedpricepergallon"`
	TotalAmountDue          float64            `json:"totalamountdue"`
}
