package services

// Nhiệm vụ: Xử lý logic đánh giá và chứng chỉ
// Liên kết:
// - Dùng internal/repositories/review_repository.go để lưu đánh giá
// Vai trò trong flow:
// - Được gọi từ handlers/reviews.go để mentor/doanh nghiệp đánh giá sinh viên

import (
	"context"
	"errors"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReviewService struct {
	// db để truy cập MongoDB database
	db *mongo.Database
}

// NewReviewService khởi tạo ReviewService với dependency
// Input: db (*mongo.Database)
// Return: *ReviewService - con trỏ đến ReviewService
func NewReviewService(db *mongo.Database) *ReviewService {
	return &ReviewService{db}
}

// SubmitReview gửi đánh giá cho task
// Input: userID (string), taskID (string), score (int), comment (string)
// Return: *models.Review (review vừa tạo), error (nếu có lỗi)
func (s *ReviewService) SubmitReview(userID, taskID string, score int, comment string) (*models.Review, error) {
	// Kiểm tra input hợp lệ
	if userID == "" || taskID == "" || score < 1 || score > 10 {
		return nil, errors.New("invalid review data")
	}

	// Tạo review mới
	review := &models.Review{
		ID:        utils.GenerateUUID(),
		UserID:    userID,
		TaskID:    taskID,
		Score:     score,
		Comment:   comment,
		CreatedAt: time.Now(),
	}

	// Lưu review vào database
	reviewRepo := repositories.NewReviewRepository(s.db)
	err := reviewRepo.InsertReview(context.Background(), review)
	if err != nil {
		return nil, err
	}

	// Trả về review
	return review, nil
}

// GetReviewByID lấy chi tiết review
// Input: reviewID (string)
// Return: *models.Review (review), error (nếu có lỗi)
func (s *ReviewService) GetReviewByID(reviewID string) (*models.Review, error) {
	// Kiểm tra reviewID hợp lệ
	if reviewID == "" {
		return nil, errors.New("review ID cannot be empty")
	}

	// Tạo repository và tìm review
	reviewRepo := repositories.NewReviewRepository(s.db)
	review, err := reviewRepo.FindReviewByID(context.Background(), reviewID)
	if err != nil || review == nil {
		return nil, errors.New("review not found")
	}

	// Trả về review
	return review, nil
}
