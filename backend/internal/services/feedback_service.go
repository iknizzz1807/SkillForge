package services

import (
	"context"
	"errors"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/utils"
)

type FeedbackService struct {
	feedbackRepo *repositories.FeedbackRepository
	projectRepo  *repositories.ProjectRepository
	userRepo     *repositories.UserRepository
}

func NewFeedbackService(feedbackRepo *repositories.FeedbackRepository, projectRepo *repositories.ProjectRepository, userRepo *repositories.UserRepository) *FeedbackService {
	return &FeedbackService{
		feedbackRepo: feedbackRepo,
		projectRepo:  projectRepo,
		userRepo:     userRepo,
	}
}

func (s *FeedbackService) CreateFeedback(projectID, fromID, toID, feedbackType, content string, rating int) (*models.Feedback, error) {
	if rating < 0 || rating > 5 {
		return nil, errors.New("rating must be between 0 and 5")
	}
	if feedbackType != "business" && feedbackType != "student" {
		return nil, errors.New("invalid feedback type")
	}
	if fromID == toID {
		return nil, errors.New("cannot give feedback to yourself")
	}

	ctx := context.Background()
	_, err := s.projectRepo.FindProjectByID(ctx, projectID)
	if err != nil {
		return nil, errors.New("project not found")
	}
	user, err2 := s.userRepo.FindUserByID(ctx, fromID)
	if err2 != nil || user == nil {
		return nil, errors.New("user not found")
	}

	feedback := &models.Feedback{
		ID:        utils.GenerateUUID(),
		ProjectID: projectID,
		FromID:    fromID,
		ToID:      toID,
		Type:      feedbackType,
		Rating:    rating,
		Content:   content,
		CreatedAt: time.Now(),
	}

	err = s.feedbackRepo.CreateFeedback(context.Background(), feedback)
	if err != nil {
		return nil, err
	}

	return feedback, nil
}

func (s *FeedbackService) GetUserFeedbacks(userID string, feedbackType string) ([]*models.Feedback, error) {
	if feedbackType != "business" && feedbackType != "student" {
		return nil, errors.New("invalid feedback type")
	}
	return s.feedbackRepo.GetFeedbacksForUser(context.Background(), userID, feedbackType)
}

func (s *FeedbackService) GetAverageRating(userID string) (float64, error) {
	if userID == "" {
		return 0, errors.New("user ID is required")
	}
	return s.feedbackRepo.GetAverageRating(context.Background(), userID)
}
