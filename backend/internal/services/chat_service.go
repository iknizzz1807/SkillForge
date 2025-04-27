package services

import (
	"context"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/utils"
)

type ChatService struct {
	chatRepo *repositories.ChatRepository
}

func NewChatService(chatRepo *repositories.ChatRepository) *ChatService {
	return &ChatService{
		chatRepo: chatRepo,
	}
}

func (c *ChatService) GetGroups(ctx context.Context, userID string) ([]*models.Group, error) {
	return c.chatRepo.GetGroups(ctx, userID)
}

func (c *ChatService) GetGroupInfo(groupID string) ([]*models.Message, []*models.User, error) {
	// return messages and members of the group
	Messages, err := c.chatRepo.GetGroupMessages(groupID)
	if err != nil {
		return nil, nil, err
	}
	Members, err := c.chatRepo.GetGroupMembers(groupID)
	if err != nil {
		return nil, nil, err
	}
	return Messages, Members, nil
}

func (c *ChatService) InsertMessage(ctx context.Context, userID, projectID, content string) (*models.Message, error) {
	message := &models.Message{
		ID:       utils.GenerateUUID(),
		SenderID: userID,
		GroupID:  projectID,
		Content:  content,
	}
	return c.chatRepo.InsertMessage(ctx, message)
}
