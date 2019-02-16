package gin_engine

import (
	"github.com/gin-gonic/gin"
	ac "github.com/mfslog/DecorationBackend/app/admin/controller"
	ec "github.com/mfslog/DecorationBackend/app/explore/controller"
	"github.com/mfslog/DecorationBackend/config"
)

func Init() *gin.Engine {
	r := gin.Default()
	if config.ReleaseFlag() {
		gin.SetMode(gin.ReleaseMode)
	}

	ec.RegisterRouter(r)
	ac.RegisterRouter(r)

	return r
}
