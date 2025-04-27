package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

type BusinessInfoHandler struct {
	businessInfoService *services.BusinessInfoService
}

// NewBusinessInfoHandler creates a new BusinessInfoHandler
func NewBusinessInfoHandler(businessInfoService *services.BusinessInfoService) *BusinessInfoHandler {
	return &BusinessInfoHandler{businessInfoService}
}

// GetBusinessInfo handles GET /api/business-info
func (h *BusinessInfoHandler) GetBusinessInfo(c *gin.Context) {
	userID := c.GetString("userID")
	role := c.GetString("role")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Nếu không phải business thì không trả về gì cả
	if role != "business" {
		c.JSON(http.StatusOK, nil)
		return
	}

	businessInfo, err := h.businessInfoService.GetBusinessInfo(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, businessInfo)
}

// UpdateBusinessInfo handles PUT /api/business-info
func (h *BusinessInfoHandler) UpdateBusinessInfo(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Extract data from form
	companyType := c.PostForm("companyType")
	founded := c.PostForm("founded")
	companySize := c.PostForm("companySize")
	website := c.PostForm("website")
	aboutUs := c.PostForm("aboutUs")

	// Update business info
	businessInfo, err := h.businessInfoService.UpdateBusinessInfo(
		userID, companyType, founded, companySize, website, aboutUs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, businessInfo)
}
