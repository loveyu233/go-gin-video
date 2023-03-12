package mongoModel

import (
	"go-gin-video/domain/vo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Vid       uint               `json:"vid" bson:"vid"` //视频ID
	CreatedAt int64              `json:"created_at" bson:"created_at"`
	Content   string             `json:"content" bson:"content"` //内容
	Uid       uint               `json:"uid" bson:"uid"`         //用户ID
	Author    vo.UserVo          `json:"author" bson:"author"`
	Reply     []Reply            `json:"reply" bson:"reply"`
	At        []string           `json:"at" bson:"at"`
	IsDelete  bool               `json:"is_delete" bson:"is_delete"`
}

type Reply struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
	Content   string             `json:"content" bson:"content"` //内容
	Uid       uint               `json:"uid" bson:"uid"`         //用户ID
	Author    vo.UserVo          `json:"author" bson:"author"`
	At        []string           `json:"at" bson:"at"`
	IsDelete  bool               `json:"is_delete" bson:"is_delete"`
}

func CommentModelListToCommentVoList(comments []Comment) []vo.CommentVo {
	length := len(comments)
	newComments := make([]vo.CommentVo, length)
	for i := 0; i < length; i++ {
		newComments[i].ID = comments[i].ID
		newComments[i].Content = comments[i].Content
		newComments[i].CreatedAt = comments[i].CreatedAt
		newComments[i].Reply = ReplyModelListToReplyVoList(comments[i].Reply)
		newComments[i].At = comments[i].At
	}

	return newComments
}

func ReplyModelListToReplyVoList(replies []Reply) []vo.ReplyVo {
	length := len(replies)
	newReplies := make([]vo.ReplyVo, length)
	for i := 0; i < length; i++ {
		newReplies[i].ID = replies[i].ID
		newReplies[i].Content = replies[i].Content
		newReplies[i].CreatedAt = replies[i].CreatedAt
		newReplies[i].At = replies[i].At
	}
	return newReplies
}
