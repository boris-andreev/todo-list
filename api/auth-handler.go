package api

import (
	_ "net/http"
	_ "strings"
	"todo-list/internal/service"

	"github.com/gin-gonic/gin"
)

func HandleAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/login" {
			c.Next()
			return
		}

		c.Set(service.UserIdKey, int32(1))

		/*
			tokenString := c.GetHeader("Authorization")
			if tokenString == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "No token provided"})
				c.Abort()

				return
			}

			tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")


			_, err := jwt.ParseToken(tokenString)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
				c.Abort()

				return
			}
		*/

		c.Next()
	}
}
