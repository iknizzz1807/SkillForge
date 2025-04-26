package repositories

import (
	"context"
	"errors"
	// "time"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NotificationRepository struct {
	collection *mongo.Collection
}

// NewNotificationRepository khởi tạo NotificationRepository với database
func NewNotificationRepository(db *mongo.Database) *NotificationRepository {
	return &NotificationRepository{
		collection: db.Collection("notifications"),
	}
}

// InsertNotification thêm notification mới vào collection
func (r *NotificationRepository) InsertNotification(ctx context.Context, notification *models.Notification) error {
	if ctx == nil {
		ctx = context.Background()
	}

	_, err := r.collection.InsertOne(ctx, notification)
	if err != nil {
		return err
	}

	return nil
}

// GetUserNotifications lấy tất cả thông báo của một user
func (r *NotificationRepository) GetUserNotifications(ctx context.Context, userID string) ([]*models.Notification, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	// Tạo options để sort theo thời gian tạo giảm dần (mới nhất lên đầu)
	opts := options.Find().SetSort(bson.M{"created_at": -1})

	cursor, err := r.collection.Find(ctx, bson.M{"to": userID}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var notifications []*models.Notification
	if err = cursor.All(ctx, &notifications); err != nil {
		return nil, err
	}

	return notifications, nil
}

// MarkAsRead đánh dấu notification là đã đọc
// func (r *NotificationRepository) MarkAsRead(ctx context.Context, notificationID string) error {
// 	if ctx == nil {
// 		ctx = context.Background()
// 	}

// 	result, err := r.collection.UpdateOne(
// 		ctx,
// 		bson.M{"_id": notificationID},
// 		bson.M{
// 			"$set": bson.M{
// 				"read":       true,
// 				"updated_at": time.Now(),
// 			},
// 		},
// 	)

// 	if err != nil {
// 		return err
// 	}

// 	if result.MatchedCount == 0 {
// 		return errors.New("notification not found")
// 	}

// 	return nil
// }

// DeleteNotification xóa một notification
func (r *NotificationRepository) DeleteNotification(ctx context.Context, notificationID string) error {
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": notificationID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("notification not found")
	}

	return nil
}
