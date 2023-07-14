package main

import (
	"os"
	routes "web-service-gin/Routes"

	"github.com/gin-gonic/gin"
)

func port() string {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}

	return ":" + port
}

func main() {

	//database.ConnectDB()
	router := gin.Default()
	router.RedirectTrailingSlash = false

	// called as /quotes/{id}
	router.GET("quotes/:quoteId", routes.ReadOnePost)

	router.Run(port())
}
