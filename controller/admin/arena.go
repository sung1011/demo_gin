package admin

import (
	"github.com/gin-gonic/gin"
)

type Notify struct {
	Base
}

func (self *Notify) Email(c *gin.Context) {
	self.success(c)
}
func (self *Notify) SMS(c *gin.Context) {
	self.success(c)
}
