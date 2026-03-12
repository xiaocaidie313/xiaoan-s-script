package qa

import (
	"xiaoan-s-script/goApi"
)

// GetOrCreateSession 获取或创建问答会话
// 需要先调用 client.SetToken() 设置认证 Token
func GetOrCreateSession(client *goApi.Client) (*goApi.CommonResponse[SessionType], error) {
	var result goApi.CommonResponse[SessionType]
	err := client.Get("/api/qa/get-or-create-session", nil, &result)
	return &result, err
}

// StreamQa 流式问答接口
// 注意：此为 SSE 流式接口，当前实现为普通 POST，如需流式处理需另行实现
func StreamQa(client *goApi.Client, params StreamQaParams) (*goApi.CommonResponse[map[string]any], error) {
	var result goApi.CommonResponse[map[string]any]
	err := client.Post("/api/qa/sse/ask", params, &result)
	return &result, err
}

// GetQaInfo 获取问答信息
func GetQaInfo(client *goApi.Client, params GetQaInfoParams) (*goApi.CommonResponse[GetQaInfoSuccess], error) {
	var result goApi.CommonResponse[GetQaInfoSuccess]
	err := client.Get("/api/qa/get-messages", params, &result)
	return &result, err
}

// GetConversationList 获取对话信息列表
func GetConversationList(client *goApi.Client, params GetConversationListParams) (*goApi.CommonResponse[GetConversationListSuccess], error) {
	var result goApi.CommonResponse[GetConversationListSuccess]
	err := client.Get("/api/qa/get-sessions", params, &result)
	return &result, err
}
