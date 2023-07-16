package token

import (
	"net/http"
	"web-service-gin/api/users"

	"github.com/gin-gonic/gin"
)

func GetAuthStrFromCookie(c *gin.Context) (string, error) {
	return c.Cookie(users.COOKIE_NAME)
}

// GetToken godoc
//
//	@Summary		Get Token
//	@Description	Returns your authorization cookie, if it exists
//	@Tags			token
//	@Produce		plain
//	@Success		200	{string}	string	"Cookie Exists"
//	@Failure		404	{string}	string	"No Cookie"
//	@Failure		500	{string}	string	"Unknown internal server error"
//	@Router			/token [get]
func GetToken(c *gin.Context) {
	cookie, err := GetAuthStrFromCookie(c)

	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
	}

	c.String(http.StatusOK, "%s", cookie)
}
