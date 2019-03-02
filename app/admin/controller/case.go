package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/db"
	"github.com/mfslog/DecorationBackend/form"
	"github.com/mfslog/DecorationBackend/models"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

//GetCase handler 获得一个case
func GetCase(c *gin.Context) {
	caseIDStr := c.Param("id")
	caseID, err := strconv.Atoi(caseIDStr)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	var caseObj *models.TCase
	caseObj, err = models.GetCaseByID(caseID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if caseObj == nil {
		c.Status(http.StatusNoContent)
		return
	}

	ret := form.ComplexCaseCategory{}
	ret.CaseInfo.ID = caseObj.ID
	ret.CaseInfo.Name = caseObj.Name
	ret.CaseInfo.Type = caseObj.Type
	ret.CaseInfo.PhoneNumber = caseObj.PhoneNumber
	ret.CaseInfo.OwnerName = caseObj.OwnerName
	ret.CaseInfo.Price = caseObj.Price
	ret.CaseInfo.Addr = caseObj.Addr
	var categoryRet *[]*models.TCaseCategory
	categoryRet, err = models.GetCategoryByCaseID(ret.CaseInfo.ID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	for _, item := range *categoryRet {
		tmp := form.CaseCategory{}
		tmp.CategoryID = item.CategoryID
		tmp.RID = item.ID

		ret.CategoryInfo = append(ret.CategoryInfo, tmp)
	}

	c.JSON(http.StatusOK, ret)
	return
}

//DelCase 逻辑删除一个case
func DelCase(c *gin.Context) {
	caseIDStr := c.Param("id")
	caseID, err := strconv.Atoi(caseIDStr)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	err = models.DelCaseByID(caseID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.Status(http.StatusOK)
}

//UpdateCaseInfo 更新case信息
func UpdateCaseInfo(c *gin.Context) {
	caseIDStr := c.Param("id")
	caseInfo := form.Case{}

	caseID, err := strconv.Atoi(caseIDStr)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = c.BindJSON(&caseInfo)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	tcase := models.TCase{}
	tcase.Name = caseInfo.Name
	tcase.Addr = caseInfo.Addr
	tcase.Price = caseInfo.Price
	tcase.PhoneNumber = caseInfo.PhoneNumber
	tcase.OwnerName = caseInfo.OwnerName
	tcase.Type = caseInfo.Type

	err = models.UpdateCaseByID(caseID, &tcase)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
	return
}

//CreateCase 创建一个case对象
func CreateCase(c *gin.Context) {

	info := form.ComplexCaseCategory{}
	err := c.BindJSON(&info)
	dbCase := models.TCase{}
	engine := db.DB()
	categorys := []models.TCaseCategory{}
	var cnt int64
	if err != nil {
		logrus.Errorf("get error request %v", err)
		goto FAILED
	}
	dbCase.Name = info.CaseInfo.Name
	dbCase.Addr = info.CaseInfo.Addr
	dbCase.OwnerName = info.CaseInfo.OwnerName
	dbCase.Price = info.CaseInfo.Price
	dbCase.PhoneNumber = info.CaseInfo.PhoneNumber
	dbCase.Type = info.CaseInfo.Type
	cnt, err = engine.InsertOne(&dbCase)
	if err != nil {
		logrus.Errorf("%v", err)
		goto FAILED
	}

	logrus.Infof("insert case id is %d", dbCase.ID)

	for _, item := range info.CategoryInfo {
		category := models.TCaseCategory{}
		category.CaseID = dbCase.ID
		category.CategoryID = item.CategoryID
		categorys = append(categorys, category)
	}
	cnt, err = engine.Insert(categorys)

	if err != nil {
		logrus.Errorf("%v", err)
	}

	logrus.Info("insert %d to case category %d", cnt)
	return
FAILED:
	c.Status(http.StatusBadRequest)

}
