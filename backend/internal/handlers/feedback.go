package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

type FeedbackHandler struct {
	feedbackService *services.FeedbackService
}

func NewFeedbackHandler(feedbackService *services.FeedbackService) *FeedbackHandler {
	return &FeedbackHandler{feedbackService: feedbackService}
}

// POST /feedbacks
func (h *FeedbackHandler) CreateFeedback(c *gin.Context) {
	var req struct {
		ProjectID string `json:"project_id" binding:"required"`
		ToID      string `json:"to_id" binding:"required"`
		Type      string `json:"type" binding:"required"`
		Rating    int    `json:"rating" binding:"required"`
		Content   string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	fromID := c.GetString("userID")
	if fromID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	feedback, err := h.feedbackService.CreateFeedback(
		req.ProjectID,
		fromID,
		req.ToID,
		req.Type,
		req.Content,
		req.Rating,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, feedback)
}

// GET /feedbacks/student/:userID
func (h *FeedbackHandler) GetStudentFeedbacks(c *gin.Context) {
	userID := c.Param("userID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	feedbacks, err := h.feedbackService.GetUserFeedbacks(userID, "student")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, feedbacks)
}

// GET /feedbacks/business/:userID
func (h *FeedbackHandler) GetBusinessFeedbacks(c *gin.Context) {
	userID := c.Param("userID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	feedbacks, err := h.feedbackService.GetUserFeedbacks(userID, "business")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, feedbacks)
}
