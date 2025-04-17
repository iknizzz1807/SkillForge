package utils

// Nhiệm vụ: Cung cấp các hàm hỗ trợ chung (time, string, v.v.)
// Liên kết:
// - Được gọi từ các service hoặc handler khi cần xử lý thời gian, chuỗi
// Vai trò trong flow:
// - Hỗ trợ các tác vụ nhỏ, tái sử dụng trong hệ thống

import (
	"strings"
	"time"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// FormatTime định dạng thời gian thành chuỗi
// Input: t (time.Time), layout (string) - kiểu định dạng (ví dụ: "2006-01-02")
// Return: string - thời gian đã định dạng
func FormatTime(t time.Time, layout string) string {
	return t.Format(layout)
}

// TruncateString cắt ngắn chuỗi nếu vượt quá độ dài
// Input: s (string), maxLen (int) - độ dài tối đa
// Return: string - chuỗi đã cắt ngắn
func TruncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

// Contains kiểm tra xem chuỗi có trong slice không
// Input: slice ([]string), target (string)
// Return: bool - true nếu tồn tại, false nếu không
func Contains(slice []string, target string) bool {
	for _, item := range slice {
		if strings.EqualFold(item, target) {
			return true
		}
	}
	return false
}

// GenerateUUID creates a new random UUID and returns it as a string.
// This function is used to generate unique identifiers for various entities
// in the application, such as projects, users, etc.
//
// Returns:
//   - string: A string representation of the generated UUID
func GenerateUUID() string {
	// Generate a random UUID (version 4)
	id := uuid.New()

	// Return the string representation of the UUID
	return id.String()
}

// HashPassword converts a plain text password into a hash
// Input: password (string) - The plain text password to hash
// Return: string - The hashed password
func HashPassword(password string) string {
	// Generate a hash using bcrypt with default cost
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// In a real application, you might want to handle this error differently
		// For simplicity, we'll return an empty string if hashing fails
		return ""
	}

	// Convert the hashed bytes to string and return
	return string(hashedBytes)
}

// VerifyPassword checks if the provided password matches the stored hash
// Input: hashedPassword (string) - The stored hashed password
//
//	password (string) - The plain text password to verify
//
// Return: bool - true if the password matches, false otherwise
func VerifyPassword(hashedPassword, password string) bool {
	// Compare the stored hash with the provided password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	// Return true if the comparison is successful (no error)
	return err == nil
}

func RemoveDuplicates(strList []string) []string {
    list := []string{}
    for _, item := range strList {
        fmt.Println(item)
		if !Contains(list, item) {
            list = append(list, item)
        }
    }
    return list
}