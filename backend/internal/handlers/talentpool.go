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

// Fix the service of this handler
func (h *TalentPoolHandler) GetTalentPool(c *gin.Context) {
	businessID := c.Param("id")

	StudentInfos, err := h.talentPoolService.GetTalentPoolByBusinessID(businessID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, StudentInfos)
}

// Dùng để serve POST request cho endpoint /api/talentpool/:id với id là id của student
// Return StudentInfo với các trường id name và skills là danh sách của các string
func (h *TalentPoolHandler) AddStudentToTalentPool(c *gin.Context) {
	studentId := c.Param("id")
	businessId := c.GetString("userID")

	// Call function from talent pool service to deal with this request
	StudentInfo, err := h.talentPoolService.AddStudentToTalentPool(studentId, businessId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot add student to talent pool"})
	}
	c.JSON(http.StatusCreated, StudentInfo)
}
