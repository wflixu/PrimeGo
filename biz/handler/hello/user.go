package hello

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// User 用户结构体
type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Email  string `json:"email"`
	Status bool   `json:"status"`
}

// Response 统一响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// success 返回成功响应
func success(data interface{}) Response {
	return Response{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
}

// fail 返回失败响应
func fail(code int, msg string) Response {
	return Response{
		Code: code,
		Msg:  msg,
	}
}

// CreateUser 创建用户 - POST /users
func CreateUser(ctx context.Context, c *app.RequestContext) {
	var user User
	// 打印原始请求数据
	hlog.Info("CreateUser called, raw body:", string(c.Request.Body()))

	if err := c.BindJSON(&user); err != nil {
		hlog.Error("Failed to bind JSON:", err.Error())
		c.JSON(400, fail(400, err.Error()))
		return
	}

	// 打印解析后的用户数据
	hlog.Info("User data parsed successfully:", user)

	// 处理创建用户逻辑...
	c.JSON(201, success(user))
}

// GetUser 获取单个用户 - GET /users/:id
func GetUser(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	// 使用id进行查询
	c.JSON(200, success(map[string]interface{}{
		"id":      id,
		"message": fmt.Sprintf("Getting user with ID: %s", id),
	}))
}

// ListUsers 获取用户列表 - GET /users?page=1&size=10
func ListUsers(ctx context.Context, c *app.RequestContext) {
	page := c.Query("page")
	size := c.Query("size")
	// 使用分页参数
	c.JSON(200, success(map[string]interface{}{
		"page":    page,
		"size":    size,
		"message": fmt.Sprintf("Listing users with page: %s, size: %s", page, size),
	}))
}

// UpdateUser 更新用户 - PUT /users/:id
func UpdateUser(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	c.JSON(200, success(map[string]interface{}{
		"id":      id,
		"message": fmt.Sprintf("Updating user with ID: %s", id),
	}))
}

// DeleteUser 删除用户 - DELETE /users/:id
func DeleteUser(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	c.JSON(200, success(map[string]interface{}{
		"id":      id,
		"message": fmt.Sprintf("Deleting user with ID: %s", id),
	}))
}

// SearchUsers 搜索用户 - GET /users/search?keyword=xxx&type=name
func SearchUsers(ctx context.Context, c *app.RequestContext) {
	keyword := c.Query("keyword")
	searchType := c.Query("type")
	c.JSON(200, success(map[string]interface{}{
		"keyword": keyword,
		"type":    searchType,
		"message": fmt.Sprintf("Searching users with keyword: %s, type: %s", keyword, searchType),
	}))
}

// UpdateStatus 更新用户状态 - PATCH /users/:id/status
func UpdateStatus(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	c.JSON(200, success(map[string]interface{}{
		"id":      id,
		"message": fmt.Sprintf("Updating status for user with ID: %s", id),
	}))
}
