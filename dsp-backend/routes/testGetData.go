package routes

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

// TestGetData 处理调用 Python 脚本并返回生成文件的逻辑
func TestGetData(c *gin.Context) {
	// 执行 Python 脚本
	cmd := exec.Command("E:/DSP-course-design/dsp-backend/dspOfPython/.venv/Scripts/python.exe", "E:/DSP-course-design/dsp-backend/dspOfPython/main.py")

	// 获取 Python 脚本的输出
	err := cmd.Run()
	if err != nil {
		// 如果出错，返回错误信息
		c.JSON(501, gin.H{"error": fmt.Sprintf("Failed to execute Python script: %v", err)})
		return
	}

	// 等待生成的 .npy 文件
	filePath := "./dist/time.npy"

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(500, gin.H{"error": "Generated file not found"})
		return
	}

	// 等待文件生成完成（模拟等待文件生成）
	// time.Sleep(2 * time.Second) // 可根据实际情况调整等待时间

	// 发送文件给前端
	c.File(filePath)
}
