package middleware

// // Nhiệm vụ: Middleware ghi log request và response
// // Liên kết:
// // - Dùng github.com/iknizzz1807/SkillForge/internal/utils để lấy logger
// // - Gắn vào router trong internal/app/routes.go
// // Vai trò trong flow:
// // - Ghi lại thông tin mỗi request (method, path, status, v.v.) để debug

// import (
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/iknizzz1807/SkillForge/internal/utils"
// )

// func LoggingMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Lấy logger từ utils
// 		logger := utils.GetLogger()

// 		// Ghi lại thời gian bắt đầu
// 		start := time.Now()

// 		// Lấy thông tin request
// 		method := c.Request.Method
// 		path := c.Request.URL.Path

// 		// Tiếp tục xử lý request
// 		c.Next()

// 		// Tính latency
// 		latency := time.Since(start)

// 		// Lấy status code
// 		status := c.Writer.Status()

// 		// Ghi log
// 		logger.WithFields(map[string]interface{}{
// 			"method":  method,
// 			"path":    path,
// 			"status":  status,
// 			"latency": latency.String(),
// 			"ip":      c.ClientIP(),
// 		}).Info("Request processed")
// 	}
// }
