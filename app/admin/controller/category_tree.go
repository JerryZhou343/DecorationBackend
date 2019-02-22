package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/app/admin/form"
	"github.com/mfslog/DecorationBackend/models"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func selectCategoryTree(c *gin.Context) {
	parentIdStr := c.Param("parent_id")
	var parentId int
	var err error
	var result []form.Category
	var dbRet []models.TCategory

	parentId, err = strconv.Atoi(parentIdStr)
	if err != nil || parentId == 0 {
		logrus.Error("prament zero")
		goto FAILED
	}
	dbRet, err = models.GetChildCategoryByParentId(parentId)
	if err != nil {
		//TODO:记录错误原因
		logrus.Errorf("%v\n", err)
		goto FAILED
	}
	for _, itr := range dbRet {
		category := form.Category{}
		category.Id = itr.Id
		category.Name = itr.Name
		category.Priority = itr.Priority
		category.ParentId = parentId
		result = append(result, category)
	}

	c.JSON(http.StatusOK, result)
	return
FAILED:
	c.Status(http.StatusNotFound)
}
