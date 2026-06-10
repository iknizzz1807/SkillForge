package handlers_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/iknizzz1807/SkillForge/internal/handlers"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/services"
	"github.com/iknizzz1807/SkillForge/tests/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newAuthHandler() *handlers.AuthHandler {
	userRepo := repositories.NewUserRepository(helpers.TestDB)
	fileService := services.NewFileService(userRepo)
	authService := services.NewAuthService(userRepo, fileService)
	return handlers.NewAuthHandler(authService)
}

func TestRegister_Success(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	handler := newAuthHandler()

	c, w := helpers.SetupMultipartContext("POST", "/api/auth/register", map[string]string{
		"email":    "newuser@test.com",
		"name":     "New User",
		"password": "password123",
		"role":     "student",
		"title":    "Go Developer",
	})

	handler.Register(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.NotEmpty(t, resp["token"])
	assert.NotEmpty(t, resp["user"])
}

func TestRegister_DuplicateEmail(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	_ = helpers.CreateTestUser(t, "student")
	handler := newAuthHandler()

	c, w := helpers.SetupMultipartContext("POST", "/api/auth/register", map[string]string{
		"email":    "student_test@test.com",
		"name":     "Dup User",
		"password": "password123",
		"role":     "student",
		"title":    "Dev",
	})
	// Override email to match the auto-generated one — actually we need to know the email
	// Let's use a different approach: create the user manually and try to register with same email
	helpers.CleanupDB(t)
	u := helpers.CreateTestUser(t, "student")

	c, w = helpers.SetupMultipartContext("POST", "/api/auth/register", map[string]string{
		"email":    u.Email,
		"name":     "Dup User",
		"password": "password123",
		"role":     "student",
		"title":    "Dev",
	})

	handler.Register(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Contains(t, resp["error"], "email already exists")
}

func TestRegister_EmptyFields(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	handler := newAuthHandler()

	c, w := helpers.SetupMultipartContext("POST", "/api/auth/register", map[string]string{
		"email":    "",
		"name":     "",
		"password": "",
		"role":     "",
		"title":    "",
	})

	handler.Register(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestRegister_InvalidRole(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	handler := newAuthHandler()

	c, w := helpers.SetupMultipartContext("POST", "/api/auth/register", map[string]string{
		"email":    "badrole@test.com",
		"name":     "Bad Role",
		"password": "password123",
		"role":     "admin",
		"title":    "Hacker",
	})

	handler.Register(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestRegister_ShortPassword(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	handler := newAuthHandler()

	c, w := helpers.SetupMultipartContext("POST", "/api/auth/register", map[string]string{
		"email":    "shortpw@test.com",
		"name":     "Short PW",
		"password": "123",
		"role":     "student",
		"title":    "Dev",
	})

	handler.Register(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLogin_Success(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	_ = helpers.CreateTestUser(t, "student")
	handler := newAuthHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/auth/login", map[string]string{
		"email":    "student_test@test.com",
		"password": "password123",
	})
	// Override email to match generated — need to get actual email from created user
	helpers.CleanupDB(t)
	u := helpers.CreateTestUser(t, "student")

	c, w = helpers.SetupAuthContext("POST", "/api/auth/login", map[string]string{
		"email":    u.Email,
		"password": "password123",
	})

	handler.Login(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NotEmpty(t, resp["token"])
}

func TestLogin_WrongPassword(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	u := helpers.CreateTestUser(t, "student")
	handler := newAuthHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/auth/login", map[string]string{
		"email":    u.Email,
		"password": "wrongpassword",
	})

	handler.Login(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestLogin_NonexistentUser(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	handler := newAuthHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/auth/login", map[string]string{
		"email":    "nobody@test.com",
		"password": "password123",
	})

	handler.Login(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
