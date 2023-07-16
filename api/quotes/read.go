package quotes

import (
	"net/http"
	db "web-service-gin/db"
	models "web-service-gin/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetQuote godoc
//
//	@Summary		Get Quote
//	@Description	gets a quote from its id
//	@Tags			quotes
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"Quote ID"
//	@Success		200	{object}	models.Quote	"Queried Quote"
//	@Failure		400	{object}	string			"Invalid ID format"
//	@Failure		404	{object}	string			"Quote not found"
//	@Failure		500	{object}	string			"Unknown internal server error"
//	@Router			/quotes/{id} [get]
func GetQuote(c *gin.Context) {
	var postCollection = db.GetCollection(db.Quotes)

	postId := c.Param("id")
	var result models.Quote

	objId, iderr := primitive.ObjectIDFromHex(postId)

	if iderr != nil {
		_ = c.AbortWithError(http.StatusBadRequest, iderr)
		return
	}

	ctx, cancel := db.DefaultContext()

	defer cancel()

	finderr := postCollection.FindOne(ctx, bson.M{"_id": objId}).
		Decode(&result)

	if finderr != nil {
		_ = c.AbortWithError(http.StatusNotFound, finderr)
		return
	}

	result.ID = objId
	c.JSON(http.StatusOK, result)
}
