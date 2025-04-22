package services

// Nhiệm vụ: Xử lý logic quản lý thông tin người dùng (lấy, cập nhật)
// Liên kết:
// - Dùng internal/repositories/user_repository.go để tương tác với database
// Vai trò trong flow:
// - Được gọi từ handlers/users.go để lấy hoặc cập nhật thông tin user

import (
	"errors"
	"time"

	"context"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	// db để truy cập MongoDB database
	db *mongo.Database
}

// NewUserService khởi tạo UserService với dependency db
// Input: db (*mongo.Database)
// Return: *UserService - con trỏ đến UserService
func NewUserService(db *mongo.Database) *UserService {
	return &UserService{db}
}

// GetUserByID lấy thông tin user theo ID
// Input: userID (string)
// Return: *models.User (user), error (nếu có lỗi)
func (s *UserService) GetUserByID(userID string) (*models.User, error) {
	// Kiểm tra userID hợp lệ
	if userID == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	// Tạo repository và tìm user
	userRepo := repositories.NewUserRepository(s.db)
	user, err := userRepo.FindUserByID(context.Background(), userID)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	// Trả về user
	return user, nil
}

// UpdateUser cập nhật thông tin user
// Input: userID (string), skills ([]string)
// Return: *models.User (user đã cập nhật), error (nếu có lỗi)
func (s *UserService) UpdateUser(userID, name, email, title string) (*models.User, error) {
	// Kiểm tra userID hợp lệ
	if userID == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	// Tạo repository và tìm user
	userRepo := repositories.NewUserRepository(s.db)
	user, err := userRepo.FindUserByID(context.Background(), userID)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	user.Name = name
	user.Email = email
	user.Title = title
	user.UpdatedAt = time.Now()

	// Lưu thay đổi vào database
	err = userRepo.UpdateUser(context.Background(), user)
	if err != nil {
		return nil, err
	}

	// Trả về user đã cập nhật
	return user, nil
}

// UpdateUserWithAvatar cập nhật thông tin user bao gồm avatar
// Input: userID (string), name (string), email (string), avatarFilename (string)
// Return: *models.User (user đã cập nhật), error (nếu có lỗi)
func (s *UserService) UpdateUserWithAvatar(userID, name, email, title, avatarFilename string) (*models.User, error) {
	// Kiểm tra userID hợp lệ
	if userID == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	// Tạo repository và tìm user
	userRepo := repositories.NewUserRepository(s.db)
	user, err := userRepo.FindUserByID(context.Background(), userID)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	// Cập nhật thông tin user
	user.Name = name
	user.Email = email
	user.Title = title
	user.AvatarName = avatarFilename
	user.UpdatedAt = time.Now()

	// Lưu thay đổi vào database
	err = userRepo.UpdateUser(context.Background(), user)
	if err != nil {
		return nil, err
	}

	// Trả về user đã cập nhật
	return user, nil
}

// GetUserRepository returns the user repository instance
// This is a helper to allow FileService to access the user repository
func (s *UserService) GetUserRepository() *repositories.UserRepository {
	return repositories.NewUserRepository(s.db)
}
