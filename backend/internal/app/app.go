package app

// Nhiệm vụ: Khởi tạo toàn bộ ứng dụng, kết nối các thành phần (Gin, MongoDB, integrations), và chạy server.
// Liên kết:
// - Sử dụng internal/config để lấy cấu hình.
// - Khởi tạo các integrations (email, realtime, storage, v.v.).
// - Khởi tạo services và truyền vào handlers.
// - Gọi routes.go để đăng ký các route API.
// Vai trò trong flow:
// - Là "bộ điều phối chính" của backend, đảm bảo mọi thành phần sẵn sàng trước khi nhận request.
// - Được gọi từ cmd/main.go để khởi động server.

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/config"

	// "github.com/iknizzz1807/SkillForge/internal/handlers"
	"github.com/iknizzz1807/SkillForge/internal/integrations"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/services"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Run() {
	// Load cấu hình từ biến môi trường
	cfg := config.Load()

	// Kết nối MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())
	db := client.Database("skillforge")

	// Khởi tạo các integrations
	emailClient := integrations.NewEmailClient(cfg.EmailHost, cfg.EmailPort, cfg.EmailUser, cfg.EmailPass)
	realtimeClient := integrations.NewRealtimeClient()
	// storageClient := integrations.NewStorageClient(cfg.StorageConfig)
	aiClient := integrations.NewAIClient(cfg.AIURL)
	githubClient := integrations.NewGitHubClient(cfg.GitHubToken)
	// paymentClient := integrations.NewPaymentClient(cfg.StripeKey)
	webrtcClient := integrations.NewWebRTCClient()

	// Khởi tạo các services
	notificationService := services.NewNotificationService(realtimeClient, db)
	userService := services.NewUserService(db)
	projectService := services.NewProjectService(db, notificationService, aiClient, githubClient, emailClient)
	applicationService := services.NewApplicationService(db, notificationService, emailClient)
	taskService := services.NewTaskService(db, notificationService)
	reviewService := services.NewReviewService(db)
	messageService := services.NewMessageService(db, realtimeClient, webrtcClient)
	portfolioService := services.NewPortfolioService(db)
	analyticsService := services.NewAnalyticsService(db)
	authService := services.NewAuthService(repositories.NewUserRepository(db), services.NewFileService(repositories.NewUserRepository(db)))
	badgeService := services.NewBadgeService(repositories.NewBadgeRepository(db))
	talentPoolService := services.NewTalentPoolService(db)
	fileService := services.NewFileService(repositories.NewUserRepository(db))
	businessInfoService := services.NewBusinessInfoService(repositories.NewBusinessInfoRepository(db), userService.GetUserRepository())
	feedbackService := services.NewFeedbackService(repositories.NewFeedbackRepository(db))
	gamificationService := services.NewGamificationService(repositories.NewGamificationRepository(db), userService)
	matchingService := services.NewMatchingService(db, aiClient)
	// Khởi tạo Gin router
	r := gin.Default()

	// Đăng ký middleware (nếu cần)
	// Ví dụ: r.Use(middleware.LogMiddleware())

	// Đăng ký các route từ routes.go
	RegisterRoutes(r, userService, projectService, applicationService, taskService, reviewService, messageService, portfolioService, analyticsService, authService, notificationService, badgeService, talentPoolService, fileService, businessInfoService, feedbackService, gamificationService, matchingService, realtimeClient)

	// Chạy server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
