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

type FollowService struct{}

func (FollowService) IsFollow(uid, followUid string) (bool, error) {
	var queryFollow = mysqlService.Follow
	count, err := queryFollow.Where(queryFollow.UID.Eq(int32(strconv.StringToInt(uid))), queryFollow.FollowUID.Eq(int32(strconv.StringToInt(followUid)))).Count()
	if err != nil {
		zap.L().Error(err.Error())
		return false, errors.New("查询失败")
	}
	return count > 0, nil
}

func (f FollowService) AddFollow(uid, followUid string) error {
	if b, err := f.IsFollow(uid, followUid); err != nil || b {
		if b {
			return errors.New("已关注无需再次关注")
		}
		if err != nil {
			return errors.New("查询失败")
		}
	}
	var queryFollow = mysqlService.Follow
	err := queryFollow.Create(&model.Follow{
		UID:       int32(strconv.StringToInt(uid)),
		FollowUID: int32(strconv.StringToInt(followUid)),
	})
	if err != nil {
		zap.L().Error(err.Error())
		return errors.New("关注失败")
	}
	return nil
}

func (f FollowService) DelFollow(uid, followUid string) error {
	if b, err := f.IsFollow(uid, followUid); err != nil || !b {
		if !b {
			return errors.New("没有关注无需取消")
		}
		if err != nil {
			return errors.New("查询失败")
		}
	}
	var queryFollow = mysqlService.Follow
	result, err := queryFollow.Where(queryFollow.UID.Eq(int32(strconv.StringToInt(uid))), queryFollow.FollowUID.Eq(int32(strconv.StringToInt(followUid)))).Delete()
	if err != nil || result.RowsAffected <= 0 {
		zap.L().Error(err.Error())
		return errors.New("取消关注失败")
	}
	return nil
}

// 查询被关注的数量
func (FollowService) GetFollowMeCount(uid string) (int64, error) {
	var queryFollow = mysqlService.Follow
	count, err := queryFollow.Where(queryFollow.UID.Eq(int32(strconv.StringToInt(uid)))).Count()
	if err != nil {
		zap.L().Error(err.Error())
		return 0, errors.New("查询失败")
	}
	return count, nil
}

// 查询被关注的人
func (FollowService) GetFollowMeUser(uid string) ([]*vo.UserVo, error) {
	var queryFollow = mysqlService.Follow
	follows, err := queryFollow.Where(queryFollow.UID.Eq(int32(strconv.StringToInt(uid)))).Find()
	if err != nil {
		zap.L().Error(err.Error())
		return nil, errors.New("查询失败")
	}
	followId := dto.GetFollowModelFollowId(follows)
	queryUser := mysqlService.User
	followUser, err := queryUser.Where(queryFollow.UID.In(followId...)).Find()
	if err != nil {
		zap.L().Error(err.Error())
		return nil, errors.New("查询失败")
	}
	toUserVo := dto.UserModelListToUserVoList(followUser)
	return toUserVo, nil
}

// 查询关注的数量
func (FollowService) GetMeFollowCount(uid string) (int64, error) {
	var queryFollow = mysqlService.Follow
	count, err := queryFollow.Where(queryFollow.FollowUID.Eq(int32(strconv.StringToInt(uid)))).Count()
	if err != nil {
		zap.L().Error(err.Error())
		return 0, errors.New("查询失败")
	}
	return count, nil
}

// 查询关注的人
func (FollowService) GetMeFollowUser(uid string) ([]*vo.UserVo, error) {
	var queryFollow = mysqlService.Follow
	follows, err := queryFollow.Where(queryFollow.FollowUID.Eq(int32(strconv.StringToInt(uid)))).Find()
	if err != nil {
		zap.L().Error(err.Error())
		return nil, errors.New("查询失败")
	}
	followId := dto.GetFollowModelFollowId(follows)
	queryUser := mysqlService.User
	followUser, err := queryUser.Where(queryFollow.UID.In(followId...)).Find()
	if err != nil {
		zap.L().Error(err.Error())
		return nil, errors.New("查询失败")
	}
	toUserVo := dto.UserModelListToUserVoList(followUser)
	return toUserVo, nil
}
