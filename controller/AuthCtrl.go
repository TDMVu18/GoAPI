package controller

import (
	"GoAPI/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func Auth(ctx *gin.Context) {
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
	var input model.AuthInput

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Xử lý request từ form đăng ký không thành công",
			"error":   err.Error(),
		})
		return
	}
	// Kiểm tra username đã tồn tại hay chưa
	existingUser, err := model.FindAccountByUserName(input.UserName)
	if err == nil && existingUser != (model.Account{}) {
		ctx.JSON(http.StatusConflict, gin.H{
			"code":    http.StatusConflict,
			"message": "Tên người dùng đã tồn tại. Vui lòng chọn tên người dùng khác.",
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
			"code":    http.StatusBadRequest,
			"message": "Lưu thông tin đăng ký không thành công",
			"error":   err.Error(),
		})
		return
	}
	fmt.Println(savedUser)
	ctx.Redirect(http.StatusFound, "/auth")
}

func Login(ctx *gin.Context) {
	var input model.AuthInput

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Xử lý request từ form đăng nhập không thành công",
			"error":   err.Error(),
		})
		return
	}
	entry, err := model.FindAccountByUserName(input.UserName)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Không tìm thấy người dùng",
			"error":   err.Error(),
		})
		return
	}
	err = entry.ValidatePassword(input.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Sai mật khẩu, vui lòng kiểm tra lại",
			"error":   err.Error(),
		})
		return
	}

}
