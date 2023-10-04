package middleware

import (
	"GoAPI/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := helper.ValidateJWT(ctx)
		if err != nil {
			ctx.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Authentication required"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
