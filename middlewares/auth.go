package middlewares

import (
	"blogos/auth"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerString := c.Request.Header.Get("Authorization")
		var tokenString string
		if len(strings.Split(bearerString, " ")) == 2 {
			tokenString = strings.Split(bearerString, " ")[1]
			user_id, err := auth.ValidatToken(tokenString)
			if err != nil {
				c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "no authorized",
				})
				c.Abort()
			}
			c.Set("user_id", user_id)
			fmt.Println(user_id,"sag")
			c.Next()
		} else {
			c.JSON(http.StatusBadRequest, map[string]string{
				"message": "bad auth request",
			})
			c.Abort()
		}
	}
}
