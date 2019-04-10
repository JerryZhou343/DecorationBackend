package controller

import "github.com/gin-gonic/gin"

//RegisterHTTPRouter 注册explore app  路由
func RegisterRouter(r *gin.Engine) {
	r.Group("/app/")
}
