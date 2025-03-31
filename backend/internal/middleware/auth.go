package middleware

// Nhiệm vụ: Middleware xác thực JWT cho các route cần bảo vệ
// Liên kết:
// - Dùng github.com/golang-jwt/jwt/v5 để xác minh token
// - Dùng github.com/iknizzz1807/SkillForge/internal/utils để lấy secret key
// - Gắn vào router trong internal/app/routes.go
// Vai trò trong flow:
// - Kiểm tra token trong header Authorization, gán userID và role vào context

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/iknizzz1807/SkillForge/internal/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy token từ header Authorization (Bearer <token>)
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Tách "Bearer " để lấy token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}
		tokenString := parts[1]

		// Parse và xác minh token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Kiểm tra phương thức ký
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			// Trả về secret key từ utils
			return []byte(utils.GetJWTSecret()), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Lấy claims từ token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Gán userID và role vào context để handler sử dụng
		userID, _ := claims["user_id"].(string)
		role, _ := claims["role"].(string)
		userName, _ := claims["name"].(string)
		c.Set("userID", userID)
		c.Set("role", role)
		c.Set("name", userName)

		// Tiếp tục xử lý request
		c.Next()
	}
}
