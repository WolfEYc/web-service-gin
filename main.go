package main

import (
	"os"
	routes "web-service-gin/Routes"
	"web-service-gin/docs"

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

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

func GetScheme() []string {
	scheme, hasScheme := os.LookupEnv("SCHEME")

	if hasScheme {
		return []string{scheme}
	} else {
		return []string{"https"}
	}
}

func InitSwag() {
	docs.SwaggerInfo.Title = "Fuelquote API"
	docs.SwaggerInfo.Description = "This is a fuelquote project API"
	docs.SwaggerInfo.Version = "1.0"

	//docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = GetScheme()
}

func main() {

	//database.ConnectDB()
	InitSwag()
	router := gin.Default()
	router.RedirectTrailingSlash = false

	v1 := router.Group(docs.SwaggerInfo.BasePath)
	{
		quotes := v1.Group("/quotes")
		{
			quotes.GET(":id", routes.GetQuote)
		}
	}

	router.GET("swagger/*any", ginswagger.WrapHandler(swaggerFiles.Handler))

	_ = router.Run(port())
}
