package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-video/domain/resp"
	"go-gin-video/service"
)

type CollectApi struct {
}

var collectService = service.ServiceApp.CollectService

// @Summary 判断用户是否收藏该视频
// @Tags collec
// @Router /collec/is [get]
func (CollectApi) IsCollect(c *gin.Context) {
	vid := c.Query("vid")
	uid := c.GetString("userid")
	err, is := collectService.IsCollect(vid, uid)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, is)
}

// @Summary 添加收藏该视频
// @Tags collec
// @Router /collec/add [get]
func (CollectApi) AddCollect(c *gin.Context) {
	vid := c.Query("vid")
	uid := c.GetString("userid")
	err := collectService.AddCollect(vid, uid)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
	}
	resp.OkResponse(c, "ok")
}

// @Summary 删除收藏该视频
// @Tags collec
// @Router /collec/del [get]
func (CollectApi) DelCollect(c *gin.Context) {
	vid := c.Query("vid")
	uid := c.GetString("userid")
	err := collectService.DelCollect(vid, uid)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
	}
	resp.OkResponse(c, "ok")
}

// @Summary 统计该视频的收藏量
// @Tags collec
// @Router /collec/count [get]
func (CollectApi) CollectCount(c *gin.Context) {
	vid := c.Query("vid")
	err, count := collectService.CollectCount(vid)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, count)
}

// @Summary 获取用户收藏的视频
// @Tags collec
// @Router /collec/usercollec [get]
func (CollectApi) GetUserCollect(c *gin.Context) {
	uid := c.GetString("userid")
	err, vos := collectService.GetUserCollect(uid)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, vos)
}
