package utils

// Nhiệm vụ: Cung cấp các hàm để tạo và lấy secret key cho JWT
// Liên kết:
// - Dùng github.com/golang-jwt/jwt/v5 để tạo token
// - Được gọi từ services/auth_service.go để tạo token
// - Được gọi từ middleware/auth.go để xác minh token
// Vai trò trong flow:
// - Tạo JWT token khi đăng ký/đăng nhập, cung cấp secret key cho xác thực

import (
	"log"
	"sync"
	"time"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iknizzz1807/SkillForge/internal/config"
)

const (
	// defaultJWTSecret là giá trị mặc định nếu không có biến môi trường
	defaultJWTSecret = "skillforge-secret-key"
	defaultJWTExpiry = 24 // hours
)

var (
	jwtSecret string
	jwtExpiry time.Duration
	jwtOnce   sync.Once
)

func initJWT() {
	cfg := config.Load()
	secret := cfg.JWTSecret
	if secret == "" || secret == defaultJWTSecret {
		log.Println("WARNING: JWT_SECRET is empty or using default value. This is insecure for production!")
	}
	if secret == "" {
		secret = defaultJWTSecret
	}
	jwtSecret = secret

	hours := cfg.JWTExpiryHours
	if hours <= 0 {
		hours = defaultJWTExpiry
	}
	jwtExpiry = time.Duration(hours) * time.Hour
}

// GenerateJWT tạo token JWT cho user
// Input: userID (string), role (string)
// Return: string (token), error (nếu có lỗi)
func GenerateJWT(userID, email, name, role string) (string, error) {
	// Tạo claims cho token
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"name":    name,
		"role":    role,
		"exp":     time.Now().Add(GetJWTExpiry()).Unix(),
		"iat":     time.Now().Unix(),
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

// GetJWTSecret lấy secret key từ config (Env struct) hoặc dùng mặc định
// Return: string - secret key
func GetJWTSecret() string {
	jwtOnce.Do(initJWT)
	return jwtSecret
}

// GetJWTExpiry lấy thời gian hết hạn token từ config
// Return: time.Duration - thời gian token có hiệu lực
func GetJWTExpiry() time.Duration {
	jwtOnce.Do(initJWT)
	return jwtExpiry
}

// ParseToken parses and validates a JWT token string
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(GetJWTSecret()), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token claims")
}
