package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/form"
	"github.com/mfslog/DecorationBackend/middleware/jwtauth"
	"github.com/mfslog/DecorationBackend/models"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

//Login 登录接口
func Login(c *gin.Context) {
	info := form.Login{}
	err := c.BindJSON(&info)
	var uid int64
	var flag bool
	if err == nil &&
		info.UserName != "" &&
		info.Password != "" {
		flag, uid = models.CheckPassport(info.UserName, info.Password)
		if !flag {
			FailedByLogin(c)
			return
		}
	} else {
		FailedByParam(c)
		return
	}

	//TODO:登录成功，签发token
	user := jwtauth.CustomClaims{
		ID:    uid,
		Name:  info.UserName,
		Email: "",
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: time.Now().Add(time.Duration(180) * time.Second).Unix(),
			Id:        strconv.Itoa(int(uid)),
			IssuedAt:  0,
			Issuer:    "",
			NotBefore: 0,
			Subject:   "",
		},
	}
	jwt := jwtauth.NewJWT()
	token, err := jwt.CreateToken(user)
	if err != nil {
		FailedByOp(c)
		logrus.Errorf("generate token failed [%v]", err)
		return
	}

	Success(c, token)

}

//UpdatePwd 修改密码接口
//TODO:增加修改密码接口
func UpdatePwd(c *gin.Context) {
	info := form.NewPassWord{}
	token := c.DefaultQuery("token", "")
	if token == "" {
		FailedByParam(c)
		return
	}
	err := c.BindJSON(&info)
	if err != nil {
		FailedByParam(c)
		logrus.Errorf("convert json not correct [%v]", err)
		return
	}

	jwt := jwtauth.NewJWT()
	cl, err := jwt.ParseToken(token)
	if err != nil {
		FailedByParam(c)
		logrus.Errorf("token not correct [%v]", err)
		return
	}

	err = models.ChangePassword(cl.Name, info.OldPassword, info.NewPassWord)
	if err != nil {
		logrus.Errorf("modify password failed [%v]", err)
		FailedByOp(c)
		return
	}

	Success(c, nil)
}
