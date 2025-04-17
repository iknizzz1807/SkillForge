package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

type TalentPoolHandler struct {
	talentPoolService *services.TalentPoolService
}

func NewTalentPoolHandler(talentPoolService *services.TalentPoolService) *TalentPoolHandler {
	return &TalentPoolHandler{talentPoolService}
}

func (h *TalentPoolHandler) GetTalentPool(c *gin.Context) {
	businessID := c.Param("id")

	StudentInfos, err := h.talentPoolService.GetTalentPoolByBusinessID(businessID)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, StudentInfos)
}