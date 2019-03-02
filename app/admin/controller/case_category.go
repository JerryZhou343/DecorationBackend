package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/form"
	"github.com/mfslog/DecorationBackend/models"
	"net/http"
	"strconv"
)

//AddCaseCategory 添加一个分类
func AddCaseCategory(c *gin.Context) {
	caseIDStr := c.Param("id")
	caseID, err := strconv.Atoi(caseIDStr)
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
	tcategory.CategoryID = categoryInfo.CategoryID
	tcategory.ID = caseID

	err = models.InsertOneCaseCategory(&tcategory)
	if err == nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
	return
}

//GetCaseCategory  获得一个case的分类
func GetCaseCategory(c *gin.Context) {
	caseIDStr := c.Param("id")
	caseID, err := strconv.Atoi(caseIDStr)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	ret, err := models.GetCategoryByCaseID(caseID)
	result := []form.CaseCategory{}
	for _, item := range *ret {
		tmp := form.CaseCategory{}
		tmp.CategoryID = item.CategoryID
		tmp.RID = item.ID

		result = append(result, tmp)
	}

	c.JSON(http.StatusOK, result)
}

//DelCaseCategory 删除一个case的分类
func DelCaseCategory(c *gin.Context) {
	caseIDStr := c.Param("id")
	caseID, err := strconv.Atoi(caseIDStr)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	categoryIDStr := c.Query("category_id")
	categoryID, err := strconv.Atoi(categoryIDStr)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = models.DelCaseCategoryByID(caseID, categoryID)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}
