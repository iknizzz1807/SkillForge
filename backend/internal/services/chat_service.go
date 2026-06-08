package services

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/utils"
)

const ChatFilesPath = "./storage/chat_files"

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

func (c *ChatService) GetGroupInfo(groupID, userID string) ([]*models.Message, []*models.User, error) {
	hasAccess, err := c.chatRepo.CheckProjectAccess(groupID, userID)
	if err != nil {
		return nil, nil, err
	}
	if !hasAccess {
		return nil, nil, errors.New("access denied")
	}
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

func (c *ChatService) InsertMessage(ctx context.Context, userID, projectID, content, msgType, fileURL, fileName string) (*models.Message, error) {
	message := &models.Message{
		ID:       utils.GenerateUUID(),
		SenderID: userID,
		GroupID:  projectID,
		Content:  content,
		Type:     msgType,
		FileURL:  fileURL,
		FileName: fileName,
	}
	return c.chatRepo.InsertMessage(message)
}

func (s *ChatService) SaveChatFile(file multipart.File, header *multipart.FileHeader) (string, string, error) {
	if file == nil || header == nil {
		return "", "", errors.New("file data is missing")
	}

	ext := strings.ToLower(filepath.Ext(header.Filename))

	allowedExts := map[string]bool{
		".jpg": true, ".jpeg": true, ".png": true, ".gif": true,
		".pdf": true, ".doc": true, ".docx": true, ".txt": true,
		".zip": true, ".rar": true, ".csv": true, ".xlsx": true, ".xls": true,
	}
	if !allowedExts[ext] {
		return "", "", errors.New("invalid file type")
	}

	if header.Size > 10*1024*1024 {
		return "", "", errors.New("file too large (max 10MB)")
	}

	newFilename := utils.GenerateUUID() + ext

	err := os.MkdirAll(ChatFilesPath, os.ModePerm)
	if err != nil {
		return "", "", fmt.Errorf("could not create storage directory: %w", err)
	}

	filePath := filepath.Join(ChatFilesPath, newFilename)

	dst, err := os.Create(filePath)
	if err != nil {
		return "", "", fmt.Errorf("could not save file: %w", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		os.Remove(filePath)
		return "", "", fmt.Errorf("could not write file content: %w", err)
	}

	return newFilename, header.Filename, nil
}
