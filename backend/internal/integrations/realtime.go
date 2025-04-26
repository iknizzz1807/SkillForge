package integrations

// Nhiệm vụ: Tích hợp thông báo realtime qua WebSocket
// Liên kết:
// - Dùng github.com/gorilla/websocket để xử lý WebSocket
// - Được gọi từ services/message_service.go để gửi tin nhắn realtime
// Vai trò trong flow:
// - Gửi thông báo tức thời (ví dụ: tin nhắn mới) tới user qua WebSocket

import (
	"errors"
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/iknizzz1807/SkillForge/internal/models"
)

type RealtimeClient struct {
	// connections lưu trữ danh sách kết nối WebSocket theo userID
	connections map[string]*websocket.Conn
}

// NewRealtimeClient khởi tạo RealtimeClient
// Return: *RealtimeClient - con trỏ đến RealtimeClient
func NewRealtimeClient() *RealtimeClient {
	return &RealtimeClient{
		connections: make(map[string]*websocket.Conn),
	}
}

// AddConnection thêm kết nối WebSocket cho user
// Input: userID (string), conn (*websocket.Conn)
// Return: None
func (c *RealtimeClient) AddConnection(userID string, conn *websocket.Conn) {
	// Lưu kết nối vào map
	c.connections[userID] = conn
}

// RemoveConnection xóa kết nối WebSocket của user
// Input: userID (string)
// Return: None
func (c *RealtimeClient) RemoveConnection(userID string) {
	// Xóa kết nối khỏi map
	delete(c.connections, userID)
}

// SendNotification gửi thông báo qua WebSocket tới user
// Input: userID (string), message (string)
// Return: error (nếu có lỗi)
func (c *RealtimeClient) SendNotification(userID string, notifications []*models.Notification) error {
	// Tìm kết nối của user
	conn, exists := c.connections[userID]
	if !exists {
		return errors.New("user not connected")
	}

	response, _ := json.Marshal(notifications)

	// Gửi message qua WebSocket
	err := conn.WriteMessage(websocket.TextMessage, []byte(response))
	if err != nil {
		return err
	}

	// Trả về nil nếu thành công
	return nil
}

func (c *RealtimeClient) SendMessage(userID, message string) error {
	// Tìm kết nối của user
	conn, exists := c.connections[userID]
	if !exists {
		return errors.New("user not connected")
	}

	// Gửi message qua WebSocket
	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return err
	}

	// Trả về nil nếu thành công
	return nil
}