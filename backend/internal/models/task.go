package models

// Nhiệm vụ: Định nghĩa struct Task để biểu diễn công việc trong dự án
// Liên kết:
// - Được dùng trong repositories/task_repository.go để lưu/lấy từ MongoDB
// - Được dùng trong services/task_service.go để xử lý logic
// Vai trò trong flow:
// - Là mô hình dữ liệu cho các task trong dự án, hỗ trợ Kanban board

import "time"

// Task biểu diễn thông tin công việc trong dự án
type Task struct {
	// ID là định danh duy nhất của task (UUID)
	ID string `json:"id" bson:"_id"`

	// ProjectID là ID của dự án chứa task
	ProjectID string `json:"project_id" bson:"project_id"`

	// Description là mô tả chi tiết task
	Description string `json:"description" bson:"description"`

	// Status là trạng thái task: "todo", "in_progress", "review", "done"
	Status string `json:"status" bson:"status"`

	// CreatedAt là thời gian tạo task
	CreatedAt time.Time `json:"created_at" bson:"created_at"`

	// UpdatedAt là thời gian cập nhật task (nếu có)
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
