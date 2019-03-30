package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/code"
	"net/http"
)

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
		"msg":  "success",
	})
}

func Failed(ctx *gin.Context, code int, msg string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": code,
		"data": "",
		"msg":  msg,
	})
}

func FailedByParam(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": code.ParamError,
		"data": "",
		"msg":  "param not correct",
	})
}

func FailedByOp(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": code.ServiceError,
		"data": "",
		"msg":  "operation Failed",
	})
}

func FailedByNotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": code.NotFound,
		"data": "",
		"msg":  "not found data",
	})
}

func FailedByLogin(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": code.LoginFailed,
		"data": "",
		"msg":  "user name or password error",
	})
}
