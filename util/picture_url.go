package util

import (
	"github.com/mfslog/DecorationBackend/config"
	"strings"
)

func GetPicFullUrl(path string) string {
	relativePath := strings.Trim(config.PicURLRelativePath(), "/")
	return config.PicURLPrefix() + "/" + relativePath + "/" + path
}
