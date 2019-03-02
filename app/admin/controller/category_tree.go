package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/form"
	"github.com/mfslog/DecorationBackend/models"
	"github.com/sirupsen/logrus"
	"net/http"
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
		logrus.Error("prament zero")
		goto FAILED
	}
	dbRet, err = models.GetChildCategoryByParentID(parentID)
	if err != nil {
		//TODO:记录错误原因
		logrus.Errorf("%v\n", err)
		goto FAILED
	}
	for _, itr := range dbRet {
		category := form.Category{}
		category.ID = itr.ID
		category.Name = itr.Name
		category.Priority = itr.Priority
		category.ParentID = parentID
		result = append(result, category)
	}

	c.JSON(http.StatusOK, result)
	return
FAILED:
	c.Status(http.StatusNotFound)
}
