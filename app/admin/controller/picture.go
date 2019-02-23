package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/form"
	"github.com/mfslog/DecorationBackend/config"
	"github.com/mfslog/DecorationBackend/models"
	"github.com/mfslog/DecorationBackend/util"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//获得一张图片
func GetPicture(c *gin.Context) {
	picIdStr := c.Param("id")
	picId, err := strconv.Atoi(picIdStr)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	picInfo, err := models.GetPictureById(picId)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if picInfo == nil {
		c.Status(http.StatusNoContent)
		return
	}

	formInfo := form.Picture{
		PicId:  picId,
		Name:   picInfo.Name,
		Remark: picInfo.Remark,
		Addr:   picInfo.Addr,
	}

	c.JSON(http.StatusOK, formInfo)
}

//添加一张图片
func CreatePicture(c *gin.Context) {
	caseIdStr := c.PostForm("caseId")
	caseId, err := strconv.Atoi(caseIdStr)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	picName := c.PostForm("picName")
	name := c.PostForm("name")
	picFileIO, err := c.FormFile("picFile")
	remark := c.PostForm("remark")
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	//计算文件名
	if picName == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	nameSlice := strings.Split(picName, ".")
	if len(nameSlice) < 2 {
		c.Status(http.StatusBadRequest)
		return
	}
	suffix := nameSlice[(len(nameSlice) - 1)]
	md5Value := util.GetMD5(caseIdStr + picName)
	fileName := md5Value + "." + suffix
	absFileName := config.GetPicPath() + "/" + fileName
	relativePath := config.PicPath() + "/" + fileName
	f, err := util.PathExists(absFileName)
	if err != nil {
		logrus.Errorf("%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	//删除目标文件
	if f {
		err = os.Remove(absFileName)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
	}
	err = c.SaveUploadedFile(picFileIO, absFileName)

	if err != nil {
		logrus.Errorf("save file [%s] error:[%v]", absFileName, err)
		c.Status(http.StatusInternalServerError)
		return
	}
	logrus.Infof("save file [%s] success", absFileName)

	picInfo := models.TPicture{}
	picInfo.Addr = relativePath
	picInfo.Name = name
	picInfo.Remark = remark
	picInfo.CaseId = caseId

	err = models.InsertOnePicture(&picInfo)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

//软删除图片
func DelPicture(c *gin.Context) {
	picIdStr := c.Param("id")
	picId, err := strconv.Atoi(picIdStr)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	models.GetCategoryByCaseId(picId)

	c.Status(http.StatusOK)
}

//更新图片的描述信息
func UpdatePicture(c *gin.Context) {
	picIdStr := c.Param("id")
	picId, err := strconv.Atoi(picIdStr)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	info := form.Picture{}
	err = c.BindJSON(&info)

	picInfo := models.TPicture{
		Name:   info.Name,
		Remark: info.Remark,
	}
	err = models.UpdateOnePicture(picId, &picInfo)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}
