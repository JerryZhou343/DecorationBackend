package logfile

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"unsafe"
)

//LogIO 日志输出的IO对象定义
type LogIO struct {
	curFile      *os.File
	fileName     string
	sizeFlag     bool
	dateFlag     bool
	compressFlag bool
	filePath     string
	sizeValue    int64
	todayDate    string
	msgQueue     chan string
	closed       bool
	cnt          int
}

//Option 日志配置函数签名定义
type Option func(*LogIO)

//NewLogFile 创建新的日志文件
func NewLogFile(opts ...Option) *LogIO {
	IOIns := LogIO{
		curFile:      os.Stdout,
		fileName:     "",
		sizeFlag:     false,
		dateFlag:     false,
		compressFlag: false,
		filePath:     "../log/",
		sizeValue:    10240,
		todayDate:    getCurrentDate(),
		closed:       false,
		cnt:          1,
		msgQueue:     make(chan string, 1024),
	}

	for _, o := range opts {
		o(&IOIns)
	}

	//潜在的创建目录错误
	absPath, _ := filepath.Abs(IOIns.filePath)
	os.MkdirAll(absPath, 0666)

	if IOIns.fileName != "" {
		file, err := os.OpenFile(IOIns.filePath+IOIns.fileName,
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			fmt.Println(err.Error())
		}
		IOIns.curFile = file
	}

	go IOIns.worker()

	return &IOIns
}

//FileName 设置文件名
func FileName(fileName string) Option {
	return func(o *LogIO) {
		o.fileName = fileName
	}
}

//FilePath 设置文件路径
func FilePath(path string) Option {
	return func(o *LogIO) {
		var slash = string(os.PathSeparator)
		dir, _ := filepath.Abs(path)
		o.filePath = dir + slash
	}
}

//FileSize 设置文件切割大小,单位为M
func FileSize(size int64) Option {
	return func(o *LogIO) {
		o.sizeFlag = true
		o.sizeValue = size * 1024 * 1024
	}
}

//FileDate 按照天来切割
func FileDate(flag bool) Option {
	return func(o *LogIO) {
		o.dateFlag = flag
	}
}

//FileCompress 设置压缩标志
func FileCompress(flag bool) Option {
	return func(o *LogIO) {
		o.compressFlag = flag
	}
}

//Write 向IO 输出内容，实现Writer接口
func (log *LogIO) Write(p []byte) (n int, err error) {
	str := (*string)(unsafe.Pointer(&p))
	log.msgQueue <- (*str)
	return len(p), nil
}

//doRotate 切割文件
func (log *LogIO) doRotate() {

	defer func() {
		rec := recover()
		if rec != nil {
			fmt.Println("doRotate %v", rec)
		}
	}()

	if log.curFile == nil {
		fmt.Println("doRotate curFile nil,return")
		return
	}

	prefile := log.curFile
	_, err := prefile.Stat()
	var prefileName = ""
	if err == nil {
		filePath := log.filePath + log.fileName
		log.closed = true
		err := prefile.Close()
		if err != nil {
			fmt.Println("doRotate close err", err.Error())
		}
		y, m, d := time.Now().Date()
		prefileName = filePath + "." + fmt.Sprintf("%.4d%.2d%.2d", y, m, d) + strconv.Itoa(log.cnt)
		log.cnt++
		err = os.Rename(filePath, prefileName)
	}

	if log.fileName != "" {
		nextFile, err := os.OpenFile(log.filePath+log.fileName,
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

		if err != nil {
			fmt.Println(err.Error())
		}
		log.closed = false
		log.curFile = nextFile
		nowDate := getCurrentDate()
		log.todayDate = nowDate
	}
	if log.compressFlag {
		go log.compressFile(prefileName, prefileName+".gz")
	}
}

//worker 输出内容到文件
func (log *LogIO) worker() {
	for {
		select {
		case msg := <-log.msgQueue:
			{
				log.curFile.WriteString(msg)
				if log.sizeFlag == true {
					curInfo, _ := os.Stat(log.filePath + log.fileName)
					if curInfo.Size() >= log.sizeValue {
						log.doRotate()
					}
				}

				nowDate := getCurrentDate()

				if log.dateFlag == true &&
					nowDate != log.todayDate {
					log.cnt = 1
					log.doRotate()
				}
			}
		}

	}

}

//compressFile 压缩日志文件
func (log *LogIO) compressFile(Src string, Dst string) error {
	defer func() {
		rec := recover()
		if rec != nil {
			fmt.Println(rec)
		}
	}()
	newfile, err := os.Create(Dst)
	if err != nil {
		return err
	}
	defer newfile.Close()

	file, err := os.Open(Src)
	if err != nil {
		return err
	}

	zw := gzip.NewWriter(newfile)

	filestat, err := file.Stat()
	if err != nil {
		return nil
	}

	zw.Name = filestat.Name()
	zw.ModTime = filestat.ModTime()
	_, err = io.Copy(zw, file)
	if err != nil {
		return nil
	}

	zw.Flush()
	if err := zw.Close(); err != nil {
		return nil
	}
	file.Close()
	os.Remove(Src)
	return nil
}

//getCurrentDate 格式化年月日
func getCurrentDate() string {
	year, month, day := time.Now().Date()
	return fmt.Sprintf("%.4d%.2d%.2d", year, month, day)
}
