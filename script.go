/**
 * 小安计划 API 测试脚本 (Go)
 * 使用方式: 设置环境变量后运行 go run .
 *   BASE_URL=https://your-api.com EMAIL=xxx@example.com PASSWORD=xxx go run .
 */

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"xiaoan-s-script/goApi"
	"xiaoan-s-script/goApi/auth"
	"xiaoan-s-script/goApi/content"
)

func env(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}

func main() {
	baseURL := env("BASE_URL", "http://localhost:3000")
	email := env("EMAIL", "")
	password := env("PASSWORD", "")

	if email == "" || password == "" {
		fmt.Println("请设置环境变量: EMAIL, PASSWORD (可选 BASE_URL)")
		os.Exit(1)
	}

	client := goApi.NewClient(baseURL)
	fmt.Printf("[1/4] 连接: %s\n", baseURL)

	// 1. 登录
	fmt.Println("[2/4] 登录...")
	loginResp, err := auth.Login(client, auth.LoginParams{
		Email:    email,
		Password: password,
		Type:     auth.LoginTypeEmail,
	})
	if err != nil {
		fmt.Printf("登录请求失败: %v\n", err)
		os.Exit(1)
	}
	if loginResp.Code != 0 {
		fmt.Printf("登录失败: code=%d, message=%s\n", loginResp.Code, loginResp.Message)
		os.Exit(1)
	}
	client.SetToken(loginResp.Data.Token)
	fmt.Println("登录成功, token 已设置")

	// 2. 获取用户信息
	fmt.Println("[3/4] 获取用户信息...")
	userResp, err := auth.GetUserInfo(client, auth.GetUserInfoParams{UserID: 0})
	if err != nil {
		fmt.Printf("获取用户信息失败: %v\n", err)
		os.Exit(1)
	}
	if userResp.Code != 0 {
		fmt.Printf("获取用户信息失败: code=%d, message=%s\n", userResp.Code, userResp.Message)
		os.Exit(1)
	}
	userJSON, _ := json.MarshalIndent(userResp.Data, "", "  ")
	fmt.Printf("用户信息: %s\n", string(userJSON))

	// 3. 获取最新文章列表
	fmt.Println("[4/4] 获取最新文章...")
	articlesResp, err := content.GetNewArticles(client, content.CursorPageParams{
		PageSize: 10,
		Cursor:   0,
	})
	if err != nil {
		fmt.Printf("获取文章列表失败: %v\n", err)
		os.Exit(1)
	}
	if articlesResp.Code != 0 {
		fmt.Printf("获取文章列表失败: code=%d, message=%s\n", articlesResp.Code, articlesResp.Message)
		os.Exit(1)
	}
	fmt.Printf("文章数量: %d\n", articlesResp.Data.Total)

	fmt.Println("\n全部测试通过.")
}
