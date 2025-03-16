package repositories

// Nhiệm vụ: Tương tác với collection "tasks" trong MongoDB (CRUD operations)
// Liên kết:
// - Dùng go.mongodb.org/mongo-driver/mongo để truy vấn database
// - Trả về models từ github.com/iknizzz1807/SkillForge/internal/models
// Vai trò trong flow:
// - Được gọi từ services/task_service.go để lưu/lấy task

import (
	"context"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	// collection để truy cập collection "tasks" trong MongoDB
	collection *mongo.Collection
}

// NewTaskRepository khởi tạo TaskRepository với database
// Input: db (*mongo.Database)
// Return: *TaskRepository - con trỏ đến TaskRepository
func NewTaskRepository(db *mongo.Database) *TaskRepository {
	return &TaskRepository{
		collection: db.Collection("tasks"),
	}
}

// InsertTask thêm task mới vào collection "tasks"
// Input: ctx (context.Context), task (*models.Task)
// Return: error (nếu có lỗi)
func (r *TaskRepository) InsertTask(ctx context.Context, task *models.Task) error {
	// Chèn task vào collection
	_, err := r.collection.InsertOne(ctx, task)
	if err != nil {
		return err
	}
	// Trả về nil nếu thành công
	return nil
}

// FindTaskByID tìm task theo ID
// Input: ctx (context.Context), taskID (string)
// Return: *models.Task (task), error (nếu có lỗi)
func (r *TaskRepository) FindTaskByID(ctx context.Context, taskID string) (*models.Task, error) {
	// Tạo filter để tìm theo ID
	filter := bson.M{"_id": taskID}
	var task models.Task

	// Truy vấn database
	err := r.collection.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Không tìm thấy task
		}
		return nil, err
	}

	// Trả về task
	return &task, nil
}

// UpdateTask cập nhật thông tin task
// Input: ctx (context.Context), task (*models.Task)
// Return: error (nếu có lỗi)
func (r *TaskRepository) UpdateTask(ctx context.Context, task *models.Task) error {
	// Tạo filter để tìm task theo ID
	filter := bson.M{"_id": task.ID}
	// Tạo update payload
	update := bson.M{"$set": task}

	// Cập nhật trong database
	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	// Trả về nil nếu thành công
	return nil
}
