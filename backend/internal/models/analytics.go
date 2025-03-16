package models

// Nhiệm vụ: Định nghĩa struct cho dữ liệu phân tích (kỹ năng, dự án)
// Liên kết:
// - Được dùng trong repositories/analytics_repository.go để lưu/lấy từ MongoDB
// - Được dùng trong services/analytics_service.go để xử lý logic
// Vai trò trong flow:
// - Là mô hình dữ liệu cho báo cáo phân tích (skill analytics, project analytics)

import "time"

// SkillAnalytics biểu diễn phân tích kỹ năng của user
type SkillAnalytics struct {
	// ID là định danh duy nhất của bản ghi phân tích (UUID)
	ID string `json:"id" bson:"_id"`

	// UserID là ID của user được phân tích
	UserID string `json:"user_id" bson:"user_id"`

	// Skill là tên kỹ năng được phân tích
	Skill string `json:"skill" bson:"skill"`

	// Proficiency là mức độ thành thạo (ví dụ: 1-100)
	Proficiency int `json:"proficiency" bson:"proficiency"`

	// LastUpdated là thời gian cập nhật phân tích
	LastUpdated time.Time `json:"last_updated" bson:"last_updated"`
}

// ProjectAnalytics biểu diễn phân tích dự án
type ProjectAnalytics struct {
	// ID là định danh duy nhất của bản ghi phân tích (UUID)
	ID string `json:"id" bson:"_id"`

	// ProjectID là ID của dự án được phân tích
	ProjectID string `json:"project_id" bson:"project_id"`

	// CompletionRate là tỷ lệ hoàn thành (0-100%)
	CompletionRate float64 `json:"completion_rate" bson:"completion_rate"`

	// AverageScore là điểm trung bình từ các review
	AverageScore float64 `json:"average_score" bson:"average_score"`

	// LastUpdated là thời gian cập nhật phân tích
	LastUpdated time.Time `json:"last_updated" bson:"last_updated"`
}
