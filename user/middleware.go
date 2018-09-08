package user

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-api-todo/common"
	"net/http"
	"reflect"
	"strings"
)

func JWTAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth_token := c.Request.Header.Get("Authorization")
		if auth_token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Authorization header not provided"})
			c.Abort()
			return
		}
		token_header := strings.Split(auth_token, " ")
		if len(token_header) != 2 {
			c.JSON(http.StatusForbidden, gin.H{"message": "Usage: JWT <token_header>"})
			c.Abort()
			return
		}
		if token_header[0] != "JWT" {
			c.JSON(http.StatusForbidden, gin.H{"message": "Authorization header need to start with 'JWT'"})
			c.Abort()
			return
		}
		token, err := jwt.Parse(token_header[1], common.KeyFunc)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"message": "Token is invalid"})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"message": "Token is invalid"})
			c.Abort()
			return
		}
		fmt.Println(claims["id"], reflect.TypeOf(claims["id"]))
		username := claims["id"].(string)
		user, err := GetUserByUsername(username)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"message": "User not found"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
