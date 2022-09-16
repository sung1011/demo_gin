package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Log(c *gin.Context) {
	// main函数阻塞, 所以不需要waitgroup等终止条件
	cCp := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("log::::", cCp.Request.URL.Path)
	}()
	c.Next()
}
