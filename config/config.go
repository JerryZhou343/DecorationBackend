package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type config struct {
	Port        int         `yaml:"port"`
	ReleaseFlag bool        `yaml:"releaseFlag"`
	Log         logConfig   `yaml:"log"`
	MySQL       mySQLConfig `yaml:"mySQL"`
}

type logConfig struct {
	Size         int64  `yaml:"size"`
	DateFlag     bool   `yaml:"dateFlag"`
	Path         string `yaml:"path"`
	CompressFlag bool   `yaml:"compressFlag"`
	Name         string `yaml:"fileName"`
	Level        int    `yaml:"level"`
}

type mySQLConfig struct {
	Name      string `yaml:"name"`
	Addr      string `yaml:"addr"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	PoolLimit int    `yaml:"poolLimit"`
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

//db 地址
func MySQLAddr() string {
	return defaultIns.MySQL.Addr
}

//返回mySQLDB数据库用户名
func MySQLUser() string {
	return defaultIns.MySQL.User
}

//db 密码
func MySQLPassword() string {
	return defaultIns.MySQL.Password
}

//db 连接数
func MySQLPoolLimit() int {
	return defaultIns.MySQL.PoolLimit
}

func MySQLDBName() string {
	return defaultIns.MySQL.Name
}
