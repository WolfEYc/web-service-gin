package quotes

import (
	"net/http"
	database "web-service-gin/database"
	"web-service-gin/httputil"
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
//	@Param			id	path		string				true	"Quote ID"
//	@Success		200	{object}	models.Quote		"Queried Quote"
//	@Failure		400	{object}	httputil.HTTPError	"Invalid ID format"
//	@Failure		404	{object}	httputil.HTTPError	"Quote not found"
//	@Failure		500	{object}	httputil.HTTPError	"Unknown internal server error"
//	@Router			/quotes/{id} [get]
func GetQuote(apictx *gin.Context) {
	var postCollection = database.GetCollection("quotes")

	postId := apictx.Param("id")
	var result models.Quote

	objId, iderr := primitive.ObjectIDFromHex(postId)

	if iderr != nil {
		httputil.NewError(apictx, http.StatusBadRequest, iderr)
		return
	}

	ctx, cancel := database.DefaultContext()

	defer cancel()

	finderr := postCollection.FindOne(ctx, bson.M{"_id": objId}).
		Decode(&result)

	if finderr != nil {
		httputil.NewError(apictx, http.StatusNotFound, finderr)
		return
	}

	result.ID = objId
	apictx.JSON(http.StatusOK, result)
}
