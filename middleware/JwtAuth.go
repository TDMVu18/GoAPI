package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		fmt.Println(authHeader)
		if authHeader == "" {
			authToken := ctx.DefaultQuery("authToken", "")
			// Sử dụng authToken từ Local Storage
			tokenString := authToken
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Phương thức không hợp lệ: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("JWT_PRIVATE_KEY")), nil
			})
			if err != nil {
				ctx.Redirect(http.StatusFound, "/auth")
				ctx.Abort()
				return
			}
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx.Set("userID", claims["id"])
				ctx.Next()
			} else {
				fmt.Println("Token không hợp lệ")
				ctx.Redirect(http.StatusFound, "/auth")
				ctx.Abort()
				return
			}
			ctx.Redirect(http.StatusOK, "/person/info")
		}
	}
}
