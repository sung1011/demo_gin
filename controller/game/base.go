package game

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Base struct {
}

func (self *Base) success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
		"_t":  time.Now().Format("2006-01-02 15:04:05"),
	})
}
func (self *Base) error(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{
		"err": err.Error(),
		"_t":  time.Now().Format("2006-01-02 15:04:05"),
	})
}
