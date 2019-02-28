package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/config"
	"github.com/mfslog/DecorationBackend/db"
	"github.com/mfslog/DecorationBackend/gin-engine"
	"github.com/mfslog/DecorationBackend/logfile"
	log "github.com/sirupsen/logrus"
	"io"
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
	logwriter := io.MultiWriter(logfile.NewLogFile(logfile.FileCompress(config.LogCompressFlag()),
		logfile.FileDate(config.LogDateFlag()),
		logfile.FileName(config.LogFileName()),
		logfile.FilePath(config.LogPath()),
		logfile.FileSize(config.LogSize()),
	), os.Stdout)

	log.SetOutput(logwriter)
	l, err := log.ParseLevel(config.LogLevel())
	if err != nil {
		os.Exit(-1)
	}
	log.SetLevel(l)

	gin.DefaultWriter = logwriter
	//创建图片保存目录
	var absPath string
	absPath, err = filepath.Abs(config.PicSavePath())
	if absPath == "" || err != nil {
		os.Exit(-1)
	}

	err = os.MkdirAll(absPath, 0777)
	if err != nil {
		os.Exit(-1)
	}

	config.SetPicPath(absPath)

	//连接数据库
	err = db.Init()
	if err != nil {
		log.Printf("%e", err)
		os.Exit(-1)
	}
	//初始化路由
	httpRouter, httpsRouter := engine.Init()
	//开启https 服务
	if config.CRTPath() == "" ||
		config.PrivateKeyPath() == "" {
		log.Fatalf("certificate can't empty")
		os.Exit(-1)
	}
	ch := make(chan error)
	//开启https 服务
	go func() {
		err := httpsRouter.RunTLS(fmt.Sprintf(":%d", config.ListenTLSPort()),
			config.CRTPath(), config.PrivateKeyPath())
		ch <- err
	}()
	//开启服务
	go func() {
		err := httpRouter.Run(fmt.Sprintf(":%d", config.ListenPort()))
		ch <- err
	}()
	//异常退出
	err = <-ch
	log.Errorf("%v", err)
}
