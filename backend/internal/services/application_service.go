// internal/services/application_service.go
package services

// Nhiệm vụ: Xử lý logic ứng tuyển dự án
// Liên kết:
// - Dùng internal/repositories/application_repository.go để lưu ứng tuyển
// - Dùng internal/repositories/project_repository.go để kiểm tra project
// - Dùng notification_service.go để thông báo
// Vai trò trong flow:
// - Được gọi từ handlers/applications.go để sinh viên ứng tuyển dự án

import (
	"context"
	"errors"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApplicationService struct {
	// db để truy cập MongoDB database
	db *mongo.Database
	// notificationService để gửi thông báo
	notificationService *NotificationService
	// projectRepo để truy vấn thông tin project
	projectRepo *repositories.ProjectRepository // Thêm project repo
	// applicationRepo để truy vấn application
	applicationRepo *repositories.ApplicationRepository // Thêm application repo

}

// NewApplicationService khởi tạo ApplicationService với dependency
// Input: db (*mongo.Database), notificationService (*NotificationService)
// Return: *ApplicationService - con trỏ đến ApplicationService
func NewApplicationService(db *mongo.Database, notificationService *NotificationService) *ApplicationService {
	return &ApplicationService{
		db:                  db,
		notificationService: notificationService,
		projectRepo:         repositories.NewProjectRepository(db),     // Khởi tạo project repo
		applicationRepo:     repositories.NewApplicationRepository(db), // Khởi tạo application repo
	}
}

// ApplyProject xử lý ứng tuyển dự án
// Input: userID (string), projectID (string), motivation (string), detailedProposal (string)
// Return: *models.Application (ứng tuyển), error (nếu có lỗi)
// --- CHANGE START ---
// Thay đổi signature hàm
func (s *ApplicationService) ApplyProject(userID, projectID, motivation, detailedProposal string) (*models.Application, error) {
	// Kiểm tra input hợp lệ
	if userID == "" || projectID == "" || motivation == "" || detailedProposal == "" {
		return nil, errors.New("invalid application data: user ID, project ID, motivation, and detailed proposal are required")
	}

	ctx := context.Background() // Tạo context

	// Kiểm tra xem project có tồn tại và đang mở không?
	project, err := s.projectRepo.FindProjectByID(ctx, projectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) || err.Error() == "project not found" {
			return nil, errors.New("project not found")
		}
		return nil, errors.New("failed to check project existence: " + err.Error())
	}
	if project.Status != "open" {
		return nil, errors.New("project is not open for applications")
	}

	// Fetch user information to get the user name
	userRepo := repositories.NewUserRepository(s.db)
	user, err := userRepo.FindUserByID(ctx, userID)
	if err != nil {
		return nil, errors.New("failed to fetch user information: " + err.Error())
	}

	// Create application with project name and user name
	application := &models.Application{
		ID:               utils.GenerateUUID(),
		UserID:           userID,
		UserName:         user.Name, // Added user name for frontend to display
		ProjectID:        projectID,
		ProjectName:      project.Title, // Added project name for frontend to display
		Motivation:       motivation,
		DetailedProposal: detailedProposal,
		Status:           "pending",
		CreatedAt:        time.Now(),
	}

	// Lưu ứng tuyển vào database
	err = s.applicationRepo.InsertApplication(ctx, application)
	if err != nil {
		return nil, errors.New("failed to save application: " + err.Error())
	}

	// Gửi thông báo đến doanh nghiệp (người tạo project)
	if project != nil && s.notificationService != nil { // Kiểm tra nil safety
		go func() { // Gửi bất đồng bộ để không block response
			// Lấy thông tin business user nếu cần (ví dụ: email)
			// userRepo := repositories.NewUserRepository(s.db)
			// businessUser, _ := userRepo.FindUserByID(ctx, project.CreatedByID)
			// if businessUser != nil {
			// s.notificationService.SendEmail(businessUser.Email, "New Application Received", fmt.Sprintf("A student has applied to your project '%s'.", project.Title))
			// s.notificationService.SendNotification(project.CreatedByID, "New Application", fmt.Sprintf("New application for project: %s", project.Title), true) // Gửi realtime notification
			// }
		}()
	}

	return application, nil
}

