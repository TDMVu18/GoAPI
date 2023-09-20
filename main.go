package main

import (
	"GoAPI/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.CreateRouter(r)
	r.LoadHTMLGlob("templates/*")
	r.Run(":3000")
}
