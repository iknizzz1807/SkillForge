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

func newInvitationHandler() *handlers.InvitationHandler {
	invService := services.NewInvitationService(helpers.TestDB, nil, nil)
	return handlers.NewInvitationHandler(invService)
}

func newInvitationHandlerFull() *handlers.InvitationHandler {
	invService := services.NewInvitationService(helpers.TestDB, helpers.NewBadgeService(), helpers.NewGamificationService())
	return handlers.NewInvitationHandler(invService)
}

func TestCreateInvitation(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newInvitationHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/invitations", gin.H{
		"student_id": student.ID,
		"project_id": projectID,
	})
	helpers.SetAuthContext(c, business)

	handler.CreateInvitation(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "pending", resp["status"])
}

func TestCreateInvitation_Duplicate(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newInvitationHandler()

	c, _ := helpers.SetupAuthContext("POST", "/api/invitations", gin.H{
		"student_id": student.ID,
		"project_id": projectID,
	})
	helpers.SetAuthContext(c, business)
	handler.CreateInvitation(c)

	c2, w2 := helpers.SetupAuthContext("POST", "/api/invitations", gin.H{
		"student_id": student.ID,
		"project_id": projectID,
	})
	helpers.SetAuthContext(c2, business)
	handler.CreateInvitation(c2)

	assert.Equal(t, http.StatusBadRequest, w2.Code)
}

func TestCreateInvitation_FullProject(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newInvitationHandlerFull()

	helpers.AddStudentToProject(t, projectID, "s1")
	helpers.AddStudentToProject(t, projectID, "s2")
	helpers.AddStudentToProject(t, projectID, "s3")
	helpers.AddStudentToProject(t, projectID, "s4")
	helpers.AddStudentToProject(t, projectID, "s5")

	c, w := helpers.SetupAuthContext("POST", "/api/invitations", gin.H{
		"student_id": student.ID,
		"project_id": projectID,
	})
	helpers.SetAuthContext(c, business)
	handler.CreateInvitation(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateInvitation_InvalidInput(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	handler := newInvitationHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/invitations", gin.H{})
	helpers.SetAuthContext(c, business)

	handler.CreateInvitation(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetMyInvitations(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newInvitationHandler()

	c, _ := helpers.SetupAuthContext("POST", "/api/invitations", gin.H{
		"student_id": student.ID,
		"project_id": projectID,
	})
	helpers.SetAuthContext(c, business)
	handler.CreateInvitation(c)

	c2, w2 := helpers.SetupAuthContext("GET", "/api/invitations", nil)
	helpers.SetAuthContext(c2, student)
	handler.GetMyInvitations(c2)

	assert.Equal(t, http.StatusOK, w2.Code)
	var invites []interface{}
	json.Unmarshal(w2.Body.Bytes(), &invites)
	assert.Len(t, invites, 1)
}

func TestRespondToInvitation_Accept(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newInvitationHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/invitations", gin.H{
		"student_id": student.ID,
		"project_id": projectID,
	})
	helpers.SetAuthContext(c, business)
	handler.CreateInvitation(c)

	var createResp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &createResp)
	invitationID := createResp["id"].(string)

	c2, w2 := helpers.SetupAuthContext("PUT", "/api/invitations/"+invitationID+"/respond", gin.H{
		"status": "accepted",
	})
	helpers.SetAuthContext(c2, student)
	c2.Params = []gin.Param{{Key: "id", Value: invitationID}}
	handler.RespondToInvitation(c2)

	assert.Equal(t, http.StatusOK, w2.Code)
}

func TestRespondToInvitation_Reject(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newInvitationHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/invitations", gin.H{
		"student_id": student.ID,
		"project_id": projectID,
	})
	helpers.SetAuthContext(c, business)
	handler.CreateInvitation(c)

	var createResp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &createResp)
	invitationID := createResp["id"].(string)

	c2, w2 := helpers.SetupAuthContext("PUT", "/api/invitations/"+invitationID+"/respond", gin.H{
		"status": "rejected",
	})
	helpers.SetAuthContext(c2, student)
	c2.Params = []gin.Param{{Key: "id", Value: invitationID}}
	handler.RespondToInvitation(c2)

	assert.Equal(t, http.StatusOK, w2.Code)
}

func TestRespondToInvitation_AlreadyResponded(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newInvitationHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/invitations", gin.H{
		"student_id": student.ID,
		"project_id": projectID,
	})
	helpers.SetAuthContext(c, business)
	handler.CreateInvitation(c)

	var createResp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &createResp)
	invitationID := createResp["id"].(string)

	c2, _ := helpers.SetupAuthContext("PUT", "/api/invitations/"+invitationID+"/respond", gin.H{
		"status": "accepted",
	})
	helpers.SetAuthContext(c2, student)
	c2.Params = []gin.Param{{Key: "id", Value: invitationID}}
	handler.RespondToInvitation(c2)

	c3, w3 := helpers.SetupAuthContext("PUT", "/api/invitations/"+invitationID+"/respond", gin.H{
		"status": "rejected",
	})
	helpers.SetAuthContext(c3, student)
	c3.Params = []gin.Param{{Key: "id", Value: invitationID}}
	handler.RespondToInvitation(c3)

	assert.Equal(t, http.StatusBadRequest, w3.Code)
}

func TestRespondToInvitation_NotForYou(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student1 := helpers.CreateTestUser(t, "student")
	student2 := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newInvitationHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/invitations", gin.H{
		"student_id": student1.ID,
		"project_id": projectID,
	})
	helpers.SetAuthContext(c, business)
	handler.CreateInvitation(c)

	var createResp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &createResp)
	invitationID := createResp["id"].(string)

	c2, w2 := helpers.SetupAuthContext("PUT", "/api/invitations/"+invitationID+"/respond", gin.H{
		"status": "accepted",
	})
	helpers.SetAuthContext(c2, student2)
	c2.Params = []gin.Param{{Key: "id", Value: invitationID}}
	handler.RespondToInvitation(c2)

	assert.Equal(t, http.StatusBadRequest, w2.Code)
}
