package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mfslog/DecorationBackend/logfile"
	"io"
	"os"
)

func main() {
	gin.DefaultWriter = io.MultiWriter(logfile.NewLogFile(logfile.FileName("DecorationServer.log")), os.Stdout)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
		//log.Print("aaaaa")
	})

	router.Run(":8080")
}
