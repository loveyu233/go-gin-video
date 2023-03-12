package dto

import (
	"go-gin-video/domain/mongoModel"
	"go-gin-video/domain/vo"
	"go-gin-video/service/mysqlService"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CommentDTO struct {
	Vid     uint
	Content string
	At      []string
}

type ReplyDTO struct {
	Vid          uint
	Content      string
	ParentID     primitive.ObjectID
	At           []string
	ReplyUserID  uint
	ReplyContent string
}

type DeleteReplyDTO struct {
	CommentID primitive.ObjectID
	ReplyID   primitive.ObjectID
}

func (r *ReplyDTO) ToReplyModel(uid uint) *mongoModel.Reply {
	return &mongoModel.Reply{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now().Unix(),
		Content:   r.Content,
		Uid:       uid,
		Author:    vo.UserVo{},
		At:        r.At,
		IsDelete:  false,
	}
}

func (c *CommentDTO) ToCommentModel(uid uint) *mongoModel.Comment {
	return &mongoModel.Comment{
		ID:        primitive.NewObjectID(),
		Vid:       c.Vid,
		CreatedAt: time.Now().Unix(),
		Content:   c.Content,
		Uid:       uid,
		Author:    vo.UserVo{},
		Reply:     []mongoModel.Reply{},
		At:        nil,
		IsDelete:  false,
	}
}

func CommentModelToCommentVo(cm *mongoModel.Comment) *vo.CommentVo {
	queryUser := mysqlService.User
	userModel, _ := queryUser.Where(queryUser.UserID.Eq(int32(cm.Uid))).First()
	userVo := UserModelToUserVo(userModel)
	return &vo.CommentVo{
		ID:        cm.ID,
		Content:   cm.Content,
		Author:    *userVo,
		Reply:     mongoModel.ReplyModelListToReplyVoList(cm.Reply),
		CreatedAt: cm.CreatedAt,
		At:        cm.At,
	}
}
