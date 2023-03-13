package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-video/domain/resp"
	"go-gin-video/service/redisService"
)

type LikeApi struct {
}

// @Summary 判断用户是否点赞该视频
// @Tags like
// @Router /like/is [get]
func (LikeApi) IsLike(c *gin.Context) {
	vid := c.Query("vid")
	uid := c.GetString("userid")
	if redisService.IsLike(vid, uid) {
		resp.OkResponse(c, true)
		return
	}
	resp.OkResponse(c, false)
}

// @Summary 添加点赞
// @Tags like
// @Router /like/add [get]
func (LikeApi) AddLike(c *gin.Context) {
	vid := c.Query("vid")
	uid := c.GetString("userid")
	if redisService.IsLike(vid, uid) {
		resp.ErrorResponse(c, "已经点过赞了")
		return
	}
	redisService.AddLike(vid, uid)
	resp.OkResponse(c, "点赞成功")
}

// @Summary 删除点赞
// @Tags like
// @Router /like/del [get]
func (LikeApi) DelLike(c *gin.Context) {
	vid := c.Query("vid")
	uid := c.GetString("userid")
	if !redisService.IsLike(vid, uid) {
		resp.ErrorResponse(c, "未点赞无需取消")
		return
	}
	redisService.DelLike(vid, uid)
	resp.OkResponse(c, "取消点赞成功")
}

// @Summary 获取视频点赞个数
// @Tags like
// @Router /like/count [get]
func (LikeApi) LikeCount(c *gin.Context) {
	vid := c.Query("vid")
	resp.OkResponse(c, redisService.LikeCount(vid))
}
