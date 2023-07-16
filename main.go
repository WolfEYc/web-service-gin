package main

import (
	"os"
	"web-service-gin/api/quotes"
	"web-service-gin/api/users"
	"web-service-gin/docs"
	"web-service-gin/httputil"

	_ "web-service-gin/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

func port() string {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}

	return ":" + port
}

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							cookie
//	@name						Authorization
//	@description				Basic password login token auth, stored in a cookie for convenience
func InitSwag() {
	docs.SwaggerInfo.Title = "Fuelquote API"
	docs.SwaggerInfo.Description = "This is a fuelquote project API"
	docs.SwaggerInfo.Version = "1.0"

	//docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = httputil.GetURLProtocol()
}

func main() {

	//database.ConnectDB()
	InitSwag()
	router := gin.Default()
	router.RedirectTrailingSlash = false

	v1 := router.Group(docs.SwaggerInfo.BasePath)
	{
		quotesRoute := v1.Group("/quotes")
		{
			quotesRoute.GET(":id", quotes.GetQuote)
		}
		usersRoute := v1.Group("/users")
		{
			usersRoute.GET(":id", users.ReadUser)
			usersRoute.POST("", users.CreateUser)
			usersRoute.POST("/login", users.LoginUser)
		}
	}

	router.GET(
		"swagger/*any",
		ginswagger.WrapHandler(swaggerFiles.Handler))

	_ = router.Run(port())
}
