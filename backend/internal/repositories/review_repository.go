package repositories

// Nhiệm vụ: Tương tác với collection "reviews" trong MongoDB (CRUD operations)
// Liên kết:
// - Dùng go.mongodb.org/mongo-driver/mongo để truy vấn database
// - Trả về models từ github.com/iknizzz1807/SkillForge/internal/models
// Vai trò trong flow:
// - Được gọi từ services/review_service.go để lưu/lấy review

import (
	"context"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReviewRepository struct {
	// collection để truy cập collection "reviews" trong MongoDB
	collection *mongo.Collection
}

// NewReviewRepository khởi tạo ReviewRepository với database
// Input: db (*mongo.Database)
// Return: *ReviewRepository - con trỏ đến ReviewRepository
func NewReviewRepository(db *mongo.Database) *ReviewRepository {
	return &ReviewRepository{
		collection: db.Collection("reviews"),
	}
}

// InsertReview thêm review mới vào collection "reviews"
// Input: ctx (context.Context), review (*models.Review)
// Return: error (nếu có lỗi)
func (r *ReviewRepository) InsertReview(ctx context.Context, review *models.Review) error {
	// Chèn review vào collection
	_, err := r.collection.InsertOne(ctx, review)
	if err != nil {
		return err
	}
	// Trả về nil nếu thành công
	return nil
}

// FindReviewByID tìm review theo ID
// Input: ctx (context.Context), reviewID (string)
// Return: *models.Review (review), error (nếu có lỗi)
func (r *ReviewRepository) FindReviewByID(ctx context.Context, reviewID string) (*models.Review, error) {
	// Tạo filter để tìm theo ID
	filter := bson.M{"_id": reviewID}
	var review models.Review

	// Truy vấn database
	err := r.collection.FindOne(ctx, filter).Decode(&review)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Không tìm thấy review
		}
		return nil, err
	}

	// Trả về review
	return &review, nil
}
