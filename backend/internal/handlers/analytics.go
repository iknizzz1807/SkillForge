package handlers

// Nhiệm vụ: Xử lý các endpoint báo cáo phân tích (kỹ năng, dự án)
// Liên kết:
// - Gọi internal/services/analytics_service.go để xử lý logic
// Vai trò trong flow:
// - Được gọi từ /api/analytics để cung cấp báo cáo cho sinh viên/doanh nghiệp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

type AnalyticsHandler struct {
	// analyticsService xử lý logic liên quan đến phân tích
	analyticsService *services.AnalyticsService
}

// NewAnalyticsHandler khởi tạo handler với analyticsService
// Input: analyticsService (*services.AnalyticsService)
// Return: *AnalyticsHandler - con trỏ đến AnalyticsHandler
func NewAnalyticsHandler(analyticsService *services.AnalyticsService) *AnalyticsHandler {
	return &AnalyticsHandler{analyticsService}
}

// GetSkillAnalytics xử lý endpoint GET /api/analytics/skills/:userID
// Return: Trả về JSON phân tích kỹ năng hoặc lỗi
func (h *AnalyticsHandler) GetSkillAnalytics(c *gin.Context) {
	userID := c.Param("userID")

	// Gọi service để lấy phân tích kỹ năng
	skills, err := h.analyticsService.GetSkillAnalytics(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Trả về phân tích kỹ năng
	c.JSON(http.StatusOK, skills)
}

// GetProjectAnalytics xử lý endpoint GET /api/analytics/projects/:projectID
// Return: Trả về JSON phân tích dự án hoặc lỗi
func (h *AnalyticsHandler) GetProjectAnalytics(c *gin.Context) {
	projectID := c.Param("projectID")

	// Gọi service để lấy phân tích dự án
	analytics, err := h.analyticsService.GetProjectAnalytics(projectID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Trả về phân tích dự án
	c.JSON(http.StatusOK, analytics)
}
