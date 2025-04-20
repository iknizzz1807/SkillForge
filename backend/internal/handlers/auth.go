// internal/handlers/auth_handler.go
package handlers

import (
	"mime/multipart" // Thêm multipart
	"net/http"
	"strings" // Thêm strings để kiểm tra role

	"github.com/gin-gonic/gin"
	// Bỏ constants nếu không dùng nữa, hoặc đảm bảo nó tồn tại
	"github.com/iknizzz1807/SkillForge/internal/services"
)

type AuthHandler struct {
	authService *services.AuthService
	// Không cần fileService ở đây nữa
}

// NewAuthHandler khởi tạo handler (không cần fileService)
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

// Register xử lý endpoint POST /auth/register (nhận multipart/form-data)
func (h *AuthHandler) Register(c *gin.Context) {
	// Parse multipart form
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing form"})
		return
	}

	// Extract form values
	email := c.PostForm("email")
	name := c.PostForm("name")
	password := c.PostForm("password")
	role := c.PostForm("role")
	website := c.PostForm("website") // Get website field

	// Basic validation
	if email == "" || name == "" || password == "" || role == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email, name, password, and role are required"})
		return
	}

	// Normalize role
	roleLower := strings.ToLower(role)
	if roleLower != "student" && roleLower != "business" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role must be 'student' or 'business'"})
		return
	}

	// Password validation
	if len(password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 6 characters"})
		return
	}

	// Process avatar file (if any)
	var avatarFile multipart.File
	var fileHeader *multipart.FileHeader
	avatarFile, fileHeader, err = c.Request.FormFile("avatar")
	if err != nil && err != http.ErrMissingFile {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error processing avatar file"})
		return
	}
	if avatarFile != nil {
		defer avatarFile.Close()
	}

	// Process website for business accounts
	if roleLower == "business" && website != "" {
		// You can add additional website validation here if needed
		// For example, check if it starts with http:// or https://
		if !strings.HasPrefix(website, "http://") && !strings.HasPrefix(website, "https://") {
			website = "https://" + website
		}
	}

	// Register the user
	user, token, err := h.authService.Register(email, name, password, roleLower, website, avatarFile, fileHeader)
	if err != nil {
		errMsg := err.Error()
		statusCode := http.StatusInternalServerError
		if strings.Contains(errMsg, "email already exists") {
			statusCode = http.StatusConflict
		}
		c.JSON(statusCode, gin.H{"error": errMsg})
		return
	}

	// Return user info and token
	c.JSON(http.StatusCreated, gin.H{
		"user":  user,
		"token": token,
	})
}

// Login (Giữ nguyên)
func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
