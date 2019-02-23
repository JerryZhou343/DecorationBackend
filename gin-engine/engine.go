package gin_engine

import (
	"github.com/gin-gonic/gin"
	ar "github.com/mfslog/DecorationBackend/app/admin/route"
	er "github.com/mfslog/DecorationBackend/app/explore/route"
	"github.com/mfslog/DecorationBackend/config"
)

func Init() *gin.Engine {
	r := gin.Default()
	if config.ReleaseFlag() {
		gin.SetMode(gin.ReleaseMode)
	}

	er.RegisterRouter(r)
	ar.RegisterRouter(r)

	return r
}
