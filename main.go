package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sung1011/test/middleware"
	"github.com/sung1011/test/router"
)

func main() {
	// r := gin.Default()
	r := gin.New()
	r.Use(
		gin.Logger(),
		gin.Recovery(),
		middleware.BanIP,
		middleware.Session,
	)

	router.InitDefault(r)
	router.InitStatc(r)
	router.InitGlobal(r)
	router.InitGame(r)
	router.InitAdmin(r)

	router.InitDemo(r)

	r.Run(":8088")
}
