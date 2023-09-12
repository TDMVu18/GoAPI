package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	//ket noi voi bien moi truong
	_ = godotenv.Load(".env")

	//nguon: https://gorm.io/docs/connecting_to_the_database.html
	dsn := os.Getenv("DB_CONNECT")
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)

	now := time.Now().UTC()

	//code mẫu của youtube {Lập trình REST API với Golang - Viet Tran}
	item := TodoItem{
		Id:          "",
		Title:       "This is Item 1",
		Description: "Description of Item 1",
		Status:      "Doing	",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	r := gin.Default() //tạo 1 instance của gin.Engine - dùng để định nghĩa các router

	//CRUD

	v1 := r.Group("/v1")

	{
		item := v1.Group("/item")
		{
			item.POST("", CreateItem(db))
			item.GET("")
			item.GET("/:id")
			item.PATCH("/:id")
			item.DELETE("/:id")
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
	r.Run(":3000")

}

type TodoItem struct {
	Id          string     `json:"id" gorm:"column:id;"`
	Title       string     `json:"title" gorm:"column:title;"`
	Description string     `json:"description" gorm:"column:description;"`
	Status      string     `json:"status" gorm:"column:status;"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at"`
}

type TodoItemCreation struct {
	Id          string    `json:"id" gorm:"column:id;"`
	Title       string    `json:"title" gorm:"column:title;"`
	Description string    `json:"description" gorm:"column:description;"`
	Status      string    `json:"status" gorm:"column:status;"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (TodoItem) TableName() string { return "todo_items" }

func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data TodoItemCreation
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		data.Id = uuid.NewV4().String()
		data.Status = "Doing"
		data.CreatedAt = time.Now()
		data.UpdatedAt = time.Now()
		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"nofi": "Created Successfully",
		})
	}
}

// get item từ db
func GetItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data TodoItem
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		data.Id = string(id)
		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data.Id,
		})
	}
}
