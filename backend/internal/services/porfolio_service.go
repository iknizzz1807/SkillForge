package services

// Nhiệm vụ: Xử lý logic quản lý portfolio sinh viên
// Liên kết:
// - Dùng internal/repositories/portfolio_repository.go để lưu/lấy portfolio
// Vai trò trong flow:
// - Được gọi từ handlers/portfolios.go để lấy portfolio

import (
	"errors"

	"context"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type PortfolioService struct {
	// db để truy cập MongoDB database
	db *mongo.Database
}

// NewPortfolioService khởi tạo PortfolioService với dependency
// Input: db (*mongo.Database)
// Return: *PortfolioService - con trỏ đến PortfolioService
func NewPortfolioService(db *mongo.Database) *PortfolioService {
	return &PortfolioService{db}
}

// GetPortfolio lấy portfolio của sinh viên
// Input: userID (string)
// Return: *models.Portfolio (portfolio), error (nếu có lỗi)
func (s *PortfolioService) GetPortfolio(userID string) (*models.Portfolio, error) {
	// Kiểm tra userID hợp lệ
	if userID == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	// Tạo repository và tìm portfolio
	portfolioRepo := repositories.NewPortfolioRepository(s.db)
	portfolio, err := portfolioRepo.FindPortfolioByUserID(context.Background(), userID)
	if err != nil || portfolio == nil {
		return nil, errors.New("portfolio not found")
	}

	// Trả về portfolio
	return portfolio, nil
}
