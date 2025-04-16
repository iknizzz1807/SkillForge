// internal/repositories/application_repository.go
package repositories

// Nhiệm vụ: Tương tác với collection "applications" trong MongoDB (CRUD operations)
// Liên kết:
// - Dùng go.mongodb.org/mongo-driver/mongo để truy vấn database
// - Trả về models từ github.com/iknizzz1807/SkillForge/internal/models
// Vai trò trong flow:
// - Được gọi từ services/application_service.go để lưu/lấy/cập nhật ứng tuyển

import (
	"context"
	"errors"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/models" // Import utils nếu cần logger
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options" // Import options nếu cần
)

type ApplicationRepository struct {
	// collection để truy cập collection "applications" trong MongoDB
	collection *mongo.Collection
}

// NewApplicationRepository khởi tạo ApplicationRepository với database
// Input: db (*mongo.Database)
// Return: *ApplicationRepository - con trỏ đến ApplicationRepository
func NewApplicationRepository(db *mongo.Database) *ApplicationRepository {
	return &ApplicationRepository{
		collection: db.Collection("applications"), // Tên collection thường là số nhiều
	}
}

// InsertApplication thêm ứng tuyển mới vào collection "applications"
// Input: ctx (context.Context), application (*models.Application)
// Return: error (nếu có lỗi)
func (r *ApplicationRepository) InsertApplication(ctx context.Context, application *models.Application) error {
	if ctx == nil {
		ctx = context.Background() // Đảm bảo context không nil
	}
	// Chèn application vào collection
	_, err := r.collection.InsertOne(ctx, application)
	if err != nil {
		// Kiểm tra lỗi duplicate key nếu có index phù hợp
		if mongo.IsDuplicateKeyError(err) {
			return errors.New("application creation failed due to duplicate key")
		}
		// utils.GetLogger().Errorf("Error inserting application: %v", err) // Log lỗi
		return err
	}
	// Trả về nil nếu thành công
	return nil
}

// FindApplicationByID tìm ứng tuyển theo ID
// Input: ctx (context.Context), applicationID (string)
// Return: *models.Application (ứng tuyển), error (nếu có lỗi)
func (r *ApplicationRepository) FindApplicationByID(ctx context.Context, applicationID string) (*models.Application, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	// Tạo filter để tìm theo ID (_id là trường mặc định của MongoDB)
	filter := bson.M{"_id": applicationID}
	var application models.Application

	// Truy vấn database
	err := r.collection.FindOne(ctx, filter).Decode(&application)
	if err != nil {
		// Trả về lỗi chuẩn mongo.ErrNoDocuments khi không tìm thấy
		// Service layer sẽ xử lý lỗi này thành "application not found"
		// if err == mongo.ErrNoDocuments {
		// 	return nil, nil // Hoặc trả về lỗi cụ thể: return nil, errors.New("application not found")
		// }
		// utils.GetLogger().Warnf("Error finding application by ID %s: %v", applicationID, err) // Log warning nếu không phải ErrNoDocuments
		return nil, err // Trả về lỗi gốc
	}

	// Trả về application
	return &application, nil
}

// --- NEW FUNCTION ---
// FindByUserAndProject tìm application dựa trên userID và projectID
// Input: ctx (context.Context), userID (string), projectID (string)
// Return: *models.Application, error
func (r *ApplicationRepository) FindByUserAndProject(ctx context.Context, userID, projectID string) (*models.Application, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	filter := bson.M{
		"user_id":    userID,
		"project_id": projectID,
	}
	var application models.Application
	err := r.collection.FindOne(ctx, filter).Decode(&application)
	if err != nil {
		// if err == mongo.ErrNoDocuments { // Không tìm thấy là trường hợp bình thường
		// 	return nil, nil
		// }
		// Log lỗi khác
		if err != mongo.ErrNoDocuments {
			// utils.GetLogger().Warnf("Error finding application by user %s and project %s: %v", userID, projectID, err)
		}
		return nil, err // Trả về lỗi gốc (bao gồm cả ErrNoDocuments)
	}
	return &application, nil
}

// --- END NEW FUNCTION ---

// FindByUserID tìm tất cả applications của một user (student)
func (r *ApplicationRepository) FindByUserID(userID string) ([]models.Application, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Thêm timeout
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		// utils.GetLogger().Errorf("Error finding applications by user ID %s: %v", userID, err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var applications []models.Application
	if err := cursor.All(ctx, &applications); err != nil {
		// utils.GetLogger().Errorf("Error decoding applications for user ID %s: %v", userID, err)
		return nil, err
	}

	// Không cần trả về nil nếu slice rỗng, trả về slice rỗng là đúng
	// if applications == nil {
	//     return []models.Application{}, nil
	// }
	return applications, nil
}

// FindByProjectIDs tìm tất cả applications cho danh sách các projectID
func (r *ApplicationRepository) FindByProjectIDs(projectIDs []string) ([]models.Application, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Xử lý trường hợp projectIDs rỗng để tránh query không cần thiết
	if len(projectIDs) == 0 {
		return []models.Application{}, nil
	}

	cursor, err := r.collection.Find(ctx, bson.M{"project_id": bson.M{"$in": projectIDs}})
	if err != nil {
		// utils.GetLogger().Errorf("Error finding applications by project IDs: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var applications []models.Application
	if err := cursor.All(ctx, &applications); err != nil {
		// utils.GetLogger().Errorf("Error decoding applications for project IDs: %v", err)
		return nil, err
	}
	return applications, nil
}

// UpdateStatus cập nhật trạng thái application và thời gian updated_at
// Input: id (string), status (string)
// Return: *models.Application (application đã cập nhật), error
func (r *ApplicationRepository) UpdateStatus(id string, status string) (*models.Application, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"status":     status,
			"updated_at": time.Now(), // Luôn cập nhật updated_at khi status thay đổi
		},
	}

	// Sử dụng FindOneAndUpdate để cập nhật và trả về document mới
	options := options.FindOneAndUpdate().SetReturnDocument(options.After) // Trả về document sau khi update

	var updatedApplication models.Application
	err := r.collection.FindOneAndUpdate(ctx, filter, update, options).Decode(&updatedApplication)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) { // Lỗi nếu không tìm thấy application để update
			return nil, errors.New("application not found for update")
		}
		// utils.GetLogger().Errorf("Error updating application status for ID %s: %v", id, err)
		return nil, err
	}

	return &updatedApplication, nil // Trả về application đã được cập nhật
}

// Delete xóa một application theo ID
// Input: id (string)
// Return: error
func (r *ApplicationRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		// utils.GetLogger().Errorf("Error deleting application with ID %s: %v", id, err)
		return err
	}

	// Kiểm tra xem có document nào thực sự bị xóa không
	if result.DeletedCount == 0 {
		// utils.GetLogger().Warnf("Attempted to delete non-existent application with ID %s", id)
		return errors.New("application not found for deletion") // Trả lỗi rõ ràng hơn
	}

	return nil
}
