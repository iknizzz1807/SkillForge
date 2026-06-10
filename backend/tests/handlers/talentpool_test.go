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

func newTalentPoolHandler() *handlers.TalentPoolHandler {
	tpService := services.NewTalentPoolService(helpers.TestDB)
	return handlers.NewTalentPoolHandler(tpService)
}

func TestAddToTalentPool(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	handler := newTalentPoolHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/talentpool/"+student.ID, nil)
	helpers.SetAuthContext(c, business)
	c.Params = []gin.Param{{Key: "id", Value: student.ID}}

	handler.AddStudentToTalentPool(c)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestAddToTalentPool_NonBusiness(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	student := helpers.CreateTestUser(t, "student")
	handler := newTalentPoolHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/talentpool/"+student.ID, nil)
	helpers.SetAuthContext(c, student)
	c.Params = []gin.Param{{Key: "id", Value: student.ID}}

	handler.AddStudentToTalentPool(c)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestGetTalentPool(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	handler := newTalentPoolHandler()

	c, _ := helpers.SetupAuthContext("POST", "/api/talentpool/"+student.ID, nil)
	helpers.SetAuthContext(c, business)
	c.Params = []gin.Param{{Key: "id", Value: student.ID}}
	handler.AddStudentToTalentPool(c)

	c2, w2 := helpers.SetupAuthContext("GET", "/api/talentpool", nil)
	helpers.SetAuthContext(c2, business)

	handler.GetTalentPool(c2)

	assert.Equal(t, http.StatusOK, w2.Code)
	var pool []interface{}
	json.Unmarshal(w2.Body.Bytes(), &pool)
	assert.Len(t, pool, 1)
}

func TestGetTalentPool_NonBusiness(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	student := helpers.CreateTestUser(t, "student")
	handler := newTalentPoolHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/talentpool", nil)
	helpers.SetAuthContext(c, student)

	handler.GetTalentPool(c)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestRemoveFromTalentPool(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	handler := newTalentPoolHandler()

	c, _ := helpers.SetupAuthContext("POST", "/api/talentpool/"+student.ID, nil)
	helpers.SetAuthContext(c, business)
	c.Params = []gin.Param{{Key: "id", Value: student.ID}}
	handler.AddStudentToTalentPool(c)

	c2, w2 := helpers.SetupAuthContext("DELETE", "/api/talentpool/"+student.ID, nil)
	helpers.SetAuthContext(c2, business)
	c2.Params = []gin.Param{{Key: "id", Value: student.ID}}

	handler.RemoveFromTalentPool(c2)

	assert.Equal(t, http.StatusOK, w2.Code)
}

func TestRemoveFromTalentPool_NonBusiness(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	student := helpers.CreateTestUser(t, "student")
	handler := newTalentPoolHandler()

	c, w := helpers.SetupAuthContext("DELETE", "/api/talentpool/"+student.ID, nil)
	helpers.SetAuthContext(c, student)
	c.Params = []gin.Param{{Key: "id", Value: student.ID}}

	handler.RemoveFromTalentPool(c)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestCheckStudentInTalentPool(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	business := helpers.CreateTestUser(t, "business")
	student := helpers.CreateTestUser(t, "student")
	handler := newTalentPoolHandler()

	c, _ := helpers.SetupAuthContext("POST", "/api/talentpool/"+student.ID, nil)
	helpers.SetAuthContext(c, business)
	c.Params = []gin.Param{{Key: "id", Value: student.ID}}
	handler.AddStudentToTalentPool(c)

	c2, w2 := helpers.SetupAuthContext("GET", "/api/talentpool/"+student.ID+"/check", nil)
	helpers.SetAuthContext(c2, business)
	c2.Params = []gin.Param{{Key: "id", Value: student.ID}}

	handler.CheckStudentInTalentPool(c2)

	assert.Equal(t, http.StatusOK, w2.Code)
	var resp map[string]interface{}
	json.Unmarshal(w2.Body.Bytes(), &resp)
	assert.True(t, resp["is_in_pool"].(bool))
}
