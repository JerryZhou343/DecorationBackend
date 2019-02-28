package engine

import (
	"github.com/gin-gonic/gin"
	ar "github.com/mfslog/DecorationBackend/app/admin/route"
	er "github.com/mfslog/DecorationBackend/app/explore/route"
	"github.com/mfslog/DecorationBackend/config"
	"net/http"
)

//Init 初始化go-gin 路由,返回 httpRouter, httpsRouter
func Init() (httpRouter *gin.Engine, httpsRouter *gin.Engine) {
	httpRouter = gin.Default()
	httpsRouter = gin.Default()
	if config.ReleaseFlag() {
		gin.SetMode(gin.ReleaseMode)
	}
	//图片URL
	httpRouter.StaticFS(config.PicURLRelativePath(), http.Dir(config.GetPicPath()))
	er.RegisterHTTPRouter(httpRouter)

	ar.RegisterHTTPRouter(httpRouter)
	ar.RegisterHTTPSRouter(httpsRouter)

	return httpRouter, httpsRouter
}
