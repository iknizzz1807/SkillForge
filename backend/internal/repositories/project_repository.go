package repositories

// Nhiệm vụ: Tương tác với collection "projects" trong MongoDB (CRUD operations)
// Liên kết:
// - Dùng go.mongodb.org/mongo-driver/mongo để truy vấn database
// - Trả về models từ github.com/iknizzz1807/SkillForge/internal/models
// Vai trò trong flow:
// - Được gọi từ services/project_service.go để lưu/lấy project

import (
	"context"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectRepository struct {
	// collection để truy cập collection "projects" trong MongoDB
	collection *mongo.Collection
}

// NewProjectRepository khởi tạo ProjectRepository với database
// Input: db (*mongo.Database)
// Return: *ProjectRepository - con trỏ đến ProjectRepository
func NewProjectRepository(db *mongo.Database) *ProjectRepository {
	return &ProjectRepository{
		collection: db.Collection("projects"),
	}
}

// InsertProject thêm project mới vào collection "projects"
// Input: ctx (context.Context), project (*models.Project)
// Return: error (nếu có lỗi)
func (r *ProjectRepository) InsertProject(ctx context.Context, project *models.Project) error {
	// Chèn project vào collection
	_, err := r.collection.InsertOne(ctx, project)
	if err != nil {
		return err
	}
	// Trả về nil nếu thành công
	return nil
}

// FindProjectByID tìm project theo ID
// Input: ctx (context.Context), projectID (string)
// Return: *models.Project (project), error (nếu có lỗi)
func (r *ProjectRepository) FindProjectByID(ctx context.Context, projectID string) (*models.Project, error) {
	// Tạo filter để tìm theo ID
	filter := bson.M{"_id": projectID}
	var project models.Project

	// Truy vấn database
	err := r.collection.FindOne(ctx, filter).Decode(&project)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Không tìm thấy project
		}
		return nil, err
	}

	// Trả về project
	return &project, nil
}

// FindAllProjects lấy tất cả project
// Input: ctx (context.Context)
// Return: []*models.Project (danh sách project), error (nếu có lỗi)
func (r *ProjectRepository) FindAllProjects(ctx context.Context) ([]*models.Project, error) {
	// Tạo cursor để lấy tất cả document
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Tạo slice để lưu danh sách project
	var projects []*models.Project
	// Duyệt qua cursor và decode từng document
	for cursor.Next(ctx) {
		var project models.Project
		if err := cursor.Decode(&project); err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}

	// Trả về danh sách project
	return projects, nil
}

// UpdateProject cập nhật thông tin project
// Input: ctx (context.Context), project (*models.Project)
// Return: error (nếu có lỗi)
func (r *ProjectRepository) UpdateProject(ctx context.Context, project *models.Project) error {
	// Tạo filter để tìm project theo ID
	filter := bson.M{"_id": project.ID}
	// Tạo update payload
	update := bson.M{"$set": project}

	// Cập nhật trong database
	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	// Trả về nil nếu thành công
	return nil
}
