package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	// "strings"
)

// BadgeService xử lý logic liên quan đến badge
type BadgeService struct {
	badgeRepo   *repositories.BadgeRepository
	userRepo    *repositories.UserRepository
	projectRepo *repositories.ProjectRepository
	taskRepo    *repositories.TaskRepository
	reviewRepo  *repositories.ReviewRepository
}

// NewBadgeService khởi tạo BadgeService với các repository cần thiết
func NewBadgeService(
	badgeRepo *repositories.BadgeRepository,
	// userRepo *repositories.UserRepository,
	// projectRepo *repositories.ProjectRepository,
	// taskRepo *repositories.TaskRepository,
	// reviewRepo *repositories.ReviewRepository,
) *BadgeService {
	return &BadgeService{
		badgeRepo: badgeRepo,
		// userRepo:    userRepo,
		// projectRepo: projectRepo,
		// taskRepo:    taskRepo,
		// reviewRepo:  reviewRepo,
	}
}

// CreateBadge tạo mới một badge
func (s *BadgeService) CreateBadge(
	name string,
	description string,
	badgeType string,
	prerequisites map[string]interface{},
) (*models.Badge, error) {
	ctx := context.Background()

	// Tạo badge mới
	badge := &models.Badge{
		ID:            uuid.New().String(),
		Name:          name,
		Description:   description,
		Type:          badgeType,
		Prerequisites: prerequisites,
		CreatedAt:     time.Now(),
	}

	// Lưu badge vào database
	err := s.badgeRepo.InsertBadge(ctx, badge)
	if err != nil {
		return nil, err
	}

	return badge, nil
}

// GetAllBadges lấy tất cả badge
func (s *BadgeService) GetAllBadges() ([]*models.Badge, error) {
	ctx := context.Background()
	return s.badgeRepo.FindAllBadges(ctx)
}

// GetBadgesByType lấy badges theo type
func (s *BadgeService) GetBadgesByType(badgeType string) ([]*models.Badge, error) {
	ctx := context.Background()
	return s.badgeRepo.FindBadgesByType(ctx, badgeType)
}

// GetBadgeByID lấy badge theo ID
func (s *BadgeService) GetBadgeByID(badgeID string) (*models.Badge, error) {
	ctx := context.Background()
	return s.badgeRepo.FindBadgeByID(ctx, badgeID)
}

// GetUserBadges lấy tất cả badge của user
func (s *BadgeService) GetUserBadges(userID string) ([]*models.UserBadge, error) {
	ctx := context.Background()
	return s.badgeRepo.FindUserBadges(ctx, userID)
}

// AwardBadgeToUser gán badge cho user
func (s *BadgeService) AwardBadgeToUser(
	userID string,
	badgeID string,
	relatedEntityID string,
) (*models.UserBadge, error) {
	ctx := context.Background()

	// Kiểm tra xem user đã có badge này chưa
	hasEarned, err := s.badgeRepo.HasUserEarnedBadge(ctx, userID, badgeID)
	if err != nil {
		return nil, err
	}
	if hasEarned {
		return nil, nil // User đã có badge này, không làm gì
	}

	// Lấy thông tin badge
	badge, err := s.badgeRepo.FindBadgeByID(ctx, badgeID)
	if err != nil {
		return nil, err
	}

	// Tạo user_badge
	userBadge := &models.UserBadge{
		ID:              uuid.New().String(),
		UserID:          userID,
		BadgeID:         badgeID,
		Badge:           badge,
		AwardedAt:       time.Now(),
		RelatedEntityID: relatedEntityID,
	}

	// Lưu user_badge vào database
	err = s.badgeRepo.AwardBadgeToUser(ctx, userBadge)
	if err != nil {
		return nil, err
	}

	return userBadge, nil
}

