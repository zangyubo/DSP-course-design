package main

import (
	"dspBackend/routes"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的 Gin 路由
	r := gin.Default()

	// 启用 CORS 中间件，允许所有域名进行访问
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost"}, // 允许 localhost 的所有端口
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// 允许所有 localhost 的来源
			return strings.HasPrefix(origin, "http://localhost")
		}, // 允许携带凭证（如 cookies）
	}))

	// 注册路由
	// routes.RegisterUploadRoutes(r)
	r.GET("/testGetData", routes.TestGetData)
	r.POST("/basicCharts", routes.UploadMP3)

	// 启动服务器
	r.Run(":8080") // 监听 8080 端口
}
