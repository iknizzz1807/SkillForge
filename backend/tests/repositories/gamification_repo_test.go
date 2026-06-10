package repositories_test

import (
	"context"
	"testing"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"github.com/iknizzz1807/SkillForge/tests/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGamificationRepo_GetUserLevel_CreatesNew(t *testing.T) {
	userID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "user_levels") })

	repo := repositories.NewGamificationRepository(testutil.GetTestDB())
	level, err := repo.GetUserLevel(context.Background(), userID)
	require.NoError(t, err)
	assert.Nil(t, level)
}

func TestGamificationRepo_UpdateUserLevel_Creates(t *testing.T) {
	userID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "user_levels") })

	repo := repositories.NewGamificationRepository(testutil.GetTestDB())
	level := &models.UserLevel{
		ID:        utils.GenerateUUID(),
		UserID:    userID,
		Level:     1,
		XPNeed:    100,
		XPCurrent: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.UpdateUserLevel(context.Background(), level)
	require.NoError(t, err)

	fetched, err := repo.GetUserLevel(context.Background(), userID)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 1, fetched.Level)
	assert.Equal(t, 0, fetched.XPCurrent)
}

func TestGamificationRepo_UpdateUserLevel_UpdatesExisting(t *testing.T) {
	userID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "user_levels") })

	repo := repositories.NewGamificationRepository(testutil.GetTestDB())
	level := &models.UserLevel{
		ID:        utils.GenerateUUID(),
		UserID:    userID,
		Level:     1,
		XPNeed:    100,
		XPCurrent: 50,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := repo.UpdateUserLevel(context.Background(), level)
	require.NoError(t, err)

	level.Level = 2
	level.XPCurrent = 0
	level.XPNeed = 200
	err = repo.UpdateUserLevel(context.Background(), level)
	require.NoError(t, err)

	fetched, err := repo.GetUserLevel(context.Background(), userID)
	require.NoError(t, err)
	assert.Equal(t, 2, fetched.Level)
	assert.Equal(t, 0, fetched.XPCurrent)
	assert.Equal(t, 200, fetched.XPNeed)
}

func TestGamificationRepo_GetUserSkill_CreatesNew(t *testing.T) {
	userID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "user_skills") })

	repo := repositories.NewGamificationRepository(testutil.GetTestDB())
	skill, err := repo.GetUserSkill(context.Background(), userID, "Go")
	require.NoError(t, err)
	assert.Nil(t, skill)
}

func TestGamificationRepo_UpdateUserSkill_Creates(t *testing.T) {
	userID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "user_skills") })

	repo := repositories.NewGamificationRepository(testutil.GetTestDB())
	skill := &models.UserSkill{
		ID:           utils.GenerateUUID(),
		UserID:       userID,
		Skill:        "Go",
		Level:        0,
		PointNeeded:  1,
		PointCurrent: 0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := repo.UpdateUserSkill(context.Background(), skill)
	require.NoError(t, err)

	fetched, err := repo.GetUserSkill(context.Background(), userID, "Go")
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, "Go", fetched.Skill)
	assert.Equal(t, 0, fetched.PointCurrent)
}

func TestGamificationRepo_GetAllUserSkills(t *testing.T) {
	userID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "user_skills") })

	repo := repositories.NewGamificationRepository(testutil.GetTestDB())
	for _, skillName := range []string{"Go", "Python", "Rust"} {
		err := repo.UpdateUserSkill(context.Background(), &models.UserSkill{
			ID:           utils.GenerateUUID(),
			UserID:       userID,
			Skill:        skillName,
			Level:        0,
			PointNeeded:  1,
			PointCurrent: 0,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		})
		require.NoError(t, err)
	}

	skills, err := repo.GetAllUserSkills(context.Background(), userID)
	require.NoError(t, err)
	assert.Len(t, skills, 3)
}
