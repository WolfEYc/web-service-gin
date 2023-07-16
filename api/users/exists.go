package users

import (
	"web-service-gin/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserByName(username string) *mongo.SingleResult {
	users := database.GetCollection("users")

	ctx, cancel := database.DefaultContext()

	defer cancel()

	return users.FindOne(ctx, bson.M{"username": username})
}

func UserNameExists(username string) bool {
	
	user := GetUserByName(username);

	return user.Err() == nil
}
