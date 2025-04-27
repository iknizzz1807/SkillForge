package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

type matchingHandler struct {
	matchingService *services.MatchingService
}

func NewMatchingHandler(matchingService *services.MatchingService) *matchingHandler {
	return &matchingHandler{
		matchingService: matchingService,
	}
}

// GetTopMatches returns top 10 project matches for the authenticated student
func (h *matchingHandler) GetTopMatches(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	matches, err := h.matchingService.GetTopProjectMatchesForStudent(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matches)
}

// GetMatchScore returns match score for specific project
func (h *matchingHandler) GetMatchScore(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	projectID := c.Param("project_id")
	if projectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "project_id is required"})
		return
	}

	score, err := h.matchingService.GetScore(userID, projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"score": score})
}
