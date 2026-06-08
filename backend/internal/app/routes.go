package app

// Nhiệm vụ: Đăng ký tất cả các route API của ứng dụng, gắn handlers với các endpoint.
// Liên kết:
// - Sử dụng internal/handlers để lấy các handler (ProjectHandler, UserHandler, v.v.).
// - Nhận các service từ app.go để truyền vào handler.
// Vai trò trong flow:
// - Xác định cách client tương tác với backend qua các endpoint (GET /projects, POST /tasks, v.v.).
// - Là "bản đồ" của API, giúp developer dễ dàng tra cứu các route.

import (
	"net/http"

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
	// messageService *services.MessageService,
	portfolioService *services.PortfolioService,
	analyticsService *services.AnalyticsService,
	authService *services.AuthService,
	notificationService *services.NotificationService,
	badgeService *services.BadgeService,
	talentPoolService *services.TalentPoolService,
	fileService *services.FileService,
	businessInfoService *services.BusinessInfoService,
	feedbackService *services.FeedbackService,
	gamificationService *services.GamificationService,
	matchingService *services.MatchingService,
	realtimeClient *integrations.RealtimeClient,
	// paymentClient *integrations.PaymentClient,
	chatService *services.ChatService,
	invitationService *services.InvitationService,
) {
	// Khởi tạo integrations ở đây
	// realtimeClient := integrations.NewRealtimeClient()

	// Khởi tạo các handler với service tương ứng
	userHandler := handlers.NewUserHandler(userService, portfolioService)
	authHandler := handlers.NewAuthHandler(authService)
	projectHandler := handlers.NewProjectHandler(projectService)
	applicationHandler := handlers.NewApplicationHandler(applicationService)
	// taskHandler := handlers.NewTaskHandler(taskService)
	reviewHandler := handlers.NewReviewHandler(reviewService)
	// messageHandler := handlers.NewMessageHandler(messageService)
	portfolioHandler := handlers.NewPortfolioHandler(portfolioService)
	analyticsHandler := handlers.NewAnalyticsHandler(analyticsService)
	badgeHandler := handlers.NewBadgeHandler(badgeService)
	talentPoolHandler := handlers.NewTalentPoolHandler(talentPoolService)
	avatarHandler := handlers.NewAvatarHandler(fileService)
	businessInfoHandler := handlers.NewBusinessInfoHandler(businessInfoService)
	websocketNotificationHandler := handlers.NewWebSocketNotificationHandler(realtimeClient, notificationService)
	websocketTaskHanlder := handlers.NewWebSocketTaskHandler(realtimeClient, taskService)
	feedbackHandler := handlers.NewFeedbackHandler(feedbackService)
	matchingHandler := handlers.NewMatchingHandler(matchingService)
	levelHandler := handlers.NewLevelHandler(gamificationService)
	chatHandler := handlers.NewChatHandler(chatService)
	websocketChatHandler := handlers.NewWebSocketChatHandler(chatService, realtimeClient)
	invitationHandler := handlers.NewInvitationHandler(invitationService)

	// 1. CORS FIRST - trước mọi route để preflight OPTIONS luôn có CORS headers
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://skillforge.ikniz.id.vn"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// 2. Global middleware (rate limit, logging)
	r.Use(middleware.RateLimitMiddleware())
	r.Use(middleware.LoggingMiddleware())

	// 3. Public routes (không cần auth)
	r.POST("/auth/register", authHandler.Register)
	r.POST("/auth/login", authHandler.Login)
	r.GET("/ws/task/:projectID/:userID", websocketTaskHanlder.HandleConnection)
	r.GET("/ws/chats/:projectID/:userID", websocketChatHandler.HandleConnection)
	r.GET("/ws/notifi/:userID", websocketNotificationHandler.HandleNotificationConnection)

	// Public static file routes
	r.Static("/storage", "./storage")
	r.Static("/portfolios", "./storage/portfolios")

	// Public avatar routes
	r.GET("/avatars", avatarHandler.ServeAvatar)
	r.GET("/avatars/:id", avatarHandler.ServeAvatarByUserID)

	// Logout - clears auth cookie (no auth needed)
	r.POST("/api/logout", func(c *gin.Context) {
		c.SetCookie("auth_token", "", -1, "/", "", false, true)
		c.JSON(http.StatusOK, gin.H{"message": "logged out"})
	})

	// 4. Auth middleware
	r.Use(middleware.AuthMiddleware())

	// Nhóm route cần auth (dùng middleware nếu cần)
	api := r.Group("/api")

	{
		// User routes
		api.GET("/user", userHandler.GetCurrentUser)
		api.GET("/user/:id", userHandler.GetUserByID)
		api.PUT("/user", userHandler.UpdateCurrentUser)
		api.GET("/users/students", middleware.RoleMiddleware("business"), userHandler.GetAllStudents)
		api.GET("/users/:id/profile", userHandler.GetUserProfile)

		// Project routes
		api.GET("/projects", projectHandler.GetProjects)                       // này là route dành cho marketplace để get tất cả các dự án trên thị trường
		api.GET("/projects/:id", projectHandler.GetProjectMarketplace)         // dùng vào việc xem chi tiết một project nào đó trong marketplace
		api.GET("/projects/students/:id", projectHandler.GetStudentsByProject) // dùng để get các thành viên tham gia dự án hiện tại, id là projectID
		api.GET("/projects/business", middleware.RoleMiddleware("business"), projectHandler.GetProjectByBusiness)     // id của business được truyền qua context
		api.GET("/projects/student", middleware.RoleMiddleware("student"), projectHandler.GetProjectByStudent)       // id của student được truyền qua context
		api.POST("/projects", middleware.RoleMiddleware("business"), projectHandler.CreateProject)
		api.PUT("/projects/:id", middleware.RoleMiddleware("business"), projectHandler.UpdateProject)
		api.DELETE("/projects/:id", middleware.RoleMiddleware("business"), projectHandler.DeleteProject)

		// Cần thêm một nhóm route để get, delete students trong project cụ thể nào đó
		// Route này để get tất cả các students trong một dự án để có thể dẫn tới profile của họ
		// hoặc hiển thị về các thông tin các hoặc delete nếu muốn
		// api.GET("/projects/students/:id") với id là id của project
		// Hàm RemoveStudentFromProject nhận vào projectID và studentID, business thực hiện thao tác thì được truyền qua context
		// đã được truyền qua context nên cần truyền ở đây là bao gồm studentID và projectID
		api.DELETE("/projects/students/:studentID/:projectID", projectHandler.RemoveStudentFromProject)

		// Application routes
		api.POST("/applications", middleware.RoleMiddleware("student"), applicationHandler.ApplyProject)
		api.GET("/applications/:id", applicationHandler.GetApplication)

		// Cái route này dùng để:
		// - Với business: xem những students nào apply vào dự án của mình
		// - Với student: xem mình đã apply vào những project nào cùng với status của các dự án đó
		api.GET("/applications/me", applicationHandler.GetApplicationsByCurrentUser)
		api.GET("/applications/count", applicationHandler.GetApplicationCount)
		// This update for status is used for accept or reject an application, if accept add the student to the project
		api.PUT("/applications/status/:id", middleware.RoleMiddleware("business"), applicationHandler.UpdateApplicationStatus)
		// ID truyền vào ở đây là application id, còn userID thì truyền vào context
		api.DELETE("/applications/:id", middleware.RoleMiddleware("student"), applicationHandler.DeleteApplication)

		// // Task routes
		// api.GET("/tasks/:id", taskHandler.GetTasksByProjectID) // Get task bằng projectID, id ở đây là projectID chứ không phải taskID
		// api.POST("/tasks", taskHandler.CreateTasks)
		// api.PUT("/tasks/:id", taskHandler.UpdateTask)
		// api.PUT("/tasks/:id/assign", taskHandler.AssignTask)
		// api.PUT("/tasks/:id/finish", taskHandler.FinishTask)
		// api.DELETE("/tasks/:id", taskHandler.DeleteTask)

		// Review routes
		api.POST("/reviews", reviewHandler.SubmitReview)
		api.GET("/reviews/:id", reviewHandler.GetReview)

		// Message routes
		// api.POST("/messages", messageHandler.SendMessage)
		// api.GET("/messages/:id", messageHandler.GetMessages)

		// Portfolio routes
		api.GET("/portfolios/:userID", portfolioHandler.GetPortfolio)
		api.POST("/portfolios", portfolioHandler.GeneratePortfolio)

		// TalentPool route
		// id là id của user, còn id của business thì được nhận vào bằng context
		api.GET("/talentpool", middleware.RoleMiddleware("business"), talentPoolHandler.GetTalentPool)
		api.POST("/talentpool/:id", middleware.RoleMiddleware("business"), talentPoolHandler.AddStudentToTalentPool)
		api.DELETE("/talentpool/:id", middleware.RoleMiddleware("business"), talentPoolHandler.RemoveFromTalentPool)
		api.GET("/talentpool/:id/check", middleware.RoleMiddleware("business"), talentPoolHandler.CheckStudentInTalentPool)

		// Analytics routes
		api.GET("/analytics/skills/:userID", analyticsHandler.GetSkillAnalytics)
		api.GET("/analytics/projects/:projectID", analyticsHandler.GetProjectAnalytics)

		// Matching routes
		api.GET("/matches", middleware.RoleMiddleware("student"), matchingHandler.GetTopMatches)             // Get top 10 matches for authenticated user
		api.GET("/matches/:project_id", middleware.RoleMiddleware("student"), matchingHandler.GetMatchScore) // Get specific match score

		// Invitation routes
		api.POST("/invitations", middleware.RoleMiddleware("business"), invitationHandler.CreateInvitation)
		api.GET("/invitations/me", invitationHandler.GetMyInvitations)
		api.PUT("/invitations/:id/respond", middleware.RoleMiddleware("student"), invitationHandler.RespondToInvitation)

		// Badges routes
		api.GET("/badges/:userID", badgeHandler.GetUserBadges)
		// More routes if needed here...

		// Feedback routes
		api.POST("/feedbacks", feedbackHandler.CreateFeedback)
		api.GET("/feedbacks/student/:userID", feedbackHandler.GetStudentFeedbacks)
		api.GET("/feedbacks/business/:userID", feedbackHandler.GetBusinessFeedbacks)

		// === Route phục vụ business info
		api.GET("/business-info", businessInfoHandler.GetBusinessInfo)
		api.PUT("/business-info", middleware.RoleMiddleware("business"), businessInfoHandler.UpdateBusinessInfo)

		//Route cho level
		api.GET("/levels", levelHandler.GetUserLevel)

		// Route cho chat
		// :id là projectID
		api.GET("/chats", chatHandler.GetGroups)
		api.GET("/chats/:id", chatHandler.GetGroupInfo)
		api.POST("/chats/upload", chatHandler.UploadChatFile)

		// Payment routes (ví dụ)
		api.POST("/payments", func(c *gin.Context) {
			// Gọi paymentClient để xử lý thanh toán
		})
	}
}
