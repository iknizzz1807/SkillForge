package repositories

// Nhiệm vụ: Tương tác với collection "tasks" trong MongoDB (CRUD operations)
// Liên kết:
// - Dùng go.mongodb.org/mongo-driver/mongo để truy vấn database
// - Trả về models từ github.com/iknizzz1807/SkillForge/internal/models
// Vai trò trong flow:
// - Được gọi từ services/task_service.go để lưu/lấy task

import (
	"context"
	"time"
	"errors"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	// collection để truy cập collection "tasks" trong MongoDB
	collection *mongo.Collection
	activityCollection *mongo.Collection
}

// NewTaskRepository khởi tạo TaskRepository với database
// Input: db (*mongo.Database)
// Return: *TaskRepository - con trỏ đến TaskRepository
func NewTaskRepository(db *mongo.Database) *TaskRepository {
	return &TaskRepository{
		collection: db.Collection("tasks"),
		activityCollection: db.Collection("activities"),
	}
}

// InsertTasks thêm nhiều tasks mới vào collection "tasks"
// Input: ctx (context.Context), tasks ([]*models.Task)
// Return: error (nếu có lỗi)
func (r *TaskRepository) InsertTasks(ctx context.Context, tasks []*models.Task) error {
	// Kiểm tra danh sách tasks không rỗng
	if len(tasks) == 0 {
		return errors.New("danh sách tasks không được rỗng")
	}

	// Chuyển đổi danh sách tasks thành slice các interface{}
	documents := make([]interface{}, len(tasks))
	for i, task := range tasks {
		documents[i] = task
	}

	// Chèn nhiều tasks vào collection một lúc
	_, err := r.collection.InsertMany(ctx, documents)
	if err != nil {
		return err
	}

	// Trả về nil nếu thành công
	return nil
}

// FindTaskByProjectID tìm các task thuộc về một project cụ thể
// Input: ctx (context.Context), projectID (string)
// Return: []models.Task (danh sách task), error (nếu có lỗi)
func (r *TaskRepository) FindTasksByProjectID(ctx context.Context, projectID string) ([]*models.Task, error) {
	// Tạo filter để tìm các task thuộc về project
	filter := bson.M{"project_id": projectID}

	// Tạo mảng lưu kết quả
	var tasks []*models.Task

	// Truy vấn database
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Đọc tất cả các task tìm được vào mảng
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	// Nếu không tìm thấy task nào, trả về slice rỗng
	if len(tasks) == 0 {
		return []*models.Task{}, nil
	}

	// Trả về danh sách task
	return tasks, nil
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

// DeleteTask xóa task theo ID
// Input: ctx (context.Context), taskID (string)
// Return: error (nếu có lỗi)
func (r *TaskRepository) DeleteTask(ctx context.Context, taskID string) error {
	// Tạo filter để tìm theo ID
	filter := bson.M{"_id": taskID}

	// Xóa task từ collection
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	// Kiểm tra xem có bản ghi nào bị xóa không
	if result.DeletedCount == 0 {
		return errors.New("there is no task to delete")
	}

	return nil
}

func (r *TaskRepository) FindActivitiesByProjectID(ctx context.Context, projectID string) ([]*models.Activity, error) {
	// Tạo filter để tìm các hoạt động thuộc về project
	filter := bson.M{"project_id": projectID}

	// Tạo mảng lưu kết quả
	var activities []*models.Activity

	// Truy vấn database
	cursor, err := r.activityCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Đọc tất cả các hoạt động tìm được vào mảng
	if err := cursor.All(ctx, &activities); err != nil {
		return nil, err
	}

	// Nếu không tìm thấy hoạt động nào, trả về slice rỗng
	if len(activities) == 0 {
		return []*models.Activity{}, nil
	}

	// Trả về danh sách hoạt động
	return activities, nil
}

func (r *TaskRepository) InsertActivity(ctx context.Context, activity *models.Activity) error {
	// Chèn hoạt động vào collection
	activity.CreatedAt = time.Now()
	_, err := r.activityCollection.InsertOne(ctx, activity)
	if err != nil {
		return err
	}

	return nil
}