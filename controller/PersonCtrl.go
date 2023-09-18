package controller

import (
	"GoAPI/appresponse"
	"GoAPI/model"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// api tim 1 item bang id
func CreatePerson(ctx *gin.Context) {
	defer model.DisconnectDB()
	var person model.Person
	err := ctx.BindJSON(&person)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	coll := model.ConnectDB()
	result, err := coll.InsertOne(context.TODO(), person)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, appresponse.SimpleSuccessRes(result))
}
