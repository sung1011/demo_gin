package game

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Arena struct {
	Base
}

func (self *Arena) Fight(c *gin.Context) {
	uid, exist := c.Get("uid")
	if !exist {
		self.error(c, errors.New("uid not exists"))
		return
	}
	v, _ := uid.(string)
	fmt.Println("uid:", v)
	self.success(c)
}
