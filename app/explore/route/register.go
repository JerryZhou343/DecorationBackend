package controller

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.Engine) {
	r.Group("/v1")
}
