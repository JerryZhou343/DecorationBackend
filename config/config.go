package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type config struct {
	Port        int           `yaml:"port"`
	TLSPort     int           `yaml:"TLSPort"`
	ReleaseFlag bool          `yaml:"releaseFlag"`
	Picture     pictureConfig `yaml:"picture"`
	Log         logConfig     `yaml:"log"`
	MySQL       mySQLConfig   `yaml:"mySQL"`
	Auth        authConfig    `yaml:"auth"`
}

type pictureConfig struct {
	SavePath        string `yaml:"savePath"`
	URLPrefix       string `yaml:"urlPrefix"`
	URLRelativePath string `yaml:"urlRelativePath"`
}

type logConfig struct {
	Size         int64  `yaml:"size"`
	DateFlag     bool   `yaml:"dateFlag"`
	Path         string `yaml:"path"`
	CompressFlag bool   `yaml:"compressFlag"`
	Name         string `yaml:"fileName"`
	Level        string `yaml:"level"`
}

type mySQLConfig struct {
	Name      string `yaml:"name"`
	Addr      string `yaml:"addr"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	PoolLimit int    `yaml:"poolLimit"`
}

type authConfig struct {
	SignKey        string `yaml:"signKey"`
	CRTPath        string `yaml:"CRTPath"`
	PrivateKeyPath string `yaml:"privateKeyPath"`
}

var defaultIns config

//Init 初始化
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

//LogSize 设置日志按照大小切割
func LogSize() int64 {
	return defaultIns.Log.Size
}

//LogDateFlag 设置日志按照日期切割
func LogDateFlag() bool {
	return defaultIns.Log.DateFlag
}

//LogPath 设置日志路径
func LogPath() string {
	return defaultIns.Log.Path
}

//LogFileName 设置日志文件名
func LogFileName() string {
	return defaultIns.Log.Name
}

//LogCompressFlag 设置日志压缩标识
//如果设置该标识，分割文件的同时压缩文件
func LogCompressFlag() bool {
	return defaultIns.Log.CompressFlag
}

//LogLevel 设置日志等级
func LogLevel() string {
	return defaultIns.Log.Level
}

//ListenPort 获得服务监听端空监听端口
func ListenPort() int {
	return defaultIns.Port
}

//ListenTLSPort 获得TLS 监听端口
func ListenTLSPort() int {
	return defaultIns.TLSPort
}

//ReleaseFlag 获得Go gin 发布模式标签
func ReleaseFlag() bool {
	return defaultIns.ReleaseFlag
}

//PicSavePath 获得上传图片存放路径
func PicSavePath() string {
	return defaultIns.Picture.SavePath
}

//PicURLPrefix 获得图片URL 前缀
func PicURLPrefix() string {
	return defaultIns.Picture.URLPrefix
}

//PicURLRelativePath 获得图片URL相对地址
func PicURLRelativePath() string {
	return defaultIns.Picture.URLRelativePath
}

//MySQLAddr 获得MysqlDB地址
func MySQLAddr() string {
	return defaultIns.MySQL.Addr
}

//MySQLUser 获得MySQLDB数据库用户名
func MySQLUser() string {
	return defaultIns.MySQL.User
}

//MySQLPassword 获得MySQL DB 密码
func MySQLPassword() string {
	return defaultIns.MySQL.Password
}

//MySQLPoolLimit 获得MySQL 连接数
func MySQLPoolLimit() int {
	return defaultIns.MySQL.PoolLimit
}

//MySQLDBName 获得MySQL DB名称
func MySQLDBName() string {
	return defaultIns.MySQL.Name
}

//图片的绝对路径
var picturePath string

//GetPicPath 获得图片的绝对存储路径
func GetPicPath() string {
	return picturePath
}

//SetPicPath 设置图片的存储的绝对的路径
func SetPicPath(path string) {
	picturePath = path
}

// CRTPath 获得证书路径
func CRTPath() string {
	return defaultIns.Auth.CRTPath
}

// PrivateKeyPath 获得私钥路径
func PrivateKeyPath() string {
	return defaultIns.Auth.PrivateKeyPath
}
