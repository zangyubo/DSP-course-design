package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的 Gin 路由引擎
	r := gin.Default()

	// 添加 CORS 中间件（可选）
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// GET 路由
	r.GET("/hello", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			name = "world"
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, " + name,
		})
	})

	// POST 路由
	r.POST("/hello", func(c *gin.Context) {
		var request struct {
			Name string `json:"name"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, " + request.Name,
		})
	})

	// 启动 HTTP 服务，监听 8080 端口
	r.Run(":8080")
}
