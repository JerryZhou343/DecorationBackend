package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/app/admin/models"
	"net/http"
)

func login(c *gin.Context) {
	userName := c.Query("username")
	passwd := c.Query("password")
	if models.CheckPassport(userName, passwd) {
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusUnauthorized)
	}

}
