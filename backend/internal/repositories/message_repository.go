package repositories

// Nhiệm vụ: Tương tác với collection "messages" trong MongoDB (CRUD operations)
// Liên kết:
// - Dùng go.mongodb.org/mongo-driver/mongo để truy vấn database
// - Trả về models từ github.com/iknizzz1807/SkillForge/internal/models
// Vai trò trong flow:
// - Được gọi từ services/message_service.go để lưu/lấy tin nhắn

import (
	"context"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepository struct {
	// collection để truy cập collection "messages" trong MongoDB
	collection *mongo.Collection
}

// NewMessageRepository khởi tạo MessageRepository với database
// Input: db (*mongo.Database)
// Return: *MessageRepository - con trỏ đến MessageRepository
func NewMessageRepository(db *mongo.Database) *MessageRepository {
	return &MessageRepository{
		collection: db.Collection("messages"),
	}
}

// InsertMessage thêm tin nhắn mới vào collection "messages"
// Input: ctx (context.Context), message (*models.Message)
// Return: error (nếu có lỗi)
func (r *MessageRepository) InsertMessage(ctx context.Context, message *models.Message) error {
	// Chèn message vào collection
	_, err := r.collection.InsertOne(ctx, message)
	if err != nil {
		return err
	}
	// Trả về nil nếu thành công
	return nil
}

// FindMessagesBetweenUsers lấy danh sách tin nhắn giữa 2 người
// Input: ctx (context.Context), userID (string), receiverID (string)
// Return: []*models.Message (danh sách tin nhắn), error (nếu có lỗi)
func (r *MessageRepository) FindMessagesBetweenUsers(ctx context.Context, userID, receiverID string) ([]*models.Message, error) {
	// Tạo filter để tìm tin nhắn giữa 2 user (theo sender và receiver)
	filter := bson.M{
		"$or": []bson.M{
			{"sender_id": userID, "receiver_id": receiverID},
			{"sender_id": receiverID, "receiver_id": userID},
		},
	}

	// Tạo cursor để lấy tất cả document
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Tạo slice để lưu danh sách tin nhắn
	var messages []*models.Message
	// Duyệt qua cursor và decode từng document
	for cursor.Next(ctx) {
		var message models.Message
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}

	// Trả về danh sách tin nhắn
	return messages, nil
}
