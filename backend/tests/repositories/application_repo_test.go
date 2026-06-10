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

func TestApplicationRepo_InsertApplication(t *testing.T) {
	appID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "applications") })

	repo := repositories.NewApplicationRepository(testutil.GetTestDB())
	app := &models.Application{
		ID:               appID,
		UserID:           utils.GenerateUUID(),
		ProjectID:        utils.GenerateUUID(),
		ProjectName:      "Test Project",
		UserName:         "Test User",
		Motivation:       "I want to join",
		DetailedProposal: "My detailed proposal",
		Status:           "pending",
		CreatedAt:        time.Now(),
	}

	err := repo.InsertApplication(context.Background(), app)
	require.NoError(t, err)

	found, err := repo.FindApplicationByID(context.Background(), appID)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Equal(t, "pending", found.Status)
	assert.Equal(t, "I want to join", found.Motivation)
}

func TestApplicationRepo_FindApplicationByID_NotFound(t *testing.T) {
	repo := repositories.NewApplicationRepository(testutil.GetTestDB())
	app, err := repo.FindApplicationByID(context.Background(), "nonexistent")
	assert.Error(t, err)
	assert.Nil(t, app)
}

func TestApplicationRepo_FindByUserAndProject(t *testing.T) {
	userID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "applications") })

	repo := repositories.NewApplicationRepository(testutil.GetTestDB())
	app := &models.Application{
		ID:               utils.GenerateUUID(),
		UserID:           userID,
		ProjectID:        projectID,
		ProjectName:      "Test",
		UserName:         "User",
		Motivation:       "motive",
		DetailedProposal: "proposal",
		Status:           "pending",
		CreatedAt:        time.Now(),
	}
	err := repo.InsertApplication(context.Background(), app)
	require.NoError(t, err)

	found, err := repo.FindByUserAndProject(context.Background(), userID, projectID)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Equal(t, app.ID, found.ID)

	notFound, err := repo.FindByUserAndProject(context.Background(), userID, "other_project")
	assert.Error(t, err)
	assert.Nil(t, notFound)
}

func TestApplicationRepo_FindByUserID(t *testing.T) {
	userID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "applications") })

	repo := repositories.NewApplicationRepository(testutil.GetTestDB())
	for i := 0; i < 3; i++ {
		err := repo.InsertApplication(context.Background(), &models.Application{
			ID:               utils.GenerateUUID(),
			UserID:           userID,
			ProjectID:        utils.GenerateUUID(),
			ProjectName:      "Project",
			UserName:         "User",
			Motivation:       "motive",
			DetailedProposal: "proposal",
			Status:           "pending",
			CreatedAt:        time.Now(),
		})
		require.NoError(t, err)
	}

	apps, err := repo.FindByUserID(userID)
	require.NoError(t, err)
	assert.Len(t, apps, 3)
}

func TestApplicationRepo_UpdateStatus(t *testing.T) {
	appID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "applications") })

	repo := repositories.NewApplicationRepository(testutil.GetTestDB())
	app := &models.Application{
		ID:               appID,
		UserID:           utils.GenerateUUID(),
		ProjectID:        utils.GenerateUUID(),
		ProjectName:      "Test",
		UserName:         "User",
		Motivation:       "motive",
		DetailedProposal: "proposal",
		Status:           "pending",
		CreatedAt:        time.Now(),
	}
	err := repo.InsertApplication(context.Background(), app)
	require.NoError(t, err)

	updated, err := repo.UpdateStatus(appID, "approved")
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, "approved", updated.Status)
}

func TestApplicationRepo_FindByProjectIDs(t *testing.T) {
	projectID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "applications") })

	repo := repositories.NewApplicationRepository(testutil.GetTestDB())
	for i := 0; i < 2; i++ {
		err := repo.InsertApplication(context.Background(), &models.Application{
			ID:               utils.GenerateUUID(),
			UserID:           utils.GenerateUUID(),
			ProjectID:        projectID,
			ProjectName:      "Project",
			UserName:         "User",
			Motivation:       "motive",
			DetailedProposal: "proposal",
			Status:           "pending",
			CreatedAt:        time.Now(),
		})
		require.NoError(t, err)
	}

	apps, err := repo.FindByProjectIDs([]string{projectID})
	require.NoError(t, err)
	assert.Len(t, apps, 2)
}

func TestApplicationRepo_FindByProjectIDs_Empty(t *testing.T) {
	repo := repositories.NewApplicationRepository(testutil.GetTestDB())
	apps, err := repo.FindByProjectIDs([]string{})
	require.NoError(t, err)
	assert.Empty(t, apps)
}

func TestApplicationRepo_Delete(t *testing.T) {
	appID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "applications") })

	repo := repositories.NewApplicationRepository(testutil.GetTestDB())
	err := repo.InsertApplication(context.Background(), &models.Application{
		ID:               appID,
		UserID:           utils.GenerateUUID(),
		ProjectID:        utils.GenerateUUID(),
		ProjectName:      "Test",
		UserName:         "User",
		Motivation:       "motive",
		DetailedProposal: "proposal",
		Status:           "pending",
		CreatedAt:        time.Now(),
	})
	require.NoError(t, err)

	err = repo.Delete(context.Background(), appID)
	require.NoError(t, err)

	found, err := repo.FindApplicationByID(context.Background(), appID)
	assert.Error(t, err)
	assert.Nil(t, found)
}

func TestApplicationRepo_DeleteApplicationsByProjectID(t *testing.T) {
	projectID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "applications") })

	repo := repositories.NewApplicationRepository(testutil.GetTestDB())
	for i := 0; i < 3; i++ {
		err := repo.InsertApplication(context.Background(), &models.Application{
			ID:               utils.GenerateUUID(),
			UserID:           utils.GenerateUUID(),
			ProjectID:        projectID,
			ProjectName:      "Project",
			UserName:         "User",
			Motivation:       "motive",
			DetailedProposal: "proposal",
			Status:           "pending",
			CreatedAt:        time.Now(),
		})
		require.NoError(t, err)
	}

	count, err := repo.DeleteApplicationsByProjectID(context.Background(), projectID)
	require.NoError(t, err)
	assert.Equal(t, int64(3), count)
}
