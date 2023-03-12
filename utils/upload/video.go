package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-gin-video/domain/dto"
	"go-gin-video/service/mysqlService"
	"go-gin-video/utils/random"
	"go-gin-video/utils/strconv"
	"go-gin-video/utils/videoSplit"
	"go.uber.org/zap"
	"strings"
)

const (
	videoPath = "./static/resources/video/"
	imagePath = "./static/resources/images/"
)

func SaveUserIconAndBg(c *gin.Context, u *dto.UserDto) (bool, string, string) {
	userIconUuidName := uuid.New().String() + "." + strings.Split(u.UserIcon.Filename, ".")[1]
	userBgUuidName := uuid.New().String() + "." + strings.Split(u.UserBg.Filename, ".")[1]
	err := c.SaveUploadedFile(&u.UserIcon, imagePath+userIconUuidName)
	if err != nil {
		return false, "", ""
	}
	err = c.SaveUploadedFile(&u.UserBg, imagePath+userBgUuidName)
	if err != nil {
		return false, "", ""
	}
	return true, imagePath + userIconUuidName, imagePath + userBgUuidName
}

func SaveVideo(c *gin.Context, v *dto.VideoDto) bool {
	filename := random.GenerateVideoFilename()
	videoUuidName := filename + "/" + uuid.New().String() + "." + strings.Split(v.Video.Filename, ".")[1]
	videoImageUuidName := filename + "/" + uuid.New().String() + "." + strings.Split(v.VideoImage.Filename, ".")[1]
	err := c.SaveUploadedFile(&v.Video, videoPath+videoUuidName)
	if err != nil {
		zap.L().Error(err.Error())
		return false
	}
	err = c.SaveUploadedFile(&v.VideoImage, imagePath+videoImageUuidName)
	if err != nil {
		zap.L().Error(err.Error())
		return false
	}
	//分辨率,视频长度
	quality, _, err := videoSplit.PreTreatmentVideo(imagePath + videoUuidName)
	if err != nil {
		return false
	}
	// 转码
	go videoSplit.VideoTransCoding(c.GetUint("userid"), quality, "./static/resources/video/"+filename, videoUuidName)
	queryVideo := mysqlService.Video
	videoModel := dto.VideoDtoToVideoModel(v, int32(strconv.StringToInt(c.GetString("userid"))))
	videoModel.Video = filename + "/" + videoUuidName
	videoModel.VideoImage = filename + "/" + videoImageUuidName
	err = queryVideo.Save(videoModel)
	if err != nil {
		return false
	}
	return true
}
