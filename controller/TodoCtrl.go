package controller

import (
	"GoAPI/appresponse"
	"GoAPI/initializer"
	"GoAPI/model"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"net/http"
	"time"
)

var db *gorm.DB

// khai bao db
func init() {
	db = initializer.ConnectMYSQL()
}

// api tim 1 item bang id
func GetItemById(ctx *gin.Context) {
	var data model.TodoItem
	id := ctx.Param("id")

	if err := ctx.ShouldBind(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, appresponse.SimpleSuccessRes(data))
}

// api post item
func CreateItem(ctx *gin.Context) {
	var data model.TodoItemCreate
	if err := ctx.ShouldBind(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	data.Id = uuid.NewV4().String()
	data.Status = "Doing"
	now := time.Now()
	data.CreatedAt = &now
	data.UpdatedAt = &now

	if err := db.Create(&data).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, appresponse.SimpleSuccessRes(data))
}

// api update item
func UpdateItem(ctx *gin.Context) {
	var data model.TodoItemUpdate
	id := ctx.Param("id")
	now := time.Now()
	data.UpdatedAt = &now
	if err := ctx.ShouldBind(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, appresponse.SimpleSuccessRes(data))
}

// api list item, hoac search item theo key (mac dinh title)
func ListItem(ctx *gin.Context) {
	var result []model.TodoItem

	query := ctx.Query("search") //Tao query de search, param truyen vao co key la search
	dbQuery := db.Where("title LIKE ?", "%"+query+"%").Find(&result)

	if dbQuery.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": dbQuery.Error.Error(),
		})
		return
	}
	if dbQuery.RowsAffected == 0 { //kiem tra query co tra ve hang nao khong
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "not found",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})
	//render template
	//ctx.HTML(http.StatusOK, "index.html", gin.H{
	//	"data": result,
	//})
}

// api sort delete item (doi status sang deleted)
func DeleteItem(ctx *gin.Context) {
	var data model.TodoItem
	id := ctx.Param("id")
	if err := ctx.ShouldBind(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := db.Table(model.TodoItem{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"Status": "Deleted"}).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, appresponse.SimpleSuccessRes("Changed status into Deleted"))
}
