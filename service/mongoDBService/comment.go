package mongoDBService

import (
	"context"
	"errors"
	"go-gin-video/domain/mongoModel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 添加评论,返回评论的id
func InsertComment(comment *mongoModel.Comment) (primitive.ObjectID, error) {
	_, err := mongoClient.Comment().InsertOne(context.TODO(), comment)
	return comment.ID, err
}

// 添加回复,返回回复的id
func InsertReply(commentId primitive.ObjectID, reply *mongoModel.Reply) (primitive.ObjectID, error) {
	_, err := mongoClient.Comment().UpdateOne(context.TODO(), bson.M{"_id": commentId}, bson.M{
		"$addToSet": bson.M{
			"reply": reply,
		},
	})
	return reply.ID, err
}

// 只查询评论,评论后后端给前端发送查询评论的id,前端在用评论id查出这条评论进行展示
func SelectCommentById(commentId primitive.ObjectID) (mongoModel.Comment, error) {
	aggregate, err := mongoClient.Comment().Aggregate(context.TODO(), bson.A{
		bson.M{
			"$match": bson.M{
				"_id": commentId,
			},
		},
		bson.M{
			"$project": bson.M{
				"uid":        "$uid",
				"content":    "$content",
				"created_at": "$created_at",
			},
		},
		bson.M{
			"$limit": 1,
		},
	})
	if err != nil {
		return mongoModel.Comment{}, err
	}
	var comment []mongoModel.Comment
	if err := aggregate.All(context.TODO(), &comment); err != nil {
		return mongoModel.Comment{}, err
	}
	return comment[0], nil
}

// 分页查询评论
func SelectCommentList(videoId uint, page, pageSize int) ([]mongoModel.Comment, error) {
	var comments []mongoModel.Comment
	cursor, err := mongoClient.Comment().Aggregate(context.TODO(), bson.A{
		bson.M{
			"$match": bson.M{
				"vid":       videoId,
				"is_delete": false,
			},
		},
		bson.M{
			"$project": bson.M{
				"uid":        "$uid",
				"content":    "$content",
				"created_at": "$created_at",
				"reply": bson.M{
					"$filter": bson.M{
						"input": "$reply",
						"as":    "item",
						"cond": bson.M{
							"$eq": bson.A{"$$item.is_delete", false},
						},
					},
				},
			},
		},
		bson.M{
			"$project": bson.M{
				"uid":        "$uid",
				"content":    "$content",
				"created_at": "$created_at",
				"reply": bson.M{
					"$slice": bson.A{"$reply", 0, 2},
				},
			},
		},
		bson.M{
			"$skip": (page - 1) * pageSize,
		},
		bson.M{
			"$limit": pageSize,
		},
	})

	if err != nil {
		return comments, err
	}

	if err := cursor.All(context.TODO(), &comments); err != nil {
		return comments, err
	}

	return comments, nil
}

// 和查询评论一样
func SelectReplyById(commentId, replyId primitive.ObjectID) (mongoModel.Reply, error) {
	var replies []mongoModel.Reply
	cursor, err := mongoClient.Comment().Aggregate(context.TODO(), bson.A{
		bson.M{
			"$match": bson.M{
				"_id": commentId,
			},
		},
		bson.M{
			"$project": bson.M{
				"reply": bson.M{
					"$filter": bson.M{
						"input": "$reply",
						"as":    "item",
						"cond": bson.M{
							"$eq": bson.A{"$$item._id", replyId},
						},
					},
				},
			},
		},
		bson.M{
			"$unwind": "$reply",
		},
		bson.M{
			"$project": bson.M{
				"_id":        "$reply._id",
				"uid":        "$reply.uid",
				"content":    "$reply.content",
				"created_at": "$reply.created_at",
			},
		},
		bson.M{
			"$limit": 1,
		},
	})

	if err != nil {
		return mongoModel.Reply{}, err
	}

	if err := cursor.All(context.TODO(), &replies); err != nil {
		return mongoModel.Reply{}, err
	}

	if len(replies) == 0 {
		return mongoModel.Reply{}, errors.New("获取回复失败")
	}

	return replies[0], nil
}

// 分页查询回复
func SelectReplyList(commentId primitive.ObjectID, page, pageSize int) ([]mongoModel.Reply, error) {
	var replies []mongoModel.Reply
	cursor, err := mongoClient.Comment().Aggregate(context.TODO(), bson.A{
		bson.M{
			"$match": bson.M{
				"_id":       commentId,
				"is_delete": false,
			},
		},
		bson.M{
			"$project": bson.M{
				"reply": bson.M{
					"$filter": bson.M{
						"input": "$reply",
						"as":    "item",
						"cond": bson.M{
							"$eq": bson.A{"$$item.is_delete", false},
						},
					},
				},
			},
		},
		bson.M{
			"$unwind": "$reply",
		},
		bson.M{
			"$project": bson.M{
				"_id":        "$reply._id",
				"uid":        "$reply.uid",
				"content":    "$reply.content",
				"created_at": "$reply.created_at",
			},
		},
		bson.M{
			"$skip": (page-1)*pageSize + 2,
		},
		bson.M{
			"$limit": pageSize,
		},
	})

	if err != nil {
		return replies, err
	}

	if err := cursor.All(context.TODO(), &replies); err != nil {
		return replies, err
	}

	return replies, nil
}

// 删除评论
func DeleteComment(commentId primitive.ObjectID) error {
	_, err := mongoClient.Comment().UpdateOne(context.TODO(), bson.M{"_id": commentId}, bson.M{
		"$set": bson.M{
			"is_delete": true,
		},
	})
	return err
}

// 删除回复
func DeleteReply(commentId, replyId primitive.ObjectID) error {
	filter := bson.M{
		"_id":       commentId,
		"reply._id": replyId,
	}

	_, err := mongoClient.Comment().UpdateOne(context.TODO(), filter, bson.M{
		"$set": bson.M{
			"reply.$.is_delete": true,
		},
	})

	return err
}
