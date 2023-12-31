package users

import (
	"encoding/base64"
	"errors"
	"net/http"
	"web-service-gin/db"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserBody struct {
	UserName      string `json:"username" example:"zamzang"`
	Password      string `json:"password" example:"notpassword"`
	FullName      string `json:"fullname" example:"zenyatta man"`
	StreetAddress string `json:"address1" example:"coolsvile avvae"`
	City          string `json:"city" example:"Houstonian"`
	State         string `json:"state" example:"AZ"`
	Zipcode       int32  `json:"zipcode" example:"827392"`
	SecondAddress string `json:"address2,omitempty" example:"billion ave"`
}

// CreateUser godoc
//
//	@Summary		Create User
//	@Description	Creates a new user from user info
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			createuserbody	body		CreateUserBody	true	"Create User Request Body"
//	@Success		201				{string}	string			"Created User's ID"
//	@Failure		400				{string}	string			"Bad User Entry"
//	@Failure		500				{string}	string			"Unknown internal server error"
//	@Router			/users [post]
func CreateUser(c *gin.Context) {
	var body CreateUserBody

	bodyReadErr := c.Bind(&body)

	if bodyReadErr != nil {
		_ = c.AbortWithError(http.StatusBadRequest, bodyReadErr)
		return
	}

	if UserNameExists(body.UserName) {

		_ = c.AbortWithError(
			http.StatusBadRequest,
			errors.New("Username Taken"))
		return
	}

	hashword, hasherr := bcrypt.GenerateFromPassword(
		[]byte(body.Password),
		10)

	if hasherr != nil {
		_ = c.AbortWithError(http.StatusBadRequest, hasherr)
		return
	}

	body.Password = base64.URLEncoding.EncodeToString(hashword)

	var userCollection = db.GetCollection(db.Users)

	ctx, cancel := db.DefaultContext()

	defer cancel()

	result, insertErr := userCollection.InsertOne(
		ctx,
		body)

	if insertErr != nil {
		_ = c.AbortWithError(http.StatusBadRequest, insertErr)
		return
	}

	c.JSON(http.StatusCreated, result.InsertedID)
}
