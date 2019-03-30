package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/form"
	"github.com/mfslog/DecorationBackend/models"
)

//Login 登录接口
func Login(c *gin.Context) {
	info := form.Login{}
	err := c.BindJSON(&info)
	if err == nil &&
		info.UserName != "" &&
		info.Password != "" {
		if !models.CheckPassport(info.UserName, info.Password) {
			FailedByLogin(c)
			return
		}
	} else {
		FailedByParam(c)
		return
	}

	//TODO:登录成功，签发token

}

//UpdatePwd 修改密码接口
//TODO:增加修改密码接口
func UpdatePwd(c *gin.Context) {

}
