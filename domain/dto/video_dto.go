package dto

import (
	"go-gin-video/domain/model"
	"go-gin-video/domain/vo"
	"go-gin-video/service/mysqlService"
	"mime/multipart"
)

type VideoDto struct {
	VideoImage        multipart.FileHeader `gorm:"column:video_image" json:"video_image" form:"videoImage"`                      // 视频封面地址
	Video             multipart.FileHeader `gorm:"column:video" json:"video" form:"video"`                                       // 视频地址
	VideoTitle        string               `gorm:"column:video_title" json:"video_title" form:"videoTitle"`                      // 视频标题
	VideoIntroduction string               `gorm:"column:Video_introduction" json:"video_introduction" form:"videoIntroduction"` // 视频简介
}

func VideoDtoToVideoModel(vd *VideoDto, userid int32) *model.Video {
	return &model.Video{
		VideoImage:        vd.VideoImage.Filename,
		Video:             vd.Video.Filename,
		VideoTitle:        vd.VideoTitle,
		VideoIntroduction: vd.VideoIntroduction,
		UserID:            userid,
	}
}
func VideoModelToVideoVo(vm *model.Video) *vo.VideoVo {
	queryUser := mysqlService.User
	userModel, _ := queryUser.Where(queryUser.UserID.Eq(vm.UserID)).First()
	userVo := UserModelToUserVo(userModel)
	return &vo.VideoVo{
		Vid:               vm.Vid,
		VideoImage:        "http://localhost:9898/static/images/" + vm.VideoImage,
		Video:             "http://localhost:9898/static/video/" + vm.Video,
		VideoTitle:        vm.VideoTitle,
		VideoIntroduction: vm.VideoIntroduction,
		User:              userVo,
		CreatTime:         vm.CreatTime,
	}
}

func VideoModelListToVideoVoList(vs []*model.Video) []*vo.VideoVo {
	videVoList := make([]*vo.VideoVo, 0, len(vs))
	for i := range vs {
		videVoList = append(videVoList, &vo.VideoVo{
			Vid:               vs[i].Vid,
			VideoImage:        "http://localhost:9898/static/images/" + vs[i].VideoImage,
			Video:             "http://localhost:9898/static/video/" + vs[i].Video,
			VideoTitle:        vs[i].VideoTitle,
			VideoIntroduction: vs[i].VideoIntroduction,
			CreatTime:         vs[i].CreatTime,
		})
	}
	return videVoList
}
