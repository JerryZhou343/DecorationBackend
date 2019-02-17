package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/app/admin/form"
	"github.com/mfslog/DecorationBackend/app/admin/models"
	"net/http"
	"strconv"
)

func selectCategoryTree(c *gin.Context) {
	parentIdStr := c.Param("parent_id")
	fmt.Println("here")
	var parentId int
	var err error
	var result []form.Category
	var dbRet []models.TCategory

	if parentIdStr == "" {
		fmt.Printf("not parament\n")
		goto FAILED
	}

	parentId, err = strconv.Atoi(parentIdStr)
	if err != nil || parentId == 0 {
		fmt.Printf("prament zero")
		goto FAILED
	}
	dbRet, err = models.GetChildCategoryByParentId(parentId)
	if err != nil {
		//TODO:记录错误原因
		fmt.Printf("%v\n", err)
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
