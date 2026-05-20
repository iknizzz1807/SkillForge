package handlers

// Nhiệm vụ: Xử lý các endpoint quản lý portfolio sinh viên
// Liên kết:
// - Gọi internal/services/portfolio_service.go để xử lý logic
// Vai trò trong flow:
// - Được gọi từ /api/portfolios để lấy portfolio của sinh viên

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

type PortfolioHandler struct {
	// portfolioService xử lý logic liên quan đến portfolio
	portfolioService *services.PortfolioService
}

// NewPortfolioHandler khởi tạo handler với portfolioService
// Input: portfolioService (*services.PortfolioService)
// Return: *PortfolioHandler - con trỏ đến PortfolioHandler
func NewPortfolioHandler(portfolioService *services.PortfolioService) *PortfolioHandler {
	return &PortfolioHandler{portfolioService}
}

// GeneratePortfolio xử lý endpoint POST /api/portfolios
func (h *PortfolioHandler) GeneratePortfolio(c *gin.Context) {
	// Lấy userID từ context hoặc body (giả định từ Auth middleware)
	userID := c.GetString("userID")
	if userID == "" {
		// Dùng userID truyền lên từ body nếu có, hoặc trả lỗi
		var requestBody struct {
			UserID string `json:"user_id"`
		}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}
		userID = requestBody.UserID
	}

	url, err := h.portfolioService.GeneratePortfolio(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Portfolio generated successfully",
		"url": url,
	})
}

// GetPortfolio xử lý endpoint GET /api/portfolios/:userID
// Return: Trả về JSON portfolio hoặc lỗi
func (h *PortfolioHandler) GetPortfolio(c *gin.Context) {
	userID := c.Param("userID")

	// Gọi service để lấy portfolio
	portfolio, err := h.portfolioService.GetPortfolio(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Trả về portfolio
	c.JSON(http.StatusOK, portfolio)
}
