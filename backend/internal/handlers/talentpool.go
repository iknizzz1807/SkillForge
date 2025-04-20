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

// GetTalentPool handles GET /api/talentpool
// Returns the list of students in the business's talent pool
func (h *TalentPoolHandler) GetTalentPool(c *gin.Context) {
	// Get business ID from token
	businessID := c.GetString("userID")

	// Check if user is a business
	if c.GetString("role") != "business" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only business accounts can access talent pool"})
		return
	}

	studentInfos, err := h.talentPoolService.GetTalentPoolByBusinessID(businessID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, studentInfos)
}

// AddStudentToTalentPool handles POST /api/talentpool/:id
// Adds a student to the business's talent pool
func (h *TalentPoolHandler) AddStudentToTalentPool(c *gin.Context) {
	studentID := c.Param("id")
	businessID := c.GetString("userID")

	// Check if user is a business
	if c.GetString("role") != "business" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only business accounts can add to talent pool"})
		return
	}

	// Call function from talent pool service to handle request
	studentInfo, err := h.talentPoolService.AddStudentToTalentPool(studentID, businessID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, studentInfo)
}

// RemoveFromTalentPool handles DELETE /api/talentpool/:id
// Removes a student from the business's talent pool
func (h *TalentPoolHandler) RemoveFromTalentPool(c *gin.Context) {
	studentID := c.Param("id")
	businessID := c.GetString("userID")

	// Check if user is a business
	if c.GetString("role") != "business" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only business accounts can remove from talent pool"})
		return
	}

	// Call service to remove student from talent pool
	err := h.talentPoolService.RemoveFromTalentPool(studentID, businessID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student removed from talent pool successfully"})
}
