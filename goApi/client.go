package goApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// CommonResponse 通用 API 响应结构
type CommonResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

// Client API 客户端
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string
}

// NewClient 创建新的 API 客户端
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

// SetToken 设置认证 Token
func (c *Client) SetToken(token string) {
	c.Token = token
}

// doRequest 执行 HTTP 请求
func (c *Client) doRequest(method, path string, params any, body any) (*http.Response, error) {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewReader(jsonBody)
	}

	reqURL := c.BaseURL + path
	if params != nil {
		values := c.structToValues(params)
		if len(values) > 0 {
			reqURL += "?" + values.Encode()
		}
	}

	req, err := http.NewRequest(method, reqURL, reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	return c.HTTPClient.Do(req)
}

// structToValues 将结构体或 map 转为 url.Values (用于 GET/DELETE 请求参数)
func (c *Client) structToValues(v any) url.Values {
	values := url.Values{}
	if v == nil {
		return values
	}
	data, err := json.Marshal(v)
	if err != nil {
		return values
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return values
	}
	for k, val := range m {
		if val != nil {
			values.Set(k, fmt.Sprintf("%v", val))
		}
	}
	return values
}

// Get 执行 GET 请求
func (c *Client) Get(path string, params any, result any) error {
	resp, err := c.doRequest(http.MethodGet, path, params, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(result)
}

// Post 执行 POST 请求
func (c *Client) Post(path string, body any, result any) error {
	resp, err := c.doRequest(http.MethodPost, path, nil, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(result)
}

// Delete 执行 DELETE 请求
func (c *Client) Delete(path string, params any, result any) error {
	resp, err := c.doRequest(http.MethodDelete, path, params, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(result)
}
