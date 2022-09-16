package middleware

import (
	"github.com/gin-gonic/gin"
)

func Session(c *gin.Context) {
	c.Set("uid", "100001")
	c.Next()
}
