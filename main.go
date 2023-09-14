package main

import (
	"GoAPI/modules/item/model"
	"GoAPI/modules/item/transport/ginitem"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	//ket noi voi bien moi truong
	_ = godotenv.Load(".env")

	//nguồn: https://gorm.io/docs/connecting_to_the_database.html
	dsn := os.Getenv("DB_CONNECT")
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//bat loi ket noi database
	if err != nil {
		log.Fatal(err)
	}

	//code mẫu của youtube {Lập trình REST API với Golang - Viet Tran}
	var item model.TodoItem

	r := gin.Default() //tạo 1 instance của gin.Engine - dùng để định nghĩa các router

	//CRUD
	api := r.Group("/go") //dinh tuyen group URL, xac dinh cac ham nao se xu ly route nao (.<method>("/url", <ham xu ly>)
	{
		item := api.Group("/item")
		{
			item.POST("", ginitem.CreateItem(db))
			//item.GET("", ListItem(db))
			//item.GET("/:id", GetItem(db))
			//item.PATCH("/:id", UpdateItem(db))
			//item.DELETE("/:id", DeleteItem(db))
		}
	}
	//đăng ký URL ping, hàm xử lý của router là func(c *gin.Context), truyền vào 1 con trỏ
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": item,
		})
	})
	r.GET("/about", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "This is information about ***",
		})
	})
	r.Run(":3000") //port ma app se chay tren do
}

//// get item
//func GetItem(db *gorm.DB) func(*gin.Context) {
//	return func(c *gin.Context) {
//		var data TodoItem
//		id := c.Param("id")
//
//		if err := c.ShouldBind(&data); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//		//data.Id = string(id)
//		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//		c.JSON(http.StatusOK, common.SimpleSuccessRes(data))
//	}
//}
//
//func UpdateItem(db *gorm.DB) func(*gin.Context) {
//	return func(c *gin.Context) {
//		var data TodoItemUpdate
//		id := c.Param("id")
//		now := time.Now()
//		data.UpdatedAt = &now
//		if err := c.ShouldBind(&data); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//
//		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//		c.JSON(http.StatusOK, gin.H{
//			"nofi": "Updated Successfully",
//		})
//	}
//}
//
//func DeleteItem(db *gorm.DB) func(*gin.Context) {
//	return func(c *gin.Context) {
//		var data TodoItem
//		id := c.Param("id")
//
//		if err := c.ShouldBind(&data); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//		//data.Id = string(id)
//		if err := db.Table(TodoItem{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
//			"Status": "Deleted", //nếu không ghi đúng dữ liệu enumerate thì database bị bỏ trống
//		}).Error; err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//		c.JSON(http.StatusOK, common.SimpleSuccessRes("Deleted"))
//	}
//}
//
//func ListItem(db *gorm.DB) func(*gin.Context) {
//	return func(c *gin.Context) {
//
//		var paging common.Paging //package common->paging
//
//		if err := c.ShouldBind(&paging); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//
//		paging.Process()
//
//		fmt.Println(paging)
//
//		var result []TodoItem
//
//		if err := db.Table(TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//		}
//
//		db = db.Where("status <> ?", "Deleted") //SQL query
//
//		//Order("id desc") - sort
//		if err := db.Order("updated_at desc").
//			Offset((paging.Page - 1) * paging.Limit).
//			Limit(paging.Limit).Find(&result).Error; err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//		}
//		c.JSON(http.StatusOK, common.NewSuccessRes(result, paging, nil))
//	}
//}
