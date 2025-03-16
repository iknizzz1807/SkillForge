package utils

// Nhiệm vụ: Cung cấp các hàm để tạo và lấy secret key cho JWT
// Liên kết:
// - Dùng github.com/golang-jwt/jwt/v5 để tạo token
// - Được gọi từ services/auth_service.go để tạo token
// - Được gọi từ middleware/auth.go để xác minh token
// Vai trò trong flow:
// - Tạo JWT token khi đăng ký/đăng nhập, cung cấp secret key cho xác thực

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	// jwtSecretKeyEnv là tên biến môi trường chứa secret key
	jwtSecretKeyEnv = "JWT_SECRET"
	// defaultJWTSecret là giá trị mặc định nếu không có biến môi trường
	defaultJWTSecret = "skillforge-secret-key"
)

// GenerateJWT tạo token JWT cho user
// Input: userID (string), role (string)
// Return: string (token), error (nếu có lỗi)
func GenerateJWT(userID, role string) (string, error) {
	// Tạo claims cho token
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Hết hạn sau 24 giờ
		"iat":     time.Now().Unix(),                     // Thời gian phát hành
	}

	// Tạo token với claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Ký token với secret key
	tokenString, err := token.SignedString([]byte(GetJWTSecret()))
	if err != nil {
		return "", err
	}

	// Trả về token
	return tokenString, nil
}

// GetJWTSecret lấy secret key từ biến môi trường hoặc dùng mặc định
// Return: string - secret key
func GetJWTSecret() string {
	// Lấy secret từ biến môi trường
	secret := os.Getenv(jwtSecretKeyEnv)
	if secret == "" {
		// Dùng giá trị mặc định nếu không tìm thấy
		return defaultJWTSecret
	}
	// Trả về secret từ biến môi trường
	return secret
}
