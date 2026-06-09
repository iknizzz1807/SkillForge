package config

// Nhiệm vụ: Load biến môi trường từ file .env hoặc hệ thống, trả về struct Config chứa toàn bộ cấu hình
// Liên kết:
// - Dùng github.com/joho/godotenv để đọc file .env
// - Trả về struct Env từ env.go
// Vai trò trong flow:
// - Được gọi từ cmd/main.go hoặc internal/app/app.go để lấy cấu hình trước khi khởi động ứng dụng
// - Cung cấp thông tin cần thiết (MongoDB URI, API keys, v.v.) cho các thành phần khác
// Hàm chính: Load()

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Load trả về struct Env chứa toàn bộ cấu hình của ứng dụng
// Return: *Env - con trỏ đến struct chứa các biến môi trường
func Load() *Env {
	// Load file .env nếu có (không panic nếu không tìm thấy file)
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	emailPort, _ := strconv.Atoi(getEnv("EMAIL_PORT", "587"))

	cfg := &Env{
		MongoURI: getEnv("MONGO_URI", "mongodb://localhost:27017"),

		EmailHost: getEnv("EMAIL_HOST", "smtp.example.com"),
		EmailPort: emailPort,
		EmailUser: getEnv("EMAIL_USER", "user@example.com"),
		EmailPass: getEnv("EMAIL_PASS", "password"),

		AIURL: getEnv("AI_URL", "http://ai:5000"),

		GitHubToken: getEnv("GITHUB_TOKEN", ""),

		StripeKey: getEnv("STRIPE_KEY", ""),

		StorageConfig: StorageConfig{
			AWSAccessKey: getEnv("AWS_ACCESS_KEY", ""),
			AWSSecretKey: getEnv("AWS_SECRET_KEY", ""),
			AWSRegion:    getEnv("AWS_REGION", "us-east-1"),
			AWSBucket:    getEnv("AWS_BUCKET", "skillforge-storage"),
		},

		JWTSecret:      getEnv("JWT_SECRET", ""),
		JWTExpiryHours: getEnvInt("JWT_EXPIRY_HOURS", "24"),
	}

	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}

	return cfg
}

// getEnv lấy giá trị biến môi trường, trả về giá trị mặc định nếu không tồn tại
// Input: key (tên biến), defaultValue (giá trị mặc định)
// Return: string - giá trị của biến môi trường
func getEnv(key, defaultValue string) string {
	// Lấy giá trị từ os.Getenv
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	// Trả về giá trị mặc định nếu không tìm thấy
	return defaultValue
}

// getEnvInt lấy giá trị biến môi trường dạng số nguyên, trả về giá trị mặc định nếu không tồn tại hoặc không phải số
// Input: key (tên biến), defaultValue (giá trị mặc định dạng string)
// Return: int - giá trị số của biến môi trường
func getEnvInt(key, defaultValue string) int {
	val, err := strconv.Atoi(getEnv(key, defaultValue))
	if err != nil {
		val, _ = strconv.Atoi(defaultValue)
	}
	return val
}
