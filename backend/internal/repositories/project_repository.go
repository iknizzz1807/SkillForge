package repositories

// Nhiệm vụ: Tương tác với collection "projects" trong MongoDB (CRUD operations)
// Liên kết:
// - Dùng go.mongodb.org/mongo-driver/mongo để truy vấn database
// - Trả về models từ github.com/iknizzz1807/SkillForge/internal/models
// Vai trò trong flow:
// - Được gọi từ services/project_service.go để lưu/lấy project

import (
	"context"

	"time"

	"errors"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	// Check if the ID is in ObjectID format
	var filter bson.M
	if isValidObjectID(projectID) {
		objID, _ := primitive.ObjectIDFromHex(projectID)
		filter = bson.M{"_id": objID}
	} else {
		// If not a valid ObjectID, try using the string directly
		filter = bson.M{"_id": projectID}
	}

	var project models.Project

	// Truy vấn database
	err := r.collection.FindOne(ctx, filter).Decode(&project)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("project not found")
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

// UpdateProject cập nhật thông tin project trong database
// Input: project (*models.Project) - project với thông tin đã được cập nhật
// Return:
//   - *models.Project - thông tin project sau khi cập nhật từ database
//   - error - lỗi nếu có
func (r *ProjectRepository) UpdateProject(project *models.Project) (*models.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if the ID is in ObjectID format
	var filter bson.M
	if isValidObjectID(project.ID) {
		objID, _ := primitive.ObjectIDFromHex(project.ID)
		filter = bson.M{"_id": objID}
	} else {
		// If not a valid ObjectID, use the string directly
		filter = bson.M{"_id": project.ID}
	}

	// Chuẩn bị dữ liệu cập nhật
	update := bson.M{
		"$set": bson.M{
			"title":       project.Title,
			"description": project.Description,
			"skills":      project.Skills,
			"start_time":  project.StartTime,
			"end_time":    project.EndTime,
			"max_member":  project.MaxMember,
			"status":      project.Status,
		},
	}

	// Thực hiện cập nhật
	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	// Fetch và trả về project đã cập nhật
	updatedProject, err := r.FindProjectByID(ctx, project.ID)
	if err != nil {
		return nil, err
	}

	return updatedProject, nil
}

// GetProjectIDsByCreatorID lấy tất cả projectID được tạo bởi một creator (business)
// Input: businessID (string) - ID của business user
// Return: []string (danh sách project ID), error (nếu có lỗi)
func (r *ProjectRepository) GetProjectIDsByCreatorID(businessID string) ([]string, error) {
	// Tạo context với timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Tạo filter tìm theo creator_id
	filter := bson.M{"created_by_id": businessID}

	// Chỉ lấy trường _id từ các documents
	projection := bson.M{"_id": 1}

	// Truy vấn database
	cursor, err := r.collection.Find(ctx, filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Tạo slice để lưu danh sách project ID
	var projectIDs []string

	// Duyệt qua cursor và lấy ID của từng project
	for cursor.Next(ctx) {
		var result struct {
			ID string `bson:"_id"`
		}
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		projectIDs = append(projectIDs, result.ID)
	}

	// Kiểm tra lỗi từ cursor
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Trả về danh sách project ID
	return projectIDs, nil
}

// DeleteProject xóa một project từ database
// Input: projectID (string) - ID của project cần xóa
// Return: error - lỗi nếu có
func (r *ProjectRepository) DeleteProject(projectID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if the ID is in ObjectID format
	var filter bson.M
	if isValidObjectID(projectID) {
		objID, _ := primitive.ObjectIDFromHex(projectID)
		filter = bson.M{"_id": objID}
	} else {
		// If not a valid ObjectID, try using the string directly
		filter = bson.M{"_id": projectID}
	}

	// Thực hiện xóa
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("no project found with the given ID")
	}

	return nil
}

// isValidObjectID kiểm tra xem một chuỗi có phải là ObjectID hợp lệ hay không
func isValidObjectID(id string) bool {
	_, err := primitive.ObjectIDFromHex(id)
	return err == nil
}
