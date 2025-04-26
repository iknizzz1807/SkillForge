package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

// AvatarHandler xử lý các request liên quan đến avatar
type AvatarHandler struct {
	fileService *services.FileService
}

// NewAvatarHandler khởi tạo một handler mới
func NewAvatarHandler(fileService *services.FileService) *AvatarHandler {
	return &AvatarHandler{
		fileService: fileService,
	}
}

// ServeAvatar phục vụ file avatar tĩnh dựa trên ID
// Được gọi bởi route GET /avatars
func (h *AvatarHandler) ServeAvatar(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.String(http.StatusBadRequest, "User ID is required")
		return
	}

	// Sử dụng service để tìm avatar, không thực hiện logic tìm file trong handler
	avatarPath, err := h.fileService.FindAvatarByUserID(userID)
	if err != nil {
		c.String(http.StatusNotFound, "Avatar not found")
		return
	}

	// Phục vụ file bằng Gin
	c.File(avatarPath)
}

// ServeAvatarByUserID phục vụ file avatar tĩnh dựa trên user ID được truyền vào
// Được gọi bởi route GET /avatars/:id
func (h *AvatarHandler) ServeAvatarByUserID(c *gin.Context) {
	// Lấy userID từ URL parameter
	userID := c.Param("id")
	if userID == "" {
		c.String(http.StatusBadRequest, "User ID is required")
		return
	}

	// Sử dụng service để tìm avatar
	avatarPath, err := h.fileService.FindAvatarByUserID(userID)
	if err != nil {
		// Trả về status 404 nếu không tìm thấy avatar
		c.String(http.StatusNotFound, "Avatar not found")
		return
	}

	// Phục vụ file bằng Gin
	c.File(avatarPath)
}

// UploadAvatar xử lý upload avatar
func (h *AvatarHandler) UploadAvatar(c *gin.Context) {
	// Lấy userID từ context đã được authentication middleware thiết lập
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	file, header, err := c.Request.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}
	defer file.Close()

	// Sử dụng file service để lưu avatar
	filename, err := h.fileService.SaveAvatar(userID, file, header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Avatar uploaded successfully",
		"filename": filename,
	})
}

// type FileHandler struct {
// 	fileService services.FileService
// }

// func NewFileHandler(fileService services.FileService) *FileHandler {
// 	return &FileHandler{fileService: fileService}
// }

// func (h *FileHandler) UploadFile(c *gin.Context) {
// 	var req struct {
// 		FileName string `json:"name" binding:"required"`
// 		// The folder name to upload the file to. Ex: "projects, messages, etc."
// 		FolderName string `json:"folder_name" binding:"required"`
// 	}
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid upload request"})
// 		return
// 	}

// 	UploadedByID := c.GetString("userID")
// 	file, err := h.fileService.UploadFile(c, req.FileName, req.FolderName, UploadedByID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, file)
// }

// func (h *FileHandler) GetFileByID(c *gin.Context) {
// 	fileID := c.Param("id")

// 	file, err := h.fileService.GetFileByID(fileID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, file)
// }
