package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/app"
	"github.com/iknizzz1807/SkillForge/internal/integrations"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/services"
	"github.com/iknizzz1807/SkillForge/tests/testutil"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter(t *testing.T) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)
	db := testutil.GetTestDB()
	testutil.CleanAll(db)

	realtimeClient := integrations.NewRealtimeClient()
	aiClient := integrations.NewAIClient("http://localhost:9999")
	notifSvc := services.NewNotificationService(realtimeClient, db, testutil.NewDummyEmailClient())
	badgeSvc := services.NewBadgeService(repositories.NewBadgeRepository(db))
	gamificationSvc := services.NewGamificationService(
		repositories.NewGamificationRepository(db),
		services.NewUserService(db),
	)

	r := gin.New()
	app.RegisterRoutes(r,
		services.NewUserService(db),
		services.NewProjectService(db, notifSvc, aiClient, badgeSvc, gamificationSvc),
		services.NewApplicationService(db, notifSvc, badgeSvc, gamificationSvc),
		services.NewTaskService(db, notifSvc, gamificationSvc),
		services.NewReviewService(db),
		services.NewPortfolioService(db),
		services.NewAnalyticsService(db),
		services.NewAuthService(repositories.NewUserRepository(db), services.NewFileService(repositories.NewUserRepository(db))),
		notifSvc,
		badgeSvc,
		services.NewTalentPoolService(db),
		services.NewFileService(repositories.NewUserRepository(db)),
		services.NewBusinessInfoService(repositories.NewBusinessInfoRepository(db), repositories.NewUserRepository(db)),
		services.NewFeedbackService(repositories.NewFeedbackRepository(db), repositories.NewProjectRepository(db), repositories.NewUserRepository(db)),
		gamificationSvc,
		services.NewMatchingService(db, aiClient),
		realtimeClient,
		services.NewChatService(repositories.NewChatRepository(db)),
		services.NewInvitationService(db, badgeSvc, gamificationSvc),
	)
	return r
}

func TestCORSMiddleware(t *testing.T) {
	r := setupTestRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", "/api/projects", nil)
	req.Header.Set("Origin", "http://localhost:5173")
	req.Header.Set("Access-Control-Request-Method", "GET")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "http://localhost:5173", w.Header().Get("Access-Control-Allow-Origin"))
}

func TestRecoveryMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}))
	r.GET("/panic", func(c *gin.Context) {
		panic("test panic")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/panic", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "internal server error")
}

func TestPublicRoutes_Exist(t *testing.T) {
	r := setupTestRouter(t)
	routes := []string{"/auth/login", "/auth/register"}
	for _, route := range routes {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", route, nil)
		r.ServeHTTP(w, req)
		assert.NotEqual(t, http.StatusNotFound, w.Code,
			"route %s should exist", route)
	}
}

func TestProtectedRoutes_RequireAuth(t *testing.T) {
	r := setupTestRouter(t)
	routes := []string{
		"/api/user", "/api/projects", "/api/chats",
		"/api/levels", "/api/business-info", "/api/portfolios", "/api/matches",
	}
	for _, route := range routes {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", route, nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Code,
			"route %s should require auth (got %d)", route, w.Code)
	}
}

func TestRouteRegistration_CorrectPaths(t *testing.T) {
	r := setupTestRouter(t)
	routes := map[string]string{
		"POST /auth/register":        "/auth/register",
		"POST /auth/login":           "/auth/login",
		"GET /api/user":              "/api/user",
		"GET /api/projects":          "/api/projects",
		"GET /api/chats":             "/api/chats",
		"GET /api/levels":            "/api/levels",
		"GET /api/business-info":     "/api/business-info",
		"GET /api/portfolios":        "/api/portfolios",
		"GET /api/matches":           "/api/matches",
	}
	for name, path := range routes {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", path, nil)
		r.ServeHTTP(w, req)
		assert.NotEqual(t, http.StatusNotFound, w.Code,
			"route %s (%s) should be registered", name, path)
	}
}
