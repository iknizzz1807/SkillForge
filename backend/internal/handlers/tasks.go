package handlers

// Nhiệm vụ: Xử lý các endpoint quản lý task (tạo, lấy, cập nhật)
// Liên kết:
// - Gọi internal/services/task_service.go để xử lý logic
// Vai trò trong flow:
// - Được gọi từ /api/tasks để quản lý task trong dự án

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

type TaskHandler struct {
	// taskService xử lý logic liên quan đến task
	taskService *services.TaskService
}

// NewTaskHandler khởi tạo handler với taskService
// Input: taskService (*services.TaskService)
// Return: *TaskHandler - con trỏ đến TaskHandler
func NewTaskHandler(taskService *services.TaskService) *TaskHandler {
	return &TaskHandler{taskService}
}

// GetTask xử lý endpoint GET /api/tasks/:id
// Return: Trả về JSON chi tiết task hoặc lỗi
func (h *TaskHandler) GetTask(c *gin.Context) {
	taskID := c.Param("id")

	// Gọi service để lấy chi tiết task
	task, err := h.taskService.GetTaskByID(taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Trả về chi tiết task
	c.JSON(http.StatusOK, task)
}

// CreateTask xử lý endpoint POST /api/tasks
// Return: Trả về JSON task vừa tạo hoặc lỗi
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req struct {
		ProjectID   string `json:"project_id" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	// Parse và validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Gọi service để tạo task
	task, err := h.taskService.CreateTask(req.ProjectID, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về task vừa tạo
	c.JSON(http.StatusCreated, task)
}

// UpdateTask xử lý endpoint PUT /api/tasks/:id
// Return: Trả về JSON task đã cập nhật hoặc lỗi
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	taskID := c.Param("id")
	var req struct {
		Status string `json:"status" binding:"required"`
	}

	// Parse và validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Gọi service để cập nhật task
	task, err := h.taskService.UpdateTask(taskID, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về task đã cập nhật
	c.JSON(http.StatusOK, task)
}
