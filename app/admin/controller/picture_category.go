package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/form"
	"github.com/mfslog/DecorationBackend/models"
	"net/http"
	"strconv"
)

func GetPicCategory(c *gin.Context) {

	picIdStr := c.Param("id")
	picId, err := strconv.Atoi(picIdStr)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	ret, err := models.GetPicCategory(picId)
	categoryInfo := []form.PicCategory{}
	for _, item := range *ret {
		tmp := form.PicCategory{}
		tmp.CategoryId = item.CategoryId
		categoryInfo = append(categoryInfo, tmp)
	}

	c.JSON(http.StatusOK, categoryInfo)
}

func DelPicCategory(c *gin.Context) {
	picIdStr := c.Param("id")
	picId, err := strconv.Atoi(picIdStr)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	categoryIdStr := c.Param("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	err = models.DelPicCategory(picId, categoryId)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func AddPicCategory(c *gin.Context) {
	picIdStr := c.Param("id")
	picId, err := strconv.Atoi(picIdStr)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	categoryIdStr := c.Param("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	err = models.InsertOnePicCategory(picId, categoryId)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
