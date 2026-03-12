package qa

// StreamQaParams 流式问答参数
type StreamQaParams struct {
	Content         string `json:"content"`
	SessionID       int    `json:"session_id"`
	ClientMessageID string `json:"client_message_id"`
}

// GetQaInfoParams 获取问答信息参数
type GetQaInfoParams struct {
	SessionID int    `json:"session_id"`
	PageSize  int    `json:"page_size"`
	Cursor    string `json:"cursor"`
	UserID    int    `json:"user_id"`
}

// GetConversationListParams 获取对话信息列表参数
type GetConversationListParams struct {
	PageSize int    `json:"page_size"`
	Cursor   string `json:"cursor"`
	UserID   int    `json:"user_id"`
}
