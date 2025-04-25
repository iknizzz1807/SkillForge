package services

import (
	"context"
	"math"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/utils"
)

type GamificationService struct {
	repo *repositories.GamificationRepository
}

func NewGamificationService(repo *repositories.GamificationRepository) *GamificationService {
	return &GamificationService{repo: repo}
}

// Calculate XP needed for next level using a formula: base * (multiplier ^ (level-1))
// Then round to nearest multiple of 5
func calculateXPNeeded(level int) int {
	rawXP := int(100 * math.Pow(1.5, float64(level-1)))
	return int(math.Round(float64(rawXP)/5.0) * 5)
}

// Calculate points needed for next skill level
func calculatePointsNeeded(level int) int {
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
	}
	return level, nil
}

func (s *GamificationService) AddXP(userID string, xp int) (*models.UserLevel, error) {
	ctx := context.Background()
	level, err := s.GetUserLevel(userID)
	if err != nil {
		return nil, err
	}

	level.XPCurrent += xp

	// Check for level up
	for level.XPCurrent >= level.XPNeed {
		level.XPCurrent -= level.XPNeed
		level.Level++
		level.XPNeed = calculateXPNeeded(level.Level)
	}

	err = s.repo.UpdateUserLevel(ctx, level)
	if err != nil {
		return nil, err
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
	}
	return skill, nil
}

func (s *GamificationService) AddSkillPoint(userID string, skillName string) (*models.UserSkill, error) {
	ctx := context.Background()
	skill, err := s.GetSkill(userID, skillName)
	if err != nil {
		return nil, err
	}

	skill.PointCurrent++

	// Check for skill level up
	if skill.PointCurrent >= skill.PointNeeded {
		skill.PointCurrent = 0
		skill.PointNeeded = calculatePointsNeeded(skill.PointNeeded/2 + 1) // Since PointNeeded = level * 2
	}

	err = s.repo.UpdateUserSkill(ctx, skill)
	if err != nil {
		return nil, err
	}

	return skill, nil
}

func (s *GamificationService) GetAllSkills(userID string) ([]models.UserSkill, error) {
	ctx := context.Background()
	return s.repo.GetAllUserSkills(ctx, userID)
}
