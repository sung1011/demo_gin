package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sung1011/test/controller/global"
)

func InitGlobal(r *gin.Engine) {
	rg := r.Group("/global")
	{
		u := &global.User{}
		rg.GET("/user/login", u.Login)
		rg.GET("/user/add", u.Add)
		rg.GET("/user/del", u.Del)
		rg.GET("/user/upd", u.Add, u.Upd) // 可以多个, 也可以抽离出来(即中间件)
	}
}