// GetApplicationByID lấy chi tiết ứng tuyển
// Input: applicationID (string)
// Return: *models.Application (ứng tuyển), error (nếu có lỗi)
func (s *ApplicationService) GetApplicationByID(applicationID string) (*models.Application, error) {
	// Kiểm tra applicationID hợp lệ
	if applicationID == "" {
		return nil, errors.New("application ID cannot be empty")
	}

	// Tìm ứng tuyển
	application, err := s.applicationRepo.FindApplicationByID(context.Background(), applicationID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) { // Xử lý lỗi chuẩn từ driver
			return nil, errors.New("application not found")
		}
		return nil, errors.New("failed to get application: " + err.Error())
	}

	// Trả về ứng tuyển
	return application, nil
}

// GetApplicationsByUserID lấy tất cả applications của một user (student)
func (s *ApplicationService) GetApplicationsByUserID(userID string) ([]models.Application, error) {
	if userID == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	applications, err := s.applicationRepo.FindByUserID(userID)
	if err != nil {
		return nil, errors.New("failed to get user applications: " + err.Error())
	}

	// Ensure we have a slice rather than nil
	if applications == nil {
		return []models.Application{}, nil
	}

	// // For each application that doesn't have project name, fetch it from the project repository
	// Can remove this shit if made sure every application has a project name
	// for i := range applications {
	// 	if applications[i].ProjectName == "" {
	// 		project, err := s.projectRepo.FindProjectByID(context.Background(), applications[i].ProjectID)
	// 		if err == nil && project != nil {
	// 			applications[i].ProjectName = project.Title
	// 		}
	// 	}
	// }

	return applications, nil
}

// GetApplicationsByBusinessID lấy tất cả applications cho các projects của business
func (s *ApplicationService) GetApplicationsByBusinessID(businessID string) ([]models.Application, error) {
	if businessID == "" {
		return nil, errors.New("business ID cannot be empty")
	}
	// Lấy tất cả project ID của business
	projectIDs, err := s.projectRepo.GetProjectIDsByCreatorID(businessID) // Repo đã có hàm này
	if err != nil {
		return nil, errors.New("failed to get business projects: " + err.Error())
	}

	// Nếu không có project nào, trả về mảng rỗng
	if len(projectIDs) == 0 {
		return []models.Application{}, nil
	}

	// Lấy tất cả applications cho các projects đó
	applications, err := s.applicationRepo.FindByProjectIDs(projectIDs) // Repo đã có hàm này
	if err != nil {
		return nil, errors.New("failed to get project applications: " + err.Error())
	}
	// Đảm bảo trả về slice rỗng thay vì nil
	if applications == nil {
		return []models.Application{}, nil
	}
	return applications, nil
}

