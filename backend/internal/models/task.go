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

	// Title là tiêu đề của task
	Title string `json:"title" bson:"title"`

	// Description là mô tả chi tiết task
	Description string `json:"description" bson:"description"`

	// Status là trạng thái task: "todo", "in_progress", "review", "done"
	Status string `json:"status" bson:"status"`

	// Review là nhận xét của business về task đó, và có thể có hay không có
	// bởi vì task có thể được review xấu rồi từ status "in_progress" hay "review" quay về status todo
	Review string `json:"review" bson:"review"`

	// Note là một trường khác ngoài description để ghi chú thích về task
	Note string `json:"note" bson:"note"`

	// Assigned by là id của người chịu trách nhiệm hoàn thành task này
	Assigned_to string `json:"assigned_to" bson:"assigned_to"`

	// Finished_by là id của người đã hoàn thành task này, để trống nếu đang ở status
	//  "todo" hoặc "in_progress"
	Finished_by string `json:"finished_by" bson:"finished_by"`

	// CreatedAt là thời gian tạo task
	CreatedAt time.Time `json:"created_at" bson:"created_at"`

	// UpdatedAt là thời gian cập nhật task (nếu có)
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type TaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Note        string `json:"note"`
	AssignedTo  string `json:"assigned_to"`
}

type TaskUpdate struct {
	TaskID      string `json:"taskId" binding:"required"`
	Status      string `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Note        string `json:"note"`
	Review      string `json:"review"`
	Assigned_to string `json:"assigned_to"`
}