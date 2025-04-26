package handlers

// Nhiệm vụ: Xử lý các endpoint quản lý dự án (tạo, lấy danh sách, lấy chi tiết)
// Liên kết:
// - Gọi internal/services/project_service.go để xử lý logic
// Vai trò trong flow:
// - Được gọi từ /api/projects để quản lý dự án

import (
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/models"
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

	// Đảm bảo không trả về null cho mảng trống
	if projects == nil {
		projects = []*models.Project{} // Trả về mảng trống thay vì nil
	}

	// Trả về danh sách project
	c.JSON(http.StatusOK, projects)
}

// GetProject xử lý endpoint GET /api/projects/:id
// Return: Trả về JSON chi tiết project hoặc lỗi
func (h *ProjectHandler) GetProjectMarketplace(c *gin.Context) {
	projectID := c.Param("id")
	userID := c.GetString("userID")

	// Gọi service đã được cập nhật để lấy chi tiết project và trạng thái tham gia
	project, hasJoined, hasApplied, err := h.projectService.GetProjectWithUserStatus(projectID, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Trả về chi tiết project kèm thông tin participation status
	c.JSON(http.StatusOK, gin.H{
		"project":    project,
		"hasJoined":  hasJoined,  // User đã tham gia vào project này
		"hasApplied": hasApplied, // User đã apply nhưng chưa được chấp nhận
	})
}

// GetProjectByStudent xử lý endpoint GET /api/projects/student
// Return: Trả về JSON danh sách project mà student có ID được chỉ định tham gia
func (h *ProjectHandler) GetProjectByStudent(c *gin.Context) {
	studentID := c.GetString("userID")

	if studentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Student ID is required"})
		return
	}

	// Gọi service để lấy danh sách project mà student đã tham gia
	projects, err := h.projectService.GetProjectsByStudentID(studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Đảm bảo không trả về null cho mảng trống
	if projects == nil {
		projects = []*models.Project{} // Trả về mảng trống thay vì nil
	}

	// Trả về danh sách project
	c.JSON(http.StatusOK, projects)
}

// GetProjectByBusiness xử lý endpoint GET /api/projects/business
// Return: Trả về JSON danh sách project mà business có ID được chỉ định đã tạo
func (h *ProjectHandler) GetProjectByBusiness(c *gin.Context) {
	businessID := c.GetString("userID")

	if businessID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	// Gọi service để lấy danh sách project thuộc về business
	projects, err := h.projectService.GetProjectsByBusinessID(businessID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Đảm bảo không trả về null cho mảng trống
	if projects == nil {
		projects = []*models.Project{} // Trả về mảng trống thay vì nil
	}

	// Trả về danh sách project
	c.JSON(http.StatusOK, projects)
}

// CreateProject xử lý endpoint POST /api/projects
// Return: Trả về JSON project vừa tạo hoặc lỗi
func (h *ProjectHandler) CreateProject(c *gin.Context) {
	// Chỉ business mới đuọc tạo project
	if c.GetString("role") != "business" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only businesses can create project"})
		return
	}

	var req struct {
		Title       string    `json:"title" binding:"required"`
		Description string    `json:"description" binding:"required"`
		Skills      []string  `json:"skills" binding:"required"`
		StartTime   time.Time `json:"start_time" binding:"required"`
		EndTime     time.Time `json:"end_time" binding:"required"`
		MaxMember   int       `json:"max_member" binding:"required"`
		Difficulty  string    `json:"difficulty" binding:"required"`
	}

	// Parse và validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Lấy userID & userName từ context (giả sử middleware auth đã gán)
	userID := c.GetString("userID")
	userName := c.GetString("name")

	// Gọi service để tạo project
	project, err := h.projectService.CreateProject(userID, userName, req.Title, req.Description, req.Skills, req.StartTime, req.EndTime, req.MaxMember, req.Difficulty)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về project vừa tạo
	c.JSON(http.StatusCreated, project)
}

// DeleteProject xử lý endpoint DELETE /api/projects/:id
// Return: Trả về thông báo xóa thành công hoặc lỗi
func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	// Lấy project ID từ parameter
	projectID := c.Param("id")

	// Lấy userID từ context (giả sử middleware auth đã gán)
	userID := c.GetString("userID")

	// Kiểm tra quyền xóa project (chỉ người tạo hoặc admin mới được xóa)
	project, err := h.projectService.GetProjectByID(projectID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	// Kiểm tra người dùng có phải là người tạo project
	if project.CreatedByID != userID && c.GetString("role") != "business" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to delete this project"})
		return
	}

	// Gọi service để xóa project
	// userID để thể hiện ai xoá project để xem họ có quyền xoá hay không
	if err := h.projectService.DeleteProject(projectID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về thông báo xóa thành công
	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}

// UpdateProject xử lý endpoint PUT /api/projects/:id
// Return: Trả về JSON project đã cập nhật hoặc lỗi
func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	// Lấy project ID từ parameter
	projectID := c.Param("id")

	// Lấy userID từ context (giả sử middleware auth đã gán)
	userID := c.GetString("userID")

	// Kiểm tra quyền cập nhật project (chỉ người tạo mới được cập nhật)
	project, err := h.projectService.GetProjectByID(projectID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	// Kiểm tra người dùng có phải là người tạo project
	if project.CreatedByID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to update this project"})
		return
	}

	// Parse request body
	var req struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Skills      []string  `json:"skills"`
		StartTime   time.Time `json:"start_time"`
		EndTime     time.Time `json:"end_time"`
		MaxMember   int       `json:"max_member"`
		Status      string    `json:"status"`
		Difficulty  string    `json:"difficulty"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validation checks
	if !req.StartTime.IsZero() && !req.EndTime.IsZero() && req.StartTime.After(req.EndTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Start time must be before end time"})
		return
	}

	if req.MaxMember > 0 && req.MaxMember < project.CurrentMember {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot reduce max members below current member count"})
		return
	}

	// Gọi service để cập nhật project
	updatedProject, err := h.projectService.UpdateProject(
		projectID,
		req.Title,
		req.Description,
		req.Skills,
		req.StartTime,
		req.EndTime,
		req.MaxMember,
		req.Status,
		req.Difficulty,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về project đã được cập nhật
	c.JSON(http.StatusOK, updatedProject)
}

// RemoveStudentFromProject xử lý endpoint DELETE /api/projects/:id/students/:studentID
// Return: Trả về thông báo xóa thành công hoặc lỗi
func (h *ProjectHandler) RemoveStudentFromProject(c *gin.Context) {
	projectID := c.Param("projectID")
	studentID := c.Param("studentID")
	businessID := c.GetString("userID")
	role := c.GetString("role")

	// Kiểm tra quyền xóa student (chỉ business sở hữu project hoặc student tự xóa chính mình)
	project, err := h.projectService.GetProjectByID(projectID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	// Kiểm tra quyền: business sở hữu project hoặc student tự rời project
	if project.CreatedByID != businessID && (role != "student" || studentID != businessID) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You don't have permission to remove this student from the project",
		})
		return
	}

	// Gọi service để xóa student khỏi project
	err = h.projectService.RemoveStudentFromProject(projectID, studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về thông báo thành công
	c.JSON(http.StatusOK, gin.H{
		"message": "Student removed from project successfully",
	})
}

// GetStudentsByProject xử lý endpoint GET /api/projects/:projectID/students
// Return: Trả về JSON danh sách sinh viên tham gia dự án hoặc lỗi
func (h *ProjectHandler) GetStudentsByProject(c *gin.Context) {
	projectID := c.Param("id")

	if projectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}

	// Kiểm tra project có tồn tại không
	_, err := h.projectService.GetProjectByID(projectID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	// Gọi service để lấy danh sách sinh viên tham gia dự án
	students, err := h.projectService.GetStudentsByProjectID(projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Đảm bảo không trả về null cho mảng trống
	if students == nil {
		students = []models.User{} // Trả về mảng rỗng thay vì nil
	}

	// Trả về danh sách sinh viên
	c.JSON(http.StatusOK, students)
}
