package middleware

// Nhiệm vụ: Middleware ghi log request và response
// Liên kết:
// - Gắn vào router trong internal/app/routes.go
// Vai trò trong flow:
// - Ghi lại thông tin mỗi request (method, path, status, v.v.) để debug

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		method := c.Request.Method
		path := c.Request.URL.Path

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		log.Printf("[%d] %s %s - %s (%s)", status, method, path, latency, c.ClientIP())
	}
}
