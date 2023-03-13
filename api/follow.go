package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-video/domain/resp"
	"go-gin-video/service"
)

type FollowApi struct {
}

var followService = service.ServiceApp.FollowService

// @Summary 查询该是否关注了该用户
// @Tags follow
// @Router /follow/is [get]
func (FollowApi) IsFollow(c *gin.Context) {
	uid := c.Query("uid")
	followid := c.GetString("userid")
	is, err := followService.IsFollow(uid, followid)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, is)
}

// @Summary 添加关注
// @Tags follow
// @Router /follow/add [get]
func (FollowApi) AddFollow(c *gin.Context) {
	uid := c.Query("uid")
	followid := c.GetString("userid")
	err := followService.AddFollow(uid, followid)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, "关注成功")
}

// @Summary 取消关注
// @Tags follow
// @Router /follow/del [get]
func (FollowApi) DelFollow(c *gin.Context) {
	uid := c.Query("uid")
	followid := c.GetString("userid")
	err := followService.DelFollow(uid, followid)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, "取消关注成功")
}

// @Summary 查询被关注数量
// @Tags follow
// @Router /follow/followmecount [get]
func (FollowApi) GetFollowMeCount(c *gin.Context) {
	uid := c.Query("uid")
	count, err := followService.GetFollowMeCount(uid)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, count)
}

// @Summary 查询被关注的用户
// @Tags follow
// @Router /follow/followmeuser [get]
func (FollowApi) GetFollowMeUser(c *gin.Context) {
	uid := c.GetString("userid")
	users, err := followService.GetFollowMeUser(uid)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, users)
}

// @Summary 查询关注数量
// @Tags follow
// @Router /follow/mefollowcount [get]
func (FollowApi) GetMeFollowCount(c *gin.Context) {
	followid := c.GetString("userid")
	users, err := followService.GetMeFollowCount(followid)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, users)
}

// @Summary 查询被关注用户
// @Tags follow
// @Router /follow/mefollowuser [get]
func (FollowApi) GetMeFollowUser(c *gin.Context) {
	followid := c.GetString("userid")
	users, err := followService.GetMeFollowUser(followid)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, users)
}
