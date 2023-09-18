package controller

import (
	"GoAPI/appresponse"
	"GoAPI/model"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Mgr model.ManageDB

// api tim 1 item bang id
func CreatePerson(ctx *gin.Context) {
	defer model.DisconnectDB()
	var dp model.Person
	err := ctx.BindJSON(&dp)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	Mgr.Insert(dp)
	ctx.JSON(http.StatusOK, appresponse.SimpleSuccessRes(dp))
}

func Create(data interface{}) error {
	db := model.ConnectDB()
	_, err := db.InsertOne(context.TODO(), data)
	return err
}
