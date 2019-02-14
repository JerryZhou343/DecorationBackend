package controller

import "github.com/gin-gonic/gin"

func login(c *gin.Context) {
	userName := c.Query("username")
	passwd := c.Query("password")
}
