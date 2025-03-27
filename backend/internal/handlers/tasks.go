package handlers

// Nhiệm vụ: Xử lý các endpoint quản lý task (tạo, lấy, cập nhật)
// Liên kết:
// - Gọi internal/services/task_service.go để xử lý logic
// Vai trò trong flow:
// - Được gọi từ /api/tasks để quản lý task trong dự án

import (
	"net/http"

	"fmt"

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

// GetTask xử lý endpoint GET /api/tasks/:id với id là projectID
// Return: Trả về JSON chi tiết task hoặc lỗi
func (h *TaskHandler) GetTasksByProjectID(c *gin.Context) {
	taskID := c.Param("id")

	// Gọi service để lấy chi tiết task
	tasks, err := h.taskService.GetTasksByProjectID(taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Trả về chi tiết task
	c.JSON(http.StatusOK, tasks)
}

// CreateTasks xử lý endpoint POST /api/tasks
// Tạo nhiều tasks một lúc
// Return: Trả về JSON danh sách các task vừa tạo hoặc lỗi
func (h *TaskHandler) CreateTasks(c *gin.Context) {
	var req struct {
		ProjectID    string   `json:"project_id" binding:"required"`
		Descriptions []string `json:"descriptions" binding:"required"`
	}

	// Parse và validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inputs"})
		return
	}

	// Kiểm tra danh sách descriptions không rỗng
	if len(req.Descriptions) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Must have at least one task description"})
		return
	}

	// Gọi service để tạo các tasks
	tasks, err := h.taskService.CreateTasks(req.ProjectID, req.Descriptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về danh sách tasks vừa tạo
	c.JSON(http.StatusCreated, gin.H{
		"message": "Created " + fmt.Sprintf("%d", len(tasks)) + " new tasks",
		"tasks":   tasks,
	})
}

// AssignTask xử lý endpoint PUT /api/tasks/:id/assign
// Có nhiệm vụ cập nhập trường assigned_to trong mongodb thành userID
// Return: Trả về JSON task đã cập nhật hoặc lỗi
func (h *TaskHandler) AssignTask(c *gin.Context) {
	taskID := c.Param("id")

	var req struct {
		UserID string `json:"user_id" binding:"required"`
	}

	// Parse và validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id is not valid"})
		return
	}

	// Gọi service để gán task cho người dùng
	task, err := h.taskService.AssignTask(taskID, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về task đã được gán
	c.JSON(http.StatusOK, gin.H{
		"message": "Assigned task successfully",
		"task":    task,
	})
}

// FinishTask xử lý endpoint PUT /api/tasks/:id/finish
// Có nhiệm vụ cập nhập trường finished_by trong mongodb thành userID và đánh dấu task là đã hoàn thành
// Return: Trả về JSON task đã hoàn thành hoặc lỗi
func (h *TaskHandler) FinishTask(c *gin.Context) {
	taskID := c.Param("id")

	var req struct {
		UserID string `json:"user_id" binding:"required"`
	}

	// Parse và validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID người dùng không hợp lệ"})
		return
	}

	// Gọi service để đánh dấu task là hoàn thành
	task, err := h.taskService.FinishTask(taskID, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về task đã hoàn thành
	c.JSON(http.StatusOK, gin.H{
		"message": "Task đã được đánh dấu là hoàn thành",
		"task":    task,
	})
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

// DeleteTask xử lý endpoint DELETE /api/tasks/:id
// Return: Trả về thông báo xóa thành công hoặc lỗi
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	// Lấy taskID từ parameter
	taskID := c.Param("id")

	// Có thể thêm xác thực người dùng ở đây (kiểm tra xem người dùng có quyền xóa task này không)
	// userID := c.GetString("user_id") // Giả sử middleware auth đã đặt user_id

	// Gọi service để xóa task
	err := h.taskService.DeleteTask(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về thông báo thành công
	c.JSON(http.StatusOK, gin.H{
		"message": "task deleted successfully",
	})
}
