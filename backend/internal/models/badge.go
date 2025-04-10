package models

import (
	"time"
)

// Badge biểu diễn huy hiệu/thành tích mà sinh viên đạt được
type Badge struct {
	// ID là định danh duy nhất của badge
	ID string `json:"id" bson:"_id"`

	// Code là mã duy nhất để xác định badge, dùng trong code
	Code string `json:"code" bson:"code"`

	// Name là tên của badge
	Name string `json:"name" bson:"name"`

	// Description mô tả chi tiết về badge và cách đạt được
	Description string `json:"description" bson:"description"`

	// Type phân loại badge: "skill", "completion", "collaboration", "achievement"
	Type string `json:"type" bson:"type"`

	// Prerequisites là các điều kiện cần để đạt được badge
	Prerequisites map[string]interface{} `json:"prerequisites" bson:"prerequisites"`

	// CreatedAt là thời điểm tạo badge
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

// UserBadge biểu diễn mối quan hệ giữa user và badge họ đạt được
type UserBadge struct {
	// ID là định danh duy nhất của user_badge
	ID string `json:"id" bson:"_id"`

	// UserID là ID của user đạt được badge
	UserID string `json:"user_id" bson:"user_id"`

	// BadgeID là ID của badge được đạt
	BadgeID string `json:"badge_id" bson:"badge_id"`

	// Badge là thông tin chi tiết của badge
	Badge *Badge `json:"badge" bson:"badge"`

	// AwardedAt là thời điểm user đạt được badge
	AwardedAt time.Time `json:"awarded_at" bson:"awarded_at"`

	// RelatedEntityID là ID của entity liên quan đến việc đạt badge (ví dụ: ID của project hoàn thành)
	RelatedEntityID string `json:"related_entity_id,omitempty" bson:"related_entity_id,omitempty"`
}
