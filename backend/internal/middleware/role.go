package middleware

// // Nhiệm vụ: Middleware kiểm tra vai trò user để phân quyền
// // Liên kết:
// // - Dùng sau AuthMiddleware để lấy role từ context
// // - Gắn vào router trong internal/app/routes.go cho các route cần phân quyền
// // Vai trò trong flow:
// // - Đảm bảo chỉ user với vai trò phù hợp (student, business, mentor, admin) mới truy cập được route

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/iknizzz1807/SkillForge/internal/constants"
// )

// func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Lấy role từ context (được gán bởi AuthMiddleware)
// 		role, exists := c.Get("role")
// 		if !exists {
// 			c.JSON(http.StatusForbidden, gin.H{"error": "Role not found in context"})
// 			c.Abort()
// 			return
// 		}

// 		// Chuyển role thành string
// 		userRole, ok := role.(string)
// 		if !ok {
// 			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid role format"})
// 			c.Abort()
// 			return
// 		}

// 		// Kiểm tra xem role của user có trong danh sách allowedRoles không
// 		roleAllowed := false
// 		for _, allowed := range allowedRoles {
// 			if userRole == allowed {
// 				roleAllowed = true
// 				break
// 			}
// 		}

// 		// Nếu role không được phép, trả về lỗi
// 		if !roleAllowed {
// 			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
// 			c.Abort()
// 			return
// 		}

// 		// Tiếp tục xử lý request
// 		c.Next()
// 	}
// }

// // Ví dụ sử dụng trong routes.go:
// // api.GET("/projects", RoleMiddleware(constants.RoleBusiness), projectHandler.GetProjects)
// // Chỉ cho phép role "business" truy cập
