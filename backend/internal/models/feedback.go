package models

import "time"

// sinh viên feedback bussiness và ngược lại
type Feedback struct {
	ID        string    `json:"id" bson:"_id"`
	ProjectID string    `json:"project_id" bson:"project_id"`
	FromID    string    `json:"from" bson:"from"`
	ToID      string    `json:"to" bson:"to"`
	Type      string    `json:"type" bson:"type"`     // "business" feedback cho business or "student" feedback cho sinh viên 
	Rating    int       `json:"rating" bson:"rating"` // rating từ 0 đến 5 (star)
	Content   string    `json:"content" bson:"content"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
