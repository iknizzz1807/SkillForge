package services

// Nhiệm vụ: Xử lý logic quản lý task (tạo, lấy, cập nhật)
// Liên kết:
// - Dùng internal/repositories/task_repository.go để tương tác với database
// - Dùng notification_service.go để thông báo
// Vai trò trong flow:
// - Được gọi từ handlers/tasks.go để quản lý task trong dự án

import (
	"errors"
	"log"
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
func (s *TaskService) GetTasksByProjectID(projectID string) ([]*models.Task, error) {
	// Kiểm tra taskID hợp lệ
	if projectID == "" {
		return nil, errors.New("project ID cannot be empty")
	}

	// Tạo repository và tìm task
	taskRepo := repositories.NewTaskRepository(s.db)
	tasks, err := taskRepo.FindTasksByProjectID(context.Background(), projectID)
	if err != nil || tasks == nil {
		return nil, errors.New("no task found")
	}

	// Trả về task
	return tasks, nil
}

// CreateTasks tạo nhiều tasks mới cho một project
// Input: projectID (string), taskInputs ([]TaskInput)
// Return: []*models.Task (danh sách tasks vừa tạo), error (nếu có lỗi)
func (s *TaskService) CreateTasks(projectID, userID string, taskInputs []models.TaskInput) ([]*models.Task, error) {
	log.Printf("Creating tasks for project %s by user %s, inputs: %+v", projectID, userID, taskInputs)
	// Kiểm tra input hợp lệ
	if projectID == "" || len(taskInputs) == 0 {
		return nil, errors.New("invalid inputs")
	}

	// Tạo danh sách tasks mới
	var tasks []*models.Task

	for _, input := range taskInputs {
		if input.Title == "" {
			continue // Bỏ qua task không có tiêu đề
		}

		task := &models.Task{
			ID:          utils.GenerateUUID(),
			ProjectID:   projectID,
			Title:       input.Title,
			Description: input.Description,
			Note:        input.Note,
			Assigned_to: input.AssignedTo,
			Status:      "todo",
			Finished_by: "",
			CreatedAt:   time.Now(),
		}
		tasks = append(tasks, task)
	}

	// Kiểm tra xem có tasks nào được tạo không
	if len(tasks) == 0 {
		return nil, errors.New("no valid task is created")
	}

	// Lưu tasks vào database
	taskRepo := repositories.NewTaskRepository(s.db)
	err := taskRepo.InsertTasks(context.Background(), tasks)
	if err != nil {
		return nil, err
	}

	for _, task := range tasks {
		s.InsertActivity("create", userID, projectID, task.Title, "", "")
	}
	// Trả về danh sách tasks
	return tasks, nil
}

// UpdateTask cập nhật trạng thái task
// Input: taskID (string), status (string)
// Return: *models.Task (task đã cập nhật), error (nếu có lỗi)
func (s *TaskService) UpdateTask(taskID, userID string, taskUpdate *models.TaskUpdate) (*models.Task, error) {
	// Kiểm tra input hợp lệ
	if taskID == "" || taskUpdate == nil {
		return nil, errors.New("invalid inputs")
	}

	// Tạo repository và tìm task
	taskRepo := repositories.NewTaskRepository(s.db)
	task, err := taskRepo.FindTaskByID(context.Background(), taskID)
	if err != nil || task == nil {
		return nil, errors.New("task not found")
	}

	if task.Status != taskUpdate.Status {
		s.InsertActivity("Move", userID, task.ProjectID, task.Title, task.Status, taskUpdate.Status)
	}

	if task.Title != taskUpdate.Title {
		s.InsertActivity("ChangeTitle", userID, task.ProjectID, task.Title, task.Title, taskUpdate.Title)
	}

	if task.Description != taskUpdate.Description {
		s.InsertActivity("ChangeDescription", userID, task.ProjectID, task.Title, task.Description, taskUpdate.Description)
	}

	if task.Note != taskUpdate.Note {
		s.InsertActivity("ChangeNote", userID, task.ProjectID, task.Title, task.Note, taskUpdate.Note)
	}

	if task.Assigned_to != taskUpdate.Assigned_to {
		s.InsertActivity("ChangeAssignee", userID, task.ProjectID, task.Title, task.Assigned_to, taskUpdate.Assigned_to)
	}

	// Cập nhật trạng thái và thời gian
	task.Status = taskUpdate.Status
	task.Title = taskUpdate.Title
	task.Description = taskUpdate.Description
	task.Note = taskUpdate.Note
	task.Assigned_to = taskUpdate.Assigned_to
	task.UpdatedAt = time.Now()

	// Lưu thay đổi
	err = taskRepo.UpdateTask(context.Background(), task)
	if err != nil {
		return nil, err
	}

	// Trả về task đã cập nhật
	return task, nil
}

// DeleteTask xóa một task khỏi database
// Input: taskID (string)
// Return: error (nếu có lỗi)
func (s *TaskService) DeleteTask(taskID, userID string) error {
	// Kiểm tra taskID hợp lệ
	if taskID == "" {
		return errors.New("task ID cannot be empty")
	}

	// Tạo repository và tìm task để kiểm tra tồn tại
	taskRepo := repositories.NewTaskRepository(s.db)
	task, err := taskRepo.FindTaskByID(context.Background(), taskID)
	if err != nil {
		return err
	}

	if task == nil {
		return errors.New("cannot find the task to delete")
	}

	// Xóa task khỏi database
	err = taskRepo.DeleteTask(context.Background(), taskID)
	if err != nil {
		return err
	}

	// // Thông báo cho các bên liên quan (nếu cần)
	// if task.Assigned_to != "" {
	// 	s.notificationService.SendNotification(task.Assigned_to, "Task Deleted",
	// 		"Một task bạn được gán đã bị xóa: "+task.Description)
	// }

	// Trả về nil nếu thành công
	s.InsertActivity("delete", userID, task.ProjectID, task.Title, "", "")
	return nil
}

func (s *TaskService) GetActivityByProjectID(userID string, projectID string) ([]*models.Activity, error) {
	// Tạo repository và tìm task
	taskRepo := repositories.NewTaskRepository(s.db)
	return taskRepo.FindActivitiesByProjectID(context.Background(), projectID)
}

func (s *TaskService) InsertActivity(Type, DoneBy, ProjectID, Title, From, To string) error {
	// Tạo repository và tìm task
	activity := &models.Activity{
		ID:        utils.GenerateUUID(),
		Type:      Type,
		DoneBy:    DoneBy,
		ProjectID: ProjectID,
		Title:     Title,
		From:      From,
		To:        To,
	}
	taskRepo := repositories.NewTaskRepository(s.db)
	return taskRepo.InsertActivity(context.Background(), activity)
}
