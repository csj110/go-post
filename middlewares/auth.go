package middlewares

import (

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	Auth   int    `header:"Authorization"`
}

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
