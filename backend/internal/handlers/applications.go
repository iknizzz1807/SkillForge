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
	application, err := h.applicationService.ApplyProject(userID, req.ProjectID, req.Proposal)
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
