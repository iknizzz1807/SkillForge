package services

// Nhiệm vụ: Xử lý logic quản lý task (tạo, lấy, cập nhật)
// Liên kết:
// - Dùng internal/repositories/task_repository.go để tương tác với database
// - Dùng notification_service.go để thông báo
// Vai trò trong flow:
// - Được gọi từ handlers/tasks.go để quản lý task trong dự án

import (
	"errors"
	"time"

	"context"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	// db để truy cập MongoDB database
	db *mongo.Database
	// notificationService để gửi thông báo
	notificationService *NotificationService
}

// NewTaskService khởi tạo TaskService với dependency
// Input: db (*mongo.Database), notificationService (*NotificationService)
// Return: *TaskService - con trỏ đến TaskService
func NewTaskService(db *mongo.Database, notificationService *NotificationService) *TaskService {
	return &TaskService{db, notificationService}
}

// GetTaskByID lấy chi tiết task
// Input: taskID (string)
// Return: *models.Task (task), error (nếu có lỗi)
func (s *TaskService) GetTaskByID(taskID string) (*models.Task, error) {
	// Kiểm tra taskID hợp lệ
	if taskID == "" {
		return nil, errors.New("task ID cannot be empty")
	}

	// Tạo repository và tìm task
	taskRepo := repositories.NewTaskRepository(s.db)
	task, err := taskRepo.FindTaskByID(context.Background(), taskID)
	if err != nil || task == nil {
		return nil, errors.New("task not found")
	}

	// Trả về task
	return task, nil
}

// CreateTask tạo task mới
// Input: projectID (string), description (string)
// Return: *models.Task (task vừa tạo), error (nếu có lỗi)
func (s *TaskService) CreateTask(projectID, description string) (*models.Task, error) {
	// Kiểm tra input hợp lệ
	if projectID == "" || description == "" {
		return nil, errors.New("invalid task data")
	}

	// Tạo task mới
	task := &models.Task{
		ID:          utils.GenerateUUID(),
		ProjectID:   projectID,
		Description: description,
		Status:      "todo",
		Review:      "",
		Finished_by: "",
		CreatedAt:   time.Now(),
	}

	// Lưu task vào database
	taskRepo := repositories.NewTaskRepository(s.db)
	err := taskRepo.InsertTask(context.Background(), task)
	if err != nil {
		return nil, err
	}

	// Gửi thông báo (giả sử thông báo đến người tạo dự án)
	projectRepo := repositories.NewProjectRepository(s.db)
	project, _ := projectRepo.FindProjectByID(context.Background(), projectID)
	// Implement send email and send notification with sound on website uisng websocket
	// Redeisgn these services to make it send to every members in the team, not just one member
	if project != nil {
		s.notificationService.SendEmail(project.CreatedBy, "New Task", "A new task was added to your project.")
	}

	// Trả về task
	return task, nil
}

// UpdateTask cập nhật trạng thái task
// Input: taskID (string), status (string)
// Return: *models.Task (task đã cập nhật), error (nếu có lỗi)
func (s *TaskService) UpdateTask(taskID, status string) (*models.Task, error) {
	// Kiểm tra input hợp lệ
	if taskID == "" || status == "" {
		return nil, errors.New("invalid task data")
	}

	// Tạo repository và tìm task
	taskRepo := repositories.NewTaskRepository(s.db)
	task, err := taskRepo.FindTaskByID(context.Background(), taskID)
	if err != nil || task == nil {
		return nil, errors.New("task not found")
	}

	// Cập nhật trạng thái và thời gian
	task.Status = status
	task.UpdatedAt = time.Now()

	// Lưu thay đổi
	err = taskRepo.UpdateTask(context.Background(), task)
	if err != nil {
		return nil, err
	}

	// Trả về task đã cập nhật
	return task, nil
}
