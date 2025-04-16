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

	// --- CHANGE START ---
	// Proposal là nội dung đề xuất ứng tuyển (đã bị thay thế)
	// Proposal string `json:"proposal" bson:"proposal"`

	// Motivation giải thích lý do sinh viên muốn tham gia dự án
	Motivation string `json:"motivation" bson:"motivation"`

	// DetailedProposal mô tả chi tiết kế hoạch hoặc cách tiếp cận của sinh viên cho dự án
	DetailedProposal string `json:"detailed_proposal" bson:"detailed_proposal"`
	// --- CHANGE END ---

	// Status là trạng thái ứng tuyển: "pending", "approved", "rejected"
	Status string `json:"status" bson:"status"`

	// CreatedAt là thời gian tạo ứng tuyển
	CreatedAt time.Time `json:"created_at" bson:"created_at"`

	// UpdatedAt là thời gian cập nhật ứng tuyển (nếu có)
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
