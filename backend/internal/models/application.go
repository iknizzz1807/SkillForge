package models

// Nhiệm vụ: Định nghĩa struct Application để biểu diễn ứng tuyển dự án
// Liên kết:
// - Được dùng trong repositories/application_repository.go để lưu/lấy từ MongoDB
// - Được dùng trong services/application_service.go để xử lý logic
// Vai trò trong flow:
// - Là mô hình dữ liệu cho việc sinh viên ứng tuyển vào dự án

import "time"

// Application biểu diễn thông tin ứng tuyển của sinh viên vào dự án
type Application struct {
	// ID là định danh duy nhất của application (UUID)
	ID string `json:"id" bson:"_id"`

	// UserID là ID của sinh viên ứng tuyển
	UserID string `json:"user_id" bson:"user_id"`

	// ProjectID là ID của dự án được ứng tuyển
	ProjectID string `json:"project_id" bson:"project_id"`

	// ProjectName là trường tên của dự án để có thể hiển thị ở phần application ở phía frontend
	ProjectName string `json:"project_name" bson:"project_name"`

	// UserName là trường tên của người dùng apply vào dự án dùng để hiển thị trên frontend
	UserName string `json:"user_name" bson:"user_name"`

	// Motivation giải thích lý do sinh viên muốn tham gia dự án
	Motivation string `json:"motivation" bson:"motivation"`

	// DetailedProposal mô tả chi tiết kế hoạch hoặc cách tiếp cận của sinh viên cho dự án
	DetailedProposal string `json:"detailed_proposal" bson:"detailed_proposal"`

	// Status là trạng thái ứng tuyển: "pending", "approved", "rejected"
	Status string `json:"status" bson:"status"`

	// CreatedAt là thời gian tạo ứng tuyển
	CreatedAt time.Time `json:"created_at" bson:"created_at"`

	// UpdatedAt là thời gian cập nhật ứng tuyển (nếu có)
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
