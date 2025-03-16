package app

// Nhiệm vụ: Đăng ký tất cả các route API của ứng dụng, gắn handlers với các endpoint.
// Liên kết:
// - Sử dụng internal/handlers để lấy các handler (ProjectHandler, UserHandler, v.v.).
// - Nhận các service từ app.go để truyền vào handler.
// Vai trò trong flow:
// - Xác định cách client tương tác với backend qua các endpoint (GET /projects, POST /tasks, v.v.).
// - Là "bản đồ" của API, giúp developer dễ dàng tra cứu các route.

import (
	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/handlers"

	// "github.com/iknizzz1807/SkillForge/internal/integrations"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

func RegisterRoutes(
	r *gin.Engine,
	userService *services.UserService,
	projectService *services.ProjectService,
	applicationService *services.ApplicationService,
	taskService *services.TaskService,
	reviewService *services.ReviewService,
	messageService *services.MessageService,
	portfolioService *services.PortfolioService,
	analyticsService *services.AnalyticsService,
	authService *services.AuthService,
	// paymentClient *integrations.PaymentClient,
) {
	// Khởi tạo các handler với service tương ứng
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)
	projectHandler := handlers.NewProjectHandler(projectService)
	applicationHandler := handlers.NewApplicationHandler(applicationService)
	taskHandler := handlers.NewTaskHandler(taskService)
	reviewHandler := handlers.NewReviewHandler(reviewService)
	messageHandler := handlers.NewMessageHandler(messageService)
	portfolioHandler := handlers.NewPortfolioHandler(portfolioService)
	analyticsHandler := handlers.NewAnalyticsHandler(analyticsService)

	// Định nghĩa các route
	// Nhóm route không cần auth
	r.POST("/auth/register", authHandler.Register)
	r.POST("/auth/login", authHandler.Login)

	// Nhóm route cần auth (dùng middleware nếu cần)
	api := r.Group("/api")
	{
		// User routes
		api.GET("/users/:id", userHandler.GetUser)
		api.PUT("/users/:id", userHandler.UpdateUser)

		// Project routes
		api.GET("/projects", projectHandler.GetProjects)
		api.GET("/projects/:id", projectHandler.GetProject)
		api.POST("/projects", projectHandler.CreateProject)

		// Application routes
		api.POST("/applications", applicationHandler.ApplyProject)
		api.GET("/applications/:id", applicationHandler.GetApplication)

		// Task routes
		api.GET("/tasks/:id", taskHandler.GetTask)
		api.POST("/tasks", taskHandler.CreateTask)
		api.PUT("/tasks/:id", taskHandler.UpdateTask)

		// Review routes
		api.POST("/reviews", reviewHandler.SubmitReview)
		api.GET("/reviews/:id", reviewHandler.GetReview)

		// Message routes
		api.POST("/messages", messageHandler.SendMessage)
		api.GET("/messages/:id", messageHandler.GetMessages)

		// Portfolio routes
		api.GET("/portfolios/:userID", portfolioHandler.GetPortfolio)

		// Analytics routes
		api.GET("/analytics/skills/:userID", analyticsHandler.GetSkillAnalytics)
		api.GET("/analytics/projects/:projectID", analyticsHandler.GetProjectAnalytics)

		// Payment routes (ví dụ)
		api.POST("/payments", func(c *gin.Context) {
			// Gọi paymentClient để xử lý thanh toán
		})
	}
}
