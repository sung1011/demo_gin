package router

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/sung1011/test/model"
)

func InitStatc(r *gin.Engine) {
	// 自定义模板函数, 必须在加载模板LoadHTMLGlob之前调用
	r.SetFuncMap(template.FuncMap{
		"UnixToDate": func(ts int) string {
			return model.UnixToDate(int64(ts))
		},
	})
	// r.Delims("{[{", "}]}")
	r.LoadHTMLGlob("tpl/**/*")

	// static 静态文件的路由
	var path = "static"
	r.StaticFile("/favicon.ico", fmt.Sprintf("%s/favicon.ico", path))
	r.Static("static_css", fmt.Sprintf("%s/css", path))
	r.Static("static_upload", fmt.Sprintf("%s/upload", path))
	r.Static("static_pic", fmt.Sprintf("%s/pic", path))
}
