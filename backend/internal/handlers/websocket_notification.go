package handlers

import (
	"encoding/json"
	// "fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/iknizzz1807/SkillForge/internal/integrations"

	// "github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

// WebSocketHandler xử lý các kết nối WebSocket
type WebSocketNotificationHandler struct {
	upgrader       websocket.Upgrader
	realtimeClient *integrations.RealtimeClient
	notifyService  *services.NotificationService
}

func NewWebSocketNotificationHandler(
	realtimeClient *integrations.RealtimeClient,
	notifyService *services.NotificationService,
) *WebSocketNotificationHandler {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return &WebSocketNotificationHandler{
		upgrader:       upgrader,
		realtimeClient: realtimeClient,
		notifyService:  notifyService,
	}
}

// GET /ws/notifi
func (h *WebSocketNotificationHandler) HandleNotificationConnection(c *gin.Context) {
	userID := c.Param("userID")
	room := "notification"
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user ID"})
		return
	}

	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error upgrading to WebSocket: %v", err)
		return
	}

	h.realtimeClient.AddConnection(room, userID, conn)

	log.Printf("New WebSocket connection established for user: %s", userID)

	defer func() {
		h.realtimeClient.RemoveConnection(room, userID)
		log.Printf("WebSocket connection closed for user: %s", userID)
	}()

	h.handleMessages(userID, conn)
}

func (h *WebSocketNotificationHandler) handleMessages(userID string, conn *websocket.Conn) {
	// Trả về danh sách thông báo lần đầu tiên từ khi bắt đầu kết nối (initial load)
	allNotifications, err := h.notifyService.GetUserNotifications(userID)
	if err != nil {
		log.Printf("Error fetching notifications: %v", err)
		return
	}

	response, err := json.Marshal(allNotifications)
	if err != nil {
		log.Printf("Error marshalling notifications: %v", err)
		return
	}

	err = conn.WriteMessage(websocket.TextMessage, response)
	if err != nil {
		log.Printf("Error sending initial notifications: %v", err)
		return
	}

	for {

	}
}
