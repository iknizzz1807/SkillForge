package services

// Nhiệm vụ: Xử lý logic gửi thông báo (email, realtime)
// Liên kết:
// - Dùng integrations/email.go để gửi email
// - Dùng integrations/realtime.go để gửi thông báo qua WebSocket
// Vai trò trong flow:
// - Được gọi từ các service khác (project, task, v.v.) để gửi thông báo

import (
	"errors"

	"github.com/iknizzz1807/SkillForge/internal/integrations"
)

type NotificationService struct {
	// emailClient để gửi email
	emailClient *integrations.EmailClient
	// realtimeClient để gửi thông báo qua WebSocket
	realtimeClient *integrations.RealtimeClient
}

// NewNotificationService khởi tạo NotificationService với dependency
// Input: emailClient (*integrations.EmailClient), realtimeClient (*integrations.RealtimeClient)
// Return: *NotificationService - con trỏ đến NotificationService
func NewNotificationService(emailClient *integrations.EmailClient, realtimeClient *integrations.RealtimeClient) *NotificationService {
	return &NotificationService{emailClient, realtimeClient}
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
// Input: userID (string), message (string)
// Return: error (nếu có lỗi)
func (s *NotificationService) SendRealtime(userID, message string) error {
	// Kiểm tra input hợp lệ
	if userID == "" || message == "" {
		return errors.New("invalid realtime data")
	}

	// Gọi realtimeClient để gửi thông báo
	err := s.realtimeClient.SendNotification(userID, message)
	if err != nil {
		return err
	}

	// Trả về nil nếu thành công
	return nil
}
