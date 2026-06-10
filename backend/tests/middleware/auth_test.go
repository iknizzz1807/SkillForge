package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/iknizzz1807/SkillForge/internal/middleware"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("JWT_SECRET", "test-secret-key-for-skillforge-tests-123456")
}

func setupAuthTest() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	return c, w
}

func TestAuthMiddleware_ValidToken(t *testing.T) {
	c, w := setupAuthTest()
	token, _ := utils.GenerateJWT("user123", "user@test.com", "Test User", "student")
	c.Request, _ = http.NewRequest("GET", "/api/test", nil)
	c.Request.Header.Set("Authorization", "Bearer "+token)

	middleware.AuthMiddleware()(c)

	assert.Equal(t, http.StatusOK, w.Code)
	userID, _ := c.Get("userID")
	role, _ := c.Get("role")
	name, _ := c.Get("name")
	assert.Equal(t, "user123", userID)
	assert.Equal(t, "student", role)
	assert.Equal(t, "Test User", name)
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	c, w := setupAuthTest()
	c.Request, _ = http.NewRequest("GET", "/api/test", nil)
	c.Request.Header.Set("Authorization", "Bearer invalidtoken123")

	middleware.AuthMiddleware()(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid or expired token")
}

func TestAuthMiddleware_EmptyToken(t *testing.T) {
	c, w := setupAuthTest()
	c.Request, _ = http.NewRequest("GET", "/api/test", nil)
	c.Request.Header.Set("Authorization", "Bearer ")

	middleware.AuthMiddleware()(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid or expired token")
}

func TestAuthMiddleware_WrongPrefix(t *testing.T) {
	c, w := setupAuthTest()
	c.Request, _ = http.NewRequest("GET", "/api/test", nil)
	c.Request.Header.Set("Authorization", "Basic token123")

	middleware.AuthMiddleware()(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid token format")
}

func TestAuthMiddleware_ExpiredToken(t *testing.T) {
	c, w := setupAuthTest()
	claims := jwt.MapClaims{
		"user_id": "user123",
		"role":    "student",
		"exp":     time.Now().Add(-24 * time.Hour).Unix(),
		"iat":     time.Now().Add(-48 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString([]byte(utils.GetJWTSecret()))
	c.Request, _ = http.NewRequest("GET", "/api/test", nil)
	c.Request.Header.Set("Authorization", "Bearer "+tokenStr)

	middleware.AuthMiddleware()(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid or expired token")
}

func TestAuthMiddleware_MissingHeader(t *testing.T) {
	c, w := setupAuthTest()
	c.Request, _ = http.NewRequest("GET", "/api/test", nil)

	middleware.AuthMiddleware()(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Authorization header required")
}

func TestAuthMiddleware_DoubleSpaceAfterBearer(t *testing.T) {
	c, w := setupAuthTest()
	c.Request, _ = http.NewRequest("GET", "/api/test", nil)
	c.Request.Header.Set("Authorization", "Bearer  ")

	middleware.AuthMiddleware()(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid token format")
}
