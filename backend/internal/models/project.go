package models

// Nhiệm vụ: Định nghĩa struct Project để biểu diễn dự án trong hệ thống
// Liên kết:
// - Được dùng trong repositories/project_repository.go để lưu/lấy từ MongoDB
// - Được dùng trong services/project_service.go để xử lý logic
// Vai trò trong flow:
// - Là mô hình dữ liệu chính cho các dự án mà doanh nghiệp tạo và sinh viên tham gia

// * Lưu ý: json và việc decode json trong body của request được gửi tới, còn tag bson là
// định nghĩa tên trường được lưu trong MongoDB

import (
	"time"
)

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

	// StartTime là thời gian bắt đầu dự án
	StartTime time.Time `json:"start_time" bson:"start_time"`

	// EndTime là thời gian kết thúc dự án
	EndTime time.Time `json:"end_time" bson:"end_time"`

	MaxMember int `json:"max_member" bson:"max_member"`

	CurrentMember int `json:"current_member" bson:"current_member"`

	// CreatedBy là ID của user tạo dự án (doanh nghiệp)
	CreatedByID string `json:"created_by_id" bson:"created_by_id"`

	// CreatedByName là name của business tạo dự án
	CreatedByName string `json:"created_by_name" bson:"created_by_name"`

	// Status là trạng thái dự án: "open", "active", "closed"
	Status string `json:"status" bson:"status"`

	// CreatedAt là thời gian tạo dự án, nó khác với thời gian bắt đầu
	CreatedAt time.Time `json:"created_at" bson:"created_at"`

	// RepoURL là đường dẫn repository GitHub (nếu có)
	// RepoURL string `json:"repo_url,omitempty" bson:"repo_url,omitempty"`
}
