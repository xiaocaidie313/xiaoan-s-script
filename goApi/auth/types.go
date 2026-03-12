package auth

// LoginType 登录类型
type LoginType string

const (
	LoginTypeEmail LoginType = "email"
	LoginTypeSMS   LoginType = "sms"
)

// RoleType 角色类型
type RoleType string

const (
	RoleTypeStudent RoleType = "student"
	RoleTypeTeacher RoleType = "teacher"
	RoleTypeAdmin   RoleType = "admin"
)

// EmailParams 邮箱验证参数
type EmailParams struct {
	Email string `json:"email"`
}

// RegisterParams 注册参数
type RegisterParams struct {
	Email         string `json:"email"`
	EmailCode     string `json:"email_code"`
	InviteCodeUsed string `json:"invite_code_used"`
	Password      string `json:"password"`
}

// LoginParams 登录参数
type LoginParams struct {
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Type     LoginType `json:"type"`
}

// GenerateInviteCodeParams 生成邀请码参数
type GenerateInviteCodeParams struct {
	ClassID    int      `json:"class_id"`
	Department string   `json:"department"`
	ExpiresAt  int64    `json:"expires_at"`
	MaxUses    int      `json:"max_uses"`
	Remark     string   `json:"remark"`
	TargetRole RoleType `json:"target_role"`
}

// GetInviteCodeParams 获取邀请码参数
type GetInviteCodeParams struct {
	PageSize int `json:"page_size"`
	Cursor   int `json:"cursor"`
	UserID   int `json:"user_id"`
}

// GetUserInfoParams 获取用户信息参数
type GetUserInfoParams struct {
	UserID int `json:"user_id"`
}

// GenerateClassParams 生成班级参数
type GenerateClassParams struct {
	ClassName        string `json:"class_name"`
	ClassDescription string `json:"class_description"`
}

// GetClassListParams 获取班级列表参数
type GetClassListParams struct {
	PageSize int `json:"page_size"`
	Cursor   int `json:"cursor"`
	UserID   int `json:"user_id"`
}

// GetClassInfoParams 获取班级信息参数
type GetClassInfoParams struct {
	ClassID int `json:"class_id"`
}
