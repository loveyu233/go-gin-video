package service

import (
	"errors"
	"go-gin-video/domain/dto"
	"go-gin-video/domain/vo"
	"go-gin-video/service/mysqlService"
)

type VideService struct {
}

func (VideService) ByIdVide(vid int) (error, *vo.VideoVo) {
	queryVideo := mysqlService.Video
	videoModel, err := queryVideo.Where(queryVideo.Ischeck.Is(true)).Where(queryVideo.Isdelete.Is(false)).Where(queryVideo.Vid.Eq(int32(vid))).First()
	if err.Error() == "record not found" {
		return nil, nil
	}
	if err != nil {
		return errors.New("视频查询失败"), nil
	}
	return nil, dto.VideoModelToVideoVo(videoModel)
}

func (VideService) GetVideoList(page, pagesize int) (error, []*vo.VideoVo) {
	queryVideo := mysqlService.Video
	videoModelList, err := queryVideo.Where(queryVideo.Ischeck.Is(true)).Where(queryVideo.Isdelete.Is(false)).Limit(pagesize).Offset((page - 1) * pagesize).Find()
	if err != nil {
		return errors.New("视频查询失败"), nil
	}
	return nil, dto.VideoModelListToVideoVoList(videoModelList)
}

func (VideService) SearchVideoList(key string, page, pagesize int) (error, []*vo.VideoVo) {
	queryVideo := mysqlService.Video
	videoModel, err := queryVideo.Where(queryVideo.Ischeck.Is(true)).Where(queryVideo.Isdelete.Is(false)).Where(queryVideo.VideoTitle.Like("%" + key + "%")).Limit(pagesize).Offset((page - 1) * pagesize).Find()
	if err != nil {
		return errors.New("视频查询失败"), nil
	}
	return nil, dto.VideoModelListToVideoVoList(videoModel)
}

func (VideService) UserIdFindVideos(uid, page, pagesize int) (error, []*vo.VideoVo) {
	queryVideo := mysqlService.Video
	videoModelList, err := queryVideo.Where(queryVideo.Ischeck.Is(true)).Where(queryVideo.Isdelete.Is(false)).Where(queryVideo.UserID.Eq(int32(uid))).Limit(pagesize).Offset((page - 1) * pagesize).Find()
	if err != nil {
		return errors.New("视频查询失败"), nil
	}
	return nil, dto.VideoModelListToVideoVoList(videoModelList)
}

func (VideService) SelectNotCheck(uid, page, pagesize int) (error, []*vo.VideoVo) {
	queryVideo := mysqlService.Video
	videoModelList, err := queryVideo.Where(queryVideo.Ischeck.Is(false)).Where(queryVideo.Isdelete.Is(false)).Where(queryVideo.UserID.Eq(int32(uid))).Limit(pagesize).Offset((page - 1) * pagesize).Find()
	if err != nil {
		return errors.New("视频查询失败"), nil
	}
	return nil, dto.VideoModelListToVideoVoList(videoModelList)
}

func (VideService) SelectNotCheckAll(page, pagesize int) (error, []*vo.VideoVo) {
	queryVideo := mysqlService.Video
	videoModelList, err := queryVideo.Where(queryVideo.Ischeck.Is(false)).Where(queryVideo.Isdelete.Is(false)).Limit(pagesize).Offset((page - 1) * pagesize).Find()
	if err != nil {
		return errors.New("视频查询失败"), nil
	}
	return nil, dto.VideoModelListToVideoVoList(videoModelList)
}

func (VideService) CheckVideoStatus(vid, check int) error {
	queryVideo := mysqlService.Video
	switch check {
	case 1:
		_, err := queryVideo.Where(queryVideo.Vid.Eq(int32(vid))).Update(queryVideo.Ischeck, true)
		if err != nil {
			return errors.New("视频状态信息修改失败")
		}
	case 2:
		_, err := queryVideo.Where(queryVideo.Vid.Eq(int32(vid))).Update(queryVideo.Isdelete, true)
		if err != nil {
			return errors.New("视频状态信息修改失败")
		}
	default:
		return errors.New("视频状态信息不合法")
	}
	return nil
}
