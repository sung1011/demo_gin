package router

import (
	"bytes"
	"crypto/md5"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func InitDemo(r *gin.Engine) {
	rg := r.Group("/demo")
	demoRouter(rg)
	demoUpload(rg)
	demoMiddleWare(rg)
	demoInput(rg)
	demoRender(rg)
}

func demoRouter(rg *gin.RouterGroup) {
	rgRouter := rg.Group("/router")
	{
		// curl -sL 127.0.0.1:8088/demo/router/match1/wifjw
		rgRouter.GET("/match1/:category", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
				"p1":  c.Params.ByName("category"),
			})
		})
		// curl -sL 127.0.0.1:8088/demo/router/match2/wifjw/abc
		rgRouter.GET("/match2/:category/:post", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
				"p1":  c.Params.ByName("category"),
				"p2":  c.Param("post"),
			})
		})
		// curl -sL 127.0.0.1:8088/demo/router/matchall/wifjw/abc
		rgRouter.GET("/matchall/*name", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
				"p1":  c.Params.ByName("name"),
			})
		})
	}
}

func demoUpload(rg *gin.RouterGroup) {
	rgUpload := rg.Group("/upload")
	{
		rgUpload.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "admin/upload.html", gin.H{})
		})
		rgUpload.POST("/doUpload", func(c *gin.Context) {
			// 获取上传的文件
			fileHeader, err := c.FormFile("avatar")
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"err": err})
				return
			}

			// 组装文件名 {uid}_{md5sum(*file)}
			// TODO 后缀没处理
			h := md5.New()
			f, _ := fileHeader.Open()
			defer f.Close()
			if _, err := io.Copy(h, f); err != nil {
				log.Fatal(err)
			}
			md5value := fmt.Sprintf("%x", h.Sum(nil))
			var buf bytes.Buffer
			buf.WriteString("./static/upload/")
			buf.WriteString(c.GetString("uid"))
			buf.WriteString("_")
			buf.WriteString(md5value)
			uploadPath := buf.String()
			if err = c.SaveUploadedFile(fileHeader, uploadPath); err != nil {
				c.JSON(http.StatusOK, gin.H{"err": err})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}
}

func demoMiddleWare(rg *gin.RouterGroup) {
	// 含有c.Next()的 handlerFunc 是中间件
	rgMw := rg.Group("/mw")
	{
		// 多个handlerFunc
		rgMw.GET("/seq",
			func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"foo": 123,
				})
			},
			func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"bar": 456,
				})
			},
		)

		// 中间件 c.Next()执行其他handler
		mw1 := func(c *gin.Context) {
			fmt.Println("123")
			c.Next() // 相当于其他handler
			fmt.Println("789")
		}
		mw2 := func(c *gin.Context) {
			fmt.Println("abc")
			c.Next() // 相当于忽略其他handler; index直接移到handlers结尾之后(超大数)
			fmt.Println("xyz")
		}
		handler := func(c *gin.Context) {
			fmt.Println("~~~")
			c.String(http.StatusOK, "%s", "关注log")
		}
		// log: 123 ~~~ 789
		rgMw.GET("/mw", mw1, handler) // 顺序不能反, 所以不如用use注册

		// 中间件 c.Adort跳过后续handler, 只运行所有mw (含有c.Next()的handler)
		// log: 123 789
		mwAbort := func(c *gin.Context) {
			fmt.Println("111")
			c.Abort() // 相当于忽略其他handler; index直接移到handlers结尾之后(超大数)
			fmt.Println("222")
		}

		rgMw.GET("mwabort", mw1, mwAbort, handler) // 123 111 222 789
		// rgMw.GET("mwabort", mw1, handler, mwAbort) // 123 ~~~ 111 222 789; c.Abort之前的会执行

		// 多个中间件
		rgMw.GET("mws", mw1, mw2, handler) // 123 abc ~~~ xyz 789

		// use组中间件
		rgMw.Use(mw1, mw2) // 方式1

		// rgMw := rg.Group("/mw", mw1, mw2) // 方式2
		rgMw.GET("mwgroup", handler) // 123 abc ~~~ xyz 789
	}

}

