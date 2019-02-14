package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/config"
	"github.com/mfslog/DecorationBackend/logfile"
	"github.com/mfslog/DecorationBackend/gin-engine"
	"io"
	"os"
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

	//连接数据库

	//初始化路由
	r := gin_engine.Init()

	//开启服务
	r.Run(fmt.Sprintf(":%d", config.ListenPort()))
}
