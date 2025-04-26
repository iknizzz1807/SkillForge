package services

// Nhiệm vụ: Xử lý logic gửi thông báo (email, realtime)
// Liên kết:
// - Dùng integrations/email.go để gửi email
// - Dùng integrations/realtime.go để gửi thông báo qua WebSocket
// Vai trò trong flow:
// - Được gọi từ các service khác (project, task, v.v.) để gửi thông báo

import (
	"errors"
	"time"
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/integrations"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
)

type NotificationService struct {
	// emailClient để gửi email
	emailClient *integrations.EmailClient
	// realtimeClient để gửi thông báo qua WebSocket
	realtimeClient *integrations.RealtimeClient
	// db để truy cập MongoDB database
	db *mongo.Database
}

// NewNotificationService khởi tạo NotificationService với dependency
// Input: emailClient (*integrations.EmailClient), realtimeClient (*integrations.RealtimeClient)
// Return: *NotificationService - con trỏ đến NotificationService
func NewNotificationService(emailClient *integrations.EmailClient, realtimeClient *integrations.RealtimeClient, db *mongo.Database) *NotificationService {
	return &NotificationService{emailClient, realtimeClient, db}
}

// SendEmail gửi email thông báo
// Input: to (string), subject (string), body (string)
// Return: error (nếu có lỗi)
func (s *NotificationService) SendEmail(to, subject, body string) error {
	// Kiểm tra input hợp lệ
	if to == "" || subject == "" || body == "" {
		return errors.New("invalid email data")
	}

	// Gọi emailClient để gửi email
	err := s.emailClient.SendEmail(to, subject, body)
	if err != nil {
		return err
	}

	// Trả về nil nếu thành công
	return nil
}

// SendRealtime gửi thông báo realtime qua WebSocket
// Input: userID gửi đến (string), message (string)
// Return: error (nếu có lỗi)
func (s *NotificationService) SendRealtime(userID, content, Type string) error {
	// Kiểm tra input hợp lệ
	
	if userID == "" || content == "" || Type == "" {
		return errors.New("invalid realtime data")
	}

	// Tạo notification
	notification := &models.Notification{
		ID:        utils.GenerateUUID(),
		ToUserID:  userID,
		Content:   content,
		Type:      Type,
		CreatedAt: time.Now(),
	}

	// Lưu notification vào database
	notificationRepo := repositories.NewNotificationRepository(s.db)
	err := notificationRepo.InsertNotification(context.Background(), notification)
	if err != nil {
		return err
	}

	notifications, err := s.GetUserNotifications(userID)

	// Gọi realtimeClient để gửi thông báo
	err = s.realtimeClient.SendNotification(userID, notifications)
	if err != nil {
		return err
	}

	// Trả về nil nếu thành công
	return nil
}

func (s *NotificationService) GetUserNotifications(userID string) ([]*models.Notification, error) {
	notificationRepo := repositories.NewNotificationRepository(s.db)
	notifications, err := notificationRepo.GetUserNotifications(context.Background(), userID)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}
