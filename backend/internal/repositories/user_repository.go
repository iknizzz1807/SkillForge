package repositories

// Nhiệm vụ: Tương tác với collection "users" trong MongoDB (CRUD operations)
// Liên kết:
// - Dùng go.mongodb.org/mongo-driver/mongo để truy vấn database
// - Trả về models từ github.com/iknizzz1807/SkillForge/internal/models
// Vai trò trong flow:
// - Được gọi từ services/auth_service.go và user_service.go để lưu/lấy user

import (
	"context"
	"errors"
	"time"

	// "errors"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	// collection để truy cập collection "users" trong MongoDB
	collection *mongo.Collection
}

// NewUserRepository khởi tạo UserRepository với database
// Input: db (*mongo.Database)
// Return: *UserRepository - con trỏ đến UserRepository
func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

// InsertUser thêm user mới vào collection "users"
// Input: ctx (context.Context), user (*models.User)
// Return: error (nếu có lỗi)
func (r *UserRepository) InsertUser(ctx context.Context, user *models.User) error {
	// Chèn user vào collection
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	// Trả về nil nếu thành công
	return nil
}

// FindUserByID tìm user theo ID
// Input: ctx (context.Context), userID (string)
// Return: *models.User (user), error (nếu có lỗi)
func (r *UserRepository) FindUserByID(ctx context.Context, userID string) (*models.User, error) {
	// Tạo filter để tìm theo ID
	filter := bson.M{"_id": userID}
	var user models.User

	// Truy vấn database
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Không tìm thấy user
		}
		return nil, err
	}

	// Trả về user
	return &user, nil
}

// FindUserByEmail tìm user theo email
// Input: ctx (context.Context), email (string)
// Return: *models.User (user), error (nếu có lỗi)
func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	// Tạo filter để tìm theo email
	filter := bson.M{"email": email}
	var user models.User

	// Truy vấn database
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Không tìm thấy user
		}
		return nil, err
	}

	// Trả về user
	return &user, nil
}

// UpdateUser cập nhật thông tin user
// Input: ctx (context.Context), user (*models.User)
// Return: error (nếu có lỗi)
func (r *UserRepository) UpdateUser(ctx context.Context, user *models.User) error {
	// Tạo filter để tìm user theo ID
	filter := bson.M{"_id": user.ID}
	// Tạo update payload
	update := bson.M{"$set": user}

	// Cập nhật trong database
	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	// Trả về nil nếu thành công
	return nil
}

// UpdateUserAvatarURL cập nhật đường dẫn avatar cho user (Đơn giản hóa)
// Input: ctx (context.Context), userID (string), avatarFilename (string) - tên file bao gồm extension
// Return: error (nếu có lỗi)
func (r *UserRepository) UpdateUserAvatarURL(ctx context.Context, userID string, avatarFilename string) error {
	filter := bson.M{"_id": userID}
	update := bson.M{
		"$set": bson.M{
			// Lưu trực tiếp tên file (ví dụ: "userID.png") vào trường avatar_url
			"avatar_url": avatarFilename,
			"updated_at": time.Now(),
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err // Trả về lỗi trực tiếp
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found for avatar update")
	}
	return nil
}
