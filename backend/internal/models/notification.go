package models

import "time"

type Notification struct {
	ID        string    `json:"id" bson:"_id"`
	ToUserID  string    `json:"to_user_id" bson:"to_user_id"`
	Content   string    `json:"content" bson:"content"`
	Type      string    `json:"type" bson:"type"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
