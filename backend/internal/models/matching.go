package models

import "time"

type Match struct {
	StudentID  string    `bson:"student_id" json:"student_id"`
	ProjectID  string    `bson:"project_id" json:"project_id"`
	Score      float64   `bson:"score" json:"score"`
	CreatedAt  time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at" json:"updated_at"`
}
