package users

import (
	"log"
	"net/http"
	db "web-service-gin/db"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserByName(username string) *mongo.SingleResult {
	users := db.GetCollection(db.Users)

	ctx, cancel := db.DefaultContext()

	defer cancel()

	return users.FindOne(ctx, bson.M{"username": username})
}

func UserNameExists(username string) bool {

	user := GetUserByName(username)

	return user.Err() == nil
}

func GetUserByID(userid primitive.ObjectID) *mongo.SingleResult {
	users := db.GetCollection(db.Users)

	ctx, cancel := db.DefaultContext()

	defer cancel()

	return users.FindOne(ctx, bson.M{"_id": userid})
}

const REQ_USER = "user"

func GetUser(c *gin.Context) {
	user, hasUser := c.Get(REQ_USER)

	if !hasUser {
		log.Fatal("Missing user in request, despite being authorized!")
	}

	c.JSON(http.StatusOK, user)
}
