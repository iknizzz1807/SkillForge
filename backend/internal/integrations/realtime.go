package integrations

// Nhiệm vụ: Tích hợp thông báo realtime qua WebSocket
// Liên kết:
// - Dùng github.com/gorilla/websocket để xử lý WebSocket
// - Được gọi từ services/message_service.go để gửi tin nhắn realtime
// Vai trò trong flow:
// - Gửi thông báo tức thời (ví dụ: tin nhắn mới) tới user qua WebSocket

import (
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/iknizzz1807/SkillForge/internal/models"
)

type clientConn struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

type RealtimeClient struct {
	mu sync.RWMutex
	// connections lưu trữ danh sách kết nối WebSocket theo userID
	connections map[string]map[string]*clientConn
}

// NewRealtimeClient khởi tạo RealtimeClient
// Return: *RealtimeClient - con trỏ đến RealtimeClient
func NewRealtimeClient() *RealtimeClient {
	return &RealtimeClient{
		connections: make(map[string]map[string]*clientConn),
	}
}

// AddConnection thêm kết nối WebSocket cho user
// Input: userID (string), conn (*websocket.Conn)
// Return: None
func (c *RealtimeClient) AddConnection(room, userID string, conn *websocket.Conn) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.connections[room] == nil {
		c.connections[room] = make(map[string]*clientConn)
	}

	c.connections[room][userID] = &clientConn{conn: conn}
}

// RemoveConnection xóa kết nối WebSocket của user
// Input: userID (string)
// Return: None
func (c *RealtimeClient) RemoveConnection(room, userID string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Kiểm tra map tồn tại trước khi xóa
	if c.connections[room] != nil {
		delete(c.connections[room], userID)

		// Nếu room trống thì xóa luôn room
		if len(c.connections[room]) == 0 {
			delete(c.connections, room)
		}
	}
}

// SendNotification gửi thông báo qua WebSocket tới user
// Input: userID (string), message (string)
// Return: error (nếu có lỗi)
func (c *RealtimeClient) SendNotification(userID string, notifications []*models.Notification) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Kiểm tra room và userID tồn tại
	if c.connections["notification"] == nil {
		return errors.New("notification room not found")
	}

	cc, exists := c.connections["notification"][userID]
	if !exists {
		return errors.New("user not connected")
	}

	response, _ := json.Marshal(notifications)

	cc.mu.Lock()
	defer cc.mu.Unlock()
	cc.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	err := cc.conn.WriteMessage(websocket.TextMessage, []byte(response))
	if err != nil {
		return err
	}

	return nil
}

func (c *RealtimeClient) SendMessage(userID, message string) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.connections["message"] == nil {
		return errors.New("message room not found")
	}

	cc, exists := c.connections["message"][userID]
	if !exists {
		return errors.New("user not connected")
	}

	cc.mu.Lock()
	defer cc.mu.Unlock()
	cc.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	err := cc.conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return err
	}

	return nil
}

func (c *RealtimeClient) Broadcast(room string, message []byte) error {
	c.mu.RLock()

	if c.connections[room] == nil {
		c.mu.RUnlock()
		return errors.New("room not found")
	}

	type staleConn struct {
		room, userID string
	}
	var stale []staleConn

	for userID, cc := range c.connections[room] {
		cc.mu.Lock()
		cc.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		err := cc.conn.WriteMessage(websocket.TextMessage, message)
		cc.mu.Unlock()
		if err != nil {
			stale = append(stale, staleConn{room, userID})
		}
	}
	c.mu.RUnlock()

	for _, s := range stale {
		c.RemoveConnection(s.room, s.userID)
	}

	return nil
}
