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
// Input: projectID (string), descriptions ([]string)
// Return: []*models.Task (danh sách tasks vừa tạo), error (nếu có lỗi)
func (s *TaskService) CreateTasks(projectID string, descriptions []string) ([]*models.Task, error) {
	// Kiểm tra input hợp lệ
	if projectID == "" || len(descriptions) == 0 {
		return nil, errors.New("project ID hoặc danh sách mô tả không hợp lệ")
	}

	// Tạo danh sách tasks mới
	var tasks []*models.Task
	now := time.Now()

	for _, description := range descriptions {
		if description == "" {
			continue // Bỏ qua mô tả rỗng
		}

		task := &models.Task{
			ID:          utils.GenerateUUID(),
			ProjectID:   projectID,
			Description: description,
			Status:      "todo",
			Review:      "",
			Finished_by: "",
			CreatedAt:   now,
		}
		tasks = append(tasks, task)
	}

	// Kiểm tra xem có tasks nào được tạo không
	if len(tasks) == 0 {
		return nil, errors.New("không có task hợp lệ nào được tạo")
	}

	// Lưu tasks vào database
	taskRepo := repositories.NewTaskRepository(s.db)
	err := taskRepo.InsertTasks(context.Background(), tasks)
	if err != nil {
		return nil, err
	}

	// // Gửi thông báo (giả sử thông báo đến người tạo dự án)
	// projectRepo := repositories.NewProjectRepository(s.db)
	// project, _ := projectRepo.FindProjectByID(context.Background(), projectID)
	// if project != nil {
	// 	message := fmt.Sprintf("%d tasks mới đã được thêm vào dự án của bạn", len(tasks))
	// 	s.notificationService.SendEmail(project.CreatedBy, "New Tasks Added", message)

	// 	// Có thể thêm thông báo realtime qua WebSocket ở đây
	// }

	// Trả về danh sách tasks
	return tasks, nil
}

// AssignTask gán task cho người dùng cụ thể
// Input: taskID (string), userID (string)
// Return: *models.Task (task đã cập nhật), error (nếu có lỗi)
func (s *TaskService) AssignTask(taskID string, userID string) (*models.Task, error) {
	// Kiểm tra input hợp lệ
	if taskID == "" || userID == "" {
		return nil, errors.New("task ID và user ID không được để trống")
	}

	// Tạo repository và tìm task
	taskRepo := repositories.NewTaskRepository(s.db)
	task, err := taskRepo.FindTaskByID(context.Background(), taskID)
	if err != nil || task == nil {
		return nil, errors.New("không tìm thấy task")
	}

	// Cập nhật người được gán và thời gian
	task.Assigned_to = userID
	task.UpdatedAt = time.Now()

	// Nếu task đang ở trạng thái "todo", chuyển sang "in-progress"
	if task.Status == "todo" {
		task.Status = "in-progress"
	}

	// Lưu thay đổi
	err = taskRepo.UpdateTask(context.Background(), task)
	if err != nil {
		return nil, err
	}

	// // Gửi thông báo cho người được gán task
	// s.notificationService.SendNotification(userID, "Task Assignment",
	//     "You have been assigned a new task: "+task.Description)

	// // Có thể gửi email hoặc thông báo khác nếu cần
	// s.notificationService.SendEmail(userID, "New Task Assignment",
	//     "You have been assigned to task: "+task.Description)

	// Trả về task đã cập nhật
	return task, nil
}

// FinishTask đánh dấu task là đã hoàn thành bởi người dùng
// Input: taskID (string), userID (string)
// Return: *models.Task (task đã cập nhật), error (nếu có lỗi)
func (s *TaskService) FinishTask(taskID string, userID string) (*models.Task, error) {
	// Kiểm tra input hợp lệ
	if taskID == "" || userID == "" {
		return nil, errors.New("invalid inputs")
	}

	// Tạo repository và tìm task
	taskRepo := repositories.NewTaskRepository(s.db)
	task, err := taskRepo.FindTaskByID(context.Background(), taskID)
	if err != nil || task == nil {
		return nil, errors.New("cannot find the task with the task id")
	}

	// Kiểm tra xem người hoàn thành có phải là người được gán hay không
	if task.Assigned_to != "" && task.Assigned_to != userID {
		return nil, errors.New("you are not assined to this task")
	}

	// Cập nhật người hoàn thành, trạng thái và thời gian
	task.Finished_by = userID
	task.Status = "completed"
	task.UpdatedAt = time.Now()

	// Lưu thay đổi
	err = taskRepo.UpdateTask(context.Background(), task)
	if err != nil {
		return nil, err
	}

	// // Lấy thông tin dự án để thông báo cho người tạo dự án
	// projectRepo := repositories.NewProjectRepository(s.db)
	// project, err := projectRepo.FindProjectByID(context.Background(), task.ProjectID)
	// if err == nil && project != nil {
	// 	// Thông báo cho người tạo dự án
	// 	s.notificationService.SendNotification(project.CreatedBy, "Task Completed",
	// 		"Task: "+task.Description+" has been completed")
	// }

	// Trả về task đã cập nhật
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

// DeleteTask xóa một task khỏi database
// Input: taskID (string)
// Return: error (nếu có lỗi)
func (s *TaskService) DeleteTask(taskID string) error {
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
	return nil
}
