package controller

import (
	"GoAPI/helper"
	"GoAPI/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func Authen(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{
		"render": "render",
	})
}

func SignUp(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", gin.H{
		"render": "render",
	})
}

func Register(ctx *gin.Context) {
	var input model.AuthenInput

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"err":   "Authed",
		})
		return
	}
	// Kiểm tra xem tên người dùng đã tồn tại hay chưa
	existingUser, err := model.FindAccountByUserName(input.UserName)
	if err == nil && existingUser != (model.Account{}) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Tên người dùng đã tồn tại. Vui lòng chọn tên người dùng khác.",
		})
		return
	}
	account := model.Account{
		UserName: input.UserName,
		Password: input.Password,
	}
	account.ID = primitive.NewObjectID()
	account.PreProcess()
	savedUser, err := account.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"err":   "savedUser",
		})
		return
	}
	fmt.Println(savedUser)
	ctx.Redirect(http.StatusFound, "/auth")
}

func Login(ctx *gin.Context) {
	var input model.AuthenInput

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	entry, err := model.FindAccountByUserName(input.UserName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found",
		})
		return
	}
	err = entry.ValidatePassword(input.Password)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Password not correct",
		})
		return
	}
	token, err := helper.CreateJwt(entry)
	ctx.Header("Authorization", "Bearer "+token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Can't create JWT",
		})
	}
	ctx.Redirect(http.StatusFound, "/person/info")
}
