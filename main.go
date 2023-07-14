package main

import (
	routes "web-service-gin/Routes"

	"github.com/gin-gonic/gin"
)

func main() {

	//database.ConnectDB()
	router := gin.Default()

	// called as localhost:3000/quotes/{id}
	router.GET("quotes/:quoteId", routes.ReadOnePost)

	router.Run("localhost:3000")
}
