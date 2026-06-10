package services_test

import (
	"testing"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/services"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"github.com/iknizzz1807/SkillForge/tests/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createTaskService(t *testing.T) *services.TaskService {
	t.Helper()
	realtime := testutil.NewTestRealtimeClient()
	emailClient := testutil.NewTestEmailClient()
	notificationService := services.NewNotificationService(realtime, testutil.GetTestDB(), emailClient)
	return services.NewTaskService(testutil.GetTestDB(), notificationService, nil)
}

func TestTask_CreateTasks(t *testing.T) {
	userID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()
	testutil.CreateTestUser(t, userID, "User", "u@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Project", userID, "User", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "tasks", "activities", "users", "projects") })

	taskService := createTaskService(t)
	inputs := []models.TaskInput{
		{Title: "Task 1", Description: "Description 1", AssignedTo: "student1"},
		{Title: "Task 2", Description: "Description 2", AssignedTo: "student2"},
	}

	tasks, err := taskService.CreateTasks(projectID, userID, inputs)
	require.NoError(t, err)
	require.Len(t, tasks, 2)
	assert.Equal(t, "todo", tasks[0].Status)
	assert.Equal(t, projectID, tasks[0].ProjectID)
	assert.Equal(t, "Task 1", tasks[0].Title)
}

func TestTask_GetTasksByProjectID(t *testing.T) {
	userID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()
	testutil.CreateTestUser(t, userID, "User", "u@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Project", userID, "User", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "tasks", "activities", "users", "projects") })

	taskService := createTaskService(t)
	inputs := []models.TaskInput{
		{Title: "Task A", Description: "Desc A"},
	}
	_, err := taskService.CreateTasks(projectID, userID, inputs)
	require.NoError(t, err)

	tasks, err := taskService.GetTasksByProjectID(projectID)
	require.NoError(t, err)
	assert.Len(t, tasks, 1)
	assert.Equal(t, "Task A", tasks[0].Title)
}

func TestTask_UpdateTaskStatus(t *testing.T) {
	userID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()
	testutil.CreateTestUser(t, userID, "User", "u@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Project", userID, "User", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "tasks", "activities", "users", "projects") })

	taskService := createTaskService(t)
	inputs := []models.TaskInput{
		{Title: "Task 1", Description: "Desc", AssignedTo: "student1"},
	}
	created, err := taskService.CreateTasks(projectID, userID, inputs)
	require.NoError(t, err)

	updated, err := taskService.UpdateTask(created[0].ID, userID, &models.TaskUpdate{
		TaskID:      created[0].ID,
		Status:      "in_progress",
		Title:       "Task 1",
		Description: "Desc",
		AssignedTo:  "student1",
	})
	require.NoError(t, err)
	assert.Equal(t, "in_progress", updated.Status)
}

func TestTask_CompleteTask(t *testing.T) {
	userID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()
	testutil.CreateTestUser(t, userID, "User", "u@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Project", userID, "User", 5, 0)
	t.Cleanup(func() { testutil.CleanupDB(t, "tasks", "activities", "users", "projects") })

	taskService := createTaskService(t)
	inputs := []models.TaskInput{
		{Title: "Task Done", Description: "Desc", AssignedTo: "student1"},
	}
	created, err := taskService.CreateTasks(projectID, userID, inputs)
	require.NoError(t, err)

	updated, err := taskService.UpdateTask(created[0].ID, userID, &models.TaskUpdate{
		TaskID:      created[0].ID,
		Status:      "done",
		Title:       "Task Done",
		Description: "Desc",
		AssignedTo:  "student1",
	})
	require.NoError(t, err)
	assert.Equal(t, "done", updated.Status)
}

func TestTask_UpdateTask_NotFound(t *testing.T) {
	taskService := createTaskService(t)
	updated, err := taskService.UpdateTask("nonexistent", "user1", &models.TaskUpdate{
		TaskID: "nonexistent",
		Status: "done",
	})
	assert.Error(t, err)
	assert.Nil(t, updated)
}

func TestTask_CreateTasks_InvalidInput(t *testing.T) {
	taskService := createTaskService(t)

	tasks, err := taskService.CreateTasks("", "user1", []models.TaskInput{{Title: "Test"}})
	assert.Error(t, err)
	assert.Nil(t, tasks)

	tasks, err = taskService.CreateTasks("proj1", "user1", nil)
	assert.Error(t, err)
	assert.Nil(t, tasks)
}
