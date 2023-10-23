package middleware

import (
	"LotusPart2/conf"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.GetHeader("Authorization")

		tokenStr := strings.Split(bearToken, " ")
		if len(tokenStr) != 2 {
			c.JSON(http.StatusForbidden, "Missing or invalid  Token")
			return
		}

		token := tokenStr[1]

		parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(conf.LoadEnv().API_SECRET), nil
		})
		if err != nil {
			c.JSON(http.StatusForbidden, "Invalid Token")
			c.Abort()
			return
		}

		claims, ok := parseToken.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
		if ok && parseToken.Valid {
			c.Set("authId", claims["authId"])
			c.Set("userId", claims["userId"])
		}

		c.Next()
	}
}
