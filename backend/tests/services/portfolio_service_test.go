package services_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/services"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"github.com/iknizzz1807/SkillForge/tests/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupPortSvc(t *testing.T) *services.PortfolioService {
	t.Helper()
	db := testutil.GetTestDB()
	testutil.CleanAll(db)
	return services.NewPortfolioService(db)
}

func createPortUser(t *testing.T, userRepo *repositories.UserRepository) string {
	t.Helper()
	user := &models.User{
		ID:    utils.GenerateUUID(),
		Name:  "Portfolio User",
		Email: "pf@test.com",
		Role:  "student",
		Title: "Developer",
		Skills: []string{"Go", "Python"},
	}
	require.NoError(t, userRepo.InsertUser(context.Background(), user))
	return user.ID
}

func TestPortfolio_Generate(t *testing.T) {
	svc := setupPortSvc(t)
	db := testutil.GetTestDB()
	userRepo := repositories.NewUserRepository(db)
	userID := createPortUser(t, userRepo)

	url, err := svc.GeneratePortfolio(userID)
	require.NoError(t, err)
	assert.Contains(t, url, "/portfolios/")
	assert.Contains(t, url, ".html")

	path := filepath.Join("./storage/portfolios", userID+".html")
	_, err = os.Stat(path)
	assert.NoError(t, err)
	os.Remove(path)
}

func TestPortfolio_NoProjects(t *testing.T) {
	svc := setupPortSvc(t)
	db := testutil.GetTestDB()
	userRepo := repositories.NewUserRepository(db)
	userID := createPortUser(t, userRepo)

	url, err := svc.GeneratePortfolio(userID)
	require.NoError(t, err)
	assert.Contains(t, url, ".html")

	path := filepath.Join("./storage/portfolios", userID+".html")
	_, err = os.Stat(path)
	assert.NoError(t, err)
	os.Remove(path)
}

func TestPortfolio_WithBadges(t *testing.T) {
	db := testutil.GetTestDB()
	testutil.CleanAll(db)
	svc := services.NewPortfolioService(db)
	userRepo := repositories.NewUserRepository(db)
	badgeRepo := repositories.NewBadgeRepository(db)
	badgeSvc := services.NewBadgeService(badgeRepo)

	badge, err := badgeSvc.CreateBadge("Go Expert", "Expert in Go", "skill",
		map[string]interface{}{"skill": "Go"})
	require.NoError(t, err)

	userID := createPortUser(t, userRepo)
	_, err = badgeSvc.AwardBadgeToUser(userID, badge.ID, "")
	require.NoError(t, err)

	url, err := svc.GeneratePortfolio(userID)
	require.NoError(t, err)
	assert.Contains(t, url, ".html")

	path := filepath.Join("./storage/portfolios", userID+".html")
	_, err = os.Stat(path)
	assert.NoError(t, err)
	os.Remove(path)
}

func TestPortfolio_GetPath(t *testing.T) {
	svc := setupPortSvc(t)
	db := testutil.GetTestDB()
	userRepo := repositories.NewUserRepository(db)
	portfolioRepo := repositories.NewPortfolioRepository(db)

	userID := createPortUser(t, userRepo)

	portfolio, err := portfolioRepo.FindPortfolioByUserID(context.Background(), userID)
	require.NoError(t, err)
	assert.Nil(t, portfolio)

	_, err = svc.GetPortfolio(userID)
	assert.Error(t, err)

	svc.GeneratePortfolio(userID)

	path := filepath.Join("./storage/portfolios", userID+".html")
	os.Remove(path)
}
