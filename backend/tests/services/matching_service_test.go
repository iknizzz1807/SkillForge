package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/integrations"
	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/services"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"github.com/iknizzz1807/SkillForge/tests/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMatchingSvcTest(t *testing.T) (*services.MatchingService, string, string, string) {
	t.Helper()
	db := testutil.GetTestDB()
	testutil.CleanAll(db)

	userRepo := repositories.NewUserRepository(db)
	student := &models.User{
		ID:    utils.GenerateUUID(),
		Name:  "Test Student",
		Email: "matchsvc@test.com",
		Role:  "student",
		Skills: []string{"Go", "Python", "MongoDB"},
	}
	require.NoError(t, userRepo.InsertUser(context.Background(), student))

	business := &models.User{
		ID:    utils.GenerateUUID(),
		Name:  "Test Business",
		Email: "biz@test.com",
		Role:  "business",
	}
	require.NoError(t, userRepo.InsertUser(context.Background(), business))

	projectRepo := repositories.NewProjectRepository(db)
	project := &models.Project{
		ID:          utils.GenerateUUID(),
		Title:       "Go Backend Project",
		Description: "Build a Go backend with MongoDB",
		Skills:      []string{"Go", "MongoDB"},
		CreatedByID: business.ID,
		Status:      "open",
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(30 * 24 * time.Hour),
		MaxMember:   5,
	}
	require.NoError(t, projectRepo.InsertProject(context.Background(), project))

	aiClient := integrations.NewAIClient(testutil.GetAIServerURL())
	matchingService := services.NewMatchingService(db, aiClient)

	return matchingService, student.ID, project.ID, business.ID
}

func TestGetScore_Success(t *testing.T) {
	svc, userID, projectID, _ := setupMatchingSvcTest(t)
	score, err := svc.GetScore(userID, projectID)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, score, float64(0))
}

func TestGetScore_NoSkillsUser(t *testing.T) {
	db := testutil.GetTestDB()
	testutil.CleanAll(db)

	userRepo := repositories.NewUserRepository(db)
	student := &models.User{
		ID:    utils.GenerateUUID(),
		Name:  "No Skill",
		Email: "noskill@test.com",
		Role:  "student",
	}
	require.NoError(t, userRepo.InsertUser(context.Background(), student))

	project := &models.Project{
		ID:     utils.GenerateUUID(),
		Title:  "Any Project",
		Skills: []string{"Go"},
	}
	projectRepo := repositories.NewProjectRepository(db)
	require.NoError(t, projectRepo.InsertProject(context.Background(), project))

	aiClient := integrations.NewAIClient(testutil.GetAIServerURL())
	matchingService := services.NewMatchingService(db, aiClient)

	score, err := matchingService.GetScore(student.ID, project.ID)
	require.NoError(t, err)
	assert.Equal(t, float64(0), score)
}

func TestGetScore_UserNotFound(t *testing.T) {
	svc, _, projectID, _ := setupMatchingSvcTest(t)
	_, err := svc.GetScore("nonexistent-user", projectID)
	assert.Error(t, err)
}

func TestGetTopProjectMatchesForStudent(t *testing.T) {
	svc, userID, _, bizID := setupMatchingSvcTest(t)

	db := testutil.GetTestDB()
	projectRepo := repositories.NewProjectRepository(db)
	for i := 0; i < 2; i++ {
		p := &models.Project{
			ID:          utils.GenerateUUID(),
			Title:       "Project",
			Description: "Desc",
			Skills:      []string{"Go", "Python"},
			CreatedByID: bizID,
			Status:      "open",
			StartTime:   time.Now(),
			EndTime:     time.Now().Add(30 * 24 * time.Hour),
			MaxMember:   5,
		}
		require.NoError(t, projectRepo.InsertProject(context.Background(), p))
	}

	matches, err := svc.GetTopProjectMatchesForStudent(userID)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(matches), 1)
}

func TestGetScoreUserProjects(t *testing.T) {
	db := testutil.GetTestDB()
	testutil.CleanAll(db)

	userRepo := repositories.NewUserRepository(db)
	student := &models.User{
		ID:    utils.GenerateUUID(),
		Name:  "Bulk Score",
		Email: "bulk@test.com",
		Role:  "student",
		Skills: []string{"Go", "React"},
	}
	require.NoError(t, userRepo.InsertUser(context.Background(), student))

	projectRepo := repositories.NewProjectRepository(db)
	projects := make([]*models.Project, 3)
	for i := 0; i < 3; i++ {
		projects[i] = &models.Project{
			ID:     utils.GenerateUUID(),
			Title:  "Project",
			Skills: []string{"Go", "React"},
		}
		require.NoError(t, projectRepo.InsertProject(context.Background(), projects[i]))
	}

	aiClient := integrations.NewAIClient(testutil.GetAIServerURL())
	matchingService := services.NewMatchingService(db, aiClient)

	scores, err := matchingService.GetScoreUserProjects(student.ID, projects)
	require.NoError(t, err)
	assert.Len(t, scores, 3)
}
