package routes

import (
	"context"
	"net/http"
	"time"
	database "web-service-gin/Database"
	model "web-service-gin/Model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReadOnePost(apictx *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var postCollection = database.GetCollection("quotes")

	postId := apictx.Param("quoteId")
	var result model.Quote

	defer cancel()

	objId, iderr := primitive.ObjectIDFromHex(postId)

	if iderr != nil {
		apictx.JSON(http.StatusBadRequest, "Invalid Object ID")
	}

	err := postCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&result)

	if err != nil {
		println(err.Error())
		apictx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	res := map[string]interface{}{"data": result}

	apictx.JSON(http.StatusOK, res)
}
