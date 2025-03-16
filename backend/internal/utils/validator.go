package utils

// Nhiệm vụ: Cung cấp hàm validation dữ liệu đầu vào
// Liên kết:
// - Dùng github.com/go-playground/validator/v10 để validate struct
// - Được gọi từ handlers để kiểm tra request body trước khi xử lý
// Vai trò trong flow:
// - Đảm bảo dữ liệu đầu vào hợp lệ trước khi truyền vào service

import (
	"github.com/go-playground/validator/v10"
)

// ValidateStruct kiểm tra tính hợp lệ của struct dựa trên các tag
// Input: data (interface{}) - struct cần validate
// Return: error - lỗi nếu dữ liệu không hợp lệ, nil nếu hợp lệ
func ValidateStruct(data interface{}) error {
	// Khởi tạo validator
	validate := validator.New()

	// Validate struct
	err := validate.Struct(data)
	if err != nil {
		// Trả về lỗi validation
		return err
	}

	// Trả về nil nếu hợp lệ
	return nil
}
