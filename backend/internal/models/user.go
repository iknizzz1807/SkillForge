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

	// Name là name của user, dùng để hiển thị thông tin trong app
	Name string `json:"name" bson:"name"`

	// Password là mật khẩu đã được hash, không trả về trong JSON
	Password string `json:"-" bson:"password"`

	// Role xác định vai trò: "student", "business", "mentor", "admin"
	Role string `json:"role" bson:"role"`

	// Skills là danh sách kỹ năng của user (chủ yếu cho student)
	Skills []string `json:"skills" bson:"skills"`

	// Title là trường string thể hiện mô tả ngắn gọn về công ty hoặc về bản thân student, ví dụ "Web developer",
	// IT consultant company,...
	// Và trường title này chỉ có thể được thêm vào thông qua việc cập nhật hồ sơ của user chứ không phải là trong
	// quá trình tạo tài khoản
	Title string `json:"title" bson:"title"`

	// CreatedAt là thời gian tạo user
	CreatedAt time.Time `json:"created_at" bson:"created_at"`

	// AvatarUrl là địa chỉ url của avatar user để frontend có thể get request
	AvatarName string `json:"avatar_name" bson:"avatar_name"`

	// Website là string thể hiện địa chỉ web của doanh nghiệp để tăng độ uy tín và cung cấp thông tin để user tìm hiểu về doanh nghiệp đó
	Website string `json:"website" bson:"website"`

	// UpdatedAt là thời gian cập nhật user (nếu có)
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
