package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sung1011/test/controller/game"
)

func InitGame(r *gin.Engine) {
	rg := r.Group("/game")
	{
		arena := &game.Arena{}
		rg.GET("/arena/fight", arena.Fight)
	}
}
