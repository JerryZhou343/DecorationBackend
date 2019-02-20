package controller

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {

	r.POST("/login", login)

	mgt := r.Group("/mgt/")
	//admin.Use(jwtauth.JWTAuth())
	{
		//单个分类
		mgt.GET("category/:id", selectCategory)
		mgt.POST("category/", createCategory)
		mgt.PUT("category/:id", updateCategory)
		mgt.DELETE("category/:id", delCategory)

		//分类簇
		mgt.GET("category_tree/:parent_id", selectCategoryTree)

		//图片
		mgt.POST("picture/", createPicture)
		mgt.PUT("picture/:id", updatePicture)
		mgt.GET("picture/:id", selectPicture)
		mgt.DELETE("picture/:id", delPicture)

		//案例
		mgt.POST("case/", createCase)
		mgt.PUT("case/:id", updateCase)
		mgt.GET("case/:id", selectCase)
		mgt.DELETE("case/:id", delCase)

	}

}
