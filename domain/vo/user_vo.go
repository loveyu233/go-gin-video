package vo

type UserVo struct {
	Username string `json:"username"`  // 用户名
	UserIcon string `json:"user_icon"` // 用户头像
	Sing     string `json:"sing"`      // 个性签名
	UserBg   string `json:"user_bg"`   // 用户主页背景
}
