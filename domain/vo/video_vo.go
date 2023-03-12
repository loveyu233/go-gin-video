package vo

import (
	"time"
)

type VideoVo struct {
	Vid               int32     `json:"vid"`                // 视频id
	VideoImage        string    `json:"video_image"`        // 视频封面地址
	Video             string    `json:"video"`              // 视频地址
	VideoTitle        string    `json:"video_title"`        // 视频标题
	VideoIntroduction string    `json:"Video_introduction"` // 视频简介
	User              *UserVo   `json:"user"`               // 视频作者id
	CreatTime         time.Time `json:"creat_time"`         // 视频创建时间
}
