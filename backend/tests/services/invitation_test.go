package services_test

import (
	"context"
	"testing"

	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/services"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"github.com/iknizzz1807/SkillForge/tests/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInvitation_CreateInvitation_Success(t *testing.T) {
	studentID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()

	testutil.CreateTestUser(t, studentID, "Student", "student@test.com", "student", nil)
	testutil.CreateTestUser(t, businessID, "Business", "business@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Test Project", businessID, "Business", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "invitations", "users", "projects") })

	db := testutil.GetTestDB()
	invService := services.NewInvitationService(db, nil, nil)

	inv, err := invService.CreateInvitation(studentID, projectID, businessID)
	require.NoError(t, err)
	require.NotNil(t, inv)
	assert.Equal(t, "pending", inv.Status)
	assert.Equal(t, studentID, inv.StudentID)
	assert.Equal(t, projectID, inv.ProjectID)
}

func TestInvitation_CreateInvitation_StudentNotFound(t *testing.T) {
	db := testutil.GetTestDB()
	invService := services.NewInvitationService(db, nil, nil)

	inv, err := invService.CreateInvitation("nonexistent", "proj1", "biz1")
	assert.Error(t, err)
	assert.Nil(t, inv)
	assert.Contains(t, err.Error(), "student not found")
}

func TestInvitation_CreateInvitation_UserNotStudent(t *testing.T) {
	userID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()

	testutil.CreateTestUser(t, userID, "Business", "biz@test.com", "business", nil)
	testutil.CreateTestUser(t, businessID, "Business2", "biz2@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Test", businessID, "Business2", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "invitations", "users", "projects") })

	db := testutil.GetTestDB()
	invService := services.NewInvitationService(db, nil, nil)

	inv, err := invService.CreateInvitation(userID, projectID, businessID)
	assert.Error(t, err)
	assert.Nil(t, inv)
	assert.Contains(t, err.Error(), "not a student")
}

func TestInvitation_CreateInvitation_ProjectNotFound(t *testing.T) {
	studentID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()

	testutil.CreateTestUser(t, studentID, "Student", "s@test.com", "student", nil)
	testutil.CreateTestUser(t, businessID, "Business", "b@test.com", "business", nil)
	t.Cleanup(func() { testutil.CleanupDB(t, "invitations", "users", "projects") })

	db := testutil.GetTestDB()
	invService := services.NewInvitationService(db, nil, nil)

	inv, err := invService.CreateInvitation(studentID, "nonexistent", businessID)
	assert.Error(t, err)
	assert.Nil(t, inv)
}

func TestInvitation_CreateInvitation_NotProjectOwner(t *testing.T) {
	studentID := utils.GenerateUUID()
	businessID1 := utils.GenerateUUID()
	businessID2 := utils.GenerateUUID()
	projectID := utils.GenerateUUID()

	testutil.CreateTestUser(t, studentID, "Student", "s@test.com", "student", nil)
	testutil.CreateTestUser(t, businessID1, "Biz1", "b1@test.com", "business", nil)
	testutil.CreateTestUser(t, businessID2, "Biz2", "b2@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Test", businessID1, "Biz1", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "invitations", "users", "projects") })

	db := testutil.GetTestDB()
	invService := services.NewInvitationService(db, nil, nil)

	inv, err := invService.CreateInvitation(studentID, projectID, businessID2)
	assert.Error(t, err)
	assert.Nil(t, inv)
	assert.Contains(t, err.Error(), "do not own")
}

func TestInvitation_CreateInvitation_ProjectFull(t *testing.T) {
	studentID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()

	testutil.CreateTestUser(t, studentID, "Student", "s@test.com", "student", nil)
	testutil.CreateTestUser(t, businessID, "Business", "b@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Full", businessID, "Business", 2, 2)
	t.Cleanup(func() { testutil.CleanupDB(t, "invitations", "users", "projects") })

	db := testutil.GetTestDB()
	invService := services.NewInvitationService(db, nil, nil)

	inv, err := invService.CreateInvitation(studentID, projectID, businessID)
	assert.Error(t, err)
	assert.Nil(t, inv)
	assert.Contains(t, err.Error(), "full")
}

