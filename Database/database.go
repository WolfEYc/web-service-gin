package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGODB_URI = "MONGODB_URI"

func GetURI() string {
	uri, exists := os.LookupEnv(MONGODB_URI)
	if !exists {
		err := godotenv.Load(".env")

		if err != nil {
			log.Fatal(err)
		}

		return os.Getenv(MONGODB_URI)
	} else {
		return uri
	}
}

func ConnectDB() *mongo.Client {

	opts := options.Client().ApplyURI(GetURI())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

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
	collection := client.Database("fuelquote").Collection(collectionName)
	return collection
}
