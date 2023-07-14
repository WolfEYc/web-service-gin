package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Quote struct {
	ID                      primitive.ObjectID `json:"_id,omitempty"`
	UserId                  primitive.ObjectID `json:"userid,omitempty"`
	GallonsRequested        int32              `json:"gallonsrequested,omitempty"`
	DeliveryAddress         string             `json:"deliveryaddress,omitempty"`
	DeliveryDate            primitive.DateTime `json:"deliverydate,omitempty"`
	SuggestedPricePerGallon float64            `json:"suggestedpricepergallon,omitempty"`
	TotalAmountDue          float64            `json:"totalamountdue,omitempty"`
}
