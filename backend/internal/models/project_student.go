package models

import "time"

// Nhiệm vụ: Định nghĩa struct Project_student để biểu diễn dự án trong hệ thống
// Liên kết:
// - Được dùng trong repositories/project_student_repository.go để lưu/lấy từ MongoDB
// - Được dùng trong services/project_service.go để xử lý logic về việc tham gia dự án của học sinh
// Vai trò trong flow:
// - Là mô hình dữ liệu chính để thực hiện việc sinh viên tham gia dự án
// - Kết hợp với model Project và Project Service để tham gia dự án cho sinh viên

type Project_student struct {
	// ID là trường ID của dữ liệu này
	ID string `json:"id" bson:"_id"`

	// Project_id là trường string trỏ tới project id mà student thuộc về/tham gia
	Project_id string `json:"project_id" bson:"project_id"`

	// Student_id là trường string trỏ tới student id thuộc về project/tham gia project
	Student_id string `json:"student_id" bson:"student_id"`

	// CreatedAt là thời gian mà dòng dữ liệu này được tạo trong database
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
