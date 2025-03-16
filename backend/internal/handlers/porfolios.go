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
