package main

import (
	"PrimeGin/internal/config"
	"PrimeGin/internal/routes"
	"PrimeGin/internal/utils"
)

func main() {
	// 加载配置
	cfg := config.LoadConfig()

	// 初始化数据库
	utils.ConnectDB(cfg)
	utils.AutoMigrate()

	// 初始化路由
	router := routes.SetupRouter()

	// 启动服务器
	router.Run(":8899")
}
