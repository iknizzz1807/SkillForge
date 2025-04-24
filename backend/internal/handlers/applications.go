// internal/handlers/applications.go
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

	// --- CHANGE START ---
	// Cập nhật struct req để nhận motivation và detailed_proposal
	var req struct {
		ProjectID        string `json:"project_id" binding:"required"`
		Motivation       string `json:"motivation" binding:"required"`
		DetailedProposal string `json:"detailed_proposal" binding:"required"`
		// Proposal  string `json:"proposal" binding:"required"` // Trường cũ
	}
	// --- CHANGE END ---

	// Parse và validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()}) // Trả về lỗi chi tiết hơn
		return
	}

	// Lấy userID từ context
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"}) // Thêm kiểm tra userID
		return
	}

	application, err := h.applicationService.ApplyProject(userID, req.ProjectID, req.Motivation, req.DetailedProposal)

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

	// (Optional: Thêm kiểm tra quyền truy cập - ví dụ: chỉ student tạo application hoặc business sở hữu project mới được xem)
	// userID := c.GetString("userID")
	// role := c.GetString("role")

	// Gọi service để lấy chi tiết ứng tuyển
	application, err := h.applicationService.GetApplicationByID(applicationID)
	if err != nil {
		// Kiểm tra lỗi cụ thể từ service
		if err.Error() == "application not found" { // Giả sử service trả về lỗi này
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// (Optional: Kiểm tra quyền xem application)
	// if role == constants.RoleStudent && application.UserID != userID {
	// 	c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
	// 	return
	// }
	// // Tương tự kiểm tra cho business

	// Trả về chi tiết ứng tuyển
	c.JSON(http.StatusOK, application)
}

// GetApplicationsByCurrentUser xử lý endpoint GET /api/applications/me
// Lấy tất cả applications liên quan đến người dùng hiện tại (dựa vào role)
// - Nếu là student: trả về các applications mà student đã đăng ký
// - Nếu là business: trả về các applications cho projects của business
// Return: Trả về JSON danh sách applications hoặc lỗi
func (h *ApplicationHandler) GetApplicationsByCurrentUser(c *gin.Context) {
	userID := c.GetString("userID")
	role := c.GetString("role")

	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

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

	// Đảm bảo trả về mảng rỗng thay vì null nếu không có kết quả
	// (Service nên xử lý việc này, nhưng kiểm tra thêm ở đây cũng tốt)
	// Ví dụ kiểm tra cho student:
	// if role == constants.RoleStudent {
	// 	if apps, ok := applications.([]models.Application); ok && apps == nil {
	// 		applications = []models.Application{}
	// 	}
	// }
	// Tương tự cho business

	// Trả về danh sách applications
	c.JSON(http.StatusOK, applications)
}

// UpdateApplicationStatus xử lý endpoint PUT /api/applications/:id/status
// Cập nhật trạng thái của application
// Route này được dùng để approve hoặc reject student vào trong dự án
// Return: Trả về JSON application đã cập nhật hoặc lỗi
func (h *ApplicationHandler) UpdateApplicationStatus(c *gin.Context) {
	// Chỉ business mới được cập nhật status application của project họ quản lý
	role := c.GetString("role")
	businessID := c.GetString("userID")
	if role != "business" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only businesses can update application status"})
		return
	}

	applicationID := c.Param("id")
	// businessID := c.GetString("userID") // ID của business thực hiện request

	var req struct {
		Status string `json:"status" binding:"required"`
	}

	// Parse và validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Gọi service để cập nhật status (có thể truyền businessID vào để kiểm tra quyền trong service)
	application, err := h.applicationService.UpdateApplicationStatus(applicationID, req.Status, businessID)
	if err != nil {
		if err.Error() == "invalid status" { // Giả sử service trả lỗi cụ thể
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else if err.Error() == "application not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Trả về application đã cập nhật
	c.JSON(http.StatusOK, gin.H{
		"message":     "Application status updated successfully",
		"application": application,
	})
}

// DeleteApplication xử lý endpoint DELETE /api/applications/:id
// Return: Trả về thông báo xóa thành công hoặc lỗi
func (h *ApplicationHandler) DeleteApplication(c *gin.Context) {
	applicationID := c.Param("id")
	userID := c.GetString("userID")

	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	err := h.applicationService.DeleteApplication(applicationID, userID)
	if err != nil {
		// Handle specific error types
		switch err.Error() {
		case "application not found":
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case "you don't have permission to delete this application":
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Application deleted successfully",
	})
}
