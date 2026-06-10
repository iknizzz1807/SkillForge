package services

import (
	"context"
	"errors"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type FeedbackService struct {
	feedbackRepo *repositories.FeedbackRepository
	projectRepo  *repositories.ProjectRepository
	userRepo     *repositories.UserRepository
	db           *mongo.Database
}

func NewFeedbackService(feedbackRepo *repositories.FeedbackRepository, projectRepo *repositories.ProjectRepository, userRepo *repositories.UserRepository, db ...*mongo.Database) *FeedbackService {
	var database *mongo.Database
	if len(db) > 0 {
		database = db[0]
	}
	return &FeedbackService{
		feedbackRepo: feedbackRepo,
		projectRepo:  projectRepo,
		userRepo:     userRepo,
		db:           database,
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
	project, err := s.projectRepo.FindProjectByID(ctx, projectID)
	if err != nil {
		return nil, errors.New("project not found")
	}
	fromUser, err2 := s.userRepo.FindUserByID(ctx, fromID)
	if err2 != nil || fromUser == nil {
		return nil, errors.New("sender not found")
	}
	toUser, err2 := s.userRepo.FindUserByID(ctx, toID)
	if err2 != nil || toUser == nil {
		return nil, errors.New("recipient not found")
	}

	if s.db != nil {
		projectStudentRepo := repositories.NewProjectStudentRepository(s.db)
		fromHasAccess := project.CreatedByID == fromID
		if !fromHasAccess {
			membership, err := projectStudentRepo.FindByProjectAndStudent(ctx, projectID, fromID)
			if err != nil {
				return nil, err
			}
			fromHasAccess = membership != nil
		}
		toHasAccess := project.CreatedByID == toID
		if !toHasAccess {
			membership, err := projectStudentRepo.FindByProjectAndStudent(ctx, projectID, toID)
			if err != nil {
				return nil, err
			}
			toHasAccess = membership != nil
		}
		if !fromHasAccess || !toHasAccess {
			return nil, errors.New("feedback users must belong to the project")
		}
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
