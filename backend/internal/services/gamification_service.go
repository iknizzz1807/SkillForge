package services

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/utils"
)

type GamificationService struct {
	repo        *repositories.GamificationRepository
	userService *UserService
}

func NewGamificationService(repo *repositories.GamificationRepository, userService *UserService) *GamificationService {
	return &GamificationService{
		repo:        repo,
		userService: userService,
	}
}

// Calculate XP needed for next level using a formula: base * (multiplier ^ (level-1))
// Then round to nearest multiple of 5
func calculateXPNeeded(level int) int {
	rawXP := int(100 * math.Pow(1.5, float64(level-1)))
	return int(math.Round(float64(rawXP)/5.0) * 5)
}

// Calculate points needed for next skill level
func calculatePointsNeeded(level int) int {
	if level == 1 {
		return 1
	}
	return level * 2 // Each level requires 2 more projects than previous
}

func (s *GamificationService) GetUserLevel(userID string) (*models.UserLevel, error) {
	ctx := context.Background()
	level, err := s.repo.GetUserLevel(ctx, userID)
	if err != nil {
		return nil, err
	}
	if level == nil {
		level = &models.UserLevel{
			ID:        utils.GenerateUUID(),
			UserID:    userID,
			Level:     1,
			XPCurrent: 0,
			XPNeed:    calculateXPNeeded(1),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		err = s.repo.UpdateUserLevel(ctx, level)
		if err != nil {
			return nil, err
		}
		return level, nil
	}
	return level, nil
}

func (s *GamificationService) AddXP(userID string, xp int) (*models.UserLevel, error) {
	ctx := context.Background()

	level, err := s.repo.AddXPAtomic(ctx, userID, xp, calculateXPNeeded(1))
	if err != nil {
		return nil, err
	}

	for level.XPCurrent >= level.XPNeed {
		level.XPCurrent -= level.XPNeed
		level.Level++
		level.XPNeed = calculateXPNeeded(level.Level)

		err = s.repo.LevelUpAtomically(ctx, userID, level.XPCurrent, level.Level, level.XPNeed)
		if err != nil {
			return nil, err
		}
	}

	return level, nil
}

func (s *GamificationService) GetSkill(userID string, skillName string) (*models.UserSkill, error) {
	ctx := context.Background()
	skill, err := s.repo.GetUserSkill(ctx, userID, skillName)
	if err != nil {
		return nil, err
	}
	if skill == nil {
		skill = &models.UserSkill{
			ID:           utils.GenerateUUID(),
			UserID:       userID,
			Skill:        skillName,
			PointCurrent: 0,
			PointNeeded:  calculatePointsNeeded(1),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
		err = s.repo.UpdateUserSkill(ctx, skill)
		if err != nil {
			return nil, err
		}
		return skill, nil
	}
	return skill, nil
}

func (s *GamificationService) AddSkillPoint(userID string, skillName string) (*models.UserSkill, error) {
	ctx := context.Background()

	existingSkill, _ := s.repo.GetUserSkill(ctx, userID, skillName)
	isNew := existingSkill == nil

	skill, err := s.repo.AddSkillPointAtomic(ctx, userID, skillName, calculatePointsNeeded(1))
	if err != nil {
		return nil, err
	}

	for skill.PointCurrent >= skill.PointNeeded {
		skill.PointCurrent -= skill.PointNeeded
		skill.Level++
		skill.PointNeeded = calculatePointsNeeded(skill.Level)

		err = s.repo.SkillLevelUpAtomically(ctx, userID, skillName, skill.PointCurrent, skill.Level, skill.PointNeeded)
		if err != nil {
			return nil, err
		}
	}

	if isNew {
		user, err := s.userService.GetUserByID(userID)
		if err != nil || user == nil {
			return nil, errors.New("user not found")
		}
		newSkills := append(user.Skills, skillName)
		_, err = s.userService.UpdateUserSkills(userID, newSkills)
		if err != nil {
			return nil, err
		}
	}

	return skill, nil
}

func (s *GamificationService) GetAllSkills(userID string) ([]models.UserSkill, error) {
	ctx := context.Background()
	return s.repo.GetAllUserSkills(ctx, userID)
}
