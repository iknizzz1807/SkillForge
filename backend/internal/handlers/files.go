package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// Định nghĩa đường dẫn lưu trữ ở đây hoặc đọc từ config/const chung
const AvatarStoragePathForServing = "./storage/avatars"

// ServeAvatarHandler phục vụ file avatar tĩnh dựa trên ID, không yêu cầu đuôi mở rộng
// Được gọi bởi route GET /avatars/:id
func ServeAvatarHandler(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.String(http.StatusBadRequest, "User ID is required")
		return
	}

	// Security check: chống path traversal
	// Đảm bảo userID không chứa các ký tự nguy hiểm
	if strings.Contains(userID, "..") || strings.ContainsAny(userID, "/\\") {
		c.String(http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Tìm file avatar dựa trên ID, không cần biết đuôi mở rộng là gì
	avatarPath, err := findAvatarByUserID(userID)
	if err != nil {
		c.String(http.StatusNotFound, "Avatar not found")
		return
	}

	// Phục vụ file bằng Gin
	c.File(avatarPath)
}

// findAvatarByUserID tìm file avatar dựa trên ID người dùng
// Hàm này sẽ tìm kiếm trong thư mục avatars các file có tên bắt đầu bằng userID
func findAvatarByUserID(userID string) (string, error) {
	// Đọc tất cả các file trong thư mục avatars
	entries, err := os.ReadDir(AvatarStoragePathForServing)
	if err != nil {
		return "", fmt.Errorf("failed to read avatar directory: %w", err)
	}

	// Tìm file có tên bắt đầu bằng userID
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasPrefix(entry.Name(), userID) {
			return filepath.Join(AvatarStoragePathForServing, entry.Name()), nil
		}
	}

	return "", fmt.Errorf("avatar not found for user: %s", userID)
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
