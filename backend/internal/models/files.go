package models

// Nhiệm vụ: Định nghĩa struct cho dữ liệu của một file được store trong hệ thống
// Liên kết:

import "time"

type File struct {
	ID         string    `json:"id" bson:"_id"`
	FileName       string    `json:"name" bson:"name"`
	Path       string    `json:"path" bson:"path"`
	UploadedAt time.Time `json:"uploaded_at" bson:"uploaded_at"`
	UploadedByID     string    `json:"uploaded_by_id" bson:"uploaded_by_id"`
}
