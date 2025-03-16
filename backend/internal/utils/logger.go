package utils

// // Nhiệm vụ: Cấu hình và cung cấp logger cho ứng dụng
// // Liên kết:
// // - Dùng github.com/sirupsen/logrus để ghi log
// // - Được gọi từ middleware/logging.go hoặc các service để ghi log
// // Vai trò trong flow:
// // - Ghi lại thông tin request, error, hoặc sự kiện quan trọng

// import (
// 	"os"

// 	"github.com/sirupsen/logrus"
// )

// var (
// 	// logger là instance logrus dùng chung
// 	logger *logrus.Logger
// )

// // InitLogger khởi tạo logger với cấu hình mặc định
// // Return: *logrus.Logger - con trỏ đến logger
// func InitLogger() *logrus.Logger {
// 	if logger == nil {
// 		// Tạo logger mới
// 		logger = logrus.New()

// 		// Thiết lập output (stdout)
// 		logger.SetOutput(os.Stdout)

// 		// Thiết lập format (JSON để dễ parse)
// 		logger.SetFormatter(&logrus.JSONFormatter{})

// 		// Thiết lập level (Info là mặc định)
// 		logger.SetLevel(logrus.InfoLevel)
// 	}

// 	// Trả về logger
// 	return logger
// }

// // GetLogger lấy instance logger đã khởi tạo
// // Return: *logrus.Logger - con trỏ đến logger
// func GetLogger() *logrus.Logger {
// 	if logger == nil {
// 		return InitLogger()
// 	}
// 	return logger
// }
