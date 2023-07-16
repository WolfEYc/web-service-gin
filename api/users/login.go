package users

import (
	"encoding/base64"
	"net/http"
	"time"
	"web-service-gin/httputil"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type LoginUserBody struct {
	UserName string `json:"username" example:"zamzang"`
	Password string `json:"password" example:"notpassword"`
}

const TOKEN_LIFETIME = time.Hour * 24
const COOKIE_LIFETIME = 3600 * 24
const COOKIE_NAME = "Authorization"

// LoginUser godoc
//
//	@Summary		Login User
//	@Description	Gain an auth jwt cookie for a user by providing username and password
//	@Tags			users
//	@Accept			json
//	@Param			loginuserbody	body		LoginUserBody		true	"Login User Request Body"
//	@Success		200				{}			nil					"Obtained Auth"
//	@Failure		400				{object}	httputil.HTTPError	"Bad Request Formation"
//	@Failure		401				{object}	httputil.HTTPError	"Wrong Password"
//	@Failure		404				{object}	httputil.HTTPError	"User Not Found"
//	@Failure		500				{object}	httputil.HTTPError	"Unknown internal server error"
//	@Router			/users/login [post]
func LoginUser(c *gin.Context) {
	var body LoginUserBody

	bodyReadErr := c.Bind(&body)

	if bodyReadErr != nil {
		httputil.NewError(c, http.StatusBadRequest, bodyReadErr)
		return
	}

	query := GetUserByName(body.UserName)

	queryerr := query.Err()

	if queryerr != nil {
		httputil.NewError(c, http.StatusNotFound, queryerr)
		return
	}

	var user models.User

	decodeErr := query.Decode(&user)

	if decodeErr != nil {
		httputil.NewError(c, http.StatusInternalServerError, decodeErr)
	}

	hashedPassword, b64DecodeErr := base64.URLEncoding.DecodeString(user.Password)

	if b64DecodeErr != nil {
		httputil.NewError(c, http.StatusInternalServerError, b64DecodeErr)
		return
	}

	compareErr := bcrypt.CompareHashAndPassword(
		hashedPassword,
		[]byte(body.Password))

	if compareErr != nil {
		httputil.NewError(c, http.StatusUnauthorized, compareErr)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(TOKEN_LIFETIME).Unix(),
	})

	tokenStr, tokenSigningErr := token.SignedString(httputil.JWTPrivateKey)

	if tokenSigningErr != nil {
		httputil.NewError(c, http.StatusInternalServerError, tokenSigningErr)
		return
	}

	c.SetSameSite(http.SameSiteDefaultMode)
	c.SetCookie(
		COOKIE_NAME,
		tokenStr,
		COOKIE_LIFETIME,
		"",
		"",
		httputil.IsSecure,
		true)

	c.Status(http.StatusOK)
}
