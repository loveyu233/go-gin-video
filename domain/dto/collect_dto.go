package dto

import "go-gin-video/domain/model"

func GetCollectVid(c []*model.Collect) []int32 {
	vid := make([]int32, 0, len(c))
	for i := range c {
		vid = append(vid, c[i].Vid)
	}
	return vid
}
