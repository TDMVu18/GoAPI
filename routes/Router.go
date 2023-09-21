package routes

import (
	"GoAPI/controller"
	"github.com/gin-gonic/gin"
)

func CreateRouter(app *gin.Engine) {
	todo := app.Group("/todo")
	{
		item := todo.Group("/item")
		{
			item.POST("", controller.CreateItem)       //Create Item
			item.GET("", controller.ListItem)          //List Item (Search Item)
			item.GET("/:id", controller.GetItemById)   //Get Item By Id
			item.PATCH("/:id", controller.UpdateItem)  //Update Item
			item.DELETE("/:id", controller.DeleteItem) //Delete Item
		}
	}
	person := app.Group("/person")
	{
		info := person.Group("/info")
		{
			info.POST("", controller.AddPerson)
			info.GET("", controller.ListPerson)
			info.GET("/:id", controller.GetPersonById)
			info.POST("/update", controller.UpdatePersonById)
			info.POST("/delete", controller.DeletePersonById)
		}
	}
}
