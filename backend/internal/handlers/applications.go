package handlers

// Nhiệm vụ: Xử lý các endpoint ứng tuyển dự án
// Liên kết:
// - Gọi internal/services/application_service.go để xử lý logic
// Vai trò trong flow:
// - Được gọi từ /api/applications để sinh viên ứng tuyển dự án

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

type ApplicationHandler struct {
	// applicationService xử lý logic ứng tuyển
	applicationService *services.ApplicationService
}

// NewApplicationHandler khởi tạo handler với applicationService
// Input: applicationService (*services.ApplicationService)
// Return: *ApplicationHandler - con trỏ đến ApplicationHandler
func NewApplicationHandler(applicationService *services.ApplicationService) *ApplicationHandler {
	return &ApplicationHandler{applicationService}
}

// ApplyProject xử lý endpoint POST /api/applications
// Return: Trả về JSON ứng tuyển hoặc lỗi
func (h *ApplicationHandler) ApplyProject(c *gin.Context) {
	// Lấy role từ context
	role := c.GetString("role")
	if role != "student" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only students can apply"})
		return
	}

	var req struct {
		ProjectID string `json:"project_id" binding:"required"`
		Proposal  string `json:"proposal" binding:"required"`
	}

	// Parse và validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Lấy userID từ context
	userID := c.GetString("userID")

	// Gọi service để ứng tuyển
	application, err := h.applicationService.ApplyProject(userID, role, req.ProjectID, req.Proposal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về ứng tuyển vừa tạo
	c.JSON(http.StatusCreated, application)
}

// GetApplication xử lý endpoint GET /api/applications/:id
// Return: Trả về JSON chi tiết ứng tuyển hoặc lỗi
func (h *ApplicationHandler) GetApplication(c *gin.Context) {
	applicationID := c.Param("id")

	// Gọi service để lấy chi tiết ứng tuyển
	application, err := h.applicationService.GetApplicationByID(applicationID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Trả về chi tiết ứng tuyển
	c.JSON(http.StatusOK, application)
}

// GetApplicationsByCurrentUser xử lý endpoint GET
// Lấy tất cả applications liên quan đến người dùng hiện tại (dựa vào role)
// - Nếu là student: trả về các applications mà student đã đăng ký
// - Nếu là business: trả về các applications cho projects của business
// Return: Trả về JSON danh sách applications hoặc lỗi
func (h *ApplicationHandler) GetApplicationsByCurrentUser(c *gin.Context) {
	userID := c.GetString("userID")
	role := c.GetString("role")

	var applications interface{}
	var err error

	// Dựa vào role để gọi service phù hợp
	switch role {
	case "student":
		applications, err = h.applicationService.GetApplicationsByUserID(userID)
	case "business":
		applications, err = h.applicationService.GetApplicationsByBusinessID(userID)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user role"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về danh sách applications
	c.JSON(http.StatusOK, applications)
}

// Giữ lại các functions cũ để tương thích ngược (bạn có thể xóa sau khi cập nhật tất cả frontend calls)
func (h *ApplicationHandler) GetApplicationsByUser(c *gin.Context) {
	userID := c.GetString("userID")

	applications, err := h.applicationService.GetApplicationsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, applications)
}

func (h *ApplicationHandler) GetApplicationsByBusiness(c *gin.Context) {
	businessID := c.GetString("userID")

	applications, err := h.applicationService.GetApplicationsByBusinessID(businessID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, applications)
}

// UpdateApplicationStatus xử lý endpoint PUT
// Cập nhật trạng thái của application
// Return: Trả về JSON application đã cập nhật hoặc lỗi
func (h *ApplicationHandler) UpdateApplicationStatus(c *gin.Context) {
	id := c.Param("id")

	var req struct {
		Status string `json:"status" binding:"required"`
	}

	// Parse và validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Gọi service để cập nhật status
	application, err := h.applicationService.UpdateApplicationStatus(id, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về application đã cập nhật
	c.JSON(http.StatusOK, gin.H{
		"message":     "Application status updated successfully",
		"application": application,
	})
}
