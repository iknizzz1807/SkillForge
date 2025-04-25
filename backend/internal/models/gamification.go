package models

import (
	"time"
)

// UserLevel là cấu trúc dữ liệu đại diện cho cấp độ người dùng trong hệ thống gamification
type UserLevel struct {
	ID        string    `json:"id" bson:"_id"`
	UserID    string    `json:"user_id" bson:"user_id"`
	Level     int       `json:"level" bson:"level"`
	XPNeed    int       `json:"xp_needed" bson:"xp_needed"`
	XPCurrent int       `json:"xp_current" bson:"xp_current"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

// UserSkill là cấu trúc dữ liệu đại diện cho kỹ năng của người dùng trong hệ thống gamification
// Một point là một project có "skill" mà user đã tham gia
type UserSkill struct {
	ID           string    `json:"id" bson:"_id"`
	UserID       string    `json:"user_id" bson:"user_id"`
	Skill        string    `json:"skill" bson:"skill"`
	PointNeeded  int       `json:"point_needed" bson:"point_needed"`
	PointCurrent int       `json:"point_current" bson:"point_current"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}
