package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/form"
	"github.com/mfslog/DecorationBackend/models"
	"net/http"
	"strconv"
)

//CreateCategory 添加一个tag
func CreateCategory(c *gin.Context) {
	info := form.Category{}
	err := c.BindJSON(&info)
	var category models.TCategory
	if err != nil {
		goto FAILED
	}

	category.ParentID = info.ParentID
	category.Priority = info.Priority
	category.Name = info.Name
	category.OperatorID = 1
	category.Remark = info.Remark
	category.State = 1
	models.InsertCategory(&category)

	c.Status(http.StatusCreated)
	return
FAILED:
	c.Status(http.StatusBadRequest)

}

//UpdateCategory 修改一个tag
func UpdateCategory(c *gin.Context) {
	var err error
	info := form.Category{}
	err = c.BindJSON(&info)
	category := models.TCategory{}
	categoryIDStr := c.Param("id")
	var categoryID int
	if err != nil {
		fmt.Printf("here %v", err)
		goto FAILED
	}

	categoryID, err = strconv.Atoi(categoryIDStr)
	if err != nil && categoryID == 0 {
		goto FAILED
	}

	category.Remark = info.Remark
	category.Name = info.Name
	category.Priority = info.Priority

	models.UpdateCategoryInfo(categoryID, &category)
	c.Status(http.StatusAccepted)

	return
FAILED:
	c.Status(http.StatusBadRequest)
}

//DelCategory 删除一个tag
func DelCategory(c *gin.Context) {
	var err error
	categoryIDStr := c.Param("id")
	var categoryID int
	categoryID, err = strconv.Atoi(categoryIDStr)
	if err != nil && categoryID == 0 {
		goto FAILED
	}
	models.DelCategory(categoryID)
	c.Status(http.StatusAccepted)
	return
FAILED:
	c.Status(http.StatusBadRequest)
}

//GetCategory 查询一个tag
func GetCategory(c *gin.Context) {
	var err error
	var tCategory *models.TCategory
	var categoryInfo form.Category
	categoryIDStr := c.Param("id")
	var categoryID int
	categoryID, err = strconv.Atoi(categoryIDStr)
	if err != nil && categoryID == 0 {
		goto FAILED
	}

	tCategory, err = models.GetCategoryByID(categoryID)
	categoryInfo.ID = tCategory.ID
	categoryInfo.Name = tCategory.Name
	categoryInfo.Priority = tCategory.Priority
	categoryInfo.Remark = tCategory.Remark
	c.JSON(http.StatusOK, categoryInfo)
	return

FAILED:
	c.Status(http.StatusBadRequest)
}
