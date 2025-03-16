package middleware

// // Nhiệm vụ: Middleware giới hạn tốc độ request để bảo vệ server
// // Liên kết:
// // - Dùng github.com/gin-contrib/ratelimit để giới hạn request
// // - Gắn vào router trong internal/app/routes.go
// // Vai trò trong flow:
// // - Ngăn chặn abuse API bằng cách giới hạn số request trong khoảng thời gian

// import (
// 	"net/http"
// 	"time"

// 	"github.com/gin-contrib/ratelimit"
// 	"github.com/gin-gonic/gin"
// )

// func RateLimitMiddleware() gin.HandlerFunc {
// 	// Tạo rate limiter: 100 request mỗi phút cho mỗi IP
// 	limiter := ratelimit.NewRateLimiter(time.Minute, 100, func(c *gin.Context) string {
// 		// Dùng IP làm key để giới hạn
// 		return c.ClientIP()
// 	})

// 	return func(c *gin.Context) {
// 		// Kiểm tra giới hạn
// 		if limiter.LimitReached(c) {
// 			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
// 			c.Abort()
// 			return
// 		}

// 		// Tiếp tục xử lý request
// 		c.Next()
// 	}
// }
