package global

import "github.com/gin-gonic/gin"

type User struct {
	Base
}

func (self *User) Login(c *gin.Context) {
	self.success(c)
}

func (self *User) Add(c *gin.Context) {
	self.success(c)
}

func (self *User) Del(c *gin.Context) {
	self.error(c, nil)
}
func (self *User) Upd(c *gin.Context) {
	self.error(c, nil)
}
