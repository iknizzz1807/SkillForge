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

	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if email == "" {
		return nil, errors.New("email cannot be empty")
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

	if name == "" {
		name = user.Name
	}
	if email == "" {
		email = user.Email
	}
	if title == "" {
		title = user.Title
	}
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

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

// UpdateUserSkills updates the skills array for a user
func (s *UserService) UpdateUserSkills(userID string, skills []string) (*models.User, error) {
	if userID == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	// Validate skills
	if len(skills) == 0 {
		return nil, errors.New("skills array cannot be empty")
	}

	// Create repository and update skills
	userRepo := repositories.NewUserRepository(s.db)
	err := userRepo.UpdateUserSkills(context.Background(), userID, skills)
	if err != nil {
		return nil, err
	}

	// Fetch and return updated user
	return s.GetUserByID(userID)
}

func (s *UserService) GetAllStudents(page, limit int, skill string) ([]*models.User, int64, error) {
	userRepo := repositories.NewUserRepository(s.db)
	students, err := userRepo.FindAllStudents(context.Background(), page, limit, skill)
	if err != nil {
		return nil, 0, err
	}
	total, err := userRepo.CountAllStudents(context.Background(), skill)
	if err != nil {
		return nil, 0, err
	}
	for _, student := range students {
		student.Password = ""
	}
	return students, total, nil
}

type UserProfile struct {
	User      *models.User        `json:"user"`
	Badges    []*models.UserBadge `json:"badges"`
	Feedbacks []*models.Feedback  `json:"feedbacks"`
	Projects  []*models.Project   `json:"projects"`
}

func (s *UserService) GetUserProfile(userID string) (*UserProfile, error) {
	if userID == "" {
		return nil, errors.New("user ID cannot be empty")
	}
	ctx := context.Background()

	userRepo := repositories.NewUserRepository(s.db)
	user, err := userRepo.FindUserByID(ctx, userID)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}
	user.Password = ""

	badgeRepo := repositories.NewBadgeRepository(s.db)
	badges, err := badgeRepo.FindUserBadges(ctx, userID)
	if err != nil {
		badges = []*models.UserBadge{}
	}
	if badges == nil {
		badges = []*models.UserBadge{}
	}

	feedbackRepo := repositories.NewFeedbackRepository(s.db)
	feedbacks, err := feedbackRepo.GetFeedbacksForUser(ctx, userID, "student")
	if err != nil {
		feedbacks = []*models.Feedback{}
	}
	if feedbacks == nil {
		feedbacks = []*models.Feedback{}
	}

	projectStudentRepo := repositories.NewProjectStudentRepository(s.db)
	projectIDs, err := projectStudentRepo.FindProjectsByStudentID(ctx, userID)
	projects := []*models.Project{}
	if err == nil {
		projectRepo := repositories.NewProjectRepository(s.db)
		foundProjects, err := projectRepo.FindProjectsByIDs(ctx, projectIDs)
		if err == nil {
			projects = foundProjects
		}
	}
	if projects == nil {
		projects = []*models.Project{}
	}

	return &UserProfile{
		User:      user,
		Badges:    badges,
		Feedbacks: feedbacks,
		Projects:  projects,
	}, nil
}

// GetUserRepository returns the user repository instance
// This is a helper to allow FileService to access the user repository
func (s *UserService) GetUserRepository() *repositories.UserRepository {
	return repositories.NewUserRepository(s.db)
}
