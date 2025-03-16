package utils

// Nhiệm vụ: Định nghĩa các custom error cho ứng dụng
// Liên kết:
// - Được gọi từ services hoặc handlers để trả về lỗi cụ thể
// Vai trò trong flow:
// - Chuẩn hóa lỗi để dễ xử lý và trả về client

import "errors"

// Custom errors
var (
	// ErrInvalidInput khi dữ liệu đầu vào không hợp lệ
	ErrInvalidInput = errors.New("invalid input data")

	// ErrNotFound khi không tìm thấy tài nguyên
	ErrNotFound = errors.New("resource not found")

	// ErrUnauthorized khi user không có quyền truy cập
	ErrUnauthorized = errors.New("unauthorized access")

	// ErrInternalServer khi có lỗi server không xác định
	ErrInternalServer = errors.New("internal server error")
)

// NewCustomError tạo lỗi tùy chỉnh với message
// Input: message (string)
// Return: error - lỗi mới
func NewCustomError(message string) error {
	return errors.New(message)
}
