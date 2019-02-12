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

type logIO struct {
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

type option func(*logIO)

func NewLogFile(opts ...option) *logIO {
	IOIns := logIO{
		curFile:      os.Stdout,
		fileName:     "",
		sizeFlag:     false,
		dateFlag:     false,
		compressFlag: false,
		filePath:     "./log/",
		sizeValue:    10240,
		todayDate:    getCurrentDate(),
		closed:       false,
		cnt:          1,
		msgQueue:     make(chan string, 1024),
	}

	for _, o := range opts {
		o(&IOIns)
	}

	if IOIns.fileName != "" {
		file, err := os.OpenFile(IOIns.filePath+IOIns.fileName,
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(err.Error())
		}
		IOIns.curFile = file
	}

	go IOIns.worker()

	return &IOIns
}

//设置文件名
func FileName(fileName string) option {
	return func(o *logIO) {
		o.fileName = fileName
	}
}

//设置文件路径
func FilePath(path string) option {
	return func(o *logIO) {
		var slash string = string(os.PathSeparator)
		dir, _ := filepath.Abs(path)
		o.filePath = dir + slash
	}
}

//设置文件切割大小,单位为M
func FileSize(size int) option {
	return func(o *logIO) {
		o.sizeFlag = true
		o.sizeValue = int64(size) * 1024 * 1024
	}
}

//按照天来切割
func FileDate(flag bool) option {
	return func(o *logIO) {
		o.dateFlag = flag
	}
}

func FileCompress(flag bool) option {
	return func(o *logIO) {
		o.compressFlag = flag
	}
}

//向IO 输出
func (log *logIO) Write(p []byte) (n int, err error) {
	str := (*string)(unsafe.Pointer(&p))
	log.msgQueue <- (*str)
	fmt.Println(*str)
	return len(p), nil
}

//切割文件
func (log *logIO) doRotate() {

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
	var prefileName string = ""
	if err == nil {
		filePath := log.filePath + log.fileName
		log.closed = true
		err := prefile.Close()
		if err != nil {
			fmt.Println("doRotate close err", err.Error())
		}
		y, m, d := time.Now().Date()
		prefileName = filePath + "." + fmt.Sprintf("%.4d%.2d%.2d", y, m, d) + strconv.Itoa(log.cnt)
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

//输出文件
func (log *logIO) worker() {
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
					log.doRotate()
				}
			}
		}

	}

}

//压缩日志文件
func (f *logIO) compressFile(Src string, Dst string) error {
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

//格式化年月日
func getCurrentDate() string {
	year, month, day := time.Now().Date()
	return fmt.Sprintf("%.4d%.2d%.2d", year, month, day)
}
