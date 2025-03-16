package models

// Nhiệm vụ: Định nghĩa struct User để biểu diễn người dùng trong hệ thống
// Liên kết:
// - Được dùng trong repositories/user_repository.go để lưu/lấy từ MongoDB
// - Được dùng trong services/auth_service.go và user_service.go để xử lý logic
// Vai trò trong flow:
// - Là mô hình dữ liệu chính cho sinh viên, doanh nghiệp, mentor, admin

import "time"

// User biểu diễn thông tin người dùng trong hệ thống
type User struct {
	// ID là định danh duy nhất của user (UUID)
	ID string `json:"id" bson:"_id"`

	// Email là địa chỉ email của user, dùng để đăng nhập
	Email string `json:"email" bson:"email"`

	// Password là mật khẩu đã được hash, không trả về trong JSON
	Password string `json:"-" bson:"password"`

	// Role xác định vai trò: "student", "business", "mentor", "admin"
	Role string `json:"role" bson:"role"`

	// Skills là danh sách kỹ năng của user (chủ yếu cho student)
	Skills []string `json:"skills" bson:"skills"`

	// CreatedAt là thời gian tạo user
	CreatedAt time.Time `json:"created_at" bson:"created_at"`

	// UpdatedAt là thời gian cập nhật user (nếu có)
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
