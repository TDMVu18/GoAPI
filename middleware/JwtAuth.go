package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, password, _ := ctx.Request.BasicAuth()

		if user == "user" && password == "password" {
			ctx.Next()
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
