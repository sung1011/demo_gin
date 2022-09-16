package router

import (
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
	r.Static("static_css", "./static/css")
	r.Static("static_upload", "./static/upload")
	r.Static("static_pic", "./static/pic")
}
