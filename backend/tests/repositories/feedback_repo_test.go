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

func TestFeedbackRepo_CreateFeedback(t *testing.T) {
	feedbackID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "feedbacks") })

	repo := repositories.NewFeedbackRepository(testutil.GetTestDB())
	feedback := &models.Feedback{
		ID:        feedbackID,
		ProjectID: utils.GenerateUUID(),
		FromID:    utils.GenerateUUID(),
		ToID:      utils.GenerateUUID(),
		Type:      "student",
		Rating:    5,
		Content:   "Excellent work!",
		CreatedAt: time.Now(),
	}

	err := repo.CreateFeedback(context.Background(), feedback)
	require.NoError(t, err)
}

func TestFeedbackRepo_GetFeedbacksForUser(t *testing.T) {
	toID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "feedbacks") })

	repo := repositories.NewFeedbackRepository(testutil.GetTestDB())
	for i := 0; i < 3; i++ {
		err := repo.CreateFeedback(context.Background(), &models.Feedback{
			ID:        utils.GenerateUUID(),
			ProjectID: utils.GenerateUUID(),
			FromID:    utils.GenerateUUID(),
			ToID:      toID,
			Type:      "student",
			Rating:    4,
			Content:   "Great",
			CreatedAt: time.Now(),
		})
		require.NoError(t, err)
	}

	feedbacks, err := repo.GetFeedbacksForUser(context.Background(), toID, "student")
	require.NoError(t, err)
	assert.Len(t, feedbacks, 3)
}

func TestFeedbackRepo_GetFeedbacksForUser_WrongType(t *testing.T) {
	toID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "feedbacks") })

	repo := repositories.NewFeedbackRepository(testutil.GetTestDB())
	err := repo.CreateFeedback(context.Background(), &models.Feedback{
		ID:        utils.GenerateUUID(),
		ProjectID: utils.GenerateUUID(),
		FromID:    utils.GenerateUUID(),
		ToID:      toID,
		Type:      "business",
		Rating:    3,
		Content:   "OK",
		CreatedAt: time.Now(),
	})
	require.NoError(t, err)

	feedbacks, err := repo.GetFeedbacksForUser(context.Background(), toID, "student")
	require.NoError(t, err)
	assert.Empty(t, feedbacks)
}

func TestFeedbackRepo_GetAverageRating(t *testing.T) {
	toID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "feedbacks") })

	repo := repositories.NewFeedbackRepository(testutil.GetTestDB())
	for i := 0; i < 4; i++ {
		err := repo.CreateFeedback(context.Background(), &models.Feedback{
			ID:        utils.GenerateUUID(),
			ProjectID: utils.GenerateUUID(),
			FromID:    utils.GenerateUUID(),
			ToID:      toID,
			Type:      "student",
			Rating:    i + 1,
			Content:   "Review",
			CreatedAt: time.Now(),
		})
		require.NoError(t, err)
	}

	avg, err := repo.GetAverageRating(context.Background(), toID)
	require.NoError(t, err)
	assert.InDelta(t, 2.5, avg, 0.01)
}

func TestFeedbackRepo_GetAverageRating_NoFeedbacks(t *testing.T) {
	repo := repositories.NewFeedbackRepository(testutil.GetTestDB())
	avg, err := repo.GetAverageRating(context.Background(), "nobody")
	require.NoError(t, err)
	assert.Equal(t, 0.0, avg)
}
