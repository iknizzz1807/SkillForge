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

func TestBadge_CheckAndAward_BadgeAlreadyEarned(t *testing.T) {
	userID := utils.GenerateUUID()
	testutil.CreateTestUser(t, userID, "Student", "s@test.com", "student", nil)
	t.Cleanup(func() { testutil.CleanupDB(t, "badges", "user_badges", "users") })

	badgeRepo := repositories.NewBadgeRepository(testutil.GetTestDB())
	badgeService := services.NewBadgeService(badgeRepo)

	projectBadge, err := badgeService.CreateBadge("Project Master", "Complete a project", "project", nil)
	require.NoError(t, err)

	_, err = badgeService.AwardBadgeToUser(userID, projectBadge.ID, "proj1")
	require.NoError(t, err)

	result, err := badgeService.AwardBadgeToUser(userID, projectBadge.ID, "proj2")
	require.NoError(t, err)
	assert.Nil(t, result)
}

func TestBadge_AwardBadge_NewBadge(t *testing.T) {
	userID := utils.GenerateUUID()
	testutil.CreateTestUser(t, userID, "Student", "s@test.com", "student", nil)
	t.Cleanup(func() { testutil.CleanupDB(t, "badges", "user_badges", "users") })

	badgeRepo := repositories.NewBadgeRepository(testutil.GetTestDB())
	badgeService := services.NewBadgeService(badgeRepo)

	badge, err := badgeService.CreateBadge("First Project", "Complete your first project", "project", nil)
	require.NoError(t, err)

	userBadge, err := badgeService.AwardBadgeToUser(userID, badge.ID, "proj1")
	require.NoError(t, err)
	require.NotNil(t, userBadge)
	assert.Equal(t, userID, userBadge.UserID)
	assert.Equal(t, badge.ID, userBadge.BadgeID)
	assert.NotNil(t, userBadge.Badge)
	assert.Equal(t, "First Project", userBadge.Badge.Name)
}

func TestBadge_GetUserBadges(t *testing.T) {
	userID := utils.GenerateUUID()
	testutil.CreateTestUser(t, userID, "Student", "s@test.com", "student", nil)
	t.Cleanup(func() { testutil.CleanupDB(t, "badges", "user_badges", "users") })

	badgeRepo := repositories.NewBadgeRepository(testutil.GetTestDB())
	badgeService := services.NewBadgeService(badgeRepo)

	badge1, err := badgeService.CreateBadge("Badge One", "First badge", "project", nil)
	require.NoError(t, err)
	badge2, err := badgeService.CreateBadge("Badge Two", "Second badge", "skill", nil)
	require.NoError(t, err)

	_, err = badgeService.AwardBadgeToUser(userID, badge1.ID, "e1")
	require.NoError(t, err)
	_, err = badgeService.AwardBadgeToUser(userID, badge2.ID, "e2")
	require.NoError(t, err)

	userBadges, err := badgeService.GetUserBadges(userID)
	require.NoError(t, err)
	assert.Len(t, userBadges, 2)

	badgeNames := make([]string, 0, len(userBadges))
	for _, ub := range userBadges {
		badgeNames = append(badgeNames, ub.Badge.Name)
	}
	assert.Contains(t, badgeNames, "Badge One")
	assert.Contains(t, badgeNames, "Badge Two")
}

func TestBadge_CheckAndAwardProjectCompletionBadges(t *testing.T) {
	userID := utils.GenerateUUID()
	testutil.CreateTestUser(t, userID, "Student", "s@test.com", "student", nil)
	t.Cleanup(func() { testutil.CleanupDB(t, "badges", "user_badges", "users") })

	badgeRepo := repositories.NewBadgeRepository(testutil.GetTestDB())
	badgeService := services.NewBadgeService(badgeRepo)

	_, err := badgeService.CreateBadge("Project Champ", "Complete projects", "project", nil)
	require.NoError(t, err)

	err = badgeService.CheckAndAwardProjectCompletionBadges(userID, "proj1")
	require.NoError(t, err)

	userBadges, err := badgeService.GetUserBadges(userID)
	require.NoError(t, err)
	assert.Len(t, userBadges, 1)
}

func TestBadge_CheckAndAwardSkillBadges(t *testing.T) {
	userID := utils.GenerateUUID()
	testutil.CreateTestUser(t, userID, "Student", "s@test.com", "student", nil)
	t.Cleanup(func() { testutil.CleanupDB(t, "badges", "user_badges", "users") })

	badgeRepo := repositories.NewBadgeRepository(testutil.GetTestDB())
	badgeService := services.NewBadgeService(badgeRepo)

	_, err := badgeService.CreateBadge("Go Master", "Master Go programming", "skill", map[string]interface{}{
		"skill": "Go",
	})
	require.NoError(t, err)

	err = badgeService.CheckAndAwardSkillBadges(userID, "Go")
	require.NoError(t, err)

	userBadges, err := badgeService.GetUserBadges(userID)
	require.NoError(t, err)
	assert.Len(t, userBadges, 1)
}