// UpdateApplicationStatus cập nhật trạng thái application
// (Optional: Thêm businessID để kiểm tra quyền trong service)
func (s *ApplicationService) UpdateApplicationStatus(applicationID string, status string /*, businessID string*/) (*models.Application, error) {
	if applicationID == "" || status == "" {
		return nil, errors.New("application ID and status cannot be empty")
	}

	// Kiểm tra status hợp lệ (sử dụng constants)
	validStatuses := map[string]bool{
		"pending":  true,
		"approved": true,
		"rejected": true,
	}
	if !validStatuses[status] {
		return nil, errors.New("invalid application status")
	}

	ctx := context.Background()

	// 1. Lấy application hiện tại để kiểm tra
	application, err := s.applicationRepo.FindApplicationByID(ctx, applicationID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("application not found")
		}
		return nil, errors.New("failed to get application for update: " + err.Error())
	}

	// (Optional: Kiểm tra quyền của business nếu businessID được truyền vào)
	// project, err := s.projectRepo.FindProjectByID(ctx, application.ProjectID)
	// if err != nil { ... }
	// if project.CreatedByID != businessID {
	// 	return nil, errors.New("permission denied to update this application")
	// }

	// 2. Kiểm tra logic chuyển đổi trạng thái (ví dụ: không thể approve lại application đã rejected)
	if application.Status == "approved" && status != "approved" {
		// Có thể muốn ngăn việc thay đổi status sau khi đã approve, hoặc xử lý việc rút lại approval
		// return nil, errors.New("cannot change status of an already approved application")
	}
	if application.Status == "rejected" && status != "rejected" {
		// return nil, errors.New("cannot change status of an already rejected application")
	}

	// 3. Thực hiện cập nhật status
	updatedApplication, err := s.applicationRepo.UpdateStatus(applicationID, status)
	if err != nil {
		return nil, errors.New("failed to update application status: " + err.Error())
	}

	// 4. Xử lý logic sau khi cập nhật status
	// Nếu status là "approved":
	if status == "approved" && application.Status != "approved" { // Chỉ xử lý khi chuyển sang approved
		// Thêm student vào project (cần inject ProjectService hoặc gọi trực tiếp repo)
		projectService := NewProjectService(s.db, s.notificationService, nil, nil) // Tạm thời khởi tạo ở đây, nên inject vào struct
		errAdd := projectService.AddStudentToProject(updatedApplication.ProjectID, updatedApplication.UserID)
		if errAdd != nil {
			// Xử lý lỗi khi thêm student vào project (quan trọng)
			// Có thể rollback status application hoặc log lỗi nghiêm trọng
			// Ví dụ: Rollback status
			// _, _ = s.applicationRepo.UpdateStatus(applicationID, application.Status) // Rollback về status cũ
			// return nil, fmt.Errorf("application approved but failed to add student to project: %v. Status rolled back", errAdd)
			// utils.GetLogger().Errorf("Application %s approved but failed to add student %s to project %s: %v", applicationID, updatedApplication.UserID, updatedApplication.ProjectID, errAdd)
			// Không rollback, chỉ log lỗi, vì application vẫn được coi là approved
		} else {
			// Gửi thông báo cho student là application đã được approved
			// go func() {
			// 	// s.notificationService.SendNotification(updatedApplication.UserID, "Application Approved", fmt.Sprintf("Your application for project '%s' has been approved!", updatedApplication.ProjectID), true)
			// 	// s.notificationService.SendEmail(...)
			// }()
		}
	} else if status == "rejected" && application.Status != "rejected" {
		// Gửi thông báo cho student là application đã bị rejected
		// go func() {
		// 	// s.notificationService.SendNotification(updatedApplication.UserID, "Application Rejected", fmt.Sprintf("Your application for project '%s' has been rejected.", updatedApplication.ProjectID), true)
		// 	// s.notificationService.SendEmail(...)
		// }()
	}

	return updatedApplication, nil
}

// DeleteApplication xóa một application theo ID
// Input: applicationID (string), userID (string) - người thực hiện xóa
// Return: error (nếu có lỗi)
func (s *ApplicationService) DeleteApplication(applicationID string, userID string) error {
	if applicationID == "" {
		return errors.New("application ID cannot be empty")
	}

	ctx := context.Background()

	// 1. Lấy application để kiểm tra quyền
	application, err := s.applicationRepo.FindApplicationByID(ctx, applicationID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("application not found")
		}
		return errors.New("failed to get application: " + err.Error())
	}

	// 2. Kiểm tra người xóa có phải là người nộp đơn hoặc chủ dự án
	project, err := s.projectRepo.FindProjectByID(ctx, application.ProjectID)
	if err != nil {
		return errors.New("failed to get project: " + err.Error())
	}

	// Chỉ cho phép student đã nộp đơn hoặc business sở hữu project được xóa
	if application.UserID != userID && project.CreatedByID != userID {
		return errors.New("you don't have permission to delete this application")
	}

	// 3. Xóa application
	err = s.applicationRepo.Delete(ctx, applicationID)
	if err != nil {
		return errors.New("failed to delete application: " + err.Error())
	}

	return nil
}

// NOTE: cái này có thể được bỏ qua do không cần handler để gọi tới service này mà hành động
// này sẽ được thực thi bằng việc gọi thẳng tới tầng repository
// DeleteApplicationsByProjectID xóa tất cả applications của một project
// Input: projectID (string)
// Return: int (số lượng applications đã xóa), error (nếu có lỗi)
func (s *ApplicationService) DeleteApplicationsByProjectID(projectID string) (int64, error) {
	if projectID == "" {
		return 0, errors.New("project ID cannot be empty")
	}

	ctx := context.Background()

	// Xóa tất cả applications liên quan đến project
	deletedCount, err := s.applicationRepo.DeleteApplicationsByProjectID(ctx, projectID)
	if err != nil {
		return 0, errors.New("failed to delete project applications: " + err.Error())
	}

	return deletedCount, nil
}
