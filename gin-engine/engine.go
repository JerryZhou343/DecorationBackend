package engine

import (
	"github.com/gin-gonic/gin"
	ar "github.com/mfslog/DecorationBackend/app/admin/route"
	er "github.com/mfslog/DecorationBackend/app/explore/route"
	"github.com/mfslog/DecorationBackend/config"
	"net/http"
)

//Init 初始化go-gin 路由,返回 appRouter, cmsRouter
func Init() (appRouter *gin.Engine, cmsRouter *gin.Engine) {
	appRouter = gin.Default()
	cmsRouter = gin.Default()
	if config.ReleaseFlag() {
		gin.SetMode(gin.ReleaseMode)
	}
	//图片URL
	appRouter.StaticFS(config.PicURLRelativePath(), http.Dir(config.GetPicPath()))

	cmsRouter.StaticFS(config.PicURLRelativePath(), http.Dir(config.GetPicPath()))

	er.RegisterRouter(appRouter)

	ar.RegisterRouter(cmsRouter)

	return appRouter, cmsRouter
}
