package models

// Nhiệm vụ: Định nghĩa struct Project để biểu diễn dự án trong hệ thống
// Liên kết:
// - Được dùng trong repositories/project_repository.go để lưu/lấy từ MongoDB
// - Được dùng trong services/project_service.go để xử lý logic
// Vai trò trong flow:
// - Là mô hình dữ liệu chính cho các dự án mà doanh nghiệp tạo và sinh viên tham gia

import "time"

// Project biểu diễn thông tin dự án trong hệ thống
type Project struct {
	// ID là định danh duy nhất của project (UUID)
	ID string `json:"id" bson:"_id"`

	// Title là tiêu đề của dự án
	Title string `json:"title" bson:"title"`

	// Description là mô tả chi tiết dự án
	Description string `json:"description" bson:"description"`

	// Skills là danh sách kỹ năng yêu cầu cho dự án
	Skills []string `json:"skills" bson:"skills"`

	// Timeline là thời gian dự kiến hoàn thành (ví dụ: "2 weeks")
	Timeline string `json:"timeline" bson:"timeline"`

	// CreatedBy là ID của user tạo dự án (doanh nghiệp)
	CreatedBy string `json:"created_by" bson:"created_by"`

	// Status là trạng thái dự án: "open", "active", "closed"
	Status string `json:"status" bson:"status"`

	// RepoURL là đường dẫn repository GitHub (nếu có)
	RepoURL string `json:"repo_url,omitempty" bson:"repo_url,omitempty"`

	// CreatedAt là thời gian tạo dự án
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
