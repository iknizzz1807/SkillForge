package services_test

import (
	"testing"

	"github.com/iknizzz1807/SkillForge/internal/services"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"github.com/iknizzz1807/SkillForge/tests/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createApplicationService(t *testing.T) *services.ApplicationService {
	t.Helper()
	realtime := testutil.NewTestRealtimeClient()
	emailClient := testutil.NewTestEmailClient()
	notificationService := services.NewNotificationService(realtime, testutil.GetTestDB(), emailClient)
	return services.NewApplicationService(testutil.GetTestDB(), notificationService, nil, nil)
}

func TestApplication_ApplyProject_Success(t *testing.T) {
	studentID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()

	testutil.CreateTestUser(t, studentID, "Student", "s@test.com", "student", nil)
	testutil.CreateTestUser(t, businessID, "Business", "b@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Test Project", businessID, "Business", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "applications", "users", "projects") })

	appService := createApplicationService(t)
	app, err := appService.ApplyProject(studentID, projectID, "I want to join", "My detailed proposal")
	require.NoError(t, err)
	require.NotNil(t, app)
	assert.Equal(t, "pending", app.Status)
	assert.Equal(t, studentID, app.UserID)
	assert.Equal(t, projectID, app.ProjectID)
	assert.Equal(t, "Test Project", app.ProjectName)
	assert.Equal(t, "Student", app.UserName)
}

func TestApplication_ApplyProject_Duplicate(t *testing.T) {
	studentID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()

	testutil.CreateTestUser(t, studentID, "Student", "s@test.com", "student", nil)
	testutil.CreateTestUser(t, businessID, "Business", "b@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Test", businessID, "Business", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "applications", "users", "projects") })

	appService := createApplicationService(t)

	_, err := appService.ApplyProject(studentID, projectID, "motivation", "proposal")
	require.NoError(t, err)

	_, err = appService.ApplyProject(studentID, projectID, "motivation again", "proposal again")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already applied")
}

func TestApplication_GetApplicationByID(t *testing.T) {
	studentID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()

	testutil.CreateTestUser(t, studentID, "Student", "s@test.com", "student", nil)
	testutil.CreateTestUser(t, businessID, "Business", "b@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Test", businessID, "Business", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "applications", "users", "projects") })

	appService := createApplicationService(t)
	created, err := appService.ApplyProject(studentID, projectID, "motivation", "proposal")
	require.NoError(t, err)

	fetched, err := appService.GetApplicationByID(created.ID)
	require.NoError(t, err)
	assert.Equal(t, created.ID, fetched.ID)
	assert.Equal(t, studentID, fetched.UserID)
}

func TestApplication_GetApplicationsByUserID(t *testing.T) {
	studentID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	projectID1 := utils.GenerateUUID()
	projectID2 := utils.GenerateUUID()

	testutil.CreateTestUser(t, studentID, "Student", "s@test.com", "student", nil)
	testutil.CreateTestUser(t, businessID, "Business", "b@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID1, "Project A", businessID, "Business", 5, 0)
	testutil.CreateTestProject(t, projectID2, "Project B", businessID, "Business", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "applications", "users", "projects") })

	appService := createApplicationService(t)
	_, err := appService.ApplyProject(studentID, projectID1, "m1", "p1")
	require.NoError(t, err)
	_, err = appService.ApplyProject(studentID, projectID2, "m2", "p2")
	require.NoError(t, err)

	apps, err := appService.GetApplicationsByUserID(studentID)
	require.NoError(t, err)
	assert.Len(t, apps, 2)
}

func TestApplication_UpdateApplicationStatus_Approve(t *testing.T) {
	studentID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()

	testutil.CreateTestUser(t, studentID, "Student", "s@test.com", "student", nil)
	testutil.CreateTestUser(t, businessID, "Business", "b@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Test", businessID, "Business", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "applications", "users", "projects", "project_student") })

	appService := createApplicationService(t)
	created, err := appService.ApplyProject(studentID, projectID, "motivation", "proposal")
	require.NoError(t, err)

	updated, err := appService.UpdateApplicationStatus(created.ID, "approved", businessID)
	require.NoError(t, err)
	assert.Equal(t, "approved", updated.Status)
}

func TestApplication_UpdateApplicationStatus_Reject(t *testing.T) {
	studentID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()

	testutil.CreateTestUser(t, studentID, "Student", "s@test.com", "student", nil)
	testutil.CreateTestUser(t, businessID, "Business", "b@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Test", businessID, "Business", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "applications", "users", "projects") })

	appService := createApplicationService(t)
	created, err := appService.ApplyProject(studentID, projectID, "motivation", "proposal")
	require.NoError(t, err)

	updated, err := appService.UpdateApplicationStatus(created.ID, "rejected", businessID)
	require.NoError(t, err)
	assert.Equal(t, "rejected", updated.Status)
}

func TestApplication_UpdateApplicationStatus_PermissionDenied(t *testing.T) {
	studentID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	otherBusinessID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()

	testutil.CreateTestUser(t, studentID, "Student", "s@test.com", "student", nil)
	testutil.CreateTestUser(t, businessID, "Biz", "b@test.com", "business", nil)
	testutil.CreateTestUser(t, otherBusinessID, "Other", "o@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Test", businessID, "Biz", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "applications", "users", "projects") })

	appService := createApplicationService(t)
	created, err := appService.ApplyProject(studentID, projectID, "m", "p")
	require.NoError(t, err)

	app, err := appService.UpdateApplicationStatus(created.ID, "approved", otherBusinessID)
	assert.Error(t, err)
	assert.Nil(t, app)
}
