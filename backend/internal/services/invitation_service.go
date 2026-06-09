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

type InvitationService struct {
	db                  *mongo.Database
	badgeService        *BadgeService
	gamificationService *GamificationService
}

func NewInvitationService(db *mongo.Database, badgeService *BadgeService, gamificationService *GamificationService) *InvitationService {
	return &InvitationService{
		db:                  db,
		badgeService:        badgeService,
		gamificationService: gamificationService,
	}
}

func (s *InvitationService) CreateInvitation(studentID, projectID, businessID string) (*models.Invitation, error) {
	if studentID == "" || projectID == "" || businessID == "" {
		return nil, errors.New("studentID, projectID, and businessID are required")
	}

	ctx := context.Background()

	userRepo := repositories.NewUserRepository(s.db)
	student, err := userRepo.FindUserByID(ctx, studentID)
	if err != nil || student == nil {
		return nil, errors.New("student not found")
	}
	if student.Role != "student" {
		return nil, errors.New("user is not a student")
	}

	projectRepo := repositories.NewProjectRepository(s.db)
	project, err := projectRepo.FindProjectByID(ctx, projectID)
	if err != nil || project == nil {
		return nil, errors.New("project not found")
	}
	if project.CreatedByID != businessID {
		return nil, errors.New("you do not own this project")
	}

	if project.CurrentMember >= project.MaxMember {
		return nil, errors.New("project is full")
	}

	invitationRepo := repositories.NewInvitationRepository(s.db)
	existing, err := invitationRepo.FindByStudentAndProject(ctx, studentID, projectID)
	if err == nil && existing != nil {
		return nil, errors.New("invitation already sent to this student for this project")
	}

	invitation := &models.Invitation{
		ID:         utils.GenerateUUID(),
		ProjectID:  projectID,
		StudentID:  studentID,
		BusinessID: businessID,
		Status:     "pending",
		CreatedAt:  time.Now(),
	}

	if err := invitationRepo.Create(ctx, invitation); err != nil {
		return nil, err
	}

	return invitation, nil
}

func (s *InvitationService) GetInvitationsByStudent(studentID string) ([]*models.Invitation, error) {
	if studentID == "" {
		return nil, errors.New("student ID cannot be empty")
	}

	invitationRepo := repositories.NewInvitationRepository(s.db)
	invitations, err := invitationRepo.FindByStudentID(context.Background(), studentID)
	if err != nil {
		return nil, err
	}

	projectRepo := repositories.NewProjectRepository(s.db)
	userRepo := repositories.NewUserRepository(s.db)

	projectIDs := make([]string, 0, len(invitations))
	businessIDs := make([]string, 0, len(invitations))
	for _, inv := range invitations {
		projectIDs = append(projectIDs, inv.ProjectID)
		businessIDs = append(businessIDs, inv.BusinessID)
	}

	projects, _ := projectRepo.FindProjectsByIDs(context.Background(), projectIDs)
	projectMap := make(map[string]string, len(projects))
	for _, p := range projects {
		projectMap[p.ID] = p.Title
	}

	businesses, _ := userRepo.FindUsersByIDs(context.Background(), businessIDs)
	businessMap := make(map[string]string, len(businesses))
	for _, b := range businesses {
		businessMap[b.ID] = b.Name
	}

	for _, inv := range invitations {
		if title, ok := projectMap[inv.ProjectID]; ok {
			inv.ProjectTitle = title
		}
		if name, ok := businessMap[inv.BusinessID]; ok {
			inv.BusinessName = name
		}
	}

	return invitations, nil
}

func (s *InvitationService) RespondToInvitation(invitationID, studentID, status string) error {
	if invitationID == "" {
		return errors.New("invitation ID cannot be empty")
	}
	if status != "accepted" && status != "rejected" {
		return errors.New("status must be 'accepted' or 'rejected'")
	}

	ctx := context.Background()
	invitationRepo := repositories.NewInvitationRepository(s.db)

	invitation, err := invitationRepo.FindByID(ctx, invitationID)
	if err != nil {
		return err
	}
	if invitation.StudentID != studentID {
		return errors.New("this invitation is not for you")
	}
	if invitation.Status != "pending" {
		return errors.New("invitation has already been responded to")
	}

	if err := invitationRepo.UpdateStatus(ctx, invitationID, status); err != nil {
		return err
	}

	if status == "accepted" {
		projectService := NewProjectService(s.db, nil, nil, nil, nil)
		if err := projectService.AddStudentToProject(invitation.ProjectID, studentID); err != nil {
			return err
		}

		if s.badgeService != nil {
			go s.badgeService.CheckAndAwardProjectCompletionBadges(studentID, invitation.ProjectID)
		}
		if s.gamificationService != nil {
			go s.gamificationService.AddXP(studentID, 50)
		}
	}

	return nil
}
