package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-video/domain/resp"
	"go-gin-video/service"
	"go-gin-video/utils/strconv"
)

type VideoApi struct {
}

var videoService = service.ServiceApp.VideService

// @Summary 根据视频id获取视频信息
// @Tags video
// @Router /video/get [get]
func (VideoApi) ByIdVideo(c *gin.Context) {
	vid := strconv.StringToInt(c.Query("vid"))
	err, videoVo := videoService.ByIdVide(vid)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, videoVo)
}

// @Summary 分页查询视频
// @Tags video
// @Router /video/list [get]
func (VideoApi) GetVideoList(c *gin.Context) {
	pagesize := strconv.StringToInt(c.Query("pagesize"))
	if pagesize > 10 {
		resp.ErrorResponse(c, "大小查出限制")
		return
	}
	page := strconv.StringToInt(c.Query("page"))
	err, videoVos := videoService.GetVideoList(page, pagesize)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, videoVos)
}

// @Summary 关键字查询视频
// @Tags video
// @Router /video/search [get]
func (VideoApi) SearchVideoList(c *gin.Context) {
	key := c.Query("search")
	page := strconv.StringToInt(c.Query("page"))
	pagesize := strconv.StringToInt(c.Query("pagesize"))
	err, videoVos := videoService.SearchVideoList(key, page, pagesize)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, videoVos)
}

// @Summary 通过用户id查询视频
// @Tags video
// @Router /video/user [get]
func (VideoApi) UserIdFindVideo(c *gin.Context) {
	userid := strconv.StringToInt(c.Query("userid"))
	page := strconv.StringToInt(c.Query("page"))
	pagesize := strconv.StringToInt(c.Query("pagesize"))
	err, videoVos := videoService.UserIdFindVideos(userid, page, pagesize)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, videoVos)
}

// @Summary 查询用户所有未过审视频,用于用于自己查询自己的视频
// @Tags video
// @Router /video/checkfalse [get]
func (VideoApi) SelectNotCheck(c *gin.Context) {
	userid := strconv.StringToInt(c.GetString("userid"))
	page := strconv.StringToInt(c.Query("page"))
	pagesize := strconv.StringToInt(c.Query("pagesize"))
	err, videoVos := videoService.SelectNotCheck(userid, page, pagesize)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, videoVos)
}

// @Summary 查询用户所有未过审视频,用于root用户查询用来审核视频使用
// @Tags video
// @Router /root/video/status [get]
func (VideoApi) SelectNotCheckAll(c *gin.Context) {
	page := strconv.StringToInt(c.Query("page"))
	pagesize := strconv.StringToInt(c.Query("pagesize"))
	err, videoVos := videoService.SelectNotCheckAll(page, pagesize)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, videoVos)
}

// @Summary 修改视频状态
// @Tags video
// @Router /root/video/check [get]
func (VideoApi) UpdateVideoCheck(c *gin.Context) {
	vid := strconv.StringToInt(c.Query("vid"))
	check := strconv.StringToInt(c.Query("check"))
	err := videoService.CheckVideoStatus(vid, check)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, "视频状态信息修改成功")
}
