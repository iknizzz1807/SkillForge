package models

// Nhiệm vụ: Định nghĩa struct Message để biểu diễn tin nhắn giữa các user
// Liên kết:
// - Được dùng trong repositories/message_repository.go để lưu/lấy từ MongoDB
// - Được dùng trong services/message_service.go để xử lý logic
// Vai trò trong flow:
// - Là mô hình dữ liệu cho hệ thống chat giữa sinh viên, mentor, doanh nghiệp

import "time"

// Message biểu diễn thông tin tin nhắn
type Message struct {
	// ID là định danh duy nhất của message (UUID)
	ID string `json:"id" bson:"_id"`

	// SenderID là ID của người gửi
	SenderID string `json:"sender_id" bson:"sender_id"`

	// ReceiverID là ID của người nhận
	ReceiverID string `json:"receiver_id" bson:"receiver_id"`

	// Content là nội dung tin nhắn
	Content string `json:"content" bson:"content"`

	// CreatedAt là thời gian gửi tin nhắn
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
