package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/app/admin/form"
	"github.com/mfslog/DecorationBackend/app/admin/models"
	"net/http"
	"strconv"
)

func queryCategoryTree(c *gin.Context) {
	parentIdStr := c.Param("parent_id")
	var parentId int
	var err error
	var result []form.Category

	if parentIdStr == "" {
		goto FAILED
	}

	parentId, err = strconv.Atoi(parentIdStr)
	if err != nil || parentId == 0 {
		goto FAILED
	}

	for _, itr := range models.GetChildCategoryByParentId(parentId) {
		category := form.Category{}
		category.Id = itr.Id
		category.Name = itr.Name
		category.Priority = itr.Priority
		category.ParentId = parentId
		result = append(result, category)
	}

	c.JSON(http.StatusOK, result)

FAILED:
	c.Status(http.StatusNotFound)
}
