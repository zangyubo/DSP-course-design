package routes

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// UploadMP3 处理接收 MP3 文件并保存到服务器
func UploadMP3(c *gin.Context) {
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("获取文件失败: %v", err)})
		return
	}

	// 确保 dist 文件夹存在
	if _, err := os.Stat("./dist"); os.IsNotExist(err) {
		err := os.Mkdir("./dist", os.ModePerm)
		if err != nil {
			c.JSON(500, gin.H{"error": fmt.Sprintf("创建目录失败: %v", err)})
			return
		}
	}

	// 获取当前时间戳作为文件名
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	fileName := fmt.Sprintf("./dist/%s.mp3", timestamp)

	// 保存文件到 ./dist 目录
	if err := c.SaveUploadedFile(file, fileName); err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("保存文件失败: %v", err)})
		return
	}

	// 返回成功消息，包含文件的保存路径
	c.JSON(200, gin.H{
		"message": "文件上传成功",
		"file":    fileName,
	})
}
