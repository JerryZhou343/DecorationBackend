package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/app/admin/form"
	"github.com/mfslog/DecorationBackend/app/admin/models"
	"net/http"
)

//TODO:TLS
func login(c *gin.Context) {
	info := form.Login{}
	c.BindJSON(&info)
	if info.UserName != "" && info.Password != "" {
		if models.CheckPassport(info.UserName, info.Password) {
			c.Status(http.StatusOK)
		} else {
			c.Status(http.StatusUnauthorized)
		}
	} else {
		c.Status(http.StatusUnauthorized)
	}

}

//TODO:增加修改密码接口
func modifypwd(c *gin.Context) {

}
