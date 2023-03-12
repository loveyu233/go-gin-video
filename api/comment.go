package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-video/domain/dto"
	"go-gin-video/domain/mongoModel"
	"go-gin-video/domain/resp"
	"go-gin-video/service"
	"go-gin-video/service/mongoDBService"
	"go-gin-video/utils/strconv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type CommentApi struct {
}

var commentService = service.ServiceApp.CommentService

// @Summary 添加评论
// @Tags comment
// @Router /comment/insert [post]
func (CommentApi) InsertComment(c *gin.Context) {
	commentDto := &dto.CommentDTO{}
	err := c.Bind(commentDto)
	if err != nil {
		zap.L().Error(err.Error())
		resp.ErrorResponse(c, "参数绑定错误")
		return
	}
	userid := strconv.StringToInt(c.GetString("userid"))
	err, commentId := commentService.InsertComment(commentDto, userid)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, commentId)
}

// @Summary 删除评论
// @Tags comment
// @Router /comment/delete [get]
func (CommentApi) DeleteComment(c *gin.Context) {
	commentId := c.Query("commentId")
	if commentId == "" {
		resp.ErrorResponse(c, "评论id不能为空")
		return
	}
	err := commentService.DelComment(commentId)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, "评论删除成功")
}

// @Summary 添加回复
// @Tags comment
// @Router /reply/add [post]
func (CommentApi) InsertReply(c *gin.Context) {
	replyDto := &dto.ReplyDTO{}
	err := c.Bind(replyDto)
	if err != nil {
		zap.L().Error(err.Error())
		resp.ErrorResponse(c, "参数解析失败")
		return
	}
	userid := strconv.StringToInt(c.GetString("userid"))
	err, replyid := commentService.AddReply(replyDto, userid)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, replyid)
}

// @Summary 获取评论信息
// @Tags comment
// @Router /comment/get [get]
func (CommentApi) GetComment(c *gin.Context) {
	vid := strconv.StringToUint(c.DefaultQuery("vid", "0"))
	page := strconv.StringToInt(c.DefaultQuery("page", "1"))
	pagesize := strconv.StringToInt(c.DefaultQuery("pagesize", "10"))
	commentList, err := mongoDBService.SelectCommentList(vid, page, pagesize)
	if err != nil {
		zap.L().Error(err.Error())
		resp.ErrorResponse(c, "获取评论失败")
		return
	}
	resp.OkResponse(c, mongoModel.CommentModelListToCommentVoList(commentList))
}

// @Summary 删除回复信息
// @Tags comment
// @Router /reply/delete [get]
func (CommentApi) DelReply(c *gin.Context) {
	commentid, _ := primitive.ObjectIDFromHex(c.Query("commentid"))
	replyid, _ := primitive.ObjectIDFromHex(c.Query("replyid"))
	err := mongoDBService.DeleteReply(commentid, replyid)
	if err != nil {
		zap.L().Error("回复删除失败:" + err.Error())
		resp.ErrorResponse(c, "回复删除失败")
		return
	}
	resp.OkResponse(c, "回复删除成功")
}

// @Summary 通过评论id获取评论
// @Tags comment
// @Router /comment/byid [get]
func (CommentApi) ByIdComment(c *gin.Context) {
	commentid, _ := primitive.ObjectIDFromHex(c.Query("commentid"))
	comment, err := mongoDBService.SelectCommentById(commentid)
	if err != nil {
		resp.ErrorResponse(c, "评论查询失败")
		return
	}
	commentVo := dto.CommentModelToCommentVo(&comment)
	resp.OkResponse(c, commentVo)
}
