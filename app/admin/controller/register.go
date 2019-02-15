package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/middleware/jwtauth"
)

func RegisterRouter(r *gin.Engine) {

	r.POST("/login", login)

	admin := r.Group("/admin")
	admin.Use(jwtauth.JWTAuth())
	{
		admin.GET("/tag", queryTag)
		admin.GET("/tag_tree", queryTagTree)
	}

}
