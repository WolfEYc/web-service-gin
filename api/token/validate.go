package token

import (
	"fmt"
	"net/http"
	"time"
	"web-service-gin/api/users"
	"web-service-gin/httputil"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RequireAuth(c *gin.Context) {
	cookie, cookieerr := GetAuthStrFromCookie(c)
	if cookieerr != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, cookieerr)
		return
	}

	token, jwtparseerr := jwt.Parse(
		cookie,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header)
			}

			return httputil.JWTPrivateKey, nil
		})

	if jwtparseerr != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, jwtparseerr)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			_ = c.AbortWithError(
				http.StatusUnauthorized,
				jwt.NewValidationError("token expired", jwt.ValidationErrorExpired))
			return
		}

		println(claims["sub"])

		userid, userHexErr := primitive.ObjectIDFromHex((claims["sub"].(string)))

		if userHexErr != nil {
			_ = c.AbortWithError(http.StatusUnauthorized, userHexErr)
			return
		}

		query := users.GetUserByID(userid)

		queryerr := query.Err()

		if queryerr != nil {
			_ = c.AbortWithError(http.StatusUnauthorized, queryerr)
			return
		}

		var user models.User
		queryReturnErr := query.Decode(&user)

		if queryReturnErr != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, queryReturnErr)
			return
		}

		c.Set(users.REQ_USER, user)

		c.Next()
	} else {

		_ = c.AbortWithError(
			http.StatusUnauthorized,
			jwt.NewValidationError("token invalid", jwt.ValidationErrorClaimsInvalid))
	}
}

// ValidateToken godoc
//
//	@Summary		Validate Token
//	@Description	Validates your auth token from your browser cookie
//	@Tags			token
//	@Success		200	{}			nil		"Authorized"
//	@Failure		404	{string}	string	"No Cookie"
//	@Failure		401	{string}	string	"Unauthorized"
//	@Failure		500	{string}	string	"Unknown internal server error"
//	@Router			/token/validate [get]
func ValidateToken(c *gin.Context) {
	c.Status(http.StatusOK)
}
