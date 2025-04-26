package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/services"

)

// LevelHandler xử lý các endpoint liên quan đến level
type LevelHandler struct {
	GamificationService *services.GamificationService
}

func NewLevelHandler(GamificationService *services.GamificationService) *LevelHandler {
	return &LevelHandler{GamificationService: GamificationService}
}

func (h *LevelHandler) GetUserLevel(c *gin.Context) {
	userID := c.GetString("userID")

	level, err := h.GamificationService.GetUserLevel(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"level": level})
}