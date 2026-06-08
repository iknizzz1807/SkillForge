package models

import "time"

type Invitation struct {
	ID            string    `json:"id" bson:"_id"`
	ProjectID     string    `json:"project_id" bson:"project_id"`
	ProjectTitle  string    `json:"project_title,omitempty" bson:"-"`
	StudentID     string    `json:"student_id" bson:"student_id"`
	BusinessID    string    `json:"business_id" bson:"business_id"`
	BusinessName  string    `json:"business_name,omitempty" bson:"-"`
	Status        string    `json:"status" bson:"status"`
	CreatedAt     time.Time `json:"created_at" bson:"created_at"`
}
