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

func newUserHandler() *handlers.UserHandler {
	userService := services.NewUserService(helpers.TestDB)
	portfolioService := services.NewPortfolioService(helpers.TestDB)
	return handlers.NewUserHandler(userService, portfolioService)
}

func TestGetCurrentUser(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	student := helpers.CreateTestUser(t, "student")
	handler := newUserHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/user", nil)
	helpers.SetAuthContext(c, student)

	handler.GetCurrentUser(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, student.ID, resp["id"])
	assert.Equal(t, student.Name, resp["name"])
}

func TestGetCurrentUser_NotFound(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	handler := newUserHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/user", nil)

	handler.GetCurrentUser(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetUserByID(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	student := helpers.CreateTestUser(t, "student")
	handler := newUserHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/users/"+student.ID, nil)
	helpers.SetAuthContext(c, student)
	c.Params = []gin.Param{{Key: "id", Value: student.ID}}

	handler.GetUserByID(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, student.ID, resp["id"])
}

func TestGetUserByID_NotFound(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	student := helpers.CreateTestUser(t, "student")
	handler := newUserHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/users/nonexistent", nil)
	helpers.SetAuthContext(c, student)
	c.Params = []gin.Param{{Key: "id", Value: "nonexistent"}}

	handler.GetUserByID(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUpdateCurrentUser(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	student := helpers.CreateTestUser(t, "student")
	handler := newUserHandler()

	c, w := helpers.SetupMultipartContext("PUT", "/api/user", map[string]string{
		"name":  "Updated Name",
		"email": student.Email,
		"title": "Senior Dev",
	})
	helpers.SetAuthContext(c, student)

	handler.UpdateCurrentUser(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	user := resp["user"].(map[string]interface{})
	assert.Equal(t, "Updated Name", user["name"])
	assert.NotEmpty(t, resp["token"])
}

func TestUpdateCurrentUser_EmptyName(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	student := helpers.CreateTestUser(t, "student")
	handler := newUserHandler()

	c, w := helpers.SetupMultipartContext("PUT", "/api/user", map[string]string{
		"name":  "",
		"email": student.Email,
	})
	helpers.SetAuthContext(c, student)

	handler.UpdateCurrentUser(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetUserProfile(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	student := helpers.CreateTestUser(t, "student")
	handler := newUserHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/users/"+student.ID+"/profile", nil)
	helpers.SetAuthContext(c, student)
	c.Params = []gin.Param{{Key: "id", Value: student.ID}}

	handler.GetUserProfile(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	user := resp["user"].(map[string]interface{})
	assert.Equal(t, student.ID, user["id"])
}

func TestGetUserProfile_NotFound(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	handler := newUserHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/users/nonexistent/profile", nil)
	c.Params = []gin.Param{{Key: "id", Value: "nonexistent"}}

	handler.GetUserProfile(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetAllStudents(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	s1 := helpers.CreateTestUser(t, "student")
	s2 := helpers.CreateTestUser(t, "student")
	handler := newUserHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/users/students", nil)
	helpers.SetAuthContext(c, s1)

	handler.GetAllStudents(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	students := resp["students"].([]interface{})
	assert.Len(t, students, 2)
	_ = s2
}
