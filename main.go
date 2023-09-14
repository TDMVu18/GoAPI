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
	"time"
)

type ItemStatus int

const (
	ItemStatusDoing ItemStatus = iota
	ItemStatusDone
	ItemStatusDeleted
)

//func (item *ItemStatus) Scan(value interface{}) error {
//
//}

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
	var item TodoItem

	r := gin.Default() //tạo 1 instance của gin.Engine - dùng để định nghĩa các router

	//CRUD

	api := r.Group("/go") //dinh tuyen group URL, xac dinh cac ham nao se xu ly route nao (.<method>("/url", <ham xu ly>)

	{
		item := api.Group("/item")
		{
			item.POST("", CreateItem(db))
			item.GET("", ListItem(db))
			item.GET("/:id", GetItem(db))
			item.PATCH("/:id", UpdateItem(db))
			item.DELETE("/:id", DeleteItem(db))
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

// tao struct TodoItem va xac dinh truong tuong ung trong database bang gorm
type TodoItem struct {
	Id          string     `json:"id" gorm:"column:id;"`
	Title       string     `json:"title" gorm:"column:title;"`
	Description string     `json:"description" gorm:"column:description;"`
	Status      string     `json:"status" gorm:"column:status;"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

// tao method TableName cho Struct TodoItem, tra ve ten bang "todo_items"
func (TodoItem) TableName() string { return "todo_items" }

// tao struct TodoItemCreation va xac dinh truong tuong ung trong database bang gorm
type TodoItemCreation struct {
	Id          string    `json:"id" gorm:"column:id;"`
	Title       string    `json:"title" gorm:"column:title;"`
	Description string    `json:"description" gorm:"column:description;"`
	Status      string    `json:"status" gorm:"column:status;"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// tao method TableName cho Struct TodoItemCreation, tra ve ten bang "todo_items"
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
		id := c.Param("id")

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		//data.Id = string(id)
		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

// tao struct TodoItemUpdate va xac dinh truong tuong ung trong database bang gorm
type TodoItemUpdate struct {
	Id          *string   `json:"id" gorm:"column:id;"` //Dùng con trỏ, xử lý trường hợp truyền vào chuỗi rỗng thì bị bỏ qua
	Title       *string   `json:"title" gorm:"column:title;"`
	Description *string   `json:"description" gorm:"column:description;"`
	Status      *string   `json:"status" gorm:"column:status;"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// tao method TableName cho Struct TodoItemUpdate, tra ve ten bang "todo_items"
func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }

// tao struct Paging va lay du lieu tu form (POSTMAN), dung de phan trang
type Paging struct {
	Page  int   `json:"page" form:"page"` //parse từ form
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

func (p *Paging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 || p.Limit >= 100 {
		p.Limit = 10
	}
}

func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data TodoItemUpdate
		id := c.Param("id")
		data.UpdatedAt = time.Now()
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		c.JSON(http.StatusOK, gin.H{
			"nofi": "Updated Successfully",
		})
	}
}

func DeleteItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data TodoItem
		id := c.Param("id")

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		//data.Id = string(id)
		if err := db.Table(TodoItem{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
			"Status": "Deleted", //nếu không ghi đúng dữ liệu enumerate thì database bị bỏ trống
		}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		c.JSON(http.StatusOK, gin.H{
			"nofi": "Deleted Successfully",
		})
	}
}

func ListItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		var paging Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Process()

		fmt.Println(paging)

		var result []TodoItem

		if err := db.Table(TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		db = db.Where("status <> ?", "Deleted") //SQL query

		//Order("id desc") - sort
		if err := db.Order("title desc").
			Offset((paging.Page - 1) * paging.Limit).
			Limit(paging.Limit).Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"all item": result,
			"paging":   paging,
		})
	}
}
