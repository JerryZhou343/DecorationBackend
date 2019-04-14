package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/form"
	"github.com/mfslog/DecorationBackend/models"
	"github.com/sirupsen/logrus"
	"strconv"
)

//CreateCategory 添加一个tag
func CreateCategory(c *gin.Context) {
	info := form.Category{}
	err := c.BindJSON(&info)
	var category models.TCategory
	if err != nil {
		FailedByParam(c)
		return
	}

	category.ParentID = info.ParentID
	category.Priority = info.Priority
	category.Name = info.Name
	category.OperatorID = 1
	category.Remark = info.Remark
	category.State = 1
	models.InsertCategory(&category)
	Success(c, nil)
	return

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
		FailedByParam(c)
		return
	}

	categoryID, err = strconv.Atoi(categoryIDStr)
	if err != nil && categoryID == 0 {
		FailedByParam(c)
		return
	}

	category.Remark = info.Remark
	category.Name = info.Name
	category.Priority = info.Priority

	models.UpdateCategoryInfo(categoryID, &category)
	//c.Status(http.StatusAccepted)
	Success(c, nil)
	return

}

//DelCategory 删除一个tag
func DelCategory(c *gin.Context) {
	var err error
	categoryIDStr := c.Param("id")
	var categoryID int
	categoryID, err = strconv.Atoi(categoryIDStr)
	if err != nil && categoryID == 0 {
		FailedByParam(c)
		return
	}
	models.DelCategory(categoryID)
	Success(c, nil)
	return

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
		FailedByParam(c)
		return
	}

	tCategory, err = models.GetCategoryByID(categoryID)
	categoryInfo.ID = tCategory.ID
	categoryInfo.Name = tCategory.Name
	categoryInfo.Priority = tCategory.Priority
	categoryInfo.Remark = tCategory.Remark
	Success(c, categoryInfo)
	return

}

func GetCategorys(c *gin.Context) {
	var err error
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")
	parentIDStr := c.DefaultQuery("parentID", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		FailedByParam(c)
		return
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		FailedByParam(c)
		return
	}
	parentID, err := strconv.Atoi(parentIDStr)
	if err != nil {
		FailedByParam(c)
		return
	}

	dbRet, err := models.GetChildCategoryByParentID(parentID, limit, offset)

	if err != nil {
		logrus.Errorf("gin [%+v], error [%+v]\n", err)
		FailedByParam(c)
		return
	}

	var result []form.Category
	for _, itr := range dbRet {
		category := form.Category{}
		category.ID = itr.ID
		category.Name = itr.Name
		category.Priority = itr.Priority
		category.ParentID = parentID
		category.Remark = itr.Remark
		category.CreatedAt = itr.CreatedAt
		result = append(result, category)
	}

	//c.JSON(http.StatusOK, result)
	Success(c, result)
	return
}
