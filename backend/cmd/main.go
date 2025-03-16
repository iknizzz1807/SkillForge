package main

// Nhiệm vụ: Điểm khởi chạy chính của backend, khởi tạo và chạy server.
// Liên kết: Gọi internal/app/app.go để cấu hình và chạy ứng dụng.
// Vai trò trong flow:
// - Là nơi bắt đầu toàn bộ ứng dụng SKILLFORGE.
// - Kết nối các thành phần (config, app) để khởi động server Gin, MongoDB, và các integrations.
// - Đảm bảo backend sẵn sàng nhận request từ client (frontend hoặc API caller).

import (
	"github.com/iknizzz1807/SkillForge/internal/app"
)

func main() {
	// Gọi hàm Run từ internal/app để khởi động ứng dụng
	app.Run()
}
