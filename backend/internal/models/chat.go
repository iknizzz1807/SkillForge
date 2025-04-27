package models

import "time"

// Message biểu diễn thông tin tin nhắn
type Message struct {
	// ID là định danh duy nhất của message (UUID)
	ID string `json:"id" bson:"_id"`

	// SenderID là ID của người gửi
	SenderID string `json:"sender_id" bson:"sender_id"`

	// GroupID là ID của group
	GroupID string `json:"group_id" bson:"group_id"`

	// Content là nội dung tin nhắn
	Content string `json:"content" bson:"content"`

	// CreatedAt là thời gian gửi tin nhắn
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type Group struct {
	ProjectID string `json:"project_id" bson:"project_id"`
	Title string `json:"title" bson:"title"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}