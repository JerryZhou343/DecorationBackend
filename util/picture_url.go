package util

import (
	"github.com/mfslog/DecorationBackend/config"
	"strings"
)

//GetPicFullURL 获得图片的完整URL 路径
func GetPicFullURL(path string) string {
	relativePath := strings.Trim(config.PicURLRelativePath(), "/")
	return config.PicURLPrefix() + "/" + relativePath + "/" + path
}
