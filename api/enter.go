package api

import (
	"sync"
)

type ApiGroup struct {
	UserApi
	CaptchaApi
	CommentApi
	UploadApi
	VideoApi
	LikeApi
	CollectApi
	FollowApi
}

var ApiGroupApp *ApiGroup
var syn sync.Once

func init() {
	syn.Do(func() {
		ApiGroupApp = new(ApiGroup)
	})
}
