package handlers

// Nhiệm vụ: Xử lý các endpoint liên quan đến xác thực (đăng ký, đăng nhập)
// Liên kết:
// - Gọi internal/services/auth_service.go để xử lý logic xác thực
// - Dùng internal/middleware/auth.go để kiểm tra token (nếu cần)
// Vai trò trong flow:
// - Là điểm vào cho đăng ký/đăng nhập, tạo token JWT cho người dùng
// - Được gọi từ /api/auth/register hoặc /api/auth/login

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

type AuthHandler struct {
	// authService xử lý logic xác thực
	authService *services.AuthService
}

// NewAuthHandler khởi tạo handler với authService
// Input: authService (*services.AuthService)
// Return: *AuthHandler - con trỏ đến AuthHandler
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

// Register xử lý endpoint POST /auth/register
// Return: Trả về JSON với thông tin user và token hoặc lỗi
func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
		Role     string `json:"role" binding:"required"`
	}

	// Parse và validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Gọi service để đăng ký user
	user, token, err := h.authService.Register(req.Email, req.Password, req.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về thông tin user và token
	c.JSON(http.StatusCreated, gin.H{
		"user":  user,
		"token": token,
	})
}

// Login xử lý endpoint POST /auth/login
// Return: Trả về JSON với token hoặc lỗi
func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	// Parse và validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Gọi service để đăng nhập
	token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Trả về token
	c.JSON(http.StatusOK, gin.H{"token": token})
}
