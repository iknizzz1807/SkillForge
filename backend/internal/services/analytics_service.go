package services

// Nhiệm vụ: Xử lý logic phân tích (kỹ năng, dự án)
// Liên kết:
// - Dùng internal/repositories/analytics_repository.go để lấy dữ liệu phân tích
// Vai trò trong flow:
// - Được gọi từ handlers/analytics.go để cung cấp báo cáo

import (
	"errors"

	"context"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type AnalyticsService struct {
	// db để truy cập MongoDB database
	db *mongo.Database
}

// NewAnalyticsService khởi tạo AnalyticsService với dependency
// Input: db (*mongo.Database)
// Return: *AnalyticsService - con trỏ đến AnalyticsService
func NewAnalyticsService(db *mongo.Database) *AnalyticsService {
	return &AnalyticsService{db}
}

// GetSkillAnalytics lấy phân tích kỹ năng của user
// Input: userID (string)
// Return: []models.SkillAnalytics (danh sách kỹ năng), error (nếu có lỗi)
func (s *AnalyticsService) GetSkillAnalytics(userID string) ([]models.SkillAnalytics, error) {
	// Kiểm tra userID hợp lệ
	if userID == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	// Tạo repository và lấy dữ liệu phân tích
	analyticsRepo := repositories.NewAnalyticsRepository(s.db)
	skills, err := analyticsRepo.GetSkillAnalyticsByUser(context.Background(), userID)
	if err != nil {
		return nil, err
	}

	// Trả về danh sách kỹ năng
	return skills, nil
}

// GetProjectAnalytics lấy phân tích dự án
// Input: projectID (string)
// Return: *models.ProjectAnalytics (phân tích dự án), error (nếu có lỗi)
func (s *AnalyticsService) GetProjectAnalytics(projectID string) (*models.ProjectAnalytics, error) {
	// Kiểm tra projectID hợp lệ
	if projectID == "" {
		return nil, errors.New("project ID cannot be empty")
	}

	// Tạo repository và lấy dữ liệu phân tích
	analyticsRepo := repositories.NewAnalyticsRepository(s.db)
	analytics, err := analyticsRepo.GetProjectAnalytics(context.Background(), projectID)
	if err != nil || analytics == nil {
		return nil, errors.New("project analytics not found")
	}

	// Trả về phân tích dự án
	return analytics, nil
}
