package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

type FileHandler struct {
	fileService services.FileService
}

func NewFileHandler(fileService services.FileService) *FileHandler {
	return &FileHandler{fileService: fileService}
}

func (h *FileHandler) UploadFile(c *gin.Context) {
	var req struct {
		FileName string `json:"name" binding:"required"`
		// The folder name to upload the file to. Ex: "projects, messages, etc."
		FolderName string `json:"folder_name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid upload request"})
		return
	}

	UploadedByID := c.GetString("userID")
	file, err := h.fileService.UploadFile(c, req.FileName, req.FolderName, UploadedByID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, file)
}

func (h *FileHandler) GetFileByID(c *gin.Context) {
	fileID := c.Param("id")

	file, err := h.fileService.GetFileByID(fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, file)
}
