package repositories

// Nhiệm vụ: Tương tác với collection "portfolios" trong MongoDB (CRUD operations)
// Liên kết:
// - Dùng go.mongodb.org/mongo-driver/mongo để truy vấn database
// - Trả về models từ github.com/iknizzz1807/SkillForge/internal/models
// Vai trò trong flow:
// - Được gọi từ services/portfolio_service.go để lưu/lấy portfolio

import (
	"context"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PortfolioRepository struct {
	// collection để truy cập collection "portfolios" trong MongoDB
	collection *mongo.Collection
}

// NewPortfolioRepository khởi tạo PortfolioRepository với database
// Input: db (*mongo.Database)
// Return: *PortfolioRepository - con trỏ đến PortfolioRepository
func NewPortfolioRepository(db *mongo.Database) *PortfolioRepository {
	return &PortfolioRepository{
		collection: db.Collection("portfolios"),
	}
}

// FindPortfolioByUserID tìm portfolio theo userID
// Input: ctx (context.Context), userID (string)
// Return: *models.Portfolio (portfolio), error (nếu có lỗi)
func (r *PortfolioRepository) FindPortfolioByUserID(ctx context.Context, userID string) (*models.Portfolio, error) {
	// Tạo filter để tìm theo userID
	filter := bson.M{"user_id": userID}
	var portfolio models.Portfolio

	// Truy vấn database
	err := r.collection.FindOne(ctx, filter).Decode(&portfolio)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Không tìm thấy portfolio
		}
		return nil, err
	}

	// Trả về portfolio
	return &portfolio, nil
}
