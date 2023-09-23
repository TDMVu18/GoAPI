package main

import (
	"GoAPI/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"html/template"
	"time"
)

func add(num, page int) int {
	return num + 6*(page-1) + 1
}

func FormatTimestamp(timestamp primitive.DateTime) string {
	// Chuyển đổi primitive.DateTime thành time.Time
	t := time.Unix(int64(timestamp)/1000, 0) // Chia cho 1000 để chuyển đổi thành giây

	// Định dạng thời gian sang múi giờ của bạn
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	formattedTime := t.In(loc).Format("2006-01-02 15:04:05.999")

	return formattedTime
}
func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"add":             add,
		"FormatTimestamp": FormatTimestamp,
	})
	routes.CreateRouter(r)
	r.LoadHTMLGlob("templates/*")
	r.Run(":3000")

}
