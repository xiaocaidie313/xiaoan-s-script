package qa

// SessionType 会话类型
type SessionType struct {
	SessionID int `json:"session_id"`
}

// GetQaInfoSuccess 获取问答信息成功响应
type GetQaInfoSuccess struct {
	Messages []Message `json:"messages"`
	Total    int       `json:"total"`
}

// Message 消息
type Message struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	Role      string `json:"role"`
	CreatedAt int64  `json:"created_at"`
}

// GetConversationListSuccess 获取对话列表成功响应
type GetConversationListSuccess struct {
	List  []Conversation `json:"list"`
	Total int            `json:"total"`
}

// Conversation 对话
type Conversation struct {
	SessionID   int    `json:"session_id"`
	Title       string `json:"title"`
	LastMessage string `json:"last_message"`
	UpdatedAt   int64  `json:"updated_at"`
}
