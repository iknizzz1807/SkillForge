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

func TestProjectStudentRepo_InsertProjectStudent(t *testing.T) {
	projectID := utils.GenerateUUID()
	studentID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "project_student", "projects", "users") })

	testutil.CreateTestUser(t, studentID, "Student", "s@test.com", "student", nil)
	testutil.CreateTestUser(t, businessID, "Biz", "b@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Test", businessID, "Biz", 5, 0)

	repo := repositories.NewProjectStudentRepository(testutil.GetTestDB())
	ps := &models.Project_student{
		ID:         utils.GenerateUUID(),
		Project_id: projectID,
		Student_id: studentID,
		CreatedAt:  time.Now(),
	}
	err := repo.InsertProjectStudent(context.Background(), ps)
	require.NoError(t, err)
}

func TestProjectStudentRepo_InsertProjectStudent_Duplicate(t *testing.T) {
	projectID := utils.GenerateUUID()
	studentID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "project_student", "projects", "users") })

	testutil.CreateTestUser(t, studentID, "Student", "s@test.com", "student", nil)
	testutil.CreateTestUser(t, businessID, "Biz", "b@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Test", businessID, "Biz", 5, 0)

	repo := repositories.NewProjectStudentRepository(testutil.GetTestDB())
	ps := &models.Project_student{
		ID:         utils.GenerateUUID(),
		Project_id: projectID,
		Student_id: studentID,
		CreatedAt:  time.Now(),
	}
	err := repo.InsertProjectStudent(context.Background(), ps)
	require.NoError(t, err)

	err = repo.InsertProjectStudent(context.Background(), ps)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already joined")
}

func TestProjectStudentRepo_FindProjectsByStudentID(t *testing.T) {
	studentID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "project_student", "projects", "users") })

	testutil.CreateTestUser(t, studentID, "Student", "s@test.com", "student", nil)
	testutil.CreateTestUser(t, businessID, "Biz", "b@test.com", "business", nil)

	repo := repositories.NewProjectStudentRepository(testutil.GetTestDB())
	projectIDs := make([]string, 3)
	for i := 0; i < 3; i++ {
		pid := utils.GenerateUUID()
		projectIDs[i] = pid
		testutil.CreateTestProject(t, pid, "Project", businessID, "Biz", 5, 0)
		err := repo.InsertProjectStudent(context.Background(), &models.Project_student{
			ID:         utils.GenerateUUID(),
			Project_id: pid,
			Student_id: studentID,
			CreatedAt:  time.Now(),
		})
		require.NoError(t, err)
	}

	found, err := repo.FindProjectsByStudentID(context.Background(), studentID)
	require.NoError(t, err)
	assert.Len(t, found, 3)
}

func TestProjectStudentRepo_FindStudentsByProjectID(t *testing.T) {
	projectID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "project_student", "projects", "users") })

	testutil.CreateTestUser(t, businessID, "Biz", "b@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Test", businessID, "Biz", 5, 0)

	repo := repositories.NewProjectStudentRepository(testutil.GetTestDB())
	studentIDs := make([]string, 2)
	for i := 0; i < 2; i++ {
		sid := utils.GenerateUUID()
		studentIDs[i] = sid
		testutil.CreateTestUser(t, sid, "Student", "s@test.com", "student", nil)
		err := repo.InsertProjectStudent(context.Background(), &models.Project_student{
			ID:         utils.GenerateUUID(),
			Project_id: projectID,
			Student_id: sid,
			CreatedAt:  time.Now(),
		})
		require.NoError(t, err)
	}

	found, err := repo.FindStudentsByProjectID(context.Background(), projectID)
	require.NoError(t, err)
	assert.Len(t, found, 2)
}

func TestProjectStudentRepo_DeleteProjectStudent(t *testing.T) {
	projectID := utils.GenerateUUID()
	studentID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "project_student", "projects", "users") })

	testutil.CreateTestUser(t, studentID, "Student", "s@test.com", "student", nil)
	testutil.CreateTestUser(t, businessID, "Biz", "b@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Test", businessID, "Biz", 5, 0)

	repo := repositories.NewProjectStudentRepository(testutil.GetTestDB())
	err := repo.InsertProjectStudent(context.Background(), &models.Project_student{
		ID:         utils.GenerateUUID(),
		Project_id: projectID,
		Student_id: studentID,
		CreatedAt:  time.Now(),
	})
	require.NoError(t, err)

	err = repo.DeleteProjectStudent(context.Background(), projectID, studentID)
	require.NoError(t, err)

	found, err := repo.FindStudentsByProjectID(context.Background(), projectID)
	require.NoError(t, err)
	assert.Empty(t, found)
}

func TestProjectStudentRepo_DeleteProjectStudent_NotFound(t *testing.T) {
	repo := repositories.NewProjectStudentRepository(testutil.GetTestDB())
	err := repo.DeleteProjectStudent(context.Background(), "nonexistent", "nobody")
	assert.Error(t, err)
}
