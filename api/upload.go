package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-video/domain/dto"
	"go-gin-video/domain/resp"
	"go-gin-video/utils/upload"
)

type UploadApi struct {
}

// @Summary 上传视频
// @Tags upload
// @Router /upload/video [post]
func (UploadApi) UploadVideo(c *gin.Context) {
	videoDto := &dto.VideoDto{}
	err := c.ShouldBind(videoDto)
	if err != nil {
		resp.ErrorResponse(c, "参数绑定失败")
		return
	}
	if !upload.IsImageOverflow(&videoDto.VideoImage) || !upload.IsVideoOverflow(&videoDto.Video) {
		resp.ErrorResponse(c, "上传文件大小超出限制")
		return
	}
	if !upload.SaveVideo(c, videoDto) {
		resp.ErrorResponse(c, "视频信息保存失败")
		return
	}
	resp.OkResponse(c, "视频信息保存成功")
}
