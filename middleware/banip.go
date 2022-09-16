package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func BanIP(c *gin.Context) {
	if "127.0.0.1" != c.ClientIP() {
		fmt.Println("ban ip:", c.ClientIP())
		c.Abort()
	}
	c.Next()
}
