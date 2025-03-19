package routes

import (
	"PrimeGin/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 全局中间件
	router.Use(CORSMiddleware())
	// router.Use(LoggerMiddleware()) // Uncomment this line after defining LoggerMiddleware

	// API路由组
	api := router.Group("/api/v1")
	{
		userCtrl := controllers.UserController{}
		api.GET("/users", userCtrl.GetUsers)
	}

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
