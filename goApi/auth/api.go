package auth

import (
	"xiaoan-s-script/goApi"
)

// SendEmailCode 发送邮箱验证码
func SendEmailCode(client *goApi.Client, params EmailParams) (*goApi.CommonResponse[map[string]any], error) {
	var result goApi.CommonResponse[map[string]any]
	err := client.Post("/api/auth/send-email", params, &result)
	return &result, err
}

// Register 注册
func Register(client *goApi.Client, params RegisterParams) (*goApi.CommonResponse[RegisterSuccess], error) {
	var result goApi.CommonResponse[RegisterSuccess]
	err := client.Post("/api/auth/register", params, &result)
	return &result, err
}

// Login 登录
func Login(client *goApi.Client, params LoginParams) (*goApi.CommonResponse[LoginSuccess], error) {
	var result goApi.CommonResponse[LoginSuccess]
	err := client.Post("/api/auth/login", params, &result)
	return &result, err
}

// GenerateInviteCode 生成邀请码
func GenerateInviteCode(client *goApi.Client, params GenerateInviteCodeParams) (*goApi.CommonResponse[GenerateInviteCodeSuccess], error) {
	var result goApi.CommonResponse[GenerateInviteCodeSuccess]
	err := client.Post("/api/auth/generate-invite-code", params, &result)
	return &result, err
}

// GetInviteCode 获取邀请码
func GetInviteCode(client *goApi.Client, params GetInviteCodeParams) (*goApi.CommonResponse[GetInviteCodeSuccess], error) {
	var result goApi.CommonResponse[GetInviteCodeSuccess]
	err := client.Get("/api/auth/get-invite-code", params, &result)
	return &result, err
}

// GetUserInfo 获取用户信息
func GetUserInfo(client *goApi.Client, params GetUserInfoParams) (*goApi.CommonResponse[UserInfo], error) {
	var result goApi.CommonResponse[UserInfo]
	err := client.Get("/api/auth/get-user-info", params, &result)
	return &result, err
}

// UpdateUserInfo 修改用户信息
func UpdateUserInfo(client *goApi.Client, params map[string]any) (*goApi.CommonResponse[UserInfo], error) {
	var result goApi.CommonResponse[UserInfo]
	err := client.Post("/api/auth/modify-user-base-info", params, &result)
	return &result, err
}

// GenerateClass 生成班级
func GenerateClass(client *goApi.Client, params GenerateClassParams) (*goApi.CommonResponse[map[string]any], error) {
	var result goApi.CommonResponse[map[string]any]
	err := client.Post("/api/auth/generate-class", params, &result)
	return &result, err
}

// GetClassList 获取班级列表
func GetClassList(client *goApi.Client, params GetClassListParams) (*goApi.CommonResponse[GetClassListSuccess], error) {
	var result goApi.CommonResponse[GetClassListSuccess]
	err := client.Get("/api/auth/get-class-list", params, &result)
	return &result, err
}

// GetClassInfo 获取班级信息
func GetClassInfo(client *goApi.Client, params GetClassInfoParams) (*goApi.CommonResponse[map[string]any], error) {
	var result goApi.CommonResponse[map[string]any]
	err := client.Get("/api/auth/get-class-info", params, &result)
	return &result, err
}

// SwitchClass 切换班级
func SwitchClass(client *goApi.Client, inviteCode string) (*goApi.CommonResponse[map[string]any], error) {
	params := map[string]string{"invite_code": inviteCode}
	var result goApi.CommonResponse[map[string]any]
	err := client.Post("/api/auth/switch-class", params, &result)
	return &result, err
}
