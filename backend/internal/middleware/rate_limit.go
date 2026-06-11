package middleware

// Nhiệm vụ: Middleware giới hạn tốc độ request để bảo vệ server
// Liên kết:
// - Gắn vào router trong internal/app/routes.go
// Vai trò trong flow:
// - Ngăn chặn abuse API bằng cách giới hạn số request trong khoảng thời gian

import (
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type visitor struct {
	count    int
	lastSeen time.Time
}

type rateLimiter struct {
	mu       sync.Mutex
	visitors map[string]*visitor
	rate     int
	interval time.Duration
}

var limiter = &rateLimiter{
	visitors: make(map[string]*visitor),
	rate:     100,
	interval: time.Minute,
}

func init() {
	go limiter.cleanup()
}

func (rl *rateLimiter) cleanup() {
	for {
		time.Sleep(time.Minute)
		rl.mu.Lock()
		for ip, v := range rl.visitors {
			if time.Since(v.lastSeen) > rl.interval {
				delete(rl.visitors, ip)
			}
		}
		rl.mu.Unlock()
	}
}

func (rl *rateLimiter) key(ip, userID string) string {
	if userID != "" {
		return "user:" + userID
	}
	return "ip:" + ip
}

func (rl *rateLimiter) allow(ip, userID string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	k := rl.key(ip, userID)
	v, exists := rl.visitors[k]
	if !exists {
		rl.visitors[k] = &visitor{count: 1, lastSeen: time.Now()}
		return true
	}

	if time.Since(v.lastSeen) > rl.interval {
		v.count = 1
		v.lastSeen = time.Now()
		return true
	}

	v.count++
	v.lastSeen = time.Now()
	return v.count <= rl.rate
}

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if os.Getenv("DISABLE_RATE_LIMIT") == "true" {
			c.Next()
			return
		}
		ip := c.ClientIP()
		userID, _ := c.Get("userID")
		uid, _ := userID.(string)
		if !limiter.allow(ip, uid) {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}
		c.Next()
	}
}
