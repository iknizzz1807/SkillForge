package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/iknizzz1807/SkillForge/internal/integrations"
	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

// WebSocketHandler xử lý các kết nối WebSocket
type WebSocketTaskHandler struct {
	upgrader       websocket.Upgrader
	realtimeClient *integrations.RealtimeClient
	taskService    *services.TaskService
}

// NewWebSocketHandler tạo một handler mới cho WebSocket
func NewWebSocketTaskHandler(
	realtimeClient *integrations.RealtimeClient,
	taskService *services.TaskService,
) *WebSocketTaskHandler {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return &WebSocketTaskHandler{
		upgrader:       upgrader,
		realtimeClient: realtimeClient,
		taskService:    taskService,
	}
}

// GET /ws/task/:projectID
// Input: projectID (string) từ URL
// Return: Trả về JSON danh sách các task của projectID theo dạng task_model hoặc thông báo lỗi nếu có
// HandleConnection xử lý kết nối WebSocket từ client
func (h *WebSocketTaskHandler) HandleConnection(c *gin.Context) {
	// Cái này là đang test, uncommnent nếu dùng thực sự trong app
	// Xác định userID - có thể từ token auth hoặc query param
	userID := c.Query("userId")
	if userID == "" {
		// Fallback lấy userID từ JWT token
		userID = c.GetString("userID")
		if userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user identification"})
			return
		}
	}

	projectID := c.Param("projectID")
	if projectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing project ID"})
		return
	}
	// userID := "Thong dep trai 123"
	room := projectID

	// Nâng cấp connection HTTP lên WebSocket
	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error upgrading to websocket: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not establish WebSocket connection"})
		return
	}

	// Đăng ký kết nối cho user
	h.realtimeClient.AddConnection(room, userID, conn)

	// Log kết nối mới
	log.Printf("New WebSocket connection established for user: %s", userID)

	// Xử lý đóng kết nối
	defer func() {
		conn.Close()
		h.realtimeClient.RemoveConnection(room, userID)
		log.Printf("WebSocket connection closed for user: %s", userID)
	}()

	// Xử lý tin nhắn từ client
	h.handleMessages(room, userID, projectID, conn)
}

type TaskMessage struct {
	Type    string          `json:"type"`
	Content json.RawMessage `json:"content"`
}

// handleMessages xử lý các message từ client
func (h *WebSocketTaskHandler) handleMessages(room, userID, projectID string, conn *websocket.Conn) {
	// initial message - gửi danh sách task hiện tại cho client
	tasks, err := h.taskService.GetTasksByProjectID(projectID)
	if err != nil {
		log.Printf("Error fetching tasks: %v", err)
		sendErrorMessage(conn, "Failed to fetch tasks")
		return
	}
	activities, err := h.taskService.GetActivityByProjectID(userID, projectID)
	if err != nil {
		log.Printf("Error fetching activities: %v", err)
		sendErrorMessage(conn, "Failed to fetch activities")
		return
	}
	response := map[string]interface{}{
		"tasks":    tasks,
		"activities": activities,
	}
	respBytes, _ := json.Marshal(response)
	h.realtimeClient.Broadcast(room, respBytes)

	for {
		// Đọc message từ WebSocket
		_, rawMessage, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading from WebSocket: %v", err)
			break
		}

		var message TaskMessage
		err = json.Unmarshal(rawMessage, &message)
		if err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			sendErrorMessage(conn, "Failed to unmarshal message")
			return
		}

		if message.Type == "create" {
			// Xử lý tạo task mới
			var req *models.TaskInput
			err = json.Unmarshal(message.Content, &req)
			if err != nil {
				log.Printf("Error unmarshalling task input: %v", err)
				sendErrorMessage(conn, "Failed to unmarshal task input")
				return
			}
			_, err = h.taskService.CreateTasks(projectID, userID, []models.TaskInput{*req})
			if err != nil {
				log.Printf("Error creating task: %v", err)
				sendErrorMessage(conn, "Failed to create task")
				return
			}
		}
		if message.Type != "update" {
			var req *models.TaskUpdate
			err = json.Unmarshal(message.Content, &req)
			if err != nil {
				log.Printf("Error unmarshalling task update: %v", err)
				sendErrorMessage(conn, "Failed to unmarshal task update")
				return
			}
			_, err = h.taskService.UpdateTask(req.TaskID, userID, req)
			if err != nil {
				log.Printf("Error updating task: %v", err)
				sendErrorMessage(conn, "Failed to update task")
				return
			}
		}
		if message.Type == "delete" {
			// Xử lý xóa task
			var req string
			err = json.Unmarshal(message.Content, &req)
			if err != nil {
				log.Printf("Error unmarshalling task ID: %v", err)
				sendErrorMessage(conn, "Failed to unmarshal task ID")
				return
			}
			err := h.taskService.DeleteTask(req, userID)
			if err != nil {
				log.Printf("Error deleting task: %v", err)
				sendErrorMessage(conn, "Failed to delete task")
				return
			}
		}
		tasks, err := h.taskService.GetTasksByProjectID(projectID)
		if err != nil {
			log.Printf("Error fetching tasks: %v", err)
			sendErrorMessage(conn, "Failed to fetch tasks")
			return
		}
		activities, err := h.taskService.GetActivityByProjectID(userID, projectID)
		if err != nil {
			log.Printf("Error fetching activities: %v", err)
			sendErrorMessage(conn, "Failed to fetch activities")
			return
		}
		response := map[string]interface{}{
			"tasks":    tasks,
			"activities": activities,
		}
		respBytes, _ := json.Marshal(response)
		h.realtimeClient.Broadcast(room, respBytes)
	}
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
