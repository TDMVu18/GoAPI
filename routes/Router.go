package routes

import (
	"GoAPI/controller"
	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

var currentImage *imageupload.Image

func CreateRouter(app *gin.Engine) {

	auth := app.Group("/auth")
	{
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
		auth.GET("", controller.Authen)
		auth.GET("/signup", controller.SignUp)
	}
	person := app.Group("/person")
	{
		//person.Use(middleware.AuthMiddleware())
		info := person.Group("/info")
		{
			info.POST("", controller.AddPerson)
			info.GET("/profile", controller.ShowProfile)
			info.GET("", controller.ListPerson)
			info.POST("/update", controller.UpdatePersonById)
			info.POST("/delete", controller.DeletePersonById)
			info.POST("/appearance", controller.ToggleAppearance)
			info.POST("/upload", controller.Upload)
		}
	}
}
