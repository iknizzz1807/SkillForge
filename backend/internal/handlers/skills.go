package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

type SkillHandler struct {
	gamificationService *services.GamificationService
}

func NewSkillHandler(gamificationService *services.GamificationService) *SkillHandler {
	return &SkillHandler{gamificationService: gamificationService}
}

func (h *SkillHandler) GetUserSkills(c *gin.Context) {
	userID := c.GetString("userID")

	skills, err := h.gamificationService.GetAllSkills(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, skills)
}
