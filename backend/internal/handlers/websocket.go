package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/iknizzz1807/SkillForge/internal/integrations"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

// WebSocketHandler xử lý các kết nối WebSocket
type WebSocketHandler struct {
	upgrader       websocket.Upgrader
	realtimeClient *integrations.RealtimeClient
	messageService *services.MessageService
	notifyService  *services.NotificationService
	taskService    *services.TaskService
	projectService *services.ProjectService
}

// WebSocketMessage định nghĩa cấu trúc của message nhận từ client
type WebSocketMessage struct {
	Type    string          `json:"type"`
	UserID  string          `json:"userId,omitempty"`
	Target  string          `json:"target,omitempty"`
	Content json.RawMessage `json:"content,omitempty"`
}

// NewWebSocketHandler tạo một handler mới cho WebSocket
func NewWebSocketHandler(
	realtimeClient *integrations.RealtimeClient,
	messageService *services.MessageService,
	notifyService *services.NotificationService,
	taskService *services.TaskService,
	projectService *services.ProjectService,
) *WebSocketHandler {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return &WebSocketHandler{
		upgrader:       upgrader,
		realtimeClient: realtimeClient,
		messageService: messageService,
		notifyService:  notifyService,
		taskService:    taskService,
		projectService: projectService,
	}
}

// HandleConnection xử lý kết nối WebSocket từ client
func (h *WebSocketHandler) HandleConnection(c *gin.Context) {
	// Cái này là đang test, uncommnent nếu dùng thực sự trong app
	// Xác định userID - có thể từ token auth hoặc query param
	// userID := c.Query("userId")
	// if userID == "" {
	// 	// Fallback lấy userID từ JWT token
	// 	userID = c.GetString("userID")
	// 	if userID == "" {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user identification"})
	// 		return
	// 	}
	// }

	userID := "Thong dep trai 123"

	// Nâng cấp connection HTTP lên WebSocket
	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error upgrading to websocket: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not establish WebSocket connection"})
		return
	}

	// Đăng ký kết nối cho user
	h.realtimeClient.AddConnection(userID, conn)

	// Log kết nối mới
	log.Printf("New WebSocket connection established for user: %s", userID)

	// Xử lý đóng kết nối
	defer func() {
		conn.Close()
		h.realtimeClient.RemoveConnection(userID)
		log.Printf("WebSocket connection closed for user: %s", userID)
	}()

	// Xử lý tin nhắn từ client
	h.handleMessages(userID, conn)
}

// handleMessages xử lý các message từ client
func (h *WebSocketHandler) handleMessages(userID string, conn *websocket.Conn) {
	for {
		// Đọc message từ WebSocket
		_, rawMessage, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading from WebSocket: %v", err)
			break
		}

		// Parse message
		var message WebSocketMessage
		if err := json.Unmarshal(rawMessage, &message); err != nil {
			log.Printf("Error parsing WebSocket message: %v", err)
			sendErrorMessage(conn, "Invalid message format")
			continue
		}

		// Xử lý message dựa vào loại
		switch message.Type {
		case "chat":
			h.handleChatMessage(conn, message)
		case "notification_ack":
			h.handleNotificationAck(message)
		case "task_update":
			h.handleTaskUpdate(conn, message)
		case "board_update":
			h.handleBoardUpdate(conn, message)
		case "ping":
			// Xử lý tin nhắn ping để giữ kết nối
			conn.WriteMessage(websocket.TextMessage, []byte(`{"type":"pong"}`))
		default:
			fmt.Printf("Received message from user %s: %s\n", userID, string(rawMessage))
			sendErrorMessage(conn, "Unknown message type")
		}
	}
}

// Các hàm xử lý message theo loại
func (h *WebSocketHandler) handleChatMessage(conn *websocket.Conn, message WebSocketMessage) {
	// TODO: Implement chat message handling
	// 1. Lưu tin nhắn vào database
	// 2. Gửi tin nhắn tới người nhận nếu online

}

func (h *WebSocketHandler) handleNotificationAck(message WebSocketMessage) {
	// TODO: Xác nhận đã đọc thông báo
}

func (h *WebSocketHandler) handleTaskUpdate(conn *websocket.Conn, message WebSocketMessage) {
	// TODO: Xử lý cập nhật task
	// 1. Cập nhật task trong database
	// 2. Thông báo cho các client khác về thay đổi
}

func (h *WebSocketHandler) handleBoardUpdate(conn *websocket.Conn, message WebSocketMessage) {
	// TODO: Xử lý cập nhật kanban board
	// 1. Cập nhật trạng thái task trong database
	// 2. Thông báo cho các client khác về thay đổi
}

// Helper function gửi thông báo lỗi
func sendErrorMessage(conn *websocket.Conn, errorMsg string) {
	response := map[string]string{
		"type":  "error",
		"error": errorMsg,
	}
	respBytes, _ := json.Marshal(response)
	conn.WriteMessage(websocket.TextMessage, respBytes)
}
