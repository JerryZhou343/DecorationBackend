package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/config"
	"github.com/mfslog/DecorationBackend/db"
	"github.com/mfslog/DecorationBackend/gin-engine"
	"github.com/mfslog/DecorationBackend/logfile"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//初始化配置
	err := config.Init()
	if err != nil {
		fmt.Println("load config failed exit")
		os.Exit(-1)
	}

	//初始化日志
	gin.DefaultWriter = io.MultiWriter(logfile.NewLogFile(logfile.FileCompress(config.LogCompressFlag()),
		logfile.FileDate(config.LogDateFlag()),
		logfile.FileName(config.LogFileName()),
		logfile.FilePath(config.LogPath()),
		logfile.FileSize(config.LogSize()),
	), os.Stdout)

	//创建图片保存目录
	var absPath string
	absPath, err = filepath.Abs(config.PicPath())
	if absPath == "" || err != nil {
		os.Exit(-1)
	}
	err = os.MkdirAll(absPath, 0666)
	if err != nil {
		os.Exit(-1)
	}

	//连接数据库
	err = db.Init()
	if err != nil {
		log.Printf("%e", err)
		os.Exit(-1)
	}
	//初始化路由
	r := gin_engine.Init()

	//开启服务
	r.Run(fmt.Sprintf(":%d", config.ListenPort()))
}
