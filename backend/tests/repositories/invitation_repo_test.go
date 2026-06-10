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

func TestInvitationRepo_CreateInvitation(t *testing.T) {
	invID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "invitations") })

	repo := repositories.NewInvitationRepository(testutil.GetTestDB())
	inv := &models.Invitation{
		ID:         invID,
		ProjectID:  utils.GenerateUUID(),
		StudentID:  utils.GenerateUUID(),
		BusinessID: utils.GenerateUUID(),
		Status:     "pending",
		CreatedAt:  time.Now(),
	}

	err := repo.Create(context.Background(), inv)
	require.NoError(t, err)

	found, err := repo.FindByID(context.Background(), invID)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Equal(t, "pending", found.Status)
}

func TestInvitationRepo_FindByStudentID(t *testing.T) {
	studentID := utils.GenerateUUID()
	projectID1 := utils.GenerateUUID()
	projectID2 := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "invitations") })

	repo := repositories.NewInvitationRepository(testutil.GetTestDB())
	for _, pid := range []string{projectID1, projectID2} {
		err := repo.Create(context.Background(), &models.Invitation{
			ID:         utils.GenerateUUID(),
			ProjectID:  pid,
			StudentID:  studentID,
			BusinessID: utils.GenerateUUID(),
			Status:     "pending",
			CreatedAt:  time.Now(),
		})
		require.NoError(t, err)
	}

	invitations, err := repo.FindByStudentID(context.Background(), studentID)
	require.NoError(t, err)
	assert.Len(t, invitations, 2)
}

func TestInvitationRepo_FindByStudentID_OnlyPending(t *testing.T) {
	studentID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "invitations") })

	repo := repositories.NewInvitationRepository(testutil.GetTestDB())

	err := repo.Create(context.Background(), &models.Invitation{
		ID:         utils.GenerateUUID(),
		ProjectID:  utils.GenerateUUID(),
		StudentID:  studentID,
		BusinessID: utils.GenerateUUID(),
		Status:     "pending",
		CreatedAt:  time.Now(),
	})
	require.NoError(t, err)

	err = repo.Create(context.Background(), &models.Invitation{
		ID:         utils.GenerateUUID(),
		ProjectID:  utils.GenerateUUID(),
		StudentID:  studentID,
		BusinessID: utils.GenerateUUID(),
		Status:     "accepted",
		CreatedAt:  time.Now(),
	})
	require.NoError(t, err)

	invitations, err := repo.FindByStudentID(context.Background(), studentID)
	require.NoError(t, err)
	assert.Len(t, invitations, 1)
	assert.Equal(t, "pending", invitations[0].Status)
}

func TestInvitationRepo_FindByID_NotFound(t *testing.T) {
	repo := repositories.NewInvitationRepository(testutil.GetTestDB())
	inv, err := repo.FindByID(context.Background(), "nonexistent")
	assert.Error(t, err)
	assert.Nil(t, inv)
	assert.Contains(t, err.Error(), "not found")
}

func TestInvitationRepo_UpdateStatus(t *testing.T) {
	invID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "invitations") })

	repo := repositories.NewInvitationRepository(testutil.GetTestDB())
	err := repo.Create(context.Background(), &models.Invitation{
		ID:         invID,
		ProjectID:  utils.GenerateUUID(),
		StudentID:  utils.GenerateUUID(),
		BusinessID: utils.GenerateUUID(),
		Status:     "pending",
		CreatedAt:  time.Now(),
	})
	require.NoError(t, err)

	err = repo.UpdateStatus(context.Background(), invID, "accepted")
	require.NoError(t, err)

	updated, err := repo.FindByID(context.Background(), invID)
	require.NoError(t, err)
	assert.Equal(t, "accepted", updated.Status)
}

func TestInvitationRepo_UpdateStatus_NotFound(t *testing.T) {
	repo := repositories.NewInvitationRepository(testutil.GetTestDB())
	err := repo.UpdateStatus(context.Background(), "nonexistent", "accepted")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestInvitationRepo_FindByStudentAndProject(t *testing.T) {
	studentID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "invitations") })

	repo := repositories.NewInvitationRepository(testutil.GetTestDB())
	inv := &models.Invitation{
		ID:         utils.GenerateUUID(),
		ProjectID:  projectID,
		StudentID:  studentID,
		BusinessID: utils.GenerateUUID(),
		Status:     "pending",
		CreatedAt:  time.Now(),
	}
	err := repo.Create(context.Background(), inv)
	require.NoError(t, err)

	found, err := repo.FindByStudentAndProject(context.Background(), studentID, projectID)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Equal(t, inv.ID, found.ID)

	missing, err := repo.FindByStudentAndProject(context.Background(), studentID, "other_project")
	require.NoError(t, err)
	assert.Nil(t, missing)
}
