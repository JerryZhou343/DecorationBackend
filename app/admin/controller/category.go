package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/app/admin/form"
	"github.com/mfslog/DecorationBackend/app/admin/models"
	"net/http"
	"strconv"
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
	return
FAILED:
	c.Status(http.StatusBadRequest)

}

//修改一个tag
func updateCategory(c *gin.Context) {
	var err error
	info := form.Category{}
	err = c.BindJSON(&info)
	category := models.TCategory{}
	categoryIdStr := c.Param("id")
	var categoryId int
	if err != nil {
		fmt.Printf("here %v", err)
		goto FAILED
	}

	categoryId, err = strconv.Atoi(categoryIdStr)
	if err != nil && categoryId == 0 {
		goto FAILED
	}

	category.Remark = info.Remark
	category.Name = info.Name
	category.Priority = info.Priority

	models.UpdateCategoryInfo(categoryId, &category)
	c.Status(http.StatusAccepted)

	return
FAILED:
	c.Status(http.StatusBadRequest)
}

//删除一个tag
func delCategory(c *gin.Context) {
	var err error
	categoryIdStr := c.Param("id")
	var categoryId int
	categoryId, err = strconv.Atoi(categoryIdStr)
	if err != nil && categoryId == 0 {
		goto FAILED
	}
	models.DelCategory(categoryId)
	c.Status(http.StatusAccepted)
	return
FAILED:
	c.Status(http.StatusBadRequest)
}

//查询一个tag
func queryCategory(c *gin.Context) {
	var err error
	var tCategory *models.TCategory
	var categoryInfo form.Category
	categoryIdStr := c.Param("id")
	var categoryId int
	categoryId, err = strconv.Atoi(categoryIdStr)
	if err != nil && categoryId == 0 {
		goto FAILED
	}

	tCategory, err = models.GetCategoryById(categoryId)
	categoryInfo.Id = tCategory.Id
	categoryInfo.Name = tCategory.Name
	categoryInfo.Priority = tCategory.Priority
	categoryInfo.Remark = tCategory.Remark
	c.JSON(http.StatusOK, categoryInfo)
	return

FAILED:
	c.Status(http.StatusBadRequest)
}
