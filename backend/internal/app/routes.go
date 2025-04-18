package app

// Nhiệm vụ: Đăng ký tất cả các route API của ứng dụng, gắn handlers với các endpoint.
// Liên kết:
// - Sử dụng internal/handlers để lấy các handler (ProjectHandler, UserHandler, v.v.).
// - Nhận các service từ app.go để truyền vào handler.
// Vai trò trong flow:
// - Xác định cách client tương tác với backend qua các endpoint (GET /projects, POST /tasks, v.v.).
// - Là "bản đồ" của API, giúp developer dễ dàng tra cứu các route.

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/handlers"
	"github.com/iknizzz1807/SkillForge/internal/middleware"

	"github.com/iknizzz1807/SkillForge/internal/integrations"
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
	notificationService *services.NotificationService,
	badgeService *services.BadgeService,
	talentPoolService *services.TalentPoolService,
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
	badgeHandler := handlers.NewBadgeHandler(badgeService)
	talentPoolHandler := handlers.NewTalentPoolHandler(talentPoolService)
	websocketHanlder := handlers.NewWebSocketHandler(realtimeClient, messageService, notificationService,
		taskService,
		projectService)

	// Định nghĩa các route
	// Nhóm route không cần auth
	r.POST("/auth/register", authHandler.Register) // Tạo user mới
	r.POST("/auth/login", authHandler.Login)       // Đăng nhập (tồn tại user)

	// Route để server web socket client
	r.GET("/ws", websocketHanlder.HandleConnection)

	// Comment cái này nếu cần test api nhanh bằng postman hay thunder client
	r.Use(middleware.AuthMiddleware())

	// CORS type shit
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://skillforge.ikniz.site"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Nhóm route cần auth (dùng middleware nếu cần)
	api := r.Group("/api")

	{
		// User routes
		api.GET("/users/:id", userHandler.GetUser)
		api.GET("/users/name/:id", userHandler.GetUsername)
		api.PUT("/users/:id", userHandler.UpdateUser)

		// Project routes
		api.GET("/projects", projectHandler.GetProjects)
		api.GET("/projects/:id", projectHandler.GetProject)
		api.GET("/projects/busiess/:id", projectHandler.GetProjectByBusiness) // Với id params là id của business
		api.POST("/projects", projectHandler.CreateProject)
		api.PUT("/projects/:id", projectHandler.UpdateProject)
		api.DELETE("/projects/:id", projectHandler.DeleteProject)
		// Change this route because it looks like ass bruh
		api.DELETE("/projects/:id/students/:studentID", projectHandler.RemoveStudentFromProject)

		// Application routes
		api.POST("/applications", applicationHandler.ApplyProject)
		api.GET("/applications/:id", applicationHandler.GetApplication)
		api.GET("/applications/me", applicationHandler.GetApplicationsByCurrentUser)
		// This update for status is used for accept or reject an application, if accept add the student to the project
		api.PUT("/applications/:id/status", applicationHandler.UpdateApplicationStatus) // Change this because the route looking ass bro
		api.DELETE("/applications/:id", applicationHandler.DeleteApplication)

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

		// TalentPool route
		// id là id của user, còn id của business thì được nhận vào bằng context
		api.GET("/talenpool", talentPoolHandler.GetTalentPool)
		api.POST("/talentpool:id", talentPoolHandler.AddStudentToTalentPool)
		api.DELETE("/talentpool/:id", talentPoolHandler.RemoveFromTalentPool)

		// Analytics routes
		api.GET("/analytics/skills/:userID", analyticsHandler.GetSkillAnalytics)
		api.GET("/analytics/projects/:projectID", analyticsHandler.GetProjectAnalytics)

		// Badges routes
		api.GET("/badges/:userID", badgeHandler.GetUserBadges)
		// More routes if needed here...

		// Payment routes (ví dụ)
		api.POST("/payments", func(c *gin.Context) {
			// Gọi paymentClient để xử lý thanh toán
		})
	}
}
