package auth

// LoginSuccess 登录成功响应
type LoginSuccess struct {
	Token string `json:"token"`
}

// RegisterSuccess 注册成功响应
type RegisterSuccess struct {
	Token string `json:"token"`
}

// GenerateInviteCodeSuccess 生成邀请码成功响应
type GenerateInviteCodeSuccess struct {
	InviteCode string `json:"invite_code"`
}

// GetInviteCodeSuccess 获取邀请码成功响应
type GetInviteCodeSuccess struct {
	List  []any `json:"list"`
	Total int   `json:"total"`
}

// UserInfo 用户信息
type UserInfo struct {
	UserID   int    `json:"user_id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Role     string `json:"role"`
}

// GetClassListSuccess 获取班级列表成功响应
type GetClassListSuccess struct {
	List  []ClassInfo `json:"list"`
	Total int         `json:"total"`
}

// ClassInfo 班级信息
type ClassInfo struct {
	ClassID   int    `json:"class_id"`
	ClassName string `json:"class_name"`
}
