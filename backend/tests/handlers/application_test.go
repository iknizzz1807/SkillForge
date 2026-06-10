package handlers_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/handlers"
	"github.com/iknizzz1807/SkillForge/internal/services"
	"github.com/iknizzz1807/SkillForge/tests/helpers"
	"github.com/stretchr/testify/assert"
)

func newApplicationHandler() *handlers.ApplicationHandler {
	appService := services.NewApplicationService(helpers.TestDB, nil, nil, nil)
	return handlers.NewApplicationHandler(appService)
}

func newApplicationHandlerFull() *handlers.ApplicationHandler {
	appService := services.NewApplicationService(helpers.TestDB, helpers.NewNotificationService(), helpers.NewBadgeService(), helpers.NewGamificationService())
	return handlers.NewApplicationHandler(appService)
}

func TestApplyProject(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newApplicationHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/applications", gin.H{
		"project_id":        projectID,
		"motivation":        "I love this project!",
		"detailed_proposal": "I will build the backend in Go",
	})
	helpers.SetAuthContext(c, student)

	handler.ApplyProject(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "pending", resp["status"])
}

func TestApplyProject_AlreadyApplied(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newApplicationHandler()

	c, _ := helpers.SetupAuthContext("POST", "/api/applications", gin.H{
		"project_id":        projectID,
		"motivation":        "I love this project!",
		"detailed_proposal": "I will build the backend in Go",
	})
	helpers.SetAuthContext(c, student)
	handler.ApplyProject(c)

	c2, w2 := helpers.SetupAuthContext("POST", "/api/applications", gin.H{
		"project_id":        projectID,
		"motivation":        "I still love it!",
		"detailed_proposal": "I will do even better",
	})
	helpers.SetAuthContext(c2, student)
	handler.ApplyProject(c2)

	assert.Equal(t, http.StatusInternalServerError, w2.Code)
	var resp map[string]interface{}
	json.Unmarshal(w2.Body.Bytes(), &resp)
	assert.Contains(t, resp["error"], "already applied")
}

func TestApplyProject_NonStudent(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newApplicationHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/applications", gin.H{
		"project_id":        projectID,
		"motivation":        "I love this project!",
		"detailed_proposal": "I will build the backend in Go",
	})
	helpers.SetAuthContext(c, business)

	handler.ApplyProject(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetApplication(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newApplicationHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/applications", gin.H{
		"project_id":        projectID,
		"motivation":        "I love this project!",
		"detailed_proposal": "I will build the backend in Go",
	})
	helpers.SetAuthContext(c, student)
	handler.ApplyProject(c)

	var createResp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &createResp)
	appID := createResp["id"].(string)

	c2, w2 := helpers.SetupAuthContext("GET", "/api/applications/"+appID, nil)
	helpers.SetAuthContext(c2, student)
	c2.Params = []gin.Param{{Key: "id", Value: appID}}

	handler.GetApplication(c2)

	assert.Equal(t, http.StatusOK, w2.Code)
}

func TestGetApplication_NoAccess(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	otherStudent := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newApplicationHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/applications", gin.H{
		"project_id":        projectID,
		"motivation":        "I love this project!",
		"detailed_proposal": "Will build it",
	})
	helpers.SetAuthContext(c, student)
	handler.ApplyProject(c)

	var createResp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &createResp)
	appID := createResp["id"].(string)

	c2, w2 := helpers.SetupAuthContext("GET", "/api/applications/"+appID, nil)
	helpers.SetAuthContext(c2, otherStudent)
	c2.Params = []gin.Param{{Key: "id", Value: appID}}

	handler.GetApplication(c2)

	assert.Equal(t, http.StatusForbidden, w2.Code)
}

func TestUpdateApplicationStatus_Approve(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newApplicationHandlerFull()

	c, w := helpers.SetupAuthContext("POST", "/api/applications", gin.H{
		"project_id":        projectID,
		"motivation":        "I love this project!",
		"detailed_proposal": "Will build it",
	})
	helpers.SetAuthContext(c, student)
	handler.ApplyProject(c)

	var createResp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &createResp)
	appID := createResp["id"].(string)

	c2, w2 := helpers.SetupAuthContext("PUT", "/api/applications/"+appID+"/status", gin.H{
		"status": "approved",
	})
	helpers.SetAuthContext(c2, business)
	c2.Params = []gin.Param{{Key: "id", Value: appID}}

	handler.UpdateApplicationStatus(c2)

	assert.Equal(t, http.StatusOK, w2.Code)
}

func TestUpdateApplicationStatus_Reject(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newApplicationHandlerFull()

	c, w := helpers.SetupAuthContext("POST", "/api/applications", gin.H{
		"project_id":        projectID,
		"motivation":        "I love this project!",
		"detailed_proposal": "Will build it",
	})
	helpers.SetAuthContext(c, student)
	handler.ApplyProject(c)

	var createResp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &createResp)
	appID := createResp["id"].(string)

	c2, w2 := helpers.SetupAuthContext("PUT", "/api/applications/"+appID+"/status", gin.H{
		"status": "rejected",
	})
	helpers.SetAuthContext(c2, business)
	c2.Params = []gin.Param{{Key: "id", Value: appID}}

	handler.UpdateApplicationStatus(c2)

	assert.Equal(t, http.StatusOK, w2.Code)
}

func TestUpdateApplicationStatus_NonBusiness(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newApplicationHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/applications", gin.H{
		"project_id":        projectID,
		"motivation":        "I love this project!",
		"detailed_proposal": "Will build it",
	})
	helpers.SetAuthContext(c, student)
	handler.ApplyProject(c)

	var createResp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &createResp)
	appID := createResp["id"].(string)

	c2, w2 := helpers.SetupAuthContext("PUT", "/api/applications/"+appID+"/status", gin.H{
		"status": "approved",
	})
	helpers.SetAuthContext(c2, student)
	c2.Params = []gin.Param{{Key: "id", Value: appID}}

	handler.UpdateApplicationStatus(c2)

	assert.Equal(t, http.StatusForbidden, w2.Code)
}