func TestInvitation_CreateInvitation_Duplicate(t *testing.T) {
	studentID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()

	testutil.CreateTestUser(t, studentID, "Student", "s@test.com", "student", nil)
	testutil.CreateTestUser(t, businessID, "Business", "b@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Test", businessID, "Business", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "invitations", "users", "projects") })

	db := testutil.GetTestDB()
	invService := services.NewInvitationService(db, nil, nil)

	_, err := invService.CreateInvitation(studentID, projectID, businessID)
	require.NoError(t, err)

	inv, err := invService.CreateInvitation(studentID, projectID, businessID)
	assert.Error(t, err)
	assert.Nil(t, inv)
}

func TestInvitation_GetInvitationsByStudent(t *testing.T) {
	studentID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()

	testutil.CreateTestUser(t, studentID, "Student", "s@test.com", "student", nil)
	testutil.CreateTestUser(t, businessID, "BizCorp", "b@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Amazing Project", businessID, "BizCorp", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "invitations", "users", "projects") })

	db := testutil.GetTestDB()
	invService := services.NewInvitationService(db, nil, nil)

	created, err := invService.CreateInvitation(studentID, projectID, businessID)
	require.NoError(t, err)

	invitations, err := invService.GetInvitationsByStudent(studentID)
	require.NoError(t, err)
	require.Len(t, invitations, 1)
	assert.Equal(t, created.ID, invitations[0].ID)
	assert.Equal(t, "Amazing Project", invitations[0].ProjectTitle)
}

func TestInvitation_RespondToInvitation_Accept(t *testing.T) {
	studentID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()

	testutil.CreateTestUser(t, studentID, "Student", "s@test.com", "student", nil)
	testutil.CreateTestUser(t, businessID, "Business", "b@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Test", businessID, "Business", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "invitations", "users", "projects", "project_student") })

	db := testutil.GetTestDB()
	invService := services.NewInvitationService(db, nil, nil)

	created, err := invService.CreateInvitation(studentID, projectID, businessID)
	require.NoError(t, err)

	err = invService.RespondToInvitation(created.ID, studentID, "accepted")
	assert.NoError(t, err)

	invRepo := repositories.NewInvitationRepository(db)
	updated, err := invRepo.FindByID(context.Background(), created.ID)
	require.NoError(t, err)
	assert.Equal(t, "accepted", updated.Status)
}

func TestInvitation_RespondToInvitation_Reject(t *testing.T) {
	studentID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()

	testutil.CreateTestUser(t, studentID, "Student", "s@test.com", "student", nil)
	testutil.CreateTestUser(t, businessID, "Business", "b@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Test", businessID, "Business", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "invitations", "users", "projects") })

	db := testutil.GetTestDB()
	invService := services.NewInvitationService(db, nil, nil)

	created, err := invService.CreateInvitation(studentID, projectID, businessID)
	require.NoError(t, err)

	err = invService.RespondToInvitation(created.ID, studentID, "rejected")
	assert.NoError(t, err)

	invRepo := repositories.NewInvitationRepository(db)
	updated, err := invRepo.FindByID(context.Background(), created.ID)
	require.NoError(t, err)
	assert.Equal(t, "rejected", updated.Status)
}

func TestInvitation_RespondToInvitation_AlreadyResponded(t *testing.T) {
	studentID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()

	testutil.CreateTestUser(t, studentID, "Student", "s@test.com", "student", nil)
	testutil.CreateTestUser(t, businessID, "Business", "b@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Test", businessID, "Business", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "invitations", "users", "projects") })

	db := testutil.GetTestDB()
	invService := services.NewInvitationService(db, nil, nil)

	created, err := invService.CreateInvitation(studentID, projectID, businessID)
	require.NoError(t, err)

	err = invService.RespondToInvitation(created.ID, studentID, "accepted")
	require.NoError(t, err)

	err = invService.RespondToInvitation(created.ID, studentID, "rejected")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already been responded to")
}

func TestInvitation_RespondToInvitation_NotForYou(t *testing.T) {
	studentID := utils.GenerateUUID()
	otherStudentID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()

	testutil.CreateTestUser(t, studentID, "Student", "s@test.com", "student", nil)
	testutil.CreateTestUser(t, otherStudentID, "Other", "o@test.com", "student", nil)
	testutil.CreateTestUser(t, businessID, "Business", "b@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Test", businessID, "Business", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "invitations", "users", "projects") })

	db := testutil.GetTestDB()
	invService := services.NewInvitationService(db, nil, nil)

	created, err := invService.CreateInvitation(studentID, projectID, businessID)
	require.NoError(t, err)

	err = invService.RespondToInvitation(created.ID, otherStudentID, "accepted")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not for you")
}
