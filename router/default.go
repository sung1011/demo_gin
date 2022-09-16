package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sung1011/test/middleware"
)

func InitDefault(r *gin.Engine) {
	rg := r.Group("/", middleware.Log)
	{
		rg.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "hello world",
				"time": time.Now().Format("2006-01-02 15:04:05"),
			})
		})
		rg.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
	}
}
