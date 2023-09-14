package ginitem

import (
	"GoAPI/common"
	"GoAPI/modules/item/business"
	"GoAPI/modules/item/model"
	"GoAPI/modules/item/storage"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemCreation
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		data.Id = uuid.NewV4().String()
		data.Status = "Doing"
		now := time.Now()
		data.CreatedAt = &now
		data.UpdatedAt = &now

		store := storage.NewSQLStore(db)

		bs := business.NewCreateItemBiz(store)

		if err := bs.CreateNewItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(data.Id))
	}
}
