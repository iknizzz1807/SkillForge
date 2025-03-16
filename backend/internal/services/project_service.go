package services

// Nhiệm vụ: Xử lý logic quản lý dự án (tạo, lấy danh sách, lấy chi tiết)
// Liên kết:
// - Dùng internal/repositories/project_repository.go để tương tác với database
// - Dùng notification_service.go để gửi thông báo
// - Dùng integrations/ai.go để gọi AI Matching
// - Dùng integrations/github.go để tạo repository
// Vai trò trong flow:
// - Được gọi từ handlers/projects.go để quản lý dự án

import (
	"errors"
	"time"

	"context"

	"github.com/iknizzz1807/SkillForge/internal/integrations"
	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectService struct {
	// db để truy cập MongoDB database
	db *mongo.Database
	// notificationService để gửi thông báo
	notificationService *NotificationService
	// aiClient để gọi API AI
	aiClient *integrations.AIClient
	// githubClient để tạo repository GitHub
	githubClient *integrations.GitHubClient
}

// NewProjectService khởi tạo ProjectService với các dependency
// Input: db (*mongo.Database), notificationService (*NotificationService), aiClient (*integrations.AIClient), githubClient (*integrations.GitHubClient)
// Return: *ProjectService - con trỏ đến ProjectService
func NewProjectService(db *mongo.Database, notificationService *NotificationService, aiClient *integrations.AIClient, githubClient *integrations.GitHubClient) *ProjectService {
	return &ProjectService{db, notificationService, aiClient, githubClient}
}

// GetAllProjects lấy danh sách tất cả dự án
// Return: []*models.Project (danh sách project), error (nếu có lỗi)
func (s *ProjectService) GetAllProjects() ([]*models.Project, error) {
	// Tạo repository và lấy danh sách project
	projectRepo := repositories.NewProjectRepository(s.db)
	projects, err := projectRepo.FindAllProjects(context.Background())
	if err != nil {
		return nil, err
	}

	// Trả về danh sách project
	return projects, nil
}

// GetProjectByID lấy chi tiết dự án theo ID
// Input: projectID (string)
// Return: *models.Project (project), error (nếu có lỗi)
func (s *ProjectService) GetProjectByID(projectID string) (*models.Project, error) {
	// Kiểm tra projectID hợp lệ
	if projectID == "" {
		return nil, errors.New("project ID cannot be empty")
	}

	// Tạo repository và tìm project
	projectRepo := repositories.NewProjectRepository(s.db)
	project, err := projectRepo.FindProjectByID(context.Background(), projectID)
	if err != nil || project == nil {
		return nil, errors.New("project not found")
	}

	// Trả về project
	return project, nil
}

// CreateProject tạo dự án mới
// Input: userID (string), title (string), description (string), skills ([]string), timeline (string)
// Return: *models.Project (project vừa tạo), error (nếu có lỗi)
func (s *ProjectService) CreateProject(userID, title, description string, skills []string, timeline string) (*models.Project, error) {
	// Kiểm tra input hợp lệ
	if userID == "" || title == "" || description == "" || len(skills) == 0 {
		return nil, errors.New("invalid project data")
	}

	// Tạo project mới
	project := &models.Project{
		ID:          utils.GenerateUUID(),
		Title:       title,
		Description: description,
		Skills:      skills,
		Timeline:    timeline,
		CreatedBy:   userID,
		Status:      "open",
		CreatedAt:   time.Now(),
	}

	// Lưu project vào database
	projectRepo := repositories.NewProjectRepository(s.db)
	err := projectRepo.InsertProject(context.Background(), project)
	if err != nil {
		return nil, err
	}

	// Tạo repository GitHub
	repoURL, err := s.githubClient.CreateRepository(title)
	if err == nil {
		project.RepoURL = repoURL
		projectRepo.UpdateProject(context.Background(), project)
	}

	// Gửi thông báo đến user
	s.notificationService.SendEmail(userID, "Project Created", "Your project "+title+" has been created.")

	// Trả về project vừa tạo
	return project, nil
}
