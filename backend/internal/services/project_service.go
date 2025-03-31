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

	"fmt"

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
func (s *ProjectService) CreateProject(userID, userName, title, description string, skills []string, startTime time.Time, endTime time.Time, maxMember int) (*models.Project, error) {
	// Kiểm tra input hợp lệ
	if userID == "" || title == "" || description == "" || len(skills) == 0 {
		return nil, errors.New("invalid project data")
	}

	// Tạo project mới
	project := &models.Project{
		ID:            utils.GenerateUUID(),
		Title:         title,
		Description:   description,
		Skills:        skills,
		StartTime:     startTime,
		EndTime:       endTime,
		MaxMember:     maxMember,
		CurrentMember: 0,
		CreatedByID:   userID,
		CreatedByName: userName,
		// Trường status này chỉ dùng để hiển thị và tìm kiếm
		Status:    "open",
		CreatedAt: time.Now(),
	}

	// Lưu project vào database
	projectRepo := repositories.NewProjectRepository(s.db)
	err := projectRepo.InsertProject(context.Background(), project)
	if err != nil {
		return nil, err
	}

	// // Tạo repository GitHub
	// repoURL, err := s.githubClient.CreateRepository(title)
	// if err == nil {
	// 	project.RepoURL = repoURL
	// 	projectRepo.UpdateProject(context.Background(), project)
	// }

	// Gửi thông báo đến user
	s.notificationService.SendEmail(userID, "Project Created", "Your project "+title+" has been created.")

	// Trả về project vừa tạo
	return project, nil
}

// DeleteProject xóa một project theo ID
// Input: projectID (string) - ID của project cần xóa
// Return: error - lỗi nếu có
func (s *ProjectService) DeleteProject(projectID string) error {
	// Kiểm tra xem project có tồn tại không
	_, err := s.GetProjectByID(projectID)
	if err != nil {
		return err
	}

	// Xóa project từ database
	projectRepo := repositories.NewProjectRepository(s.db)
	err = projectRepo.DeleteProject(projectID)
	if err != nil {
		return err
	}

	// Có thể thêm logic để xóa các dữ liệu liên quan như tasks, applications
	// và gửi thông báo cho các thành viên trong project

	return nil
}

// UpdateProject cập nhật thông tin của một project
// Input:
//   - projectID (string) - ID của project cần cập nhật
//   - title (string) - tiêu đề mới
//   - description (string) - mô tả mới
//   - skills ([]string) - danh sách kỹ năng cần thiết
//   - startTime (time.Time) - thời gian bắt đầu mới
//   - endTime (time.Time) - thời gian kết thúc mới
//   - maxMember (int) - số lượng thành viên tối đa mới
//   - status (string) - trạng thái mới
//
// Return:
//   - *models.Project - thông tin project sau khi cập nhật
//   - error - lỗi nếu có
func (s *ProjectService) UpdateProject(
	projectID string,
	title string,
	description string,
	skills []string,
	startTime time.Time,
	endTime time.Time,
	maxMember int,
	status string,
) (*models.Project, error) {
	// Lấy thông tin project hiện tại
	project, err := s.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}

	// Chỉ cập nhật các trường được cung cấp (không null/zero)
	if title != "" {
		project.Title = title
	}
	if description != "" {
		project.Description = description
	}
	if skills != nil && len(skills) > 0 {
		project.Skills = skills
	}
	if !startTime.IsZero() {
		project.StartTime = startTime
	}
	if !endTime.IsZero() {
		project.EndTime = endTime
	}
	if maxMember > 0 {
		project.MaxMember = maxMember
	}
	if status != "" {
		// Kiểm tra status hợp lệ
		if status != "open" && status != "active" && status != "closed" {
			return nil, fmt.Errorf("invalid project status: must be 'open', 'active', or 'closed'")
		}
		project.Status = status
	}

	// Cập nhật project trong database
	projectRepo := repositories.NewProjectRepository(s.db)
	updated, err := projectRepo.UpdateProject(project)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
