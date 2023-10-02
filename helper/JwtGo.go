package helper

import (
	"GoAPI/model"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"strings"
	"time"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func CreateJwt(account model.Account) (string, error) {
	//tạo JWT cho user với id tương ứng
	privateKey := []byte(os.Getenv("JWT_PRIVATE_KEY"))
	tokenTLL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  account.ID,
		"iat": time.Now().Unix(),                                            //thời gian code được tạo
		"eat": time.Now().Add(time.Second * time.Duration(tokenTLL)).Unix(), //thời gian code expired
	})
	return token.SignedString(privateKey)
}
func ValidateJWT(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	if err != nil {
		return err
	}

	if token.Valid {
		return nil
	}

	return errors.New("invalid token provided")
}

//func getToken(ctx *gin.Context) (*jwt.Token, error) {
//	tokenString := getTokenFromRequest(ctx)
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//		}
//
//		return privateKey, nil
//	})
//	return token, err
//}

func GetTokenFromRequest(ctx *gin.Context) string {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader != "" {
		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) == 2 {
			return splitToken[1]
		}
	}
	fmt.Println("Header:", authHeader)
	return ""
}
