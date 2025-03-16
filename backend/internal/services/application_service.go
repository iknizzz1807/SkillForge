package services

// Nhiệm vụ: Xử lý logic ứng tuyển dự án
// Liên kết:
// - Dùng internal/repositories/application_repository.go để lưu ứng tuyển
// - Dùng notification_service.go để thông báo
// Vai trò trong flow:
// - Được gọi từ handlers/applications.go để sinh viên ứng tuyển dự án

import (
	"errors"
	"time"

	"context"

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
}

// NewApplicationService khởi tạo ApplicationService với dependency
// Input: db (*mongo.Database), notificationService (*NotificationService)
// Return: *ApplicationService - con trỏ đến ApplicationService
func NewApplicationService(db *mongo.Database, notificationService *NotificationService) *ApplicationService {
	return &ApplicationService{db, notificationService}
}

// ApplyProject xử lý ứng tuyển dự án
// Input: userID (string), projectID (string), proposal (string)
// Return: *models.Application (ứng tuyển), error (nếu có lỗi)
func (s *ApplicationService) ApplyProject(userID, projectID, proposal string) (*models.Application, error) {
	// Kiểm tra input hợp lệ
	if userID == "" || projectID == "" || proposal == "" {
		return nil, errors.New("invalid application data")
	}

	// Tạo ứng tuyển mới
	application := &models.Application{
		ID:        utils.GenerateUUID(),
		UserID:    userID,
		ProjectID: projectID,
		Proposal:  proposal,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	// Lưu ứng tuyển vào database
	appRepo := repositories.NewApplicationRepository(s.db)
	err := appRepo.InsertApplication(context.Background(), application)
	if err != nil {
		return nil, err
	}

	// Gửi thông báo đến doanh nghiệp (giả sử CreatedBy là doanh nghiệp)
	projectRepo := repositories.NewProjectRepository(s.db)
	project, _ := projectRepo.FindProjectByID(context.Background(), projectID)
	if project != nil {
		s.notificationService.SendEmail(project.CreatedBy, "New Application", "A student applied to your project: "+project.Title)
	}

	// Trả về ứng tuyển
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

	// Tạo repository và tìm ứng tuyển
	appRepo := repositories.NewApplicationRepository(s.db)
	application, err := appRepo.FindApplicationByID(context.Background(), applicationID)
	if err != nil || application == nil {
		return nil, errors.New("application not found")
	}

	// Trả về ứng tuyển
	return application, nil
}
