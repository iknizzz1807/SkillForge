package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/iknizzz1807/SkillForge/internal/integrations"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

type WebSocketChatHandler struct {
	upgrader       websocket.Upgrader
	realtimeClient *integrations.RealtimeClient
	chatService *services.ChatService
}

func NewWebSocketChatHandler(chatService *services.ChatService, realtimeClient *integrations.RealtimeClient) *WebSocketChatHandler {
	return &WebSocketChatHandler{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		realtimeClient: realtimeClient,
		chatService: chatService,
	}
}

func (wsh *WebSocketChatHandler) HandleConnection(c *gin.Context) {
	userID := c.GetString("userID")
	projectID := c.Param("projectID")
	conn, err := wsh.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	room := "message" + projectID
	wsh.realtimeClient.AddConnection(room, userID, conn)

	log.Printf("New WebSocket connection established for room: %s", room)

	defer func() {
		wsh.realtimeClient.RemoveConnection(room, userID)
		log.Printf("WebSocket connection closed for room: %s", room)
	}()

	wsh.handleMessages(room, userID, conn)
}

func (wsh *WebSocketChatHandler) handleMessages(room, userID string, conn *websocket.Conn) {
	for {
		_, rawMessage, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading from WebSocket: %v", err)
			break
		}

		var Content string
		err = json.Unmarshal(rawMessage, &Content)
		if err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			continue
		}

		Message, err := wsh.chatService.InsertMessage(context.Background(), room, userID, Content)
		if err != nil {
			log.Printf("Error inserting message: %v", err)
			continue
		}

		response, err := json.Marshal(Message)
		if err != nil {
			log.Printf("Error marshaling message: %v", err)
			continue
		}

		err = conn.WriteMessage(websocket.TextMessage, response)
		if err != nil {
			log.Printf("Error sending message: %v", err)
			break
		}
	}
}
