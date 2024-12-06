package routes

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

// UploadMP3 处理接收 MP3 文件并保存到服务器
func BasicChartReceiveMP3(c *gin.Context) {

	channel := c.DefaultPostForm("channel", "0")

	// 默认存储目录为 channel2
	storeURL := "./dist/basicChart"
	fileName := "signal.mp3"

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

	// 保存文件到指定的目录
	fileName = fmt.Sprintf("%s/%s", storeURL, fileName)
	if err := c.SaveUploadedFile(file, fileName); err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("保存文件失败: %v", err)})
		return
	}

	// 返回成功消息，包含文件的保存路径
	c.JSON(200, gin.H{
		"message": "文件上传成功",
	})
}

// TestGetData 处理调用 Python 脚本并返回生成文件的逻辑
func SendChartData(c *gin.Context) {
	// 获取前端发送的 channel 参数
	channel := c.DefaultQuery("channel", "1") // 默认值为 "1"
	mode := c.DefaultQuery("mode", "1")       // 默认值为 "1"

	// 执行 Python 脚本
	cmd := exec.Command("E:/DSP-course-design/dsp-backend/dspOfPython/.venv/Scripts/python.exe",
		"E:/DSP-course-design/dsp-backend/dspOfPython/src/basicChart.py", channel, mode)

	// 获取 Python 脚本的输出
	err := cmd.Run()
	if err != nil {
		// 如果出错，返回错误信息
		c.JSON(501, gin.H{"error": fmt.Sprintf("Failed to execute Python script: %v", err)})
		return
	}

	// 等待生成的 .npy 文件
	filePath := fmt.Sprintf("./dist/basicChart/channel%s/channel%s_data.npy", channel, channel)

	if mode == "1" {
		filePath = fmt.Sprintf("./dist/basicChart/channel%s/channel%s_freXdata.npy", channel, channel)
	} else if mode == "2" {
		filePath = fmt.Sprintf("./dist/basicChart/channel%s/channel%s_freYdata.npy", channel, channel)
	} else if mode == "3" {
		filePath = fmt.Sprintf("./dist/basicChart/channel%s/channel%s_TimeXdata.npy", channel, channel)
	} else if mode == "4" {
		filePath = fmt.Sprintf("./dist/basicChart/channel%s/channel%s_TimeYdata.npy", channel, channel)
	}

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(500, gin.H{"error": "Generated file not found"})
		return
	}

	// 发送文件给前端
	c.File(filePath)
}
