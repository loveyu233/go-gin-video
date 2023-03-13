package service

import (
	"errors"
	"go-gin-video/domain/dto"
	"go-gin-video/domain/model"
	"go-gin-video/domain/vo"
	"go-gin-video/service/mysqlService"
	"go-gin-video/utils/strconv"
	"go.uber.org/zap"
)

type CollectService struct {
}

func (CollectService) IsCollect(vid, uid string) (error, bool) {
	var queryCollect = mysqlService.Collect
	collects, err := queryCollect.Where(queryCollect.Vid.Eq(int32(strconv.StringToInt(vid)))).Where(queryCollect.UID.Eq(int32(strconv.StringToInt(uid)))).Find()
	if err != nil {
		zap.L().Error(err.Error())
		return err, false
	}
	return nil, len(collects) > 0
}

func (c CollectService) AddCollect(vid, uid string) error {
	err, is := c.IsCollect(vid, uid)
	if err != nil {
		return err
	}
	if is {
		return errors.New("已经添加过收藏了")
	}
	var queryCollect = mysqlService.Collect
	err = queryCollect.Create(&model.Collect{
		UID: int32(strconv.StringToInt(uid)),
		Vid: int32(strconv.StringToInt(vid)),
	})
	if err != nil {
		return errors.New("添加收藏失败")
	}
	return nil
}

func (c CollectService) DelCollect(vid, uid string) error {
	err, is := c.IsCollect(vid, uid)
	if err != nil {
		return err
	}
	if !is {
		return errors.New("未收藏无需取消")
	}
	var queryCollect = mysqlService.Collect
	result, err := queryCollect.Where(queryCollect.Vid.Eq(int32(strconv.StringToInt(vid)))).Where(queryCollect.UID.Eq(int32(strconv.StringToInt(uid)))).Delete()
	if err != nil {
		zap.L().Error(err.Error())
		return errors.New("取消收藏失败")
	}
	if result.RowsAffected == 0 {
		return errors.New("取消收藏失败")
	}
	return nil
}

func (CollectService) CollectCount(vid string) (error, int64) {
	var queryCollect = mysqlService.Collect
	count, err := queryCollect.Where(queryCollect.Vid.Eq(int32(strconv.StringToInt(vid)))).Count()
	if err != nil {
		zap.L().Error(err.Error())
		return errors.New("查询失败"), 0
	}
	return nil, count
}

func (CollectService) GetUserCollect(uid string) (error, []*vo.VideoVo) {
	var queryCollect = mysqlService.Collect
	collects, err := queryCollect.Where(queryCollect.UID.Eq(int32(strconv.StringToInt(uid)))).Find()
	if err != nil {
		zap.L().Error(err.Error())
		return errors.New("查询失败"), nil
	}
	queryVideo := mysqlService.Video
	video, err := queryVideo.Where(queryVideo.Vid.In(dto.GetCollectVid(collects)...)).Find()
	if err != nil {
		zap.L().Error(err.Error())
		return errors.New("查询失败"), nil
	}
	videoVoList := dto.VideoModelListToVideoVoList(video)
	return nil, videoVoList
}
