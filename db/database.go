package db

import (
	"context"
	"fmt"
	"log"
	"time"
	"web-service-gin/httputil"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Quotes  = "quotes"
	Users   = "users"
	DB_NAME = "fuelquote"
)

func ConnectDB() *mongo.Client {

	opts := options.Client().ApplyURI(httputil.GetMongoURI())
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}

var client *mongo.Client = ConnectDB()

func GetCollection(collectionName string) *mongo.Collection {
	collection := client.Database(DB_NAME).Collection(collectionName)
	return collection
}

func DefaultContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(
		context.Background(),
		10*time.Second)
}
