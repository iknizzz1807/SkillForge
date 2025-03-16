package handlers

// Nhiệm vụ: Xử lý các endpoint quản lý thông tin người dùng (lấy, cập nhật profile)
// Liên kết:
// - Gọi internal/services/user_service.go để xử lý logic
// Vai trò trong flow:
// - Được gọi từ /api/users/:id để lấy hoặc cập nhật thông tin user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/services"
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

// GetUser xử lý endpoint GET /api/users/:id
// Return: Trả về JSON với thông tin user hoặc lỗi
func (h *UserHandler) GetUser(c *gin.Context) {
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

// UpdateUser xử lý endpoint PUT /api/users/:id
// Return: Trả về JSON với thông tin user đã cập nhật hoặc lỗi
func (h *UserHandler) UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	var req struct {
		Skills []string `json:"skills"`
	}

	// Parse và validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Gọi service để cập nhật user
	user, err := h.userService.UpdateUser(userID, req.Skills)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về user đã cập nhật
	c.JSON(http.StatusOK, user)
}
