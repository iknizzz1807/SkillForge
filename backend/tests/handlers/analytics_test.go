package handlers_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/handlers"
	"github.com/iknizzz1807/SkillForge/internal/services"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"github.com/iknizzz1807/SkillForge/tests/helpers"
	"github.com/stretchr/testify/assert"
)

func newAnalyticsHandler() *handlers.AnalyticsHandler {
	analyticsService := services.NewAnalyticsService(helpers.TestDB)
	return handlers.NewAnalyticsHandler(analyticsService)
}

func TestGetSkillAnalytics(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	student := helpers.CreateTestUser(t, "student")

	_, err := helpers.TestDB.Collection("analytics").InsertOne(context.Background(), map[string]interface{}{
		"_id":          utils.GenerateUUID(),
		"user_id":      student.ID,
		"skill":        "Go",
		"proficiency":  85,
		"last_updated": time.Now(),
	})
	assert.NoError(t, err)

	handler := newAnalyticsHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/analytics/skills/"+student.ID, nil)
	c.Params = []gin.Param{{Key: "userID", Value: student.ID}}

	handler.GetSkillAnalytics(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var skills []interface{}
	json.Unmarshal(w.Body.Bytes(), &skills)
	assert.Len(t, skills, 1)
}

func TestGetSkillAnalytics_Empty(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	handler := newAnalyticsHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/analytics/skills/nonexistent", nil)
	c.Params = []gin.Param{{Key: "userID", Value: "nonexistent"}}

	handler.GetSkillAnalytics(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var skills []interface{}
	json.Unmarshal(w.Body.Bytes(), &skills)
	assert.Len(t, skills, 0)
}

func TestGetProjectAnalytics(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)

	_, err := helpers.TestDB.Collection("analytics").InsertOne(context.Background(), map[string]interface{}{
		"_id":              utils.GenerateUUID(),
		"project_id":       projectID,
		"completion_rate":  75.0,
		"average_score":    4.5,
		"last_updated":     time.Now(),
	})
	assert.NoError(t, err)

	handler := newAnalyticsHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/analytics/projects/"+projectID, nil)
	c.Params = []gin.Param{{Key: "projectID", Value: projectID}}

	handler.GetProjectAnalytics(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, projectID, resp["project_id"])
}

func TestGetProjectAnalytics_NotFound(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	handler := newAnalyticsHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/analytics/projects/nonexistent", nil)
	c.Params = []gin.Param{{Key: "projectID", Value: "nonexistent"}}

	handler.GetProjectAnalytics(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
