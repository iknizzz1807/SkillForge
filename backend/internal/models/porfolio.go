package models

// Nhiệm vụ: Định nghĩa struct Portfolio để biểu diễn portfolio của sinh viên
// Liên kết:
// - Được dùng trong repositories/portfolio_repository.go để lưu/lấy từ MongoDB
// - Được dùng trong services/portfolio_service.go để xử lý logic
// Vai trò trong flow:
// - Là mô hình dữ liệu tổng hợp thành tích của sinh viên (dự án, kỹ năng, v.v.)

import "time"

// Portfolio biểu diễn thông tin portfolio của sinh viên
type Portfolio struct {
	// ID là định danh duy nhất của portfolio (UUID)
	ID string `json:"id" bson:"_id"`

	// UserID là ID của sinh viên sở hữu portfolio
	UserID string `json:"user_id" bson:"user_id"`

	// Projects là danh sách ID dự án đã tham gia
	Projects []string `json:"projects" bson:"projects"`

	// Skills là danh sách kỹ năng đã chứng minh
	Skills []string `json:"skills" bson:"skills"`

	// CreatedAt là thời gian tạo portfolio
	CreatedAt time.Time `json:"created_at" bson:"created_at"`

	// UpdatedAt là thời gian cập nhật portfolio (nếu có)
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
