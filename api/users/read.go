package users

import (
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

// GetUser godoc
//
//	@Summary		Get User
//	@Description	returns your user info from your auth cookie
//	@Tags			users
//	@Produce		json
//	@Success		200	{object}	models.User	"Current User"
//	@Failure		401	{object}	string		"Not authorized as any user"
//	@Failure		500	{object}	string		"Unknown internal server error"
//	@Router			/users/me [get]
func GetUser(c *gin.Context) {
	user, hasUser := c.Get(REQ_USER)

	if !hasUser {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "could not find user in request, despite being authorized!")
		return
	}

	c.JSON(http.StatusOK, user)
}
