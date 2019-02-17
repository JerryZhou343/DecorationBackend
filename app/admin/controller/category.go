package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/app/admin/form"
	"github.com/mfslog/DecorationBackend/app/admin/models"
	"net/http"
)

//添加一个tag
func addCategory(c *gin.Context) {
	info := form.Category{}
	err := c.BindJSON(&info)
	var category models.TCategory
	if err != nil {
		goto FAILED
	}

	category.ParentId = info.ParentId
	category.Priority = info.Priority
	category.Name = info.Name
	category.OperatorId = 1
	category.Remark = info.Remark
	category.State = 1
	models.InsertCategory(&category)

	c.Status(http.StatusCreated)

FAILED:
	c.Status(http.StatusBadRequest)

}

//修改一个tag
func modifyCategory(c *gin.Context) {

}

//删除一个tag
func delCategory(c *gin.Context) {

}

//查询一个tag
func queryCategory(c *gin.Context) {

}
