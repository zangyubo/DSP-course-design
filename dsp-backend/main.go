package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个默认的 Gin 路由引擎
	r := gin.Default()

	// 定义一个 GET 路由，接收 name 参数
	r.GET("/hello", func(c *gin.Context) {
		// 从查询参数中获取 name 参数
		name := c.Query("name")
		if name == "" {
			name = "world" // 如果 name 参数为空，默认为 "world"
		}

		// 返回 JSON 响应
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, " + name,
		})
	})

	// 启动 HTTP 服务，监听 8080 端口
	r.Run(":8080")
}
