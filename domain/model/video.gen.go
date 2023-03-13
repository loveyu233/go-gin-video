// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameVideo = "video"

// Video mapped from table <video>
type Video struct {
	Vid               int32     `gorm:"column:vid" json:"vid"`
	VideoImage        string    `gorm:"column:video_image" json:"video_image"`
	Video             string    `gorm:"column:video" json:"video"`
	VideoTitle        string    `gorm:"column:video_title" json:"video_title"`
	VideoIntroduction string    `gorm:"column:video_introduction" json:"video_introduction"`
	UserID            int32     `gorm:"column:user_id" json:"user_id"`
	CreateTime        time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	Ischeck           bool      `gorm:"column:ischeck" json:"ischeck"`
	Isdelete          bool      `gorm:"column:isdelete;not null" json:"isdelete"`
}

// TableName Video's table name
func (*Video) TableName() string {
	return TableNameVideo
}
