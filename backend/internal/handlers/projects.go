package handlers

// Nhiệm vụ: Xử lý các endpoint quản lý dự án (tạo, lấy danh sách, lấy chi tiết)
// Liên kết:
// - Gọi internal/services/project_service.go để xử lý logic
// Vai trò trong flow:
// - Được gọi từ /api/projects để quản lý dự án

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

type ProjectHandler struct {
	// projectService xử lý logic liên quan đến project
	projectService *services.ProjectService
}

// NewProjectHandler khởi tạo handler với projectService
// Input: projectService (*services.ProjectService)
// Return: *ProjectHandler - con trỏ đến ProjectHandler
func NewProjectHandler(projectService *services.ProjectService) *ProjectHandler {
	return &ProjectHandler{projectService}
}

// GetProjects xử lý endpoint GET /api/projects
// Return: Trả về JSON danh sách project hoặc lỗi
func (h *ProjectHandler) GetProjects(c *gin.Context) {
	// Gọi service để lấy danh sách project
	projects, err := h.projectService.GetAllProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về danh sách project
	c.JSON(http.StatusOK, projects)
}

// GetProject xử lý endpoint GET /api/projects/:id
// Return: Trả về JSON chi tiết project hoặc lỗi
func (h *ProjectHandler) GetProject(c *gin.Context) {
	projectID := c.Param("id")

	// Gọi service để lấy chi tiết project
	project, err := h.projectService.GetProjectByID(projectID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Trả về chi tiết project
	c.JSON(http.StatusOK, project)
}

// CreateProject xử lý endpoint POST /api/projects
// Return: Trả về JSON project vừa tạo hoặc lỗi
func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var req struct {
		Title       string   `json:"title" binding:"required"`
		Description string   `json:"description" binding:"required"`
		Skills      []string `json:"skills" binding:"required"`
		Timeline    string   `json:"timeline" binding:"required"`
	}

	// Parse và validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Lấy userID từ context (giả sử middleware auth đã gán)
	userID := c.GetString("userID")

	// Gọi service để tạo project
	project, err := h.projectService.CreateProject(userID, req.Title, req.Description, req.Skills, req.Timeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về project vừa tạo
	c.JSON(http.StatusCreated, project)
}
