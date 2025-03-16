package handlers

// Nhiệm vụ: Xử lý các endpoint quản lý tin nhắn (chat, video call)
// Liên kết:
// - Gọi internal/services/message_service.go để xử lý logic
// Vai trò trong flow:
// - Được gọi từ /api/messages để gửi/lấy tin nhắn hoặc khởi tạo video call

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

type MessageHandler struct {
	// messageService xử lý logic liên quan đến tin nhắn
	messageService *services.MessageService
}

// NewMessageHandler khởi tạo handler với messageService
// Input: messageService (*services.MessageService)
// Return: *MessageHandler - con trỏ đến MessageHandler
func NewMessageHandler(messageService *services.MessageService) *MessageHandler {
	return &MessageHandler{messageService}
}

// SendMessage xử lý endpoint POST /api/messages
// Return: Trả về JSON tin nhắn vừa gửi hoặc lỗi
func (h *MessageHandler) SendMessage(c *gin.Context) {
	var req struct {
		ReceiverID string `json:"receiver_id" binding:"required"`
		Content    string `json:"content" binding:"required"`
	}

	// Parse và validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Lấy userID từ context
	senderID := c.GetString("userID")

	// Gọi service để gửi tin nhắn
	message, err := h.messageService.SendMessage(senderID, req.ReceiverID, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về tin nhắn vừa gửi
	c.JSON(http.StatusCreated, message)
}

// GetMessages xử lý endpoint GET /api/messages/:id
// Return: Trả về JSON danh sách tin nhắn hoặc lỗi
func (h *MessageHandler) GetMessages(c *gin.Context) {
	userID := c.GetString("userID")
	receiverID := c.Param("id")

	// Gọi service để lấy danh sách tin nhắn
	messages, err := h.messageService.GetMessages(userID, receiverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về danh sách tin nhắn
	c.JSON(http.StatusOK, messages)
}
