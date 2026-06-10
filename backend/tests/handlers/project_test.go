package handlers_test

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/handlers"
	"github.com/iknizzz1807/SkillForge/internal/integrations"
	"github.com/iknizzz1807/SkillForge/internal/services"
	"github.com/iknizzz1807/SkillForge/tests/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newProjectHandler() *handlers.ProjectHandler {
	projService := services.NewProjectService(helpers.TestDB, nil, nil, nil, nil)
	return handlers.NewProjectHandler(projService)
}

func TestCreateProject_Success(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	handler := newProjectHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/projects", gin.H{
		"title":       "New Project",
		"description": "A test project",
		"skills":      []string{"Go", "React"},
		"start_time":  time.Now(),
		"end_time":    time.Now().Add(30 * 24 * time.Hour),
		"max_member":  5,
		"difficulty":  "intermediate",
	})
	helpers.SetAuthContext(c, business)

	handler.CreateProject(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "New Project", resp["title"])
}

func TestCreateProject_NoAuth(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	handler := newProjectHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/projects", gin.H{
		"title":       "New Project",
		"description": "A test project",
		"skills":      []string{"Go"},
		"start_time":  time.Now(),
		"end_time":    time.Now().Add(30 * 24 * time.Hour),
		"max_member":  5,
		"difficulty":  "intermediate",
	})

	handler.CreateProject(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestCreateProject_InvalidData(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	handler := newProjectHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/projects", gin.H{
		"title": "",
	})
	helpers.SetAuthContext(c, business)

	handler.CreateProject(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetProjects(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newProjectHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/projects", nil)
	helpers.SetAuthContext(c, business)

	handler.GetProjects(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var projects []interface{}
	json.Unmarshal(w.Body.Bytes(), &projects)
	assert.Len(t, projects, 1)
}

func TestGetProjectMarketplace(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newProjectHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/projects/"+projectID, nil)
	helpers.SetAuthContext(c, business)
	c.Params = []gin.Param{{Key: "id", Value: projectID}}

	handler.GetProjectMarketplace(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NotEmpty(t, resp["project"])
}

func TestGetProjectMarketplace_NotFound(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	handler := newProjectHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/projects/nonexistent", nil)
	helpers.SetAuthContext(c, business)
	c.Params = []gin.Param{{Key: "id", Value: "nonexistent"}}

	handler.GetProjectMarketplace(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetProjectByStudent(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	helpers.AddStudentToProject(t, projectID, student.ID)
	handler := newProjectHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/projects/student", nil)
	helpers.SetAuthContext(c, student)

	handler.GetProjectByStudent(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var projects []interface{}
	json.Unmarshal(w.Body.Bytes(), &projects)
	assert.Len(t, projects, 1)
}

func TestGetProjectByBusiness(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newProjectHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/projects/business", nil)
	helpers.SetAuthContext(c, business)

	handler.GetProjectByBusiness(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var projects []interface{}
	json.Unmarshal(w.Body.Bytes(), &projects)
	assert.Len(t, projects, 1)
}

func TestUpdateProject_Success(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newProjectHandler()

	c, w := helpers.SetupAuthContext("PUT", "/api/projects/"+projectID, gin.H{
		"title": "Updated Title",
	})
	helpers.SetAuthContext(c, business)
	c.Params = []gin.Param{{Key: "id", Value: projectID}}

	handler.UpdateProject(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Updated Title", resp["title"])
}

func TestUpdateProject_NotOwner(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	owner := helpers.CreateTestUser(t, "business")
	other := helpers.CreateTestUser(t, "business")
	projectID := helpers.CreateTestProject(t, owner.ID, owner.Name)
	handler := newProjectHandler()

	c, w := helpers.SetupAuthContext("PUT", "/api/projects/"+projectID, gin.H{
		"title": "Hacked Title",
	})
	helpers.SetAuthContext(c, other)
	c.Params = []gin.Param{{Key: "id", Value: projectID}}

	handler.UpdateProject(c)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestDeleteProject_Success(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newProjectHandler()

	c, w := helpers.SetupAuthContext("DELETE", "/api/projects/"+projectID, nil)
	helpers.SetAuthContext(c, business)
	c.Params = []gin.Param{{Key: "id", Value: projectID}}

	handler.DeleteProject(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteProject_NotOwner(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	owner := helpers.CreateTestUser(t, "business")
	other := helpers.CreateTestUser(t, "business")
	projectID := helpers.CreateTestProject(t, owner.ID, owner.Name)
	handler := newProjectHandler()

	c, w := helpers.SetupAuthContext("DELETE", "/api/projects/"+projectID, nil)
	helpers.SetAuthContext(c, other)
	c.Params = []gin.Param{{Key: "id", Value: projectID}}

	handler.DeleteProject(c)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestDeleteProject_NotFound(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	handler := newProjectHandler()

	c, w := helpers.SetupAuthContext("DELETE", "/api/projects/nonexistent", nil)
	helpers.SetAuthContext(c, business)
	c.Params = []gin.Param{{Key: "id", Value: "nonexistent"}}

	handler.DeleteProject(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestRemoveStudentFromProject(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	helpers.AddStudentToProject(t, projectID, student.ID)
	handler := newProjectHandler()

	c, w := helpers.SetupAuthContext("DELETE", "/api/projects/"+projectID+"/students/"+student.ID, nil)
	helpers.SetAuthContext(c, business)
	c.Params = []gin.Param{
		{Key: "projectID", Value: projectID},
		{Key: "studentID", Value: student.ID},
	}

	handler.RemoveStudentFromProject(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAddStudentToProject_Service(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)

	projService := services.NewProjectService(helpers.TestDB,
		helpers.NewNotificationService(),
		integrations.NewAIClient("http://localhost:8000"),
		helpers.NewBadgeService(),
		helpers.NewGamificationService(),
	)

	err := projService.AddStudentToProject(projectID, student.ID)
	require.NoError(t, err)
}

func TestAddStudentToProject_FullProject(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)

	projService := services.NewProjectService(helpers.TestDB,
		helpers.NewNotificationService(),
		integrations.NewAIClient("http://localhost:8000"),
		helpers.NewBadgeService(),
		helpers.NewGamificationService(),
	)

	helpers.AddStudentToProject(t, projectID, "student1")
	helpers.AddStudentToProject(t, projectID, "student2")
	helpers.AddStudentToProject(t, projectID, "student3")
	helpers.AddStudentToProject(t, projectID, "student4")
	helpers.AddStudentToProject(t, projectID, "student5")

	err := projService.AddStudentToProject(projectID, "student6")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "maximum capacity")
}

func TestGetStudentsByProject(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	helpers.AddStudentToProject(t, projectID, student.ID)
	handler := newProjectHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/projects/"+projectID+"/students", nil)
	helpers.SetAuthContext(c, business)
	c.Params = []gin.Param{{Key: "id", Value: projectID}}

	handler.GetStudentsByProject(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var students []interface{}
	json.Unmarshal(w.Body.Bytes(), &students)
	assert.Len(t, students, 1)
}
