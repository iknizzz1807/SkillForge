package services

// Nhiệm vụ: Xử lý logic xác thực (đăng ký, đăng nhập) và tạo token JWT
// Liên kết:
// - Dùng internal/repositories/user_repository.go để lưu/lấy user
// - Dùng internal/utils/jwt.go để tạo token
// Vai trò trong flow:
// - Được gọi từ handlers/auth.go để đăng ký hoặc đăng nhập người dùng
// - Tạo token JWT để xác thực các request sau này

import (
	"errors"
	"time"

	"context"
	"strings"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/utils"
)

type AuthService struct {
	// userRepo để tương tác với collection users trong MongoDB
	userRepo *repositories.UserRepository
}

// NewAuthService khởi tạo AuthService với dependency userRepo
// Input: userRepo (*repositories.UserRepository)
// Return: *AuthService - con trỏ đến AuthService
func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{userRepo}
}

// Register đăng ký người dùng mới
// Input: email (string), password (string), role (string)
// Return: *models.User (user vừa tạo), string (token JWT), error (nếu có lỗi)
func (s *AuthService) Register(email, name, password, role string) (*models.User, string, error) {
	// Kiểm tra email đã tồn tại chưa
	existingUser, _ := s.userRepo.FindUserByEmail(context.Background(), email)
	if existingUser != nil {
		return nil, "", errors.New("email already exists")
	}

	// Tạo user mới
	user := &models.User{
		ID:        utils.GenerateUUID(), // Tạo ID duy nhất
		Email:     email,
		Name:      name,
		Password:  utils.HashPassword(password), // Hash password trước khi lưu
		Role:      strings.ToLower(role),
		CreatedAt: time.Now(),
	}

	// Lưu user vào database
	err := s.userRepo.InsertUser(context.Background(), user)
	if err != nil {
		return nil, "", err
	}

	// Tạo JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Name, user.Role)
	if err != nil {
		return nil, "", err
	}

	// Trả về user và token
	return user, token, nil
}

// Login đăng nhập người dùng
// Input: email (string), password (string)
// Return: string (token JWT), error (nếu có lỗi)
func (s *AuthService) Login(email, password string) (string, error) {
	// Tìm user theo email
	user, err := s.userRepo.FindUserByEmail(context.Background(), email)
	if err != nil || user == nil {
		return "", errors.New("invalid credentials")
	}

	// Kiểm tra password
	if !utils.VerifyPassword(user.Password, password) {
		return "", errors.New("invalid credentials")
	}

	// Tạo JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Name, user.Role)
	if err != nil {
		return "", err
	}

	// Trả về token
	return token, nil
}
