package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

// BadgeHandler xử lý các endpoint liên quan đến badge
type BadgeHandler struct {
	badgeService *services.BadgeService
}

// NewBadgeHandler khởi tạo BadgeHandler với badgeService
func NewBadgeHandler(badgeService *services.BadgeService) *BadgeHandler {
	return &BadgeHandler{badgeService}
}

// GetAllBadges xử lý endpoint GET /api/badges
func (h *BadgeHandler) GetAllBadges(c *gin.Context) {
	// Lấy query param type nếu có
	badgeType := c.Query("type")

	var badges []*models.Badge
	var err error

	// Nếu có type, lấy badges theo type
	if badgeType != "" {
		badges, err = h.badgeService.GetBadgesByType(badgeType)
	} else {
		// Nếu không có type, lấy tất cả badges
		badges, err = h.badgeService.GetAllBadges()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Đảm bảo không trả về null cho mảng trống
	if badges == nil {
		badges = []*models.Badge{}
	}

	c.JSON(http.StatusOK, badges)
}

// GetBadge xử lý endpoint GET /api/badges/:id
func (h *BadgeHandler) GetBadge(c *gin.Context) {
	badgeID := c.Param("id")

	badge, err := h.badgeService.GetBadgeByID(badgeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, badge)
}

// GetUserBadges xử lý endpoint GET /api/users/:id/badges
func (h *BadgeHandler) GetUserBadges(c *gin.Context) {
	userID := c.Param("id")

	userBadges, err := h.badgeService.GetUserBadges(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Đảm bảo không trả về null cho mảng trống
	if userBadges == nil {
		userBadges = []*models.UserBadge{}
	}

	c.JSON(http.StatusOK, userBadges)
}

// // CreateBadge xử lý endpoint POST /api/badges
// func (h *BadgeHandler) CreateBadge(c *gin.Context) {
// 	// Chỉ admin mới được tạo badge
// 	if c.GetString("role") != "admin" {
// 		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can create badges"})
// 		return
// 	}

// 	var req struct {
// 		Name          string                 `json:"name" binding:"required"`
// 		Description   string                 `json:"description" binding:"required"`
// 		Type          string                 `json:"type" binding:"required"`
// 		Icon          string                 `json:"icon" binding:"required"`
// 		Prerequisites map[string]interface{} `json:"prerequisites" binding:"required"`
// 	}

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	badge, err := h.badgeService.CreateBadge(
// 		req.Name,
// 		req.Description,
// 		req.Type,
// 		req.Prerequisites,
// 	)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, badge)
// }
