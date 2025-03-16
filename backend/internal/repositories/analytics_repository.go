package repositories

// Nhiệm vụ: Tương tác với collection "analytics" trong MongoDB (lấy dữ liệu phân tích)
// Liên kết:
// - Dùng go.mongodb.org/mongo-driver/mongo để truy vấn database
// - Trả về models từ github.com/iknizzz1807/SkillForge/internal/models
// Vai trò trong flow:
// - Được gọi từ services/analytics_service.go để cung cấp dữ liệu phân tích

import (
	"context"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AnalyticsRepository struct {
	// collection để truy cập collection "analytics" trong MongoDB
	collection *mongo.Collection
}

// NewAnalyticsRepository khởi tạo AnalyticsRepository với database
// Input: db (*mongo.Database)
// Return: *AnalyticsRepository - con trỏ đến AnalyticsRepository
func NewAnalyticsRepository(db *mongo.Database) *AnalyticsRepository {
	return &AnalyticsRepository{
		collection: db.Collection("analytics"),
	}
}

// GetSkillAnalyticsByUser lấy dữ liệu phân tích kỹ năng của user
// Input: ctx (context.Context), userID (string)
// Return: []models.SkillAnalytics (danh sách kỹ năng), error (nếu có lỗi)
func (r *AnalyticsRepository) GetSkillAnalyticsByUser(ctx context.Context, userID string) ([]models.SkillAnalytics, error) {
	// Tạo filter để tìm theo userID
	filter := bson.M{"user_id": userID}

	// Tạo cursor để lấy tất cả document
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Tạo slice để lưu danh sách kỹ năng
	var skills []models.SkillAnalytics
	// Duyệt qua cursor và decode từng document
	for cursor.Next(ctx) {
		var skill models.SkillAnalytics
		if err := cursor.Decode(&skill); err != nil {
			return nil, err
		}
		skills = append(skills, skill)
	}

	// Trả về danh sách kỹ năng
	return skills, nil
}

// GetProjectAnalytics lấy dữ liệu phân tích dự án
// Input: ctx (context.Context), projectID (string)
// Return: *models.ProjectAnalytics (phân tích dự án), error (nếu có lỗi)
func (r *AnalyticsRepository) GetProjectAnalytics(ctx context.Context, projectID string) (*models.ProjectAnalytics, error) {
	// Tạo filter để tìm theo projectID
	filter := bson.M{"project_id": projectID}
	var analytics models.ProjectAnalytics

	// Truy vấn database
	err := r.collection.FindOne(ctx, filter).Decode(&analytics)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Không tìm thấy analytics
		}
		return nil, err
	}

	// Trả về analytics
	return &analytics, nil
}
