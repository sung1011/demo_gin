package main

import (
	"net/http"

	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
	"github.com/sung1011/test/middleware"
	"github.com/sung1011/test/model"
	"github.com/sung1011/test/router"
)

func main() {
	go func() {
		http.ListenAndServe(":8089", nil) // for pprof; http://localhost:8089/debug/pprof/
	}()
	model.Log()
	r := gin.New()
	r.MaxMultipartMemory = 8 << 10
	r.Use(
		gin.Logger(),
		gin.Recovery(),
		middleware.BanIP,
		middleware.Session,
		// middleware.Cookie,
	)

	router.InitDefault(r)
	router.InitStatc(r)
	router.InitGlobal(r)
	router.InitGame(r)
	router.InitAdmin(r)

	router.InitDemo(r)

	r.Run(":8088")
}
