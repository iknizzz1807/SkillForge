package services

// Nhiệm vụ: Xử lý logic tin nhắn (chat, video call)
// Liên kết:
// - Dùng internal/repositories/message_repository.go để lưu tin nhắn
// - Dùng integrations/realtime.go để gửi thông báo realtime
// - Dùng integrations/webrtc.go để khởi tạo video call
// Vai trò trong flow:
// - Được gọi từ handlers/messages.go để quản lý tin nhắn

import (
	"errors"
	"time"

	"context"

	"github.com/iknizzz1807/SkillForge/internal/integrations"
	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/iknizzz1807/SkillForge/internal/utils"
)

type MessageService struct {
	// db để truy cập MongoDB database
	db *mongo.Database
	// realtimeClient để gửi tin nhắn qua WebSocket
	realtimeClient *integrations.RealtimeClient
	// webrtcClient để xử lý video call
	webrtcClient *integrations.WebRTCClient
}

// NewMessageService khởi tạo MessageService với dependency
// Input: db (*mongo.Database), realtimeClient (*integrations.RealtimeClient), webrtcClient (*integrations.WebRTCClient)
// Return: *MessageService - con trỏ đến MessageService
func NewMessageService(db *mongo.Database, realtimeClient *integrations.RealtimeClient, webrtcClient *integrations.WebRTCClient) *MessageService {
	return &MessageService{db, realtimeClient, webrtcClient}
}

// SendMessage gửi tin nhắn
// Input: senderID (string), receiverID (string), content (string)
// Return: *models.Message (tin nhắn vừa gửi), error (nếu có lỗi)
func (s *MessageService) SendMessage(senderID, receiverID, content string) (*models.Message, error) {
	// Kiểm tra input hợp lệ
	if senderID == "" || receiverID == "" || content == "" {
		return nil, errors.New("invalid message data")
	}

	// Tạo tin nhắn mới
	message := &models.Message{
		ID:         utils.GenerateUUID(),
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
		CreatedAt:  time.Now(),
	}

	// Lưu tin nhắn vào database
	messageRepo := repositories.NewMessageRepository(s.db)
	err := messageRepo.InsertMessage(context.Background(), message)
	if err != nil {
		return nil, err
	}

	// Gửi tin nhắn qua WebSocket
	s.realtimeClient.SendMessage(receiverID, content)

	// Trả về tin nhắn
	return message, nil
}

// GetMessages lấy danh sách tin nhắn giữa 2 người
// Input: userID (string), receiverID (string)
// Return: []*models.Message (danh sách tin nhắn), error (nếu có lỗi)
func (s *MessageService) GetMessages(userID, receiverID string) ([]*models.Message, error) {
	// Kiểm tra input hợp lệ
	if userID == "" || receiverID == "" {
		return nil, errors.New("invalid user IDs")
	}

	// Tạo repository và lấy danh sách tin nhắn
	messageRepo := repositories.NewMessageRepository(s.db)
	messages, err := messageRepo.FindMessagesBetweenUsers(context.Background(), userID, receiverID)
	if err != nil {
		return nil, err
	}

	// Trả về danh sách tin nhắn
	return messages, nil
}