func demoInput(rg *gin.RouterGroup) {
	rgInput := rg.Group("/input")
	{
		/*
			POST {{url}}?uid=10001
			Content-Type: application/x-www-form-urlencoded

			name=sunji&group=c
		*/
		/*
			action=/demo/input/dopost的表单
		*/
		// post 解析参数
		rgInput.POST("/dopost", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"uid":   c.Query("uid"),
				"page":  c.DefaultQuery("page", "0"),
				"name":  c.PostForm("name"),
				"group": c.DefaultPostForm("group", "a"),
			})
		})
		type Account struct {
			Usname string `form:"un" xml:"usn" json:"username"`
			Passwd string `form:"pd" xml:"pwd" json:"password"`
		}
		// post 绑定到结构体 form关联参数, json关联输出
		rgInput.POST("/postbind", func(c *gin.Context) {
			d := &Account{}
			err := c.ShouldBind(d)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"err": err, // 参数类型错误
				})
			} else {
				c.JSON(http.StatusOK, d)
			}
		})
		/*
			POST {{url}}
			Content-Type: application/xml

			<request>
				<usn>sample</usn>
				<pwd>Wed, 21 Oct 2015 18:27:50 GMT</pwd>
			</request>
		*/
		rgInput.POST("/xmlbind", func(c *gin.Context) {
			d := &Account{}
			bs, _ := c.GetRawData()
			err := xml.Unmarshal(bs, d)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"err": err,
				})
			} else {
				c.JSON(http.StatusOK, d)
			}
		})
	}
}

func demoRender(rg *gin.RouterGroup) {
	rgRender := rg.Group("/route")
	{
		// json
		rgRender.GET("/json", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "json",
				"zh":  "你好",
			})
		})
		// json 结构体
		rgRender.GET("/jsons", func(c *gin.Context) {
			c.JSON(http.StatusOK, struct {
				Msg string `json:"msg"`
			}{Msg: "json struct"})
		})
		// jsonp
		// curl -sL "127.0.0.1:8088/jsonp?callback=xxx"
		// xxx({"msg": "jsonp"})
		rgRender.GET("/jsonp", func(c *gin.Context) {
			c.JSONP(http.StatusOK, gin.H{
				"msg": "jsonp",
			})
		})
		// xml
		rgRender.GET("/xml", func(c *gin.Context) {
			c.XML(http.StatusOK, gin.H{
				"msg": "xml",
			})
		})
		// string
		rgRender.GET("/string", func(c *gin.Context) {
			c.String(http.StatusOK, "msg: %v", "string")
		})
		// yaml
		rgRender.GET("/yaml", func(c *gin.Context) {
			c.YAML(http.StatusOK, gin.H{
				"message": "yaml",
				"status":  http.StatusOK,
			})
		})
		// protobuf
		rgRender.GET("/protobuf", func(c *gin.Context) {
			reps := []int64{int64(1), int64(2), int64(3)}
			label := "test protobuf"
			data := &protoexample.Test{
				Label: &label,
				Reps:  reps,
			}
			// 请注意，数据在响应中变为二进制数据
			// 将输出被 protoexample.Test protobuf 序列化了的数据
			c.ProtoBuf(http.StatusOK, data)
		})
		// toml
		rgRender.GET("/toml", func(c *gin.Context) {
			c.TOML(http.StatusOK, gin.H{"msg": "toml"})
		})
		// pure json
		rgRender.GET("/purejson", func(c *gin.Context) {
			c.PureJSON(http.StatusOK, gin.H{"msg": "pure json", "zh": "你好"})
		})
		// indented json
		rgRender.GET("/indentedjson", func(c *gin.Context) {
			c.IndentedJSON(http.StatusOK, gin.H{"msg": "indented json"})
		})
		// secure json
		rgRender.GET("/securejson", func(c *gin.Context) {
			c.SecureJSON(http.StatusOK, gin.H{"msg": "secure json"})
		})
		// ascii json
		rgRender.GET("/asciijson", func(c *gin.Context) {
			c.AsciiJSON(http.StatusOK, gin.H{"msg": "ascii json", "zh": "你好"})
		})
		// html
		type Foo struct {
			Name string
		}
		var fs []*Foo
		fs = append(fs, &Foo{Name: "a"}, &Foo{Name: "b"}, &Foo{Name: "c"})
		rgRender.GET("/htmlindex", func(c *gin.Context) {
			c.HTML(http.StatusOK, "default/index.html", gin.H{
				"title": "Main website",
				"t0":    "var",
				"t1":    "assignment",
				"t2":    fs,      // range
				"t3":    []int{}, // range empty
				"t4": struct {
					Name string
					Age  int
					Sex  int
				}{Name: "sun", Age: 21, Sex: 2}, // with
				"t5": 89,                     // if else
				"t6": int(time.Now().Unix()), // 自定义函数
			})
		})
		rgRender.GET("/htmladmin", func(c *gin.Context) {
			c.HTML(http.StatusOK, "admin/index.html", gin.H{
				"title": "Admin website",
				"h1":    "hello administrator",
			})
		})
	}
}
