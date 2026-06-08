package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

type InvitationHandler struct {
	invitationService *services.InvitationService
}

func NewInvitationHandler(invitationService *services.InvitationService) *InvitationHandler {
	return &InvitationHandler{invitationService: invitationService}
}

func (h *InvitationHandler) CreateInvitation(c *gin.Context) {
	var req struct {
		StudentID string `json:"student_id" binding:"required"`
		ProjectID string `json:"project_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	businessID := c.GetString("userID")

	invitation, err := h.invitationService.CreateInvitation(req.StudentID, req.ProjectID, businessID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, invitation)
}

func (h *InvitationHandler) GetMyInvitations(c *gin.Context) {
	userID := c.GetString("userID")
	invitations, err := h.invitationService.GetInvitationsByStudent(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if invitations == nil {
		invitations = []*models.Invitation{}
	}
	c.JSON(http.StatusOK, invitations)
}

func (h *InvitationHandler) RespondToInvitation(c *gin.Context) {
	invitationID := c.Param("id")
	studentID := c.GetString("userID")

	var req struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.invitationService.RespondToInvitation(invitationID, studentID, req.Status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invitation " + req.Status + " successfully"})
}
