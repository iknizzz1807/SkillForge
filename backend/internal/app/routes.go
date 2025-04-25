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
	fileService *services.FileService,
	businessInfoService *services.BusinessInfoService,
	feedbackService *services.FeedbackService,
	gamificationService *services.GamificationService,
	matchingService *services.MatchingService,
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
	avatarHandler := handlers.NewAvatarHandler(fileService)
	businessInfoHandler := handlers.NewBusinessInfoHandler(businessInfoService)
	websocketHanlder := handlers.NewWebSocketHandler(realtimeClient, messageService, notificationService,
		taskService,
		projectService)
	feedbackHandler := handlers.NewFeedbackHandler(feedbackService)
	matchingHandler := handlers.NewMatchingHandler(matchingService)

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
		AllowOrigins:     []string{"http://localhost:5173", "http://skillforge.ikniz.id.vn"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.Static("/storage", "./storage")

	// Nhóm route cần auth (dùng middleware nếu cần)
	api := r.Group("/api")

	{
		// User routes
		api.GET("/user", userHandler.GetCurrentUser)
		api.PUT("/user", userHandler.UpdateCurrentUser)

		// Project routes
		api.GET("/projects", projectHandler.GetProjects)                       // này là route dành cho marketplace để get tất cả các dự án trên thị trường
		api.GET("/projects/:id", projectHandler.GetProjectMarketplace)         // dùng vào việc xem chi tiết một project nào đó trong marketplace
		api.GET("/projects/students/:id", projectHandler.GetStudentsByProject) // dùng để get các thành viên tham gia dự án hiện tại, id là projectID
		api.GET("/projects/business", projectHandler.GetProjectByBusiness)     // id của business được truyền qua context
		api.GET("/projects/student", projectHandler.GetProjectByStudent)       // id của student được truyền qua context
		api.POST("/projects", projectHandler.CreateProject)
		api.PUT("/projects/:id", projectHandler.UpdateProject)
		api.DELETE("/projects/:id", projectHandler.DeleteProject)

		// Cần thêm một nhóm route để get, delete students trong project cụ thể nào đó
		// Route này để get tất cả các students trong một dự án để có thể dẫn tới profile của họ
		// hoặc hiển thị về các thông tin các hoặc delete nếu muốn
		// api.GET("/projects/students/:id") với id là id của project
		// Hàm RemoveStudentFromProject nhận vào projectID và studentID, business thực hiện thao tác thì được truyền qua context
		// đã được truyền qua context nên cần truyền ở đây là bao gồm studentID và projectID
		api.DELETE("/projects/students/:studentID/:projectID", projectHandler.RemoveStudentFromProject)

		// Application routes
		api.POST("/applications", applicationHandler.ApplyProject)
		api.GET("/applications/:id", applicationHandler.GetApplication)

		// Cái route này dùng để:
		// - Với business: xem những students nào apply vào dự án của mình
		// - Với student: xem mình đã apply vào những project nào cùng với status của các dự án đó
		api.GET("/applications/me", applicationHandler.GetApplicationsByCurrentUser)
		// This update for status is used for accept or reject an application, if accept add the student to the project
		api.PUT("/applications/status/:id", applicationHandler.UpdateApplicationStatus)
		// ID truyền vào ở đây là application id, còn userID thì truyền vào context
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

		// Matching routes
		api.GET("/matches", matchingHandler.GetTopMatches)             // Get top 10 matches for authenticated user
		api.GET("/matches/:project_id", matchingHandler.GetMatchScore) // Get specific match score

		// Badges routes
		api.GET("/badges/:userID", badgeHandler.GetUserBadges)
		// More routes if needed here...

		// Feedback routes
		api.POST("/feedbacks", feedbackHandler.CreateFeedback)
		api.GET("/feedbacks/student/:userID", feedbackHandler.GetStudentFeedbacks)
		api.GET("/feedbacks/business/:userID", feedbackHandler.GetBusinessFeedbacks)

		// === Route phục vụ file avatar ===
		r.GET("/avatars", avatarHandler.ServeAvatar)

		// === Route phục vụ business info
		api.GET("/business-info", businessInfoHandler.GetBusinessInfo)
		api.PUT("/business-info", businessInfoHandler.UpdateBusinessInfo)

		// Payment routes (ví dụ)
		api.POST("/payments", func(c *gin.Context) {
			// Gọi paymentClient để xử lý thanh toán
		})
	}
}
