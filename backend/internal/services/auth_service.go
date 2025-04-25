// internal/services/auth_service.go
package services

import (
	"context"
	"errors"
	"fmt"            // Thêm fmt để log lỗi nhỏ
	"mime/multipart" // Thêm multipart
	"strings"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/utils" // Giữ lại utils cho GenerateJWT và HashPassword
)

type AuthService struct {
	userRepo    *repositories.UserRepository
	fileService *FileService // Thêm FileService
}

// NewAuthService khởi tạo AuthService (cập nhật để nhận FileService)
func NewAuthService(userRepo *repositories.UserRepository, fileService *FileService) *AuthService {
	return &AuthService{
		userRepo:    userRepo,
		fileService: fileService, // Khởi tạo fileService
	}
}

// Register đăng ký người dùng mới, bao gồm cả xử lý avatar (nếu có)
// Input: email, name, password, role (string), file (multipart.File), header (*multipart.FileHeader) - file và header có thể nil
// Return: *models.User (user vừa tạo, có thể có AvatarURL), string (token JWT), error
func (s *AuthService) Register(email, name, password, role, website string, file multipart.File, header *multipart.FileHeader) (*models.User, string, error) {
	ctx := context.Background()

	// 1. Kiểm tra email tồn tại
	existingUser, _ := s.userRepo.FindUserByEmail(ctx, email)
	if existingUser != nil {
		return nil, "", errors.New("email already exists")
	}

	// 2. Tạo user mới (chưa có AvatarURL)
	user := &models.User{
		ID:        utils.GenerateUUID(), // Tạo ID duy nhất
		Email:     email,
		Name:      name,
		Password:  utils.HashPassword(password),
		Role:      strings.ToLower(role),
		Website:   website,
		CreatedAt: time.Now(),
		// AvatarURL sẽ được cập nhật sau nếu có file
	}

	// 3. Lưu user vào database
	err := s.userRepo.InsertUser(ctx, user)
	if err != nil {
		fmt.Printf("Error inserting user %s: %v\n", user.Email, err)
		return nil, "", fmt.Errorf("could not register user: %w", err)
	}
	fmt.Printf("Successfully inserted user %s with ID %s\n", user.Email, user.ID)

	// 4. Xử lý upload avatar NẾU có file được cung cấp
	var avatarFilename string = "" // Khởi tạo rỗng
	if file != nil && header != nil {
		fmt.Printf("Attempting to save avatar for user %s\n", user.ID)
		// Gọi FileService để lưu avatar, sử dụng user.ID vừa tạo
		savedFilename, err := s.fileService.SaveAvatar(user.ID, file, header)
		if err != nil {
			fmt.Printf("Warning: Failed to save avatar for user %s during registration: %v\n", user.ID, err)
			// Không return lỗi ở đây
		} else {
			fmt.Printf("Avatar saved for user %s with filename %s\n", user.ID, savedFilename)
			avatarFilename = savedFilename // Lưu lại tên file để cập nhật DB

			// Cập nhật thông tin user model
			user.AvatarName = avatarFilename
			user.UpdatedAt = time.Now()

			errUpdate := s.userRepo.UpdateUser(ctx, user)
			if errUpdate != nil {
				fmt.Printf("Warning: Failed to update user in DB for user %s after saving avatar: %v\n", user.ID, errUpdate)
			} else {
				fmt.Printf("Successfully updated user with avatar for user %s\n", user.ID)

				// Thêm log để kiểm tra lưu trữ avatar
				avatarPath, _ := s.fileService.GetAvatarFilePath(avatarFilename)
				fmt.Printf("Avatar stored at: %s\n", avatarPath)
			}
		}
	} else {
		fmt.Printf("No avatar file provided for user %s during registration.\n", user.ID)
	}

	// 5. Tạo JWT token (luôn tạo kể cả khi avatar lỗi)
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Name, user.Role)
	if err != nil {
		// Lỗi này nghiêm trọng hơn, vì user không thể đăng nhập
		fmt.Printf("Error generating JWT for user %s: %v\n", user.ID, err)
		// Có thể xem xét xóa user vừa tạo ở đây nếu muốn roll back hoàn toàn
		return nil, "", fmt.Errorf("could not generate authentication token: %w", err)
	}

	// 6. Trả về user (đã có thể cập nhật AvatarURL) và token
	return user, token, nil
}

// Login (Giữ nguyên không đổi)
func (s *AuthService) Login(email, password string) (string, error) {
	ctx := context.Background()
	user, err := s.userRepo.FindUserByEmail(ctx, email)
	if err != nil || user == nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.VerifyPassword(user.Password, password) {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.ID, user.Email, user.Name, user.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}
