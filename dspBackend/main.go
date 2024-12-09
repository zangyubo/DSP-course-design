package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的 Gin 路由器
	r := gin.Default()

	// 定义一个 GET 路由：/hello
	r.GET("/hello", func(c *gin.Context) {
		// 获取查询参数 "name"
		name := c.DefaultQuery("name", "Guest") // 如果没有提供name，则默认值为"Guest"

		// 返回一个 JSON 响应
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello " + name,
		})
	})

	// 启动 HTTP 服务器，监听在 8080 端口
	r.Run(":8080")
}
