package routes

import (
	"GoAPI/controller"
	"GoAPI/middleware"
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
		person.Use(middleware.JWTAuthMiddleware())
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
		salary := person.Group("/salary")
		{
			salary.POST("", controller.SalaryAdd)
			salary.GET("", controller.ListSalary)
			salary.POST("/delete", controller.DeleteSalaryById)
			salary.POST("/update", controller.UpdateSalaryById)
			salary.GET("/level", controller.GetSalaryLevels)
		}
		office := person.Group("/office")
		{
			office.POST("", controller.OfficeAdd)
			office.GET("", controller.ListOffice)
			office.POST("/delete", controller.DeleteOfficeById)
			office.POST("/update", controller.UpdateOfficeById)
			office.GET("/name", controller.GetOfficeName)
		}
	}
}
