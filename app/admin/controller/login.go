package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/app/admin/form"
	"github.com/mfslog/DecorationBackend/app/admin/models"
	"net/http"
)

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
