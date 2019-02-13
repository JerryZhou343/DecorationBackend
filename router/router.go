package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/config"
)

func Init() *gin.Engine {
	r := gin.Default()
	if config.ReleaseFlag() {
		gin.SetMode(gin.ReleaseMode)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}
