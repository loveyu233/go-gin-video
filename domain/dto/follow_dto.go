package dto

import "go-gin-video/domain/model"

func GetFollowModelFollowId(fm []*model.Follow) []int32 {
	followid := make([]int32, 0, len(fm))
	for i := range fm {
		followid = append(followid, fm[i].FollowUID)
	}
	return followid
}
