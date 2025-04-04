package main

import (
	"PrimeGo/biz/router"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default(
		server.WithHostPorts(":6910"),
	)
	// 注册业务路由
	router.Register(h)
	h.Spin()
}
