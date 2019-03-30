package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/config"
	"github.com/mfslog/DecorationBackend/form"
	"github.com/mfslog/DecorationBackend/models"
	"github.com/mfslog/DecorationBackend/util"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

//GetPicture 获得一张图片
func GetPicture(c *gin.Context) {
	picIDStr := c.Param("id")
	picID, err := strconv.Atoi(picIDStr)
	if err != nil {
		//c.Status(http.StatusBadRequest)
		FailedByParam(c)
		return
	}

	picInfo, err := models.GetPictureByID(picID)
	if err != nil {
		//c.Status(http.StatusBadRequest)
		FailedByParam(c)
		return
	}

	if picInfo == nil {
		//c.Status(http.StatusNoContent)
		FailedByNotFound(c)
		return
	}

	formInfo := form.Picture{
		PicID:  picID,
		Name:   picInfo.Name,
		Remark: picInfo.Remark,
		Addr:   util.GetPicFullURL(picInfo.Addr),
	}

	//c.JSON(http.StatusOK, formInfo)
	Success(c, formInfo)
}

//CreatePicture 添加一张图片
func CreatePicture(c *gin.Context) {
	caseIDStr := c.PostForm("caseID")
	caseID, err := strconv.Atoi(caseIDStr)
	if err != nil {
		//c.Status(http.StatusBadRequest)
		FailedByParam(c)
		return
	}
	picName := c.PostForm("picName")
	name := c.PostForm("name")
	picFileIO, err := c.FormFile("picFile")
	remark := c.PostForm("remark")
	if err != nil {
		//c.Status(http.StatusBadRequest)
		FailedByParam(c)
		return
	}
	//计算文件名
	if picName == "" {
		//c.Status(http.StatusBadRequest)
		FailedByParam(c)
		return
	}

	nameSlice := strings.Split(picName, ".")
	if len(nameSlice) < 2 {
		//c.Status(http.StatusBadRequest)
		FailedByParam(c)
		return
	}
	suffix := nameSlice[(len(nameSlice) - 1)]
	md5Value := util.GetMD5(caseIDStr + picName)
	fileName := md5Value + "." + suffix
	absFileName := config.GetPicPath() + "/" + fileName
	relativePath := fileName
	f, err := util.PathExists(absFileName)
	if err != nil {
		logrus.Errorf("%v", err)
		//c.Status(http.StatusInternalServerError)
		FailedByOp(c)
		return
	}

	//删除目标文件
	if f {
		err = os.Remove(absFileName)
		if err != nil {
			//c.Status(http.StatusInternalServerError)
			FailedByOp(c)
			return
		}
	}
	err = c.SaveUploadedFile(picFileIO, absFileName)

	if err != nil {
		logrus.Errorf("save file [%s] error:[%v]", absFileName, err)
		//c.Status(http.StatusInternalServerError)
		FailedByOp(c)
		return
	}
	logrus.Infof("save file [%s] success", absFileName)

	picInfo := models.TPicture{}
	picInfo.Addr = relativePath
	picInfo.Name = name
	picInfo.Remark = remark
	picInfo.CaseID = caseID

	err = models.InsertOnePicture(&picInfo)
	if err != nil {
		//c.Status(http.StatusInternalServerError)
		FailedByOp(c)
		return
	}
	//c.Status(http.StatusOK)
	/*
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"msg":    "success",
			"imgUrl": util.GetPicFullURL(relativePath),
		})
	*/
	Success(c, util.GetPicFullURL(relativePath))
}

//DelPicture 软删除图片
func DelPicture(c *gin.Context) {
	picIDStr := c.Param("id")
	picID, err := strconv.Atoi(picIDStr)
	if err != nil {
		//c.Status(http.StatusBadRequest)
		FailedByParam(c)
		return
	}

	models.GetCategoryByCaseID(picID)

	//c.Status(http.StatusOK)
	Success(c, nil)
}

//UpdatePicture 更新图片的描述信息
func UpdatePicture(c *gin.Context) {
	picIDStr := c.Param("id")
	picID, err := strconv.Atoi(picIDStr)
	if err != nil {
		//c.Status(http.StatusBadRequest)
		FailedByParam(c)
		return
	}
	info := form.Picture{}
	err = c.BindJSON(&info)

	picInfo := models.TPicture{
		Name:   info.Name,
		Remark: info.Remark,
	}
	err = models.UpdateOnePicture(picID, &picInfo)
	if err != nil {
		//c.Status(http.StatusBadRequest)
		FailedByOp(c)
		return
	}

	//c.Status(http.StatusOK)
	Success(c, nil)
}
