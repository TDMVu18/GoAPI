package middleware

import (
	"GoAPI/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func AuthMiddleware() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//
//		tokenString, exists := ctx.Get("Authorization")
//		if !exists || tokenString == nil {
//			ctx.Redirect(http.StatusFound, "/auth")
//			ctx.Abort()
//			return
//		}
//
//		err := helper.ValidateJWT(ctx)
//		if err != nil {
//			ctx.Redirect(http.StatusFound, "/auth")
//			ctx.Abort()
//			return
//		}
//		ctx.Next()
//	}
//}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := helper.GetTokenFromRequest(ctx)
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			ctx.Abort()
			return
		}

		err := helper.ValidateJWT(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
