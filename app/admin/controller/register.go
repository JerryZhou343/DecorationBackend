package controller

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {

	r.POST("/login", login)

	mgt := r.Group("/mgt/")
	//admin.Use(jwtauth.JWTAuth())
	{
		mgt.GET("category/id", queryCategory)
		mgt.POST("category/", addCategory)
		mgt.PUT("category/id", modifyCategory)
		mgt.DELETE("category/id", delCategory)

		mgt.GET("/category_tree/:parent_id", queryCategoryTree)

	}

}
