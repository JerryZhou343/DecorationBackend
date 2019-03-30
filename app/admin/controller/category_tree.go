package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/form"
	"github.com/mfslog/DecorationBackend/models"
	"github.com/sirupsen/logrus"
	"strconv"
)

//GetCategoryTree 获得一个分类子树，传入的节点ID 作为父节点ID
//查询返回其第一层子节点
func GetCategoryTree(c *gin.Context) {
	parentIDStr := c.Param("parent_id")
	var parentID int
	var err error
	var result []form.Category
	var dbRet []models.TCategory

	parentID, err = strconv.Atoi(parentIDStr)
	if err != nil || parentID == 0 {
		FailedByParam(c)
		return
	}
	dbRet, err = models.GetChildCategoryByParentID(parentID)
	if err != nil {
		logrus.Errorf("gin [%+v], error [%+v]\n", err)
		FailedByParam(c)
		return
	}
	for _, itr := range dbRet {
		category := form.Category{}
		category.ID = itr.ID
		category.Name = itr.Name
		category.Priority = itr.Priority
		category.ParentID = parentID
		result = append(result, category)
	}

	//c.JSON(http.StatusOK, result)
	Success(c, result)
	return
}
