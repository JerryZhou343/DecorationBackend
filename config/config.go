package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type config struct {
	Port        int         `yaml:"port"`
	ReleaseFlag bool        `yaml:"releaseFlag"`
	Log         logConfig   `yaml:"log"`
	MySQL       mysqlConfig `yaml:"MySQL"`
}

type logConfig struct {
	Size         int64  `yaml:"size"`
	DateFlag     bool   `yaml:"dateFlag"`
	Path         string `yaml:"path"`
	CompressFlag bool   `yaml:"compressFlag"`
	Name         string `yaml:"fileName"`
	Level        int    `yaml:"level"`
}

type mysqlConfig struct {
	Name     string `yaml:"name"`
	Addr     string `yaml:"addr"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	ConLimit int    `yaml:"conLimit"`
}

var defaultIns config

//初始化
func Init() error {
	return defaultIns.init()
}

//初始化配置
func (c *config) init() error {
	f, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}
	return yaml.Unmarshal(f, &defaultIns)
}

//日志按照大小切割
func LogSize() int64 {
	return defaultIns.Log.Size
}

//日志按照日期切割
func LogDateFlag() bool {
	return defaultIns.Log.DateFlag
}

//日志路径
func LogPath() string {
	return defaultIns.Log.Path
}

//日志文件名
func LogFileName() string {
	return defaultIns.Log.Name
}

//日志压缩标识
func LogCompressFlag() bool {
	return defaultIns.Log.CompressFlag
}

//监听端口
func ListenPort() int {
	return defaultIns.Port
}

//发布模式标签
func ReleaseFlag() bool {
	return defaultIns.ReleaseFlag
}
