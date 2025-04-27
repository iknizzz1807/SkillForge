package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iknizzz1807/SkillForge/internal/services"
)

type ChatHandler struct {
	chatService *services.ChatService
}

func NewChatHandler(chatService *services.ChatService) *ChatHandler {
	return &ChatHandler{
		chatService: chatService,
	}
}

func (ch *ChatHandler) GetGroups(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}
	groups, err := ch.chatService.GetGroups(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, groups)
}

func (ch *ChatHandler) GetGroupInfo(c *gin.Context) {
	groupID := c.Param("group_id")
	if groupID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Group ID is required"})
		return
	}
	Messages, members, err := ch.chatService.GetGroupInfo(c, groupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	info := map[string]interface{}{
		"messages": Messages,
		"members":  members,
	}
	c.JSON(http.StatusOK, info)
}
