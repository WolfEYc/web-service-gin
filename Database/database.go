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
		godotenv.Load(".env")
		return os.Getenv(MONGODB_URI)
	} else {
		return uri
	}
}

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(GetURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
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
