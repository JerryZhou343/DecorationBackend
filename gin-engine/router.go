package gin_engine

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/config"
)

func Init() *gin.Engine {
	r := gin.Default()
	if config.ReleaseFlag() {
		gin.SetMode(gin.ReleaseMode)
	}
	return r
}
