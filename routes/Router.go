package routes

import (
	"GoAPI/controller"
	"github.com/gin-gonic/gin"
)

func CreateRouter(app *gin.Engine) {
	api := app.Group("/go")
	{
		item := api.Group("/item")
		{
			item.POST("", controller.CreateItem)       //Create Item
			item.GET("", controller.ListItem)          //List Item (Search Item)
			item.GET("/:id", controller.GetItemById)   //Get Item By Id
			item.PATCH("/:id", controller.UpdateItem)  //Update Item
			item.DELETE("/:id", controller.DeleteItem) //Delete Item
		}
	}
}
