package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的 Gin 路由
	r := gin.Default()

	// 启用 CORS 中间件，允许所有域名进行访问
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5174"},                   // 允许你的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允许的 HTTP 方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 允许的请求头
		AllowCredentials: true,                                                // 允许携带凭证（如 cookies）
	}))

	// 注册路由
	// routes.RegisterUploadRoutes(r)

	// 启动服务器
	r.Run(":8080") // 监听 8080 端口
}
