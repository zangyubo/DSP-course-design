package routes

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// UploadMP3 处理接收 MP3 文件并保存到服务器
func UploadMP3(c *gin.Context) {
	// 获取是哪个通道，默认值为 "0"
	channel := c.DefaultPostForm("channel", "0")

	// 默认存储目录为 channel2
	storeURL := "./dist/basicChart"
	print(channel)
	fmt.Printf("channel 的类型是: %T\n", channel)

	// 根据 channel 参数判断存储目录
	if channel == "1" {
		storeURL = "./dist/basicChart/channel1"
	} else if channel == "2" {
		storeURL = "./dist/basicChart/channel2"
	}

	// 如果目标目录已经存在，删除其中所有文件
	if _, err := os.Stat(storeURL); !os.IsNotExist(err) {
		// 清空目录中的所有文件
		files, err := os.ReadDir(storeURL)
		if err != nil {
			c.JSON(500, gin.H{"error": fmt.Sprintf("读取目录失败: %v", err)})
			return
		}

		// 遍历删除目录中的所有文件
		for _, file := range files {
			filePath := fmt.Sprintf("%s/%s", storeURL, file.Name())
			if err := os.Remove(filePath); err != nil {
				c.JSON(500, gin.H{"error": fmt.Sprintf("删除文件失败: %v", err)})
				return
			}
		}
	}

	// 确保 dist 目录及对应的 channel 目录存在
	if err := os.MkdirAll(storeURL, os.ModePerm); err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("创建目录失败: %v", err)})
		return
	}

	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("获取文件失败: %v", err)})
		return
	}

	// 获取当前时间戳作为文件名
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	fileName := fmt.Sprintf("%s/%s.mp3", storeURL, timestamp)

	// 保存文件到指定的目录
	if err := c.SaveUploadedFile(file, fileName); err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("保存文件失败: %v", err)})
		return
	}

	// 返回成功消息，包含文件的保存路径
	c.JSON(200, gin.H{
		"message": "文件上传成功",
	})
}
