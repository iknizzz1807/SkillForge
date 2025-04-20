package models

import "time"

// TalentPool represents a favorite candidate for a business
type TalentPool struct {
	ID          string    `bson:"_id" json:"id"`
	BusinessID  string    `bson:"business_id" json:"business_id"`
	StudentID   string    `bson:"student_id" json:"student_id"`
	StudentName string    `bson:"student_name" json:"student_name"`
	Skills      []string  `bson:"skills" json:"skills"`
	Notes       string    `bson:"notes" json:"notes,omitempty"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}
