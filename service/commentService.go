package service

import (
	"errors"
	"go-gin-video/domain/dto"
	"go-gin-video/service/mongoDBService"
	"go-gin-video/service/mysqlService"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type CommentService struct {
}

func (CommentService) InsertComment(commentDto *dto.CommentDTO, uid int) (error, *primitive.ObjectID) {
	queryUser := mysqlService.User
	user, _ := queryUser.Where(queryUser.UserID.Eq(int32(uid))).First()
	commentModel := commentDto.ToCommentModel(uint(uid))
	commentModel.Author = *dto.UserModelToUserVo(user)
	commentId, err := mongoDBService.InsertComment(commentModel)
	if err != nil {
		zap.L().Error(err.Error())
		return errors.New("评论添加失败"), nil
	}
	return nil, &commentId
}

func (CommentService) DelComment(commentId string) error {
	hex, _ := primitive.ObjectIDFromHex(commentId)
	err := mongoDBService.DeleteComment(hex)
	if err != nil {
		zap.L().Error(err.Error())
		return errors.New("评论删除失败")
	}
	return nil
}

func (CommentService) AddReply(replyDto *dto.ReplyDTO, uid int) (error, *primitive.ObjectID) {
	replyModel := replyDto.ToReplyModel(uint(uid))
	queryUser := mysqlService.User
	user, _ := queryUser.Where(queryUser.UserID.Eq(int32(uid))).First()
	replyModel.Author = *dto.UserModelToUserVo(user)
	replyid, err := mongoDBService.InsertReply(replyDto.ParentID, replyModel)
	if err != nil {
		zap.L().Error(err.Error())
		return errors.New("回复添加失败"), nil
	}
	return nil, &replyid
}
