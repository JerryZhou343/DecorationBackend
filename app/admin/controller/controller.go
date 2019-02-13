package controller

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.Engine) {
	admin := r.Group("/admin/v1")
	admin.Use(jwt.JWT())
}
