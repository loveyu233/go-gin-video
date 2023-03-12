package upload

import (
	"github.com/spf13/viper"
	"go-gin-video/utils/strconv"
	"mime/multipart"
)

func IsVideoOverflow(files ...*multipart.FileHeader) bool {
	maxSize := strconv.StringToInt(viper.GetString("file.videoMaxSize"))
	for i := range files {
		if files[i].Size > int64(maxSize) {
			return true
		}
	}
	return false
}

func IsImageOverflow(files ...*multipart.FileHeader) bool {
	maxSize := strconv.StringToInt(viper.GetString("file.imageMaxSize"))
	for i := range files {
		if files[i].Size > int64(maxSize) {
			return true
		}
	}
	return false
}
