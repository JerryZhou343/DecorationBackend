package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/app/admin/form"
	"github.com/mfslog/DecorationBackend/db"
	"github.com/mfslog/DecorationBackend/models"
	"github.com/sirupsen/logrus"
	"net/http"
)

func selectCase(c *gin.Context) {

}

func delCase(c *gin.Context) {

}

func updateCase(c *gin.Context) {

}

func createCase(c *gin.Context) {
	info := form.Case{}
	err := c.BindJSON(&info)
	dbCase := models.TCase{}
	engine := db.DB()
	categorys := []models.TCaseCategory{}
	var cnt int64
	if err != nil {
		logrus.Errorf("get error request %v", err)
		goto FAILED
	}
	dbCase.Name = info.Name
	dbCase.Addr = info.Addr
	dbCase.OwnerName = info.OwnerName
	dbCase.Price = info.Price
	dbCase.PhoneNumber = info.PhoneNumber
	dbCase.Type = info.Type
	cnt, err = engine.InsertOne(&dbCase)
	if err != nil {
		logrus.Errorf("%v", err)
		goto FAILED
	}

	logrus.Infof("insert case id is %d", dbCase.Id)

	for _, item := range info.Categorys {
		category := models.TCaseCategory{}
		category.CaseId = dbCase.Id
		category.CategoryId = item.CId
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
