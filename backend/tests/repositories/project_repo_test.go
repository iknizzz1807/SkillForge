package repositories_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"github.com/iknizzz1807/SkillForge/tests/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProjectRepo_CreateProject(t *testing.T) {
	projectID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "projects") })

	repo := repositories.NewProjectRepository(testutil.GetTestDB())
	project := &models.Project{
		ID:            projectID,
		Title:         "Test Project",
		Description:   "A test project",
		Skills:        []string{"Go", "React"},
		MaxMember:     5,
		CurrentMember: 0,
		Difficulty:    "intermediate",
		CreatedByID:   utils.GenerateUUID(),
		CreatedByName: "Test Business",
		Status:        "open",
		CreatedAt:     time.Now(),
	}

	err := repo.InsertProject(context.Background(), project)
	require.NoError(t, err)

	found, err := repo.FindProjectByID(context.Background(), projectID)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Equal(t, "Test Project", found.Title)
	assert.Equal(t, "open", found.Status)
}

func TestProjectRepo_FindProjectByID_NotFound(t *testing.T) {
	repo := repositories.NewProjectRepository(testutil.GetTestDB())
	project, err := repo.FindProjectByID(context.Background(), "nonexistent")
	if err != nil && !strings.Contains(err.Error(), "not found") {
		project, err = repo.FindProjectByID(context.Background(), "nonexistent")
	}
	assert.Error(t, err)
	assert.Nil(t, project)
	assert.Contains(t, err.Error(), "not found")
}

func TestProjectRepo_FindAllProjects(t *testing.T) {
	t.Cleanup(func() { testutil.CleanupDB(t, "projects") })

	businessID := utils.GenerateUUID()

	repo := repositories.NewProjectRepository(testutil.GetTestDB())
	for i := 0; i < 3; i++ {
		project := &models.Project{
			ID:            utils.GenerateUUID(),
			Title:         "Project " + string(rune('A'+i)),
			Description:   "Desc",
			Skills:        []string{"Go"},
			MaxMember:     5,
			CurrentMember: 0,
			Difficulty:    "beginner",
			CreatedByID:   businessID,
			CreatedByName: "Biz",
			Status:        "open",
			CreatedAt:     time.Now(),
		}
		err := repo.InsertProject(context.Background(), project)
		require.NoError(t, err)
	}

	projects, err := repo.FindAllProjects(context.Background(), 1, 10)
	require.NoError(t, err)
	assert.Len(t, projects, 3)
}

func TestProjectRepo_UpdateProject(t *testing.T) {
	projectID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "projects") })

	repo := repositories.NewProjectRepository(testutil.GetTestDB())
	project := &models.Project{
		ID:            projectID,
		Title:         "Original Title",
		Description:   "Original Desc",
		Skills:        []string{"Go"},
		MaxMember:     5,
		CurrentMember: 0,
		Difficulty:    "beginner",
		CreatedByID:   utils.GenerateUUID(),
		CreatedByName: "Biz",
		Status:        "open",
		CreatedAt:     time.Now(),
	}
	err := repo.InsertProject(context.Background(), project)
	require.NoError(t, err)

	project.Title = "Updated Title"
	project.Description = "Updated Desc"
	project.Status = "close"

	updated, err := repo.UpdateProject(project)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, "Updated Title", updated.Title)
	assert.Equal(t, "Updated Desc", updated.Description)
	assert.Equal(t, "close", updated.Status)
}

func TestProjectRepo_DeleteProject(t *testing.T) {
	projectID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "projects") })

	repo := repositories.NewProjectRepository(testutil.GetTestDB())
	project := &models.Project{
		ID:            projectID,
		Title:         "To Delete",
		Description:   "Desc",
		Skills:        []string{"Go"},
		MaxMember:     5,
		CurrentMember: 0,
		Difficulty:    "beginner",
		CreatedByID:   utils.GenerateUUID(),
		CreatedByName: "Biz",
		Status:        "open",
		CreatedAt:     time.Now(),
	}
	err := repo.InsertProject(context.Background(), project)
	require.NoError(t, err)

	err = repo.DeleteProject(projectID)
	require.NoError(t, err)

	found, err := repo.FindProjectByID(context.Background(), projectID)
	assert.Error(t, err)
	assert.Nil(t, found)
}

func TestProjectRepo_DeleteProject_NotFound(t *testing.T) {
	repo := repositories.NewProjectRepository(testutil.GetTestDB())
	err := repo.DeleteProject("nonexistent")
	assert.Error(t, err)
}

func TestProjectRepo_FindProjectsByIDs(t *testing.T) {
	id1 := utils.GenerateUUID()
	id2 := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "projects") })

	repo := repositories.NewProjectRepository(testutil.GetTestDB())
	for _, id := range []string{id1, id2} {
		err := repo.InsertProject(context.Background(), &models.Project{
			ID:            id,
			Title:         "Project " + id[:8],
			Description:   "Batch test",
			Skills:        []string{"Go"},
			MaxMember:     5,
			CurrentMember: 0,
			Difficulty:    "beginner",
			CreatedByID:   utils.GenerateUUID(),
			CreatedByName: "Biz",
			Status:        "open",
			CreatedAt:     time.Now(),
		})
		require.NoError(t, err)
	}

	projects, err := repo.FindProjectsByIDs(context.Background(), []string{id1, id2})
	require.NoError(t, err)
	assert.Len(t, projects, 2)
}

func TestProjectRepo_FindProjectsByIDs_Empty(t *testing.T) {
	repo := repositories.NewProjectRepository(testutil.GetTestDB())
	projects, err := repo.FindProjectsByIDs(context.Background(), []string{})
	require.NoError(t, err)
	assert.Nil(t, projects)
}

func TestProjectRepo_GetProjectIDsByCreatorID(t *testing.T) {
	businessID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "projects") })

	repo := repositories.NewProjectRepository(testutil.GetTestDB())
	for i := 0; i < 2; i++ {
		err := repo.InsertProject(context.Background(), &models.Project{
			ID:            utils.GenerateUUID(),
			Title:         "Biz Project",
			Description:   "Desc",
			Skills:        []string{"Go"},
			MaxMember:     5,
			CurrentMember: 0,
			Difficulty:    "beginner",
			CreatedByID:   businessID,
			CreatedByName: "Biz",
			Status:        "open",
			CreatedAt:     time.Now(),
		})
		require.NoError(t, err)
	}

	ids, err := repo.GetProjectIDsByCreatorID(businessID)
	require.NoError(t, err)
	assert.Len(t, ids, 2)
}
