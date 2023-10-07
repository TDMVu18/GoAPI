package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		isLoggedIn := session.Get("isLoggedIn")

		if isLoggedIn != true {
			ctx.Redirect(http.StatusFound, "/login")
			ctx.Abort()
		} else {
			ctx.Next()
		}
	}
}
