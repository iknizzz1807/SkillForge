package models

// Nhiệm vụ: Định nghĩa struct Review để biểu diễn đánh giá task
// Liên kết:
// - Được dùng trong repositories/review_repository.go để lưu/lấy từ MongoDB
// - Được dùng trong services/review_service.go để xử lý logic
// Vai trò trong flow:
// - Là mô hình dữ liệu cho đánh giá từ mentor/doanh nghiệp về task của sinh viên

import "time"

// Review biểu diễn thông tin đánh giá task
type Review struct {
	// ID là định danh duy nhất của review (UUID)
	ID string `json:"id" bson:"_id"`

	// UserID là ID của người đánh giá (mentor/doanh nghiệp)
	UserID string `json:"user_id" bson:"user_id"`

	// TaskID là ID của task được đánh giá
	TaskID string `json:"task_id" bson:"task_id"`

	// Score là điểm số (1-10)
	Score int `json:"score" bson:"score"`

	// Comment là nhận xét (nếu có)
	Comment string `json:"comment,omitempty" bson:"comment,omitempty"`

	// CreatedAt là thời gian tạo review
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
