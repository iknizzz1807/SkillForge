package handlers

// Nhiệm vụ: Xử lý các endpoint đánh giá và chứng chỉ
// Liên kết:
// - Gọi internal/services/review_service.go để xử lý logic
// Vai trò trong flow:
// - Được gọi từ /api/reviews để mentor/doanh nghiệp đánh giá sinh viên

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

type ReviewHandler struct {
	// reviewService xử lý logic liên quan đến review
	reviewService *services.ReviewService
}

// NewReviewHandler khởi tạo handler với reviewService
// Input: reviewService (*services.ReviewService)
// Return: *ReviewHandler - con trỏ đến ReviewHandler
func NewReviewHandler(reviewService *services.ReviewService) *ReviewHandler {
	return &ReviewHandler{reviewService}
}

// SubmitReview xử lý endpoint POST /api/reviews
// Return: Trả về JSON review vừa tạo hoặc lỗi
func (h *ReviewHandler) SubmitReview(c *gin.Context) {
	var req struct {
		TaskID  string `json:"task_id" binding:"required"`
		Score   int    `json:"score" binding:"required,min=1,max=10"`
		Comment string `json:"comment"`
	}

	// Parse và validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Lấy userID từ context
	userID := c.GetString("userID")

	// Gọi service để gửi review
	review, err := h.reviewService.SubmitReview(userID, req.TaskID, req.Score, req.Comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về review vừa tạo
	c.JSON(http.StatusCreated, review)
}

// GetReview xử lý endpoint GET /api/reviews/:id
// Return: Trả về JSON chi tiết review hoặc lỗi
func (h *ReviewHandler) GetReview(c *gin.Context) {
	reviewID := c.Param("id")

	// Gọi service để lấy chi tiết review
	review, err := h.reviewService.GetReviewByID(reviewID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Trả về chi tiết review
	c.JSON(http.StatusOK, review)
}
