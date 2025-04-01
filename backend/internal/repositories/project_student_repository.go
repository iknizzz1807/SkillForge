package repositories

// Nhiệm vụ: Tương tác với collection "project_student" trong MongoDB (CRUD operations)
// Liên kết:
// - Dùng go.mongodb.org/mongo-driver/mongo để truy vấn database
// - Trả về models từ github.com/iknizzz1807/SkillForge/internal/models
// Vai trò trong flow:
// - Được gọi từ services/project_service.go để lưu/lấy quan hệ giữa project và student

import (
	"context"
	"errors"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectStudentRepository struct {
	// collection để truy cập collection "project_student" trong MongoDB
	collection *mongo.Collection
}

// NewProjectStudentRepository khởi tạo ProjectStudentRepository với database
// Input: db (*mongo.Database)
// Return: *ProjectStudentRepository - con trỏ đến ProjectStudentRepository
func NewProjectStudentRepository(db *mongo.Database) *ProjectStudentRepository {
	return &ProjectStudentRepository{
		collection: db.Collection("project_student"),
	}
}

// InsertProjectStudent thêm quan hệ project-student mới vào collection
// Input: ctx (context.Context), projectStudent (*models.Project_student)
// Return: error (nếu có lỗi)
func (r *ProjectStudentRepository) InsertProjectStudent(ctx context.Context, projectStudent *models.Project_student) error {
	// Kiểm tra xem quan hệ đã tồn tại chưa
	existing, err := r.FindByProjectAndStudent(ctx, projectStudent.Project_id, projectStudent.Student_id)
	if err == nil && existing != nil {
		return errors.New("student already joined this project")
	}

	// Chèn quan hệ vào collection
	_, err = r.collection.InsertOne(ctx, projectStudent)
	if err != nil {
		return err
	}

	// Trả về nil nếu thành công
	return nil
}

// FindByProjectAndStudent tìm quan hệ project-student theo project ID và student ID
// Input: ctx (context.Context), projectID (string), studentID (string)
// Return: *models.Project_student (quan hệ), error (nếu có lỗi)
func (r *ProjectStudentRepository) FindByProjectAndStudent(ctx context.Context, projectID, studentID string) (*models.Project_student, error) {
	var projectStudent models.Project_student

	// Tạo filter tìm theo project_id và student_id
	filter := bson.M{
		"project_id": projectID,
		"student_id": studentID,
	}

	// Truy vấn database
	err := r.collection.FindOne(ctx, filter).Decode(&projectStudent)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Không tìm thấy, trả về nil
		}
		return nil, err
	}

	// Trả về quan hệ
	return &projectStudent, nil
}

// FindProjectsByStudentID tìm tất cả project ID mà student tham gia
// Input: ctx (context.Context), studentID (string)
// Return: []string (danh sách project ID), error (nếu có lỗi)
func (r *ProjectStudentRepository) FindProjectsByStudentID(ctx context.Context, studentID string) ([]string, error) {
	// Tạo filter tìm theo student_id
	filter := bson.M{"student_id": studentID}

	// Truy vấn database
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Tạo slice để lưu danh sách project ID
	var projectIDs []string

	// Duyệt qua cursor và lấy project_id từ từng document
	for cursor.Next(ctx) {
		var projectStudent models.Project_student
		if err := cursor.Decode(&projectStudent); err != nil {
			return nil, err
		}
		projectIDs = append(projectIDs, projectStudent.Project_id)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Trả về danh sách project ID
	return projectIDs, nil
}

// FindStudentsByProjectID tìm tất cả student ID tham gia vào project
// Input: ctx (context.Context), projectID (string)
// Return: []string (danh sách student ID), error (nếu có lỗi)
func (r *ProjectStudentRepository) FindStudentsByProjectID(ctx context.Context, projectID string) ([]string, error) {
	// Tạo filter tìm theo project_id
	filter := bson.M{"project_id": projectID}

	// Truy vấn database
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Tạo slice để lưu danh sách student ID
	var studentIDs []string

	// Duyệt qua cursor và lấy student_id từ từng document
	for cursor.Next(ctx) {
		var projectStudent models.Project_student
		if err := cursor.Decode(&projectStudent); err != nil {
			return nil, err
		}
		studentIDs = append(studentIDs, projectStudent.Student_id)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Trả về danh sách student ID
	return studentIDs, nil
}

// DeleteProjectStudent xóa quan hệ project-student
// Input: ctx (context.Context), projectID (string), studentID (string)
// Return: error (nếu có lỗi)
func (r *ProjectStudentRepository) DeleteProjectStudent(ctx context.Context, projectID, studentID string) error {
	// Tạo filter tìm theo project_id và student_id
	filter := bson.M{
		"project_id": projectID,
		"student_id": studentID,
	}

	// Xóa document
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("no record found for this project and student")
	}

	// Trả về nil nếu thành công
	return nil
}
