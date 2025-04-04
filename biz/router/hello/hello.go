// Code generated by hertz generator. DO NOT EDIT.

package hello

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	hello "PrimeGo/biz/handler/hello"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {
	root := r.Group("/hello",rootMw()... )
	root.GET("/", hello.HelloMethod)
	root.GET("/new", hello.NewMethod)
	root.POST("/other", hello.OtherMethod)

	// RESTful APIs for users - 直接在根路由下定义
	root.POST("/users", hello.CreateUser)            // 创建用户
	root.GET("/users/:id", hello.GetUser)           // 获取单个用户
	root.GET("/users", hello.ListUsers)             // 查询用户列表
	root.PUT("/users/:id", hello.UpdateUser)        // 更新用户
	root.DELETE("/users/:id", hello.DeleteUser)     // 删除用户
	root.GET("/users/search", hello.SearchUsers)    // 搜索用户
	root.PATCH("/users/:id/status", hello.UpdateStatus)  // 更新状态
}
