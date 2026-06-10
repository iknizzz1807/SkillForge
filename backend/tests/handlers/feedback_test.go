package handlers_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/handlers"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/services"
	"github.com/iknizzz1807/SkillForge/tests/helpers"
	"github.com/stretchr/testify/assert"
)

func newFeedbackHandler() *handlers.FeedbackHandler {
	fbRepo := repositories.NewFeedbackRepository(helpers.TestDB)
	projRepo := repositories.NewProjectRepository(helpers.TestDB)
	userRepo := repositories.NewUserRepository(helpers.TestDB)
	fbService := services.NewFeedbackService(fbRepo, projRepo, userRepo)
	return handlers.NewFeedbackHandler(fbService)
}

func TestCreateFeedback(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	helpers.AddStudentToProject(t, projectID, student.ID)
	handler := newFeedbackHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/feedbacks", gin.H{
		"project_id": projectID,
		"to_id":      student.ID,
		"type":       "student",
		"rating":     5,
		"content":    "Great work!",
	})
	helpers.SetAuthContext(c, business)

	handler.CreateFeedback(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, float64(5), resp["rating"])
}

func TestCreateFeedback_SelfFeedback(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newFeedbackHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/feedbacks", gin.H{
		"project_id": projectID,
		"to_id":      business.ID,
		"type":       "business",
		"rating":     5,
		"content":    "Great work!",
	})
	helpers.SetAuthContext(c, business)

	handler.CreateFeedback(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateFeedback_InvalidData(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	student := helpers.CreateTestUser(t, "student")
	handler := newFeedbackHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/feedbacks", gin.H{})
	helpers.SetAuthContext(c, student)

	handler.CreateFeedback(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetStudentFeedbacks(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	helpers.AddStudentToProject(t, projectID, student.ID)
	handler := newFeedbackHandler()

	c, _ := helpers.SetupAuthContext("POST", "/api/feedbacks", gin.H{
		"project_id": projectID,
		"to_id":      student.ID,
		"type":       "student",
		"rating":     4,
		"content":    "Good work!",
	})
	helpers.SetAuthContext(c, business)
	handler.CreateFeedback(c)

	c2, w2 := helpers.SetupAuthContext("GET", "/api/feedbacks/student/"+student.ID, nil)
	helpers.SetAuthContext(c2, business)
	c2.Params = []gin.Param{{Key: "userID", Value: student.ID}}

	handler.GetStudentFeedbacks(c2)

	assert.Equal(t, http.StatusOK, w2.Code)
	var feedbacks []interface{}
	json.Unmarshal(w2.Body.Bytes(), &feedbacks)
	assert.Len(t, feedbacks, 1)
}

func TestGetBusinessFeedbacks(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	helpers.AddStudentToProject(t, projectID, student.ID)
	handler := newFeedbackHandler()

	c, _ := helpers.SetupAuthContext("POST", "/api/feedbacks", gin.H{
		"project_id": projectID,
		"to_id":      business.ID,
		"type":       "business",
		"rating":     5,
		"content":    "Great mentor!",
	})
	helpers.SetAuthContext(c, student)
	handler.CreateFeedback(c)

	c2, w2 := helpers.SetupAuthContext("GET", "/api/feedbacks/business/"+business.ID, nil)
	helpers.SetAuthContext(c2, student)
	c2.Params = []gin.Param{{Key: "userID", Value: business.ID}}

	handler.GetBusinessFeedbacks(c2)

	assert.Equal(t, http.StatusOK, w2.Code)
	var feedbacks []interface{}
	json.Unmarshal(w2.Body.Bytes(), &feedbacks)
	assert.Len(t, feedbacks, 1)
}

func TestGetStudentFeedbacks_EmptyID(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	handler := newFeedbackHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/feedbacks/student/", nil)
	helpers.SetAuthContext(c, helpers.CreateTestUser(t, "student"))

	handler.GetStudentFeedbacks(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
