package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"_id" example:"63f78fc8c3a63a346225fd96"`
	UserName      string             `json:"username" example:"zamzang"`
	Password      string             `json:"password" example:"bdeac530209b16657065661aa11fcdd85b21b01a"`
	FullName      string             `json:"fullname" example:"zenyatta man"`
	StreetAddress string             `json:"address1" example:"coolsvile avvae"`
	City          string             `json:"city" example:"Houstonian"`
	State         string             `json:"state" example:"AZ"`
	Zipcode       int32              `json:"zipcode" example:"827392"`
	QuoteCount    int32              `json:"quotes" example:"0"`
	SecondAddress string             `json:"address2,omitempty" example:"billion ave"`
}
