package main

// // Nhiệm vụ: Khởi tạo dữ liệu mẫu (seeding) cho database để hỗ trợ phát triển, kiểm thử, hoặc triển khai ban đầu.
// // Liên kết:
// // - Sử dụng internal/config để lấy cấu hình (MongoDB URI).
// // - Sử dụng internal/models để định nghĩa dữ liệu mẫu.
// // - Sử dụng internal/repositories để chèn dữ liệu vào MongoDB.
// // Vai trò trong flow:
// // - Được chạy độc lập (go run cmd/seed.go) khi cần khởi tạo dữ liệu ban đầu.
// // - Hỗ trợ developer/test team bằng cách cung cấp dữ liệu mẫu (users, projects) để test tính năng.
// // - Có thể dùng trong production để tạo tài khoản admin hoặc dữ liệu demo.

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/iknizzz1807/SkillForge/internal/config"
// 	"github.com/iknizzz1807/SkillForge/internal/models"
// 	"github.com/iknizzz1807/SkillForge/internal/repositories"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func main() {
// 	// Load cấu hình từ biến môi trường
// 	cfg := config.Load()

// 	// Kết nối đến MongoDB
// 	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoURI))
// 	if err != nil {
// 		log.Fatalf("Failed to connect to MongoDB: %v", err)
// 	}
// 	defer client.Disconnect(context.Background())

// 	db := client.Database("skillforge")

// 	// Khởi tạo repositories
// 	userRepo := repositories.NewUserRepository(db)
// 	projectRepo := repositories.NewProjectRepository(db)

// 	// Chạy các hàm seed
// 	seedUsers(userRepo)
// 	seedProjects(projectRepo)

// 	fmt.Println("Database seeded successfully!")
// }

// // seedUsers chèn dữ liệu mẫu cho users
// // Input: userRepo (*repositories.UserRepository)
// // Return: Không trả về giá trị, chỉ in log nếu lỗi
// func seedUsers(userRepo *repositories.UserRepository) {
// 	// Định nghĩa dữ liệu mẫu: 2 sinh viên, 1 doanh nghiệp
// 	users := []models.User{
// 		{
// 			ID:        "student1",
// 			Email:     "student1@example.com",
// 			Role:      "student",
// 			Skills:    []string{"JavaScript", "React"},
// 			CreatedAt: time.Now(),
// 		},
// 		{
// 			ID:        "student2",
// 			Email:     "student2@example.com",
// 			Role:      "student",
// 			Skills:    []string{"Python", "Django"},
// 			CreatedAt: time.Now(),
// 		},
// 		{
// 			ID:        "business1",
// 			Email:     "business1@example.com",
// 			Role:      "business",
// 			Skills:    []string{},
// 			CreatedAt: time.Now(),
// 		},
// 	}

// 	// Chèn từng user vào database
// 	for _, user := range users {
// 		err := userRepo.InsertUser(context.Background(), &user)
// 		if err != nil {
// 			log.Printf("Failed to seed user %s: %v", user.ID, err)
// 		} else {
// 			fmt.Printf("Seeded user: %s\n", user.ID)
// 		}
// 	}
// }

// // seedProjects chèn dữ liệu mẫu cho projects
// // Input: projectRepo (*repositories.ProjectRepository)
// // Return: Không trả về giá trị, chỉ in log nếu lỗi
// func seedProjects(projectRepo *repositories.ProjectRepository) {
// 	// Định nghĩa dữ liệu mẫu: 2 dự án
// 	projects := []models.Project{
// 		{
// 			ID:          "project1",
// 			Title:       "Build a React App",
// 			Description: "Create a simple React app with user authentication",
// 			Skills:      []string{"JavaScript", "React"},
// 			Timeline:    "2 weeks",
// 			CreatedBy:   "business1",
// 			CreatedAt:   time.Now(),
// 		},
// 		{
// 			ID:          "project2",
// 			Title:       "Develop a REST API",
// 			Description: "Build a REST API using Django",
// 			Skills:      []string{"Python", "Django"},
// 			Timeline:    "3 weeks",
// 			CreatedBy:   "business1",
// 			CreatedAt:   time.Now(),
// 		},
// 	}

// 	// Chèn từng project vào database
// 	for _, project := range projects {
// 		err := projectRepo.InsertProject(context.Background(), &project)
// 		if err != nil {
// 			log.Printf("Failed to seed project %s: %v", project.ID, err)
// 		} else {
// 			fmt.Printf("Seeded project: %s\n", project.ID)
// 		}
// 	}
// }
