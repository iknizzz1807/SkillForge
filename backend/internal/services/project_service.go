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

	"github.com/google/uuid"
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
func (s *ProjectService) CreateProject(creatorID, userName, title, description string, skills []string, startTime time.Time, endTime time.Time, maxMember int, difficulty string) (*models.Project, error) {
	// Kiểm tra input hợp lệ
	if creatorID == "" || title == "" || description == "" || len(skills) == 0 {
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
		Difficulty:    difficulty,
		CreatedByID:   creatorID,
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
	// s.notificationService.SendEmail(creatorID, "Project Created", "Your project "+title+" has been created.")

	// Trả về project vừa tạo
	return project, nil
}

// DeleteProject xóa project và tất cả dữ liệu liên quan
// Input: projectID (string), userID (string) - userID để kiểm tra quyền
// Return: error (nếu có lỗi)
func (s *ProjectService) DeleteProject(projectID string, userID string) error {
	if projectID == "" {
		return errors.New("project ID cannot be empty")
	}

	ctx := context.Background()

	projectRepo := repositories.NewProjectRepository(s.db)

	// 1. Kiểm tra project có tồn tại không và userID có quyền xóa không
	project, err := projectRepo.FindProjectByID(ctx, projectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) || err.Error() == "project not found" {
			return errors.New("project not found")
		}
		return errors.New("failed to check project existence: " + err.Error())
	}

	// Chỉ chủ dự án mới được phép xóa
	if project.CreatedByID != userID {
		return errors.New("you don't have permission to delete this project")
	}

	// 2. Khởi tạo các repositories cần thiết
	applicationService := NewApplicationService(s.db, s.notificationService)
	projectStudentRepo := repositories.NewProjectStudentRepository(s.db)

	// 3. Xóa tất cả applications liên quan đến project
	_, err = applicationService.DeleteApplicationsByProjectID(projectID)
	if err != nil {
		return errors.New("failed to delete project applications: " + err.Error())
	}

	// 4. Lấy danh sách sinh viên tham gia project
	studentIDs, err := projectStudentRepo.FindStudentsByProjectID(ctx, projectID)
	if err != nil {
		// Log error but continue
		// utils.GetLogger().Warnf("Failed to get students for project %s: %v", projectID, err)
	}

	// 5. Xóa tất cả quan hệ giữa project và student
	for _, studentID := range studentIDs {
		err := projectStudentRepo.DeleteProjectStudent(ctx, projectID, studentID)
		if err != nil {
			// Log error but continue
			// utils.GetLogger().Warnf("Failed to delete project-student relation for project %s and student %s: %v", projectID, studentID, err)
		}
	}

	// 6. Xóa project
	err = projectRepo.DeleteProject(projectID)
	if err != nil {
		return errors.New("failed to delete project: " + err.Error())
	}

	// 7. Gửi thông báo cho các sinh viên về việc dự án đã bị xóa
	for _, studentID := range studentIDs {
		if s.notificationService != nil {
			go func(id string) {
				// s.notificationService.SendNotification(
				//    id,
				//    "Project Removed",
				//    fmt.Sprintf("The project '%s' you were participating in has been deleted.", project.Title),
				//    true
				// )
			}(studentID)
		}
	}

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
	difficulty string,
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
	if difficulty != "" {
		project.Difficulty = difficulty
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
		if status != "open" && status != "close" {
			return nil, fmt.Errorf("invalid project status: must be 'open' or 'close'")
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

// GetProjectsByStudentID lấy danh sách tất cả dự án mà student đã tham gia
// Input: studentID (string) - ID của student
// Return: []*models.Project (danh sách project), error (nếu có lỗi)
func (s *ProjectService) GetProjectsByStudentID(studentID string) ([]*models.Project, error) {
	// Kiểm tra studentID hợp lệ
	if studentID == "" {
		return nil, errors.New("student ID cannot be empty")
	}

	ctx := context.Background()

	// Lấy danh sách project ID từ repository project_student
	projectStudentRepo := repositories.NewProjectStudentRepository(s.db)
	projectIDs, err := projectStudentRepo.FindProjectsByStudentID(ctx, studentID)
	if err != nil {
		return nil, err
	}

	// Nếu không có project nào, trả về mảng rỗng
	if len(projectIDs) == 0 {
		return []*models.Project{}, nil
	}

	// Lấy thông tin chi tiết từng project
	projectRepo := repositories.NewProjectRepository(s.db)
	var projects []*models.Project

	for _, projectID := range projectIDs {
		project, err := projectRepo.FindProjectByID(ctx, projectID)
		if err == nil && project != nil {
			projects = append(projects, project)
		}
	}

	// Trả về danh sách project
	return projects, nil
}

// GetProjectsByBusinessID lấy danh sách tất cả dự án được tạo bởi business
// Input: businessID (string) - ID của business
// Return: []*models.Project (danh sách project), error (nếu có lỗi)
func (s *ProjectService) GetProjectsByBusinessID(businessID string) ([]*models.Project, error) {
	// Kiểm tra businessID hợp lệ
	if businessID == "" {
		return nil, errors.New("business ID cannot be empty")
	}

	ctx := context.Background()

	// Tạo filter tìm theo creator_id (businessID)
	projectRepo := repositories.NewProjectRepository(s.db)

	// Lấy danh sách project ID của business
	projectIDs, err := projectRepo.GetProjectIDsByCreatorID(businessID)
	if err != nil {
		return nil, err
	}

	// Nếu không có project nào, trả về mảng rỗng
	if len(projectIDs) == 0 {
		return []*models.Project{}, nil
	}

	// Lấy thông tin chi tiết từng project
	var projects []*models.Project

	for _, projectID := range projectIDs {
		project, err := projectRepo.FindProjectByID(ctx, projectID)
		if err == nil && project != nil {
			projects = append(projects, project)
		}
	}

	// Trả về danh sách project
	return projects, nil
}

// AddStudentToProject thêm student vào project
// Input: projectID (string), studentID (string)
// Return: error (nếu có lỗi)
func (s *ProjectService) AddStudentToProject(projectID, studentID string) error {
	ctx := context.Background()

	// Kiểm tra project và capacity
	project, err := s.GetProjectByID(projectID)
	if err != nil {
		return err
	}

	// Kiểm tra số lượng thành viên
	projectStudentRepository := repositories.NewProjectStudentRepository(s.db)
	studentIDs, err := projectStudentRepository.FindStudentsByProjectID(ctx, projectID)
	if err != nil {
		return err
	}

	if len(studentIDs) >= project.MaxMember {
		return errors.New("project has reached maximum capacity")
	}

	// Tạo bản ghi project_student mới
	projectStudent := &models.Project_student{
		ID:         uuid.New().String(), // hoặc cách tạo ID của bạn
		Project_id: projectID,
		Student_id: studentID,
		CreatedAt:  time.Now(),
	}

	// Thêm vào database
	err = projectStudentRepository.InsertProjectStudent(ctx, projectStudent)
	if err != nil {
		return err
	}

	// Cập nhật số lượng thành viên
	return s.UpdateProjectMemberCount(projectID)
}

// RemoveStudentFromProject xóa student khỏi project
// Input: projectID (string), studentID (string)
// Return: error (nếu có lỗi)
func (s *ProjectService) RemoveStudentFromProject(projectID, studentID string) error {
	ctx := context.Background()

	projectStudentRepository := repositories.NewProjectStudentRepository(s.db)
	// Xóa bản ghi từ database
	err := projectStudentRepository.DeleteProjectStudent(ctx, projectID, studentID)
	if err != nil {
		return err
	}

	// Cập nhật số lượng thành viên
	return s.UpdateProjectMemberCount(projectID)
}

// UpdateProjectMemberCount đếm và cập nhật số lượng thành viên dự án
// Input: projectID (string)
// Return: error (nếu có lỗi)
func (s *ProjectService) UpdateProjectMemberCount(projectID string) error {
	ctx := context.Background()

	projectStudentRepository := repositories.NewProjectStudentRepository(s.db)
	// Lấy danh sách sinh viên trong dự án
	studentIDs, err := projectStudentRepository.FindStudentsByProjectID(ctx, projectID)
	if err != nil {
		return err
	}

	// Đếm số lượng sinh viên
	memberCount := len(studentIDs)

	projectRepository := repositories.NewProjectRepository(s.db)
	// Cập nhật trường currentMember
	return projectRepository.UpdateProjectMemberCount(ctx, projectID, memberCount)
}

// SyncAllProjectMemberCounts đồng bộ hóa số lượng thành viên cho tất cả các dự án
// Return: error (nếu có lỗi)
func (s *ProjectService) SyncAllProjectMemberCounts() error {
	ctx := context.Background()

	// Lấy tất cả project
	projectRepository := repositories.NewProjectRepository(s.db)
	projects, err := projectRepository.FindAllProjects(ctx)
	if err != nil {
		return err
	}

	// Cập nhật số lượng thành viên cho từng project
	for _, project := range projects {
		err := s.UpdateProjectMemberCount(project.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetProjectWithUserStatus lấy chi tiết project và thông tin tham gia của user
// Input: projectID (string), userID (string)
// Return: *models.Project (project), bool (đã tham gia), bool (đã apply), error (nếu có lỗi)
func (s *ProjectService) GetProjectWithUserStatus(projectID, userID string) (*models.Project, bool, bool, error) {
	// Kiểm tra projectID hợp lệ
	if projectID == "" {
		return nil, false, false, errors.New("project ID cannot be empty")
	}

	// Lấy thông tin project từ database
	projectRepo := repositories.NewProjectRepository(s.db)
	project, err := projectRepo.FindProjectByID(context.Background(), projectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, false, false, errors.New("project not found")
		}
		return nil, false, false, errors.New("failed to get project: " + err.Error())
	}

	// Nếu không có userID (user chưa đăng nhập), trả về project mà không có status
	if userID == "" {
		return project, false, false, nil
	}

	// Kiểm tra xem user đã tham gia project chưa bằng ProjectStudentRepository
	projectStudentRepo := repositories.NewProjectStudentRepository(s.db)
	hasJoined := false

	// Kiểm tra nếu student_id tồn tại trong collection project_student với project_id tương ứng
	ctx := context.Background()
	studentIDs, err := projectStudentRepo.FindStudentsByProjectID(ctx, projectID)
	if err == nil {
		for _, studentID := range studentIDs {
			if studentID == userID {
				hasJoined = true
				break
			}
		}
	}

	// Kiểm tra xem user đã apply vào project này chưa
	hasApplied := false
	if !hasJoined { // Chỉ kiểm tra nếu chưa tham gia
		// Sử dụng application repository để kiểm tra
		applicationRepo := repositories.NewApplicationRepository(s.db)
		application, err := applicationRepo.FindByUserAndProject(ctx, userID, projectID)

		// Nếu không có lỗi và tìm thấy application, đánh dấu là đã apply
		if err == nil && application != nil {
			hasApplied = true
		}
	}

	return project, hasJoined, hasApplied, nil
}

// GetStudentsByProjectID lấy danh sách thông tin sinh viên tham gia project
// Input: projectID (string)
// Return: []models.User (danh sách user), error (nếu có lỗi)
func (s *ProjectService) GetStudentsByProjectID(projectID string) ([]models.User, error) {
	if projectID == "" {
		return nil, errors.New("project ID cannot be empty")
	}

	ctx := context.Background()

	// Lấy danh sách student ID từ project_student repository
	projectStudentRepo := repositories.NewProjectStudentRepository(s.db)
	studentIDs, err := projectStudentRepo.FindStudentsByProjectID(ctx, projectID)
	if err != nil {
		return nil, errors.New("failed to get student IDs: " + err.Error())
	}

	// Nếu không có student nào, trả về mảng rỗng
	if len(studentIDs) == 0 {
		return []models.User{}, nil
	}

	// Lấy thông tin chi tiết từng student
	userRepo := repositories.NewUserRepository(s.db)
	var students []models.User

	for _, studentID := range studentIDs {
		student, err := userRepo.FindUserByID(ctx, studentID)
		if err == nil && student != nil {
			// Loại bỏ các thông tin nhạy cảm trước khi trả về
			student.Password = ""
			students = append(students, *student)
		}
	}

	return students, nil
}
