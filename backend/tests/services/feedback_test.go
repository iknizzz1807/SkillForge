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

func TestFeedback_CreateFeedback_Success(t *testing.T) {
	fromID := utils.GenerateUUID()
	toID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()

	testutil.CreateTestUser(t, fromID, "From", "from@test.com", "student", nil)
	testutil.CreateTestUser(t, toID, "To", "to@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Project", toID, "To", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "feedbacks", "users", "projects") })

	feedbackRepo := repositories.NewFeedbackRepository(testutil.GetTestDB())
	projectRepo := repositories.NewProjectRepository(testutil.GetTestDB())
	userRepo := repositories.NewUserRepository(testutil.GetTestDB())
	feedbackService := services.NewFeedbackService(feedbackRepo, projectRepo, userRepo)

	feedback, err := feedbackService.CreateFeedback(projectID, fromID, toID, "student", "Great work!", 5)
	require.NoError(t, err)
	require.NotNil(t, feedback)
	assert.Equal(t, 5, feedback.Rating)
	assert.Equal(t, "student", feedback.Type)
	assert.Equal(t, fromID, feedback.FromID)
	assert.Equal(t, toID, feedback.ToID)
}

func TestFeedback_CreateFeedback_SelfFeedback(t *testing.T) {
	userID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()

	testutil.CreateTestUser(t, userID, "User", "u@test.com", "student", nil)
	testutil.CreateTestProject(t, projectID, "Project", userID, "User", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "feedbacks", "users", "projects") })

	feedbackRepo := repositories.NewFeedbackRepository(testutil.GetTestDB())
	projectRepo := repositories.NewProjectRepository(testutil.GetTestDB())
	userRepo := repositories.NewUserRepository(testutil.GetTestDB())
	feedbackService := services.NewFeedbackService(feedbackRepo, projectRepo, userRepo)

	feedback, err := feedbackService.CreateFeedback(projectID, userID, userID, "student", "Self feedback", 3)
	assert.Error(t, err)
	assert.Nil(t, feedback)
	assert.Contains(t, err.Error(), "cannot give feedback to yourself")
}

func TestFeedback_CreateFeedback_InvalidRating(t *testing.T) {
	fromID := utils.GenerateUUID()
	toID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()

	testutil.CreateTestUser(t, fromID, "From", "f@test.com", "student", nil)
	testutil.CreateTestUser(t, toID, "To", "t@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Project", toID, "To", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "feedbacks", "users", "projects") })

	feedbackRepo := repositories.NewFeedbackRepository(testutil.GetTestDB())
	projectRepo := repositories.NewProjectRepository(testutil.GetTestDB())
	userRepo := repositories.NewUserRepository(testutil.GetTestDB())
	feedbackService := services.NewFeedbackService(feedbackRepo, projectRepo, userRepo)

	fb, err := feedbackService.CreateFeedback(projectID, fromID, toID, "student", "Bad", -1)
	assert.Error(t, err)
	assert.Nil(t, fb)

	fb, err = feedbackService.CreateFeedback(projectID, fromID, toID, "student", "Bad", 6)
	assert.Error(t, err)
	assert.Nil(t, fb)
}

func TestFeedback_CreateFeedback_InvalidType(t *testing.T) {
	fromID := utils.GenerateUUID()
	toID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()

	testutil.CreateTestUser(t, fromID, "From", "f@test.com", "student", nil)
	testutil.CreateTestUser(t, toID, "To", "t@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Project", toID, "To", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "feedbacks", "users", "projects") })

	feedbackRepo := repositories.NewFeedbackRepository(testutil.GetTestDB())
	projectRepo := repositories.NewProjectRepository(testutil.GetTestDB())
	userRepo := repositories.NewUserRepository(testutil.GetTestDB())
	feedbackService := services.NewFeedbackService(feedbackRepo, projectRepo, userRepo)

	fb, err := feedbackService.CreateFeedback(projectID, fromID, toID, "invalid_type", "test", 3)
	assert.Error(t, err)
	assert.Nil(t, fb)
	assert.Contains(t, err.Error(), "invalid feedback type")
}

func TestFeedback_GetUserFeedbacks(t *testing.T) {
	fromID := utils.GenerateUUID()
	toID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()

	testutil.CreateTestUser(t, fromID, "From", "f@test.com", "student", nil)
	testutil.CreateTestUser(t, toID, "To", "t@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Project", toID, "To", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "feedbacks", "users", "projects") })

	feedbackRepo := repositories.NewFeedbackRepository(testutil.GetTestDB())
	projectRepo := repositories.NewProjectRepository(testutil.GetTestDB())
	userRepo := repositories.NewUserRepository(testutil.GetTestDB())
	feedbackService := services.NewFeedbackService(feedbackRepo, projectRepo, userRepo)

	_, err := feedbackService.CreateFeedback(projectID, fromID, toID, "student", "Great!", 5)
	require.NoError(t, err)

	feedbacks, err := feedbackService.GetUserFeedbacks(toID, "student")
	require.NoError(t, err)
	assert.Len(t, feedbacks, 1)
}

func TestFeedback_GetAverageRating(t *testing.T) {
	fromID1 := utils.GenerateUUID()
	fromID2 := utils.GenerateUUID()
	toID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()

	testutil.CreateTestUser(t, fromID1, "F1", "f1@test.com", "student", nil)
	testutil.CreateTestUser(t, fromID2, "F2", "f2@test.com", "student", nil)
	testutil.CreateTestUser(t, toID, "To", "t@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Project", toID, "To", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "feedbacks", "users", "projects") })

	feedbackRepo := repositories.NewFeedbackRepository(testutil.GetTestDB())
	projectRepo := repositories.NewProjectRepository(testutil.GetTestDB())
	userRepo := repositories.NewUserRepository(testutil.GetTestDB())
	feedbackService := services.NewFeedbackService(feedbackRepo, projectRepo, userRepo)

	_, err := feedbackService.CreateFeedback(projectID, fromID1, toID, "student", "Good", 4)
	require.NoError(t, err)
	_, err = feedbackService.CreateFeedback(projectID, fromID2, toID, "student", "OK", 2)
	require.NoError(t, err)

	avg, err := feedbackService.GetAverageRating(toID)
	require.NoError(t, err)
	assert.InDelta(t, 3.0, avg, 0.01)
}
