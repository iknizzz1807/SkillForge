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

func TestTaskRepo_InsertTasks(t *testing.T) {
	projectID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "tasks") })

	repo := repositories.NewTaskRepository(testutil.GetTestDB())
	tasks := []*models.Task{
		{
			ID:        utils.GenerateUUID(),
			ProjectID: projectID,
			Title:     "Task 1",
			Status:    "todo",
			CreatedAt: time.Now(),
		},
		{
			ID:        utils.GenerateUUID(),
			ProjectID: projectID,
			Title:     "Task 2",
			Status:    "todo",
			CreatedAt: time.Now(),
		},
	}

	err := repo.InsertTasks(context.Background(), tasks)
	require.NoError(t, err)

	found, err := repo.FindTasksByProjectID(context.Background(), projectID)
	require.NoError(t, err)
	assert.Len(t, found, 2)
}

func TestTaskRepo_FindTasksByProjectID_Empty(t *testing.T) {
	repo := repositories.NewTaskRepository(testutil.GetTestDB())
	tasks, err := repo.FindTasksByProjectID(context.Background(), "nonexistent")
	require.NoError(t, err)
	assert.Empty(t, tasks)
}

func TestTaskRepo_FindTaskByID(t *testing.T) {
	taskID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "tasks") })

	repo := repositories.NewTaskRepository(testutil.GetTestDB())
	task := &models.Task{
		ID:        taskID,
		ProjectID: projectID,
		Title:     "Find Me",
		Status:    "todo",
		CreatedAt: time.Now(),
	}
	err := repo.InsertTasks(context.Background(), []*models.Task{task})
	require.NoError(t, err)

	found, err := repo.FindTaskByID(context.Background(), taskID)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Equal(t, "Find Me", found.Title)
}

func TestTaskRepo_FindTaskByID_NotFound(t *testing.T) {
	repo := repositories.NewTaskRepository(testutil.GetTestDB())
	task, err := repo.FindTaskByID(context.Background(), "nonexistent")
	require.NoError(t, err)
	assert.Nil(t, task)
}

func TestTaskRepo_UpdateTask(t *testing.T) {
	taskID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "tasks") })

	repo := repositories.NewTaskRepository(testutil.GetTestDB())
	task := &models.Task{
		ID:        taskID,
		ProjectID: projectID,
		Title:     "Original",
		Status:    "todo",
		CreatedAt: time.Now(),
	}
	err := repo.InsertTasks(context.Background(), []*models.Task{task})
	require.NoError(t, err)

	task.Status = "done"
	err = repo.UpdateTask(context.Background(), task)
	require.NoError(t, err)

	updated, err := repo.FindTaskByID(context.Background(), taskID)
	require.NoError(t, err)
	assert.Equal(t, "done", updated.Status)
}

func TestTaskRepo_DeleteTask(t *testing.T) {
	taskID := utils.GenerateUUID()
	projectID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "tasks") })

	repo := repositories.NewTaskRepository(testutil.GetTestDB())
	task := &models.Task{
		ID:        taskID,
		ProjectID: projectID,
		Title:     "Delete Me",
		Status:    "todo",
		CreatedAt: time.Now(),
	}
	err := repo.InsertTasks(context.Background(), []*models.Task{task})
	require.NoError(t, err)

	err = repo.DeleteTask(context.Background(), taskID)
	require.NoError(t, err)

	found, err := repo.FindTaskByID(context.Background(), taskID)
	require.NoError(t, err)
	assert.Nil(t, found)
}

func TestTaskRepo_DeleteTask_NotFound(t *testing.T) {
	repo := repositories.NewTaskRepository(testutil.GetTestDB())
	err := repo.DeleteTask(context.Background(), "nonexistent")
	assert.Error(t, err)
}

func TestTaskRepo_InsertActivity(t *testing.T) {
	projectID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "activities") })

	repo := repositories.NewTaskRepository(testutil.GetTestDB())
	activity := &models.Activity{
		ID:        utils.GenerateUUID(),
		Type:      "create",
		DoneBy:    utils.GenerateUUID(),
		ProjectID: projectID,
		Title:     "Task created",
	}
	err := repo.InsertActivity(context.Background(), activity)
	require.NoError(t, err)
}

func TestTaskRepo_FindActivitiesByProjectID(t *testing.T) {
	projectID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "activities") })

	repo := repositories.NewTaskRepository(testutil.GetTestDB())
	for i := 0; i < 3; i++ {
		err := repo.InsertActivity(context.Background(), &models.Activity{
			ID:        utils.GenerateUUID(),
			Type:      "create",
			DoneBy:    utils.GenerateUUID(),
			ProjectID: projectID,
			Title:     "Activity",
		})
		require.NoError(t, err)
	}

	activities, err := repo.FindActivitiesByProjectID(context.Background(), projectID)
	require.NoError(t, err)
	assert.Len(t, activities, 3)
}

func TestTaskRepo_DeleteTasksByProjectID(t *testing.T) {
	projectID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "tasks") })

	repo := repositories.NewTaskRepository(testutil.GetTestDB())
	tasks := make([]*models.Task, 3)
	for i := 0; i < 3; i++ {
		tasks[i] = &models.Task{
			ID:        utils.GenerateUUID(),
			ProjectID: projectID,
			Title:     "Task",
			Status:    "todo",
			CreatedAt: time.Now(),
		}
	}
	err := repo.InsertTasks(context.Background(), tasks)
	require.NoError(t, err)

	err = repo.DeleteTasksByProjectID(context.Background(), projectID)
	require.NoError(t, err)

	found, err := repo.FindTasksByProjectID(context.Background(), projectID)
	require.NoError(t, err)
	assert.Empty(t, found)
}
