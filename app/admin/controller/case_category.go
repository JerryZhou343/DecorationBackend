package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/form"
	"github.com/mfslog/DecorationBackend/models"
	"net/http"
	"strconv"
)

func AddCaseCategory(c *gin.Context) {
	caseIdStr := c.Param("id")
	caseId, err := strconv.Atoi(caseIdStr)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	categoryInfo := form.CaseCategory{}
	err = c.BindJSON(&categoryInfo)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	//todo:校验category 是否存在
	tcategory := models.TCaseCategory{}
	tcategory.CategoryId = categoryInfo.CategoryId
	tcategory.Id = caseId

	err = models.InsertOneCaseCategory(&tcategory)
	if err == nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
	return
}

func GetCaseCategory(c *gin.Context) {
	caseIdStr := c.Param("id")
	caseId, err := strconv.Atoi(caseIdStr)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	ret, err := models.GetCategoryByCaseId(caseId)
	result := []form.CaseCategory{}
	for _, item := range *ret {
		tmp := form.CaseCategory{}
		tmp.CategoryId = item.CategoryId
		tmp.RId = item.Id

		result = append(result, tmp)
	}

	c.JSON(http.StatusOK, result)
}

func DelCaseCategory(c *gin.Context) {
	caseIdStr := c.Param("id")
	caseId, err := strconv.Atoi(caseIdStr)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	categoryIdStr := c.Query("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = models.DelCaseCategoryById(caseId, categoryId)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}
