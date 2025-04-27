package handlers

// Nhiệm vụ: Xử lý các endpoint quản lý thông tin người dùng (lấy, cập nhật profile)
// Liên kết:
// - Gọi internal/services/user_service.go để xử lý logic
// Vai trò trong flow:
// - Được gọi từ /api/users/:id để lấy hoặc cập nhật thông tin user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/services"
	"github.com/iknizzz1807/SkillForge/internal/utils"
)

type UserHandler struct {
	// userService xử lý logic liên quan đến user
	userService *services.UserService
}

// NewUserHandler khởi tạo handler với userService
// Input: userService (*services.UserService)
// Return: *UserHandler - con trỏ đến UserHandler
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService}
}

// GetUser xử lý endpoint GET /api/user
// Return: Trả về JSON với thông tin user hoặc lỗi
func (h *UserHandler) GetCurrentUser(c *gin.Context) {
	userID := c.GetString("userID")

	// Gọi service để lấy thông tin user
	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Trả về thông tin user
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	// Gọi service để lấy thông tin user
	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Trả về thông tin user
	c.JSON(http.StatusOK, user)
}

// UpdateUser xử lý endpoint PUT /api/user
// Cập nhật thông tin user và avatar (nếu có)
// Return: Trả về JSON với thông tin user đã cập nhật và token JWT mới
func (h *UserHandler) UpdateCurrentUser(c *gin.Context) {
	userID := c.GetString("userID")

	// Đảm bảo rằng userID có giá trị
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	// Sử dụng FormFile thay vì ShouldBindJSON để xử lý multipart form
	name := c.PostForm("name")
	email := c.PostForm("email")
	title := c.PostForm("title")

	// Khởi tạo user trước khi cập nhật
	var user *models.User
	var err error

	// Kiểm tra xem có file avatar được upload không
	file, header, err := c.Request.FormFile("avatar")
	if err == nil {
		// Có file được upload, xử lý avatar
		defer file.Close()

		// Khởi tạo FileService
		fileService := services.NewFileService(h.userService.GetUserRepository())

		// Lưu avatar vào storage
		avatarFilename, err := fileService.SaveAvatar(userID, file, header)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to save avatar: %v", err)})
			return
		}

		// Cập nhật thông tin user với avatarFilename mới
		user, err = h.userService.UpdateUserWithAvatar(userID, name, email, title, avatarFilename)
	} else {
		// Không có file được upload, chỉ cập nhật thông tin cơ bản
		user, err = h.userService.UpdateUser(userID, name, email, title)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generate a new JWT token with updated user information
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Name, user.Role)
	if err != nil {
		// Even if token generation fails, we still want to return the updated user
		c.JSON(http.StatusOK, gin.H{
			"user":  user,
			"error": "Token generation failed but user was updated",
		})
		return
	}

	// Trả về user đã cập nhật và token mới
	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}
