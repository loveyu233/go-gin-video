package vo

type FollowVo struct {
	UID       int32 `gorm:"column:uid" json:"uid"`
	FollowUID int32 `gorm:"column:follow_uid" json:"follow_uid"`
}
