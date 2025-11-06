package api

import (
	"net/http"
	"strings"
	"todo-list/internal/service"
	"todo-list/internal/utils/jwt"

	"github.com/gin-gonic/gin"
)

func HandleAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/login" {
			c.Next()
			return
		}

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "No token provided"})
			c.Abort()

			return
		}

		token, err := jwt.ParseToken(strings.ReplaceAll(tokenString, "Bearer ", ""))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error(), "AAAA": "AAAA"})
			c.Abort()

			return
		}

		c.Set(service.UserIdKey, token.UserId)

		c.Next()
	}
}
