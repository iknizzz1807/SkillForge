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

func newChatHandler() *handlers.ChatHandler {
	chatRepo := repositories.NewChatRepository(helpers.TestDB)
	chatService := services.NewChatService(chatRepo)
	return handlers.NewChatHandler(chatService)
}

func TestGetGroups(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	helpers.AddStudentToProject(t, projectID, student.ID)
	handler := newChatHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/chat/groups", nil)
	helpers.SetAuthContext(c, student)

	handler.GetGroups(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var groups []interface{}
	json.Unmarshal(w.Body.Bytes(), &groups)
	assert.Len(t, groups, 1)
}

func TestGetGroups_Unauthorized(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	handler := newChatHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/chat/groups", nil)

	handler.GetGroups(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestGetGroupInfo(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	helpers.AddStudentToProject(t, projectID, student.ID)
	handler := newChatHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/chat/groups/"+projectID, nil)
	helpers.SetAuthContext(c, student)
	c.Params = []gin.Param{{Key: "id", Value: projectID}}

	handler.GetGroupInfo(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Contains(t, resp, "members")
	assert.Contains(t, resp, "messages")
}

func TestGetGroupInfo_NoAccess(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	otherUser := helpers.CreateTestUser(t, "student")
	projectID := helpers.CreateTestProject(t, business.ID, business.Name)
	handler := newChatHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/chat/groups/"+projectID, nil)
	helpers.SetAuthContext(c, otherUser)
	c.Params = []gin.Param{{Key: "id", Value: projectID}}

	handler.GetGroupInfo(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetGroupInfo_MissingID(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	student := helpers.CreateTestUser(t, "student")
	handler := newChatHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/chat/groups/", nil)
	helpers.SetAuthContext(c, student)

	handler.GetGroupInfo(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUploadChatFile_NoFile(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	student := helpers.CreateTestUser(t, "student")
	handler := newChatHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/chat/upload", nil)
	helpers.SetAuthContext(c, student)

	handler.UploadChatFile(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
