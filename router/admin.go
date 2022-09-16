package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sung1011/test/controller/admin"
	"github.com/sung1011/test/controller/global"
)

func InitAdmin(r *gin.Engine) {
	rg := r.Group("/admin")
	{
		u := &global.User{}
		rg.GET("/user/add", u.Add)
		rg.GET("/user/del", u.Del)

		admin := admin.Notify{}
		rg.GET("/notify/email", admin.Email)
		rg.GET("/notify/sms", admin.SMS)
	}
}
