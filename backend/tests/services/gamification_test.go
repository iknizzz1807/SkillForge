package services_test

import (
	"testing"

	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/services"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"github.com/iknizzz1807/SkillForge/tests/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGamification_GetUserLevel_CreatesNewLevel(t *testing.T) {
	userID := utils.GenerateUUID()
	testutil.CreateTestUser(t, userID, "Test User", "test@example.com", "student", nil)
	t.Cleanup(func() { testutil.CleanupDB(t, "user_levels", "user_skills", "users") })

	gamificationRepo := repositories.NewGamificationRepository(testutil.GetTestDB())
	userService := services.NewUserService(testutil.GetTestDB())
	gamificationService := services.NewGamificationService(gamificationRepo, userService)

	level, err := gamificationService.GetUserLevel(userID)
	require.NoError(t, err)
	require.NotNil(t, level)
	assert.Equal(t, 1, level.Level)
	assert.Equal(t, 0, level.XPCurrent)
	assert.NotZero(t, level.XPNeed)
}

func TestGamification_AddXP_IncreasesXP(t *testing.T) {
	userID := utils.GenerateUUID()
	testutil.CreateTestUser(t, userID, "Test User", "test@example.com", "student", nil)
	t.Cleanup(func() { testutil.CleanupDB(t, "user_levels", "user_skills", "users") })

	gamificationRepo := repositories.NewGamificationRepository(testutil.GetTestDB())
	userService := services.NewUserService(testutil.GetTestDB())
	gamificationService := services.NewGamificationService(gamificationRepo, userService)

	level, err := gamificationService.AddXP(userID, 50)
	require.NoError(t, err)
	assert.Equal(t, 50, level.XPCurrent)
	assert.Equal(t, 1, level.Level)
}

func TestGamification_AddXP_LevelUp(t *testing.T) {
	userID := utils.GenerateUUID()
	testutil.CreateTestUser(t, userID, "Test User", "test@example.com", "student", nil)
	t.Cleanup(func() { testutil.CleanupDB(t, "user_levels", "user_skills", "users") })

	gamificationRepo := repositories.NewGamificationRepository(testutil.GetTestDB())
	userService := services.NewUserService(testutil.GetTestDB())
	gamificationService := services.NewGamificationService(gamificationRepo, userService)

	level, err := gamificationService.AddXP(userID, 200)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, level.Level, 2)
	assert.Less(t, level.XPCurrent, level.XPNeed)
}

func TestGamification_AddXP_MultipleLevelUps(t *testing.T) {
	userID := utils.GenerateUUID()
	testutil.CreateTestUser(t, userID, "Test User", "test@example.com", "student", nil)
	t.Cleanup(func() { testutil.CleanupDB(t, "user_levels", "user_skills", "users") })

	gamificationRepo := repositories.NewGamificationRepository(testutil.GetTestDB())
	userService := services.NewUserService(testutil.GetTestDB())
	gamificationService := services.NewGamificationService(gamificationRepo, userService)

	level, err := gamificationService.AddXP(userID, 10000)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, level.Level, 5)
	assert.Less(t, level.XPCurrent, level.XPNeed)
}

func TestGamification_GetSkill_CreatesNewSkill(t *testing.T) {
	userID := utils.GenerateUUID()
	testutil.CreateTestUser(t, userID, "Test User", "test@example.com", "student", nil)
	t.Cleanup(func() { testutil.CleanupDB(t, "user_levels", "user_skills", "users") })

	gamificationRepo := repositories.NewGamificationRepository(testutil.GetTestDB())
	userService := services.NewUserService(testutil.GetTestDB())
	gamificationService := services.NewGamificationService(gamificationRepo, userService)

	skill, err := gamificationService.GetSkill(userID, "Go")
	require.NoError(t, err)
	require.NotNil(t, skill)
	assert.Equal(t, "Go", skill.Skill)
	assert.Equal(t, 0, skill.PointCurrent)
}

func TestGamification_AddSkillPoint_IncreasesPoints(t *testing.T) {
	userID := utils.GenerateUUID()
	testutil.CreateTestUser(t, userID, "Test User", "test@example.com", "student", nil)
	t.Cleanup(func() { testutil.CleanupDB(t, "user_levels", "user_skills", "users") })

	gamificationRepo := repositories.NewGamificationRepository(testutil.GetTestDB())
	userService := services.NewUserService(testutil.GetTestDB())
	gamificationService := services.NewGamificationService(gamificationRepo, userService)

	// Points 1-2 trigger level-ups (PointNeeded=1 at lv1, 1 at lv1->lv2).
	// At level 2, PointNeeded=4, so 3rd point stays: PointCurrent=1.
	for i := 0; i < 3; i++ {
		_, err := gamificationService.AddSkillPoint(userID, "Python")
		require.NoError(t, err)
	}
	skill, err := gamificationService.GetSkill(userID, "Python")
	require.NoError(t, err)
	assert.Equal(t, 1, skill.PointCurrent)
}

func TestGamification_AddSkillPoint_LevelUp(t *testing.T) {
	userID := utils.GenerateUUID()
	testutil.CreateTestUser(t, userID, "Test User", "test@example.com", "student", nil)
	t.Cleanup(func() { testutil.CleanupDB(t, "user_levels", "user_skills", "users") })

	gamificationRepo := repositories.NewGamificationRepository(testutil.GetTestDB())
	userService := services.NewUserService(testutil.GetTestDB())
	gamificationService := services.NewGamificationService(gamificationRepo, userService)

	for i := 0; i < 5; i++ {
		_, err := gamificationService.AddSkillPoint(userID, "Java")
		require.NoError(t, err)
	}
	skill, err := gamificationService.GetSkill(userID, "Java")
	require.NoError(t, err)
	assert.GreaterOrEqual(t, skill.Level, 1)
}

func TestGamification_AddSkillPoint_AddsToUserSkills(t *testing.T) {
	userID := utils.GenerateUUID()
	testutil.CreateTestUser(t, userID, "Test User", "test@example.com", "student", nil)
	t.Cleanup(func() { testutil.CleanupDB(t, "user_levels", "user_skills", "users") })

	gamificationRepo := repositories.NewGamificationRepository(testutil.GetTestDB())
	userService := services.NewUserService(testutil.GetTestDB())
	gamificationService := services.NewGamificationService(gamificationRepo, userService)

	_, err := gamificationService.AddSkillPoint(userID, "Rust")
	require.NoError(t, err)

	user, err := userService.GetUserByID(userID)
	require.NoError(t, err)
	assert.Contains(t, user.Skills, "Rust")
}

func TestGamification_GetAllSkills(t *testing.T) {
	userID := utils.GenerateUUID()
	testutil.CreateTestUser(t, userID, "Test User", "test@example.com", "student", nil)
	t.Cleanup(func() { testutil.CleanupDB(t, "user_levels", "user_skills", "users") })

	gamificationRepo := repositories.NewGamificationRepository(testutil.GetTestDB())
	userService := services.NewUserService(testutil.GetTestDB())
	gamificationService := services.NewGamificationService(gamificationRepo, userService)

	_, err := gamificationService.AddSkillPoint(userID, "Go")
	require.NoError(t, err)
	_, err = gamificationService.AddSkillPoint(userID, "Python")
	require.NoError(t, err)

	skills, err := gamificationService.GetAllSkills(userID)
	require.NoError(t, err)
	assert.Len(t, skills, 2)
}
