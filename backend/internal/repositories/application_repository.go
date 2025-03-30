package repositories

// Nhiệm vụ: Tương tác với collection "applications" trong MongoDB (CRUD operations)
// Liên kết:
// - Dùng go.mongodb.org/mongo-driver/mongo để truy vấn database
// - Trả về models từ github.com/iknizzz1807/SkillForge/internal/models
// Vai trò trong flow:
// - Được gọi từ services/application_service.go để lưu/lấy ứng tuyển

import (
	"context"

	"time"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApplicationRepository struct {
	// collection để truy cập collection "applications" trong MongoDB
	collection *mongo.Collection
}

// NewApplicationRepository khởi tạo ApplicationRepository với database
// Input: db (*mongo.Database)
// Return: *ApplicationRepository - con trỏ đến ApplicationRepository
func NewApplicationRepository(db *mongo.Database) *ApplicationRepository {
	return &ApplicationRepository{
		collection: db.Collection("applications"),
	}
}

// InsertApplication thêm ứng tuyển mới vào collection "applications"
// Input: ctx (context.Context), application (*models.Application)
// Return: error (nếu có lỗi)
func (r *ApplicationRepository) InsertApplication(ctx context.Context, application *models.Application) error {
	// Chèn application vào collection
	_, err := r.collection.InsertOne(ctx, application)
	if err != nil {
		return err
	}
	// Trả về nil nếu thành công
	return nil
}

// FindApplicationByID tìm ứng tuyển theo ID
// Input: ctx (context.Context), applicationID (string)
// Return: *models.Application (ứng tuyển), error (nếu có lỗi)
func (r *ApplicationRepository) FindApplicationByID(ctx context.Context, applicationID string) (*models.Application, error) {
	// Tạo filter để tìm theo ID
	filter := bson.M{"_id": applicationID}
	var application models.Application

	// Truy vấn database
	err := r.collection.FindOne(ctx, filter).Decode(&application)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Không tìm thấy application
		}
		return nil, err
	}

	// Trả về application
	return &application, nil
}

// FindByUserID tìm tất cả applications của một user
func (r *ApplicationRepository) FindByUserID(userID string) ([]models.Application, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var applications []models.Application
	if err := cursor.All(ctx, &applications); err != nil {
		return nil, err
	}
	return applications, nil
}

// FindByProjectIDs tìm tất cả applications cho danh sách các projectID
func (r *ApplicationRepository) FindByProjectIDs(projectIDs []string) ([]models.Application, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"project_id": bson.M{"$in": projectIDs}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var applications []models.Application
	if err := cursor.All(ctx, &applications); err != nil {
		return nil, err
	}
	return applications, nil
}

// UpdateStatus cập nhật trạng thái application
func (r *ApplicationRepository) UpdateStatus(id string, status string) (*models.Application, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"status":     status,
			"updated_at": time.Now(),
		},
	}

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return nil, err
	}

	return r.FindApplicationByID(ctx, id)
}

// Delete xóa một application
func (r *ApplicationRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
