package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/form"
	"github.com/mfslog/DecorationBackend/models"
	"strconv"
)

//GetPicCategory 获得图片的分类
func GetPicCategory(c *gin.Context) {

	picIDStr := c.Param("id")
	picID, err := strconv.Atoi(picIDStr)
	if err != nil {
		//c.Status(http.StatusBadRequest)
		FailedByParam(c)
		return
	}

	ret, err := models.GetPicCategory(picID)
	categoryInfo := []form.PicCategory{}
	for _, item := range *ret {
		tmp := form.PicCategory{}
		tmp.CategoryID = item.CategoryID
		categoryInfo = append(categoryInfo, tmp)
	}

	//c.JSON(http.StatusOK, categoryInfo)
	Success(c, categoryInfo)
}

//DelPicCategory 删除图片的分类
func DelPicCategory(c *gin.Context) {
	picIDStr := c.Param("id")
	picID, err := strconv.Atoi(picIDStr)
	if err != nil {
		//c.Status(http.StatusBadRequest)
		FailedByParam(c)
		return
	}

	categoryIDStr := c.Param("category_id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		//c.Status(http.StatusBadRequest)
		FailedByParam(c)
		return
	}
	err = models.DelPicCategory(picID, categoryID)
	if err != nil {
		//c.Status(http.StatusInternalServerError)
		FailedByParam(c)
		return
	}

	//c.Status(http.StatusOK)
	Success(c, nil)
}

//AddPicCategory 添加图片分类
func AddPicCategory(c *gin.Context) {
	picIDStr := c.Param("id")
	picID, err := strconv.Atoi(picIDStr)
	if err != nil {
		//c.Status(http.StatusBadRequest)
		FailedByParam(c)
		return
	}

	categoryIDStr := c.Param("category_id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		//c.Status(http.StatusBadRequest)
		FailedByParam(c)
		return
	}
	err = models.InsertOnePicCategory(picID, categoryID)
	if err != nil {
		//c.Status(http.StatusInternalServerError)
		FailedByOp(c)
		return
	}

	//c.Status(http.StatusOK)
	Success(c, nil)
}
