package route

import (
	"github.com/gin-gonic/gin"
	. "github.com/mfslog/DecorationBackend/app/admin/controller"
	"github.com/mfslog/DecorationBackend/middleware/jwtauth"
)

//RegisterHTTPRouter 注册admin app HTTP 路由
func RegisterHTTPRouter(r *gin.Engine) {

	mgt := r.Group("/mgt/")
	mgt.Use(jwtauth.JWTAuth())
	{
		//单个分类
		mgt.GET("category/:id", GetCategory)
		mgt.POST("category/", CreateCategory)
		mgt.PUT("category/:id", UpdateCategory)
		mgt.DELETE("category/:id", DelCategory)

		//分类簇
		mgt.GET("category_tree/:parent_id", GetCategoryTree)

		//图片
		mgt.POST("picture/", CreatePicture)
		mgt.PUT("picture/:id", UpdatePicture)
		mgt.GET("picture/:id", GetPicture)
		mgt.DELETE("picture/:id", DelPicture)

		//图片分类
		mgt.POST("pic_category/", AddPicCategory)
		mgt.GET("pic_category/:id", GetPicCategory)
		mgt.DELETE("pic_category/:id", DelPicCategory)

		//案例
		mgt.POST("case/", CreateCase)
		mgt.PUT("case/:id", UpdateCaseInfo)
		mgt.GET("case/:id", GetCase)
		mgt.DELETE("case/:id", DelCase)

		//案例分类
		mgt.POST("case_category/:id", AddCaseCategory)
		mgt.DELETE("case_category/:id", DelCaseCategory)
		mgt.GET("case_category/:id", GetCaseCategory)
	}

}

//RegisterHTTPSRouter 注册 admin app HTTPS 路由
func RegisterHTTPSRouter(r *gin.Engine) {

	r.POST("/login", Login)
	r.PUT("/changePassword", UpdatePwd)
}