// CheckAndAwardProjectCompletionBadges kiểm tra và trao badge hoàn thành dự án
func (s *BadgeService) CheckAndAwardProjectCompletionBadges(userID string, projectID string) error {
	ctx := context.Background()

	// Lấy tất cả badge thuộc loại "completion"
	completionBadges, err := s.badgeRepo.FindBadgesByType(ctx, "completion")
	if err != nil {
		return err
	}

	// Đếm số dự án user đã hoàn thành
	// TODO: Triển khai hàm này trong ProjectRepository
	completedProjectCount := 1 // Giả sử đây là dự án đầu tiên

	// Kiểm tra từng badge
	for _, badge := range completionBadges {
		// Kiểm tra xem user đã đủ điều kiện nhận badge chưa
		if count, ok := badge.Prerequisites["completed_projects"].(float64); ok {
			if float64(completedProjectCount) >= count {
				// Trao badge cho user
				_, err := s.AwardBadgeToUser(userID, badge.ID, projectID)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// CheckAndAwardSkillBadges kiểm tra và trao badge kỹ năng
func (s *BadgeService) CheckAndAwardSkillBadges(userID string, skill string) error {
	ctx := context.Background()

	// Lấy tất cả badge thuộc loại "skill"
	skillBadges, err := s.badgeRepo.FindBadgesByType(ctx, "skill")
	if err != nil {
		return err
	}

	// Kiểm tra từng badge
	for _, badge := range skillBadges {
		// Kiểm tra xem badge có yêu cầu kỹ năng này không
		if badgeSkill, ok := badge.Prerequisites["skill"].(string); ok && badgeSkill == skill {
			// Trao badge cho user
			_, err := s.AwardBadgeToUser(userID, badge.ID, "")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// // CheckFirstApplicationBadge kiểm tra và trao badge cho lần đầu ứng tuyển
// func (s *BadgeService) CheckFirstApplicationBadge(userID string) {
//     ctx := context.Background()

//     // Tìm badge có code "FIRST_APPLICATION"
//     badge, err := s.badgeRepo.FindBadgeByCode(ctx, "FIRST_APPLICATION")
//     if err != nil || badge == nil {
//         return // Không tìm thấy badge hoặc có lỗi, thoát
//     }

//     // Đếm số lượng applications của user này
//     count, err := s.applicationRepo.CountApplicationsByUserID(ctx, userID)
//     if err != nil {
//         return
//     }

//     // Nếu đây là application đầu tiên
//     if count == 1 {
//         // Trao badge và không quan tâm kết quả
//         _, _ = s.AwardBadgeToUser(userID, badge.ID, "")

//         // Tạo thông báo về badge mới (nếu có NotificationService)
//         if s.notificationService != nil {
//             _ = s.notificationService.CreateNotification(userID, "badge",
//                 "Bạn đã nhận được huy hiệu mới: "+badge.Name,
//                 map[string]interface{}{"badge_id": badge.ID})
//         }
//     }
// }

// // CheckProjectCompletionBadge kiểm tra và trao badge khi hoàn thành dự án
// func (s *BadgeService) CheckProjectCompletionBadge(userID string) {
//     ctx := context.Background()

//     // Đếm số lượng dự án đã hoàn thành
//     completedCount, err := s.projectRepo.CountCompletedProjectsByMemberID(ctx, userID)
//     if err != nil {
//         return
//     }

//     // Danh sách badge cần kiểm tra
//     badgeCodes := []struct {
//         Code  string
//         Count int
//     }{
//         {"FIRST_PROJECT", 1},
//         {"FIVE_PROJECTS", 5},
//         {"TEN_PROJECTS", 10},
//     }

//     // Kiểm tra từng loại badge
//     for _, bc := range badgeCodes {
//         if completedCount >= bc.Count {
//             // Tìm badge theo code
//             badge, err := s.badgeRepo.FindBadgeByCode(ctx, bc.Code)
//             if err != nil || badge == nil {
//                 continue
//             }

//             // Trao badge nếu chưa có
//             _, _ = s.AwardBadgeToUser(userID, badge.ID, "")
//         }
//     }
// }

// // CheckSkillBadge kiểm tra và trao badge dựa trên kỹ năng sử dụng
// func (s *BadgeService) CheckSkillBadge(userID string, skill string) {
//     ctx := context.Background()

//     // Tìm badge tương ứng với kỹ năng
//     badgeCode := "SKILL_" + strings.ToUpper(skill)
//     badge, err := s.badgeRepo.FindBadgeByCode(ctx, badgeCode)
//     if err != nil || badge == nil {
//         return
//     }

//     // Trao badge
//     _, _ = s.AwardBadgeToUser(userID, badge.ID, "")
// }
