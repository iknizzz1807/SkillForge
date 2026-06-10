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

func TestBadgeRepo_InsertBadge(t *testing.T) {
	badgeID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "badges") })

	repo := repositories.NewBadgeRepository(testutil.GetTestDB())
	badge := &models.Badge{
		ID:            badgeID,
		Code:          "badge-project-" + badgeID,
		Name:          "Test Badge",
		Description:   "A test badge",
		Type:          "project",
		Prerequisites: map[string]interface{}{"skill": "Go"},
		CreatedAt:     time.Now(),
	}

	err := repo.InsertBadge(context.Background(), badge)
	require.NoError(t, err)

	found, err := repo.FindBadgeByID(context.Background(), badgeID)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Equal(t, "Test Badge", found.Name)
	assert.Equal(t, "project", found.Type)
}

func TestBadgeRepo_FindBadgeByID_NotFound(t *testing.T) {
	repo := repositories.NewBadgeRepository(testutil.GetTestDB())
	badge, err := repo.FindBadgeByID(context.Background(), "nonexistent")
	assert.Error(t, err)
	assert.Nil(t, badge)
	assert.Contains(t, err.Error(), "not found")
}

func TestBadgeRepo_FindAllBadges(t *testing.T) {
	t.Cleanup(func() { testutil.CleanupDB(t, "badges") })

	repo := repositories.NewBadgeRepository(testutil.GetTestDB())
	for i := 0; i < 3; i++ {
		err := repo.InsertBadge(context.Background(), &models.Badge{
			ID:   utils.GenerateUUID(),
			Code: "badge-" + string(rune('0'+i)),
			Name: "Badge",
			Type: "skill",
		})
		require.NoError(t, err)
	}

	badges, err := repo.FindAllBadges(context.Background())
	require.NoError(t, err)
	assert.Len(t, badges, 3)
}

func TestBadgeRepo_FindBadgesByType(t *testing.T) {
	t.Cleanup(func() { testutil.CleanupDB(t, "badges") })

	repo := repositories.NewBadgeRepository(testutil.GetTestDB())
	err := repo.InsertBadge(context.Background(), &models.Badge{
		ID:   utils.GenerateUUID(),
		Code: "badge-1",
		Name: "Project Badge",
		Type: "project",
	})
	require.NoError(t, err)
	err = repo.InsertBadge(context.Background(), &models.Badge{
		ID:   utils.GenerateUUID(),
		Code: "badge-2",
		Name: "Skill Badge",
		Type: "skill",
	})
	require.NoError(t, err)

	projectBadges, err := repo.FindBadgesByType(context.Background(), "project")
	require.NoError(t, err)
	assert.Len(t, projectBadges, 1)
	assert.Equal(t, "Project Badge", projectBadges[0].Name)
}

func TestBadgeRepo_AwardBadgeToUser(t *testing.T) {
	userID := utils.GenerateUUID()
	badgeID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "badges", "user_badges") })

	repo := repositories.NewBadgeRepository(testutil.GetTestDB())
	err := repo.InsertBadge(context.Background(), &models.Badge{
		ID:   badgeID,
		Code: "badge-" + badgeID,
		Name: "Awarded Badge",
		Type: "project",
	})
	require.NoError(t, err)

	userBadge := &models.UserBadge{
		ID:              utils.GenerateUUID(),
		UserID:          userID,
		BadgeID:         badgeID,
		AwardedAt:       time.Now(),
		RelatedEntityID: "entity1",
	}
	err = repo.AwardBadgeToUser(context.Background(), userBadge)
	require.NoError(t, err)
}

func TestBadgeRepo_HasUserEarnedBadge(t *testing.T) {
	userID := utils.GenerateUUID()
	badgeID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "badges", "user_badges") })

	repo := repositories.NewBadgeRepository(testutil.GetTestDB())
	err := repo.InsertBadge(context.Background(), &models.Badge{
		ID:   badgeID,
		Code: "badge-" + badgeID,
		Name: "Earned Badge",
		Type: "project",
	})
	require.NoError(t, err)

	hasBadge, err := repo.HasUserEarnedBadge(context.Background(), userID, badgeID)
	require.NoError(t, err)
	assert.False(t, hasBadge)

	err = repo.AwardBadgeToUser(context.Background(), &models.UserBadge{
		ID:        utils.GenerateUUID(),
		UserID:    userID,
		BadgeID:   badgeID,
		AwardedAt: time.Now(),
	})
	require.NoError(t, err)

	hasBadge, err = repo.HasUserEarnedBadge(context.Background(), userID, badgeID)
	require.NoError(t, err)
	assert.True(t, hasBadge)
}

func TestBadgeRepo_FindUserBadges(t *testing.T) {
	userID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "badges", "user_badges") })

	repo := repositories.NewBadgeRepository(testutil.GetTestDB())
	badgeIDs := make([]string, 2)
	for i := 0; i < 2; i++ {
		bid := utils.GenerateUUID()
		badgeIDs[i] = bid
		err := repo.InsertBadge(context.Background(), &models.Badge{
			ID:   bid,
			Code: "badge-" + bid,
			Name: "Badge",
			Type: "project",
		})
		require.NoError(t, err)
		err = repo.AwardBadgeToUser(context.Background(), &models.UserBadge{
			ID:        utils.GenerateUUID(),
			UserID:    userID,
			BadgeID:   bid,
			AwardedAt: time.Now(),
		})
		require.NoError(t, err)
	}

	userBadges, err := repo.FindUserBadges(context.Background(), userID)
	require.NoError(t, err)
	assert.Len(t, userBadges, 2)
}
