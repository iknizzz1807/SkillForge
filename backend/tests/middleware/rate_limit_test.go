package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/middleware"
	"github.com/stretchr/testify/assert"
)

func setupRateLimitTest() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/api/test", nil)
	return c, w
}

func TestRateLimit_UnderLimit(t *testing.T) {
	for i := 0; i < 50; i++ {
		c, w := setupRateLimitTest()
		c.Request.RemoteAddr = "10.0.0.1:12345"

		middleware.RateLimitMiddleware()(c)

		assert.Equal(t, http.StatusOK, w.Code, "request %d should be allowed", i+1)
	}
}

func TestRateLimit_LimitExact(t *testing.T) {
	for i := 0; i < 100; i++ {
		c, w := setupRateLimitTest()
		c.Request.RemoteAddr = "10.0.0.2:12345"

		middleware.RateLimitMiddleware()(c)

		assert.Equal(t, http.StatusOK, w.Code, "request %d should be allowed", i+1)
	}
}

func TestRateLimit_DifferentIPs(t *testing.T) {
	for i := 0; i < 10; i++ {
		c, w := setupRateLimitTest()
		c.Request.RemoteAddr = "10.0.0.100:12345"

		middleware.RateLimitMiddleware()(c)
		assert.Equal(t, http.StatusOK, w.Code)

		c2, w2 := setupRateLimitTest()
		c2.Request.RemoteAddr = "10.0.0.200:12345"

		middleware.RateLimitMiddleware()(c2)
		assert.Equal(t, http.StatusOK, w2.Code)
	}
}

func TestRateLimit_SameUserDifferentIP(t *testing.T) {
	userID := "test-user-rate"
	for i := 0; i < 10; i++ {
		c, w := setupRateLimitTest()
		c.Request.RemoteAddr = "10.0.0.50:12345"
		c.Set("userID", userID)

		middleware.RateLimitMiddleware()(c)
		assert.Equal(t, http.StatusOK, w.Code)
	}
}
