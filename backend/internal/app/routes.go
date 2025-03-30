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

	"github.com/iknizzz1807/SkillForge/internal/integrations"
	"github.com/iknizzz1807/SkillForge/internal/services"
)

// // WebSocket upgrader cấu hình
// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// 	// Cho phép tất cả các origin để dễ phát triển
// 	// Trong production nên giới hạn origin cụ thể
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

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
	notificationService *services.NotificationService,
	// paymentClient *integrations.PaymentClient,
) {
	// Khởi tạo integrations ở đây
	realtimeClient := integrations.NewRealtimeClient()

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
	websocketHanlder := handlers.NewWebSocketHandler(realtimeClient, messageService, notificationService,
		taskService,
		projectService)

	// Định nghĩa các route
	// Nhóm route không cần auth
	r.POST("/auth/register", authHandler.Register) // Tạo user mới
	r.POST("/auth/login", authHandler.Login)       // Đăng nhập (tồn tại user)

	// Route để server web socket client
	// Todo: chia cái này thành handler riêng để dễ xử lý và code sạch, cân nhắc tên là websocket hanlder
	r.GET("/ws", websocketHanlder.HandleConnection)

	// Nhóm route cần auth (dùng middleware nếu cần)
	api := r.Group("/api")
	// Comment cái này nếu cần test api nhanh bằng postman hay thunder client
	// r.Use(middleware.AuthMiddleware())

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

		// Endpoint mới hợp nhất dựa vào role
		api.GET("/applications/me", applicationHandler.GetApplicationsByCurrentUser)

		// Giữ lại các endpoints cũ để tương thích ngược (có thể xóa sau khi cập nhật tất cả frontend calls)
		// Đánh dấu deprecated để team biết sẽ loại bỏ trong tương lai
		api.GET("/applications/user", applicationHandler.GetApplicationsByUser)         // Deprecated
		api.GET("/applications/business", applicationHandler.GetApplicationsByBusiness) // Deprecated

		api.PUT("/applications/:id/status", applicationHandler.UpdateApplicationStatus) // Change this because the route looking ass bro

		// Task routes
		api.GET("/tasks/:id", taskHandler.GetTasksByProjectID) // Get task bằng projectID, id ở đây là projectID chứ không phải taskID
		api.POST("/tasks", taskHandler.CreateTasks)
		api.PUT("/tasks/:id", taskHandler.UpdateTask)
		api.PUT("/tasks/:id/assign", taskHandler.AssignTask)
		api.PUT("/tasks/:id/finish", taskHandler.FinishTask)
		api.DELETE("/tasks/:id", taskHandler.DeleteTask)

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
