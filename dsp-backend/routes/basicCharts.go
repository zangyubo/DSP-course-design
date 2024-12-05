package routes

import (
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// RegisterUploadRoutes registers file upload routes
func RegisterUploadRoutes(r *gin.Engine) {
	r.POST("/basicCharts", func(c *gin.Context) {
		// Define a struct to receive the file information
		type RequestData struct {
			File *multipart.FileHeader `form:"file" binding:"required"`
		}
		var requestData RequestData
		// Bind the uploaded file
		if err := c.ShouldBind(&requestData); err != nil {
			// If binding fails, return a 400 error
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "File is required",
			})
			return
		}
		// Ensure the dist folder exists
		if _, err := os.Stat("./dist"); os.IsNotExist(err) {
			err := os.Mkdir("./dist", os.ModePerm)
			if err != nil {
				log.Fatalf("Failed to create dist folder: %v", err)
			}
		}
		// Save the file to the dist directory
		dst := filepath.Join("./dist", requestData.File.Filename)
		if err := c.SaveUploadedFile(requestData.File, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to save the file",
			})
			return
		}
		// Return the response
		c.JSON(http.StatusOK, gin.H{
			"message": "File uploaded successfully",
			"file":    requestData.File.Filename,
		})
	})
}
