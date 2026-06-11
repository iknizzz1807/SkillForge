package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/config"
	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const demoPassword = "Demo@123456"

type demoUser struct {
	ID      string
	Email   string
	Name    string
	Role    string
	Title   string
	Website string
	Skills  []string
}

type demoProject struct {
	ID          string
	Title       string
	Description string
	Skills      []string
	Business    demoUser
	Status      string
	Difficulty  string
	MaxMember   int
	Members     []demoUser
}

func main() {
	reset := flag.Bool("reset", false, "drop demo collections before seeding")
	dbName := flag.String("db", "skillforge", "MongoDB database name")
	flag.Parse()

	cfg := config.Load()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatalf("connect MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("ping MongoDB: %v", err)
	}

	db := client.Database(*dbName)
	if *reset {
		resetCollections(ctx, db)
	}

	seed := buildSeed()
	now := time.Now()

	users := seedUsers(ctx, db, seed.users, now)
	seedBusinessInfo(ctx, db, seed.businesses, now)
	projects := seedProjects(ctx, db, seed.projects, now)
	seedMemberships(ctx, db, seed.projects, now)
	seedApplications(ctx, db, seed.projects, seed.students, now)
	seedInvitations(ctx, db, seed.projects, seed.students, now)
	seedTasks(ctx, db, seed.projects, now)
	seedChat(ctx, db, seed.projects, now)
	badges := seedBadges(ctx, db, now)
	seedGamification(ctx, db, seed.students, badges, now)
	seedTalentPool(ctx, db, seed.businesses, seed.students, now)
	seedFeedback(ctx, db, seed.projects, now)
	seedNotifications(ctx, db, seed.users, now)
	seedPortfolios(ctx, db, seed.projects, seed.students, now)

	fmt.Println("\nSkillForge demo seed complete.")
	fmt.Printf("Database: %s\n", *dbName)
	fmt.Printf("Users: %d, Projects: %d\n", len(users), len(projects))
	fmt.Println("\nDemo accounts, password for all:", demoPassword)
	for _, u := range seed.users {
		fmt.Printf("- %-8s %-28s %s\n", u.Role, u.Email, u.Name)
	}
	fmt.Println("\nSuggested demo path:")
	fmt.Println("1. Login business@skillforge.demo and show dashboard/project/applications/talent search")
	fmt.Println("2. Login an@skillforge.demo and show marketplace/matching/profile/portfolio")
	fmt.Println("3. Open chat + project workspace for 'AI Resume Screening Platform'")
}

type seedData struct {
	users      []demoUser
	businesses []demoUser
	students   []demoUser
	projects   []demoProject
}

func buildSeed() seedData {
	businesses := []demoUser{
		{ID: "demo-biz-aurora", Email: "business@skillforge.demo", Name: "Aurora Labs", Role: "business", Title: "AI Product Studio", Website: "https://auroralabs.demo"},
		{ID: "demo-biz-nova", Email: "nova@skillforge.demo", Name: "Nova Commerce", Role: "business", Title: "E-commerce Growth Team", Website: "https://novacommerce.demo"},
	}
	students := []demoUser{
		{ID: "demo-stu-an", Email: "an@skillforge.demo", Name: "Nguyen Minh An", Role: "student", Title: "Full Stack AI Engineer", Skills: []string{"React", "Svelte", "Go", "Python", "FastAPI", "MongoDB", "Docker", "AI"}},
		{ID: "demo-stu-linh", Email: "linh@skillforge.demo", Name: "Tran Gia Linh", Role: "student", Title: "Frontend Developer", Skills: []string{"React", "Svelte", "TypeScript", "Tailwind", "UI/UX", "Testing"}},
		{ID: "demo-stu-khoa", Email: "khoa@skillforge.demo", Name: "Pham Dang Khoa", Role: "student", Title: "Backend Developer", Skills: []string{"Go", "Node.js", "MongoDB", "PostgreSQL", "Docker", "WebSocket"}},
		{ID: "demo-stu-mai", Email: "mai@skillforge.demo", Name: "Le Hoang Mai", Role: "student", Title: "Data Scientist", Skills: []string{"Python", "Machine Learning", "NLP", "Pandas", "FastAPI", "MLOps"}},
		{ID: "demo-stu-quang", Email: "quang@skillforge.demo", Name: "Do Quang Huy", Role: "student", Title: "DevOps Engineer", Skills: []string{"Docker", "Kubernetes", "CI/CD", "Linux", "Monitoring", "AWS"}},
	}
	users := append([]demoUser{}, businesses...)
	users = append(users, students...)
	projects := []demoProject{
		{
			ID: "demo-project-ai-resume", Business: businesses[0], Status: "active", Difficulty: "expert", MaxMember: 5,
			Title:       "AI Resume Screening Platform",
			Description: "Build an AI-assisted resume screening workflow for recruiters: upload CVs, extract skills, rank candidates, and explain matching decisions. Investor demo shows AI + marketplace + task collaboration in one flow.",
			Skills:      []string{"React", "Go", "Python", "FastAPI", "MongoDB", "AI", "Docker"},
			Members:     []demoUser{students[0], students[1], students[2]},
		},
		{
			ID: "demo-project-realtime-crm", Business: businesses[0], Status: "open", Difficulty: "intermediate", MaxMember: 4,
			Title:       "Realtime Customer Success CRM",
			Description: "Create a CRM workspace with real-time notifications, account health scores, and a collaborative task board for customer success teams.",
			Skills:      []string{"Svelte", "Go", "WebSocket", "MongoDB", "Analytics"},
			Members:     []demoUser{students[2]},
		},
		{
			ID: "demo-project-commerce-insights", Business: businesses[1], Status: "open", Difficulty: "beginer", MaxMember: 6,
			Title:       "E-commerce Insights Dashboard",
			Description: "Analyze orders, products, and conversion funnels. Build clean dashboards and AI-generated recommendations for store operators.",
			Skills:      []string{"React", "TypeScript", "Python", "Data Visualization", "UI/UX"},
			Members:     []demoUser{students[3]},
		},
		{
			ID: "demo-project-devops-launchpad", Business: businesses[1], Status: "close", Difficulty: "expert", MaxMember: 3,
			Title:       "Cloud DevOps Launchpad",
			Description: "Design a production-grade deployment pipeline with container builds, health checks, monitoring, and rollback playbooks.",
			Skills:      []string{"Docker", "Kubernetes", "CI/CD", "Linux", "Monitoring"},
			Members:     []demoUser{students[4], students[0]},
		},
	}
	return seedData{users: users, businesses: businesses, students: students, projects: projects}
}

func resetCollections(ctx context.Context, db *mongo.Database) {
	collections := []string{
		"users", "business_info", "projects", "project_student", "applications", "invitations",
		"messages", "groups", "tasks", "activities", "feedbacks", "notifications", "talent_pool",
		"badges", "user_badges", "user_levels", "user_skills", "portfolios",
	}
	for _, name := range collections {
		if err := db.Collection(name).Drop(ctx); err != nil {
			log.Printf("warning: drop %s: %v", name, err)
		}
	}
}

func seedUsers(ctx context.Context, db *mongo.Database, users []demoUser, now time.Time) []models.User {
	password := utils.HashPassword(demoPassword)
	docs := make([]interface{}, 0, len(users))
	result := make([]models.User, 0, len(users))
	for _, u := range users {
		model := models.User{ID: u.ID, Email: u.Email, Name: u.Name, Password: password, Role: u.Role, Skills: u.Skills, Title: u.Title, Website: u.Website, CreatedAt: now.Add(-14 * 24 * time.Hour), UpdatedAt: now.Add(-2 * time.Hour)}
		result = append(result, model)
		docs = append(docs, model)
	}
	upsertMany(ctx, db.Collection("users"), docs)
	return result
}

func seedBusinessInfo(ctx context.Context, db *mongo.Database, businesses []demoUser, now time.Time) {
	docs := []interface{}{
		models.BusinessInfo{UserID: businesses[0].ID, CompanyType: "AI SaaS / HR Tech", Founded: "2022", CompanySize: "25-50", Website: businesses[0].Website, AboutUs: "Aurora Labs builds applied AI products for hiring, operations, and knowledge work. We use SkillForge to discover high-signal student builders through real project outcomes.", CreatedAt: now.Add(-12 * 24 * time.Hour), UpdatedAt: now},
		models.BusinessInfo{UserID: businesses[1].ID, CompanyType: "E-commerce Technology", Founded: "2020", CompanySize: "50-100", Website: businesses[1].Website, AboutUs: "Nova Commerce helps online merchants turn data into growth decisions. Our teams collaborate with emerging technical talent on analytics and infrastructure projects.", CreatedAt: now.Add(-12 * 24 * time.Hour), UpdatedAt: now},
	}
	upsertMany(ctx, db.Collection("business_info"), docs)
}

func seedProjects(ctx context.Context, db *mongo.Database, projects []demoProject, now time.Time) []models.Project {
	docs := make([]interface{}, 0, len(projects))
	result := make([]models.Project, 0, len(projects))
	for i, p := range projects {
		model := models.Project{ID: p.ID, Title: p.Title, Description: p.Description, Skills: p.Skills, StartTime: now.Add(time.Duration(i-1) * 7 * 24 * time.Hour), EndTime: now.Add(time.Duration(i+5) * 7 * 24 * time.Hour), MaxMember: p.MaxMember, CurrentMember: len(p.Members), Difficulty: p.Difficulty, CreatedByID: p.Business.ID, CreatedByName: p.Business.Name, Status: p.Status, CreatedAt: now.Add(time.Duration(-10+i) * 24 * time.Hour)}
		result = append(result, model)
		docs = append(docs, model)
	}
	upsertMany(ctx, db.Collection("projects"), docs)
	return result
}

func seedMemberships(ctx context.Context, db *mongo.Database, projects []demoProject, now time.Time) {
	docs := []interface{}{}
	for _, p := range projects {
		for _, m := range p.Members {
			docs = append(docs, models.Project_student{ID: fmt.Sprintf("demo-member-%s-%s", p.ID, m.ID), Project_id: p.ID, Student_id: m.ID, CreatedAt: now.Add(-5 * 24 * time.Hour)})
		}
	}
	upsertMany(ctx, db.Collection("project_student"), docs)
}

func seedApplications(ctx context.Context, db *mongo.Database, projects []demoProject, students []demoUser, now time.Time) {
	docs := []interface{}{
		models.Application{ID: "demo-app-linh-ai", UserID: students[1].ID, ProjectID: projects[0].ID, ProjectName: projects[0].Title, UserName: students[1].Name, Motivation: "I want to refine frontend workflows for explainable AI and recruiter productivity.", DetailedProposal: "I will implement candidate ranking views, bias-check UI states, and reusable review components with Svelte and TypeScript.", Status: "approved", CreatedAt: now.Add(-7 * 24 * time.Hour), UpdatedAt: now.Add(-6 * 24 * time.Hour)},
		models.Application{ID: "demo-app-mai-crm", UserID: students[3].ID, ProjectID: projects[1].ID, ProjectName: projects[1].Title, UserName: students[3].Name, Motivation: "I want to apply ML scoring and analytics to customer success workflows.", DetailedProposal: "I can build account-health scoring prototypes and create explainable analytics cards for the dashboard.", Status: "pending", CreatedAt: now.Add(-36 * time.Hour)},
		models.Application{ID: "demo-app-quang-crm", UserID: students[4].ID, ProjectID: projects[1].ID, ProjectName: projects[1].Title, UserName: students[4].Name, Motivation: "I can help ship the realtime CRM reliably with Docker and observability.", DetailedProposal: "I will containerize services, add health checks, and create local demo scripts for the engineering team.", Status: "pending", CreatedAt: now.Add(-20 * time.Hour)},
		models.Application{ID: "demo-app-an-commerce", UserID: students[0].ID, ProjectID: projects[2].ID, ProjectName: projects[2].Title, UserName: students[0].Name, Motivation: "I want to connect AI recommendations with polished product UX.", DetailedProposal: "I will integrate backend insight APIs and build dashboard interactions for store operators.", Status: "rejected", CreatedAt: now.Add(-4 * 24 * time.Hour), UpdatedAt: now.Add(-3 * 24 * time.Hour)},
	}
	upsertMany(ctx, db.Collection("applications"), docs)
}

func seedInvitations(ctx context.Context, db *mongo.Database, projects []demoProject, students []demoUser, now time.Time) {
	docs := []interface{}{
		models.Invitation{ID: "demo-invite-an-crm", ProjectID: projects[1].ID, StudentID: students[0].ID, BusinessID: projects[1].Business.ID, Status: "pending", CreatedAt: now.Add(-8 * time.Hour)},
		models.Invitation{ID: "demo-invite-quang-devops", ProjectID: projects[3].ID, StudentID: students[4].ID, BusinessID: projects[3].Business.ID, Status: "accepted", CreatedAt: now.Add(-8 * 24 * time.Hour)},
		models.Invitation{ID: "demo-invite-linh-commerce", ProjectID: projects[2].ID, StudentID: students[1].ID, BusinessID: projects[2].Business.ID, Status: "pending", CreatedAt: now.Add(-10 * time.Hour)},
	}
	upsertMany(ctx, db.Collection("invitations"), docs)
}

func seedTasks(ctx context.Context, db *mongo.Database, projects []demoProject, now time.Time) {
	docs := []interface{}{
		models.Task{ID: "demo-task-ai-1", ProjectID: projects[0].ID, Title: "Design candidate ranking dashboard", Description: "Create recruiter-facing dashboard showing ranked candidates, matching skills, and explanation snippets.", Status: "done", Note: "Ready for demo", AssignedTo: "demo-stu-linh", FinishedBy: "demo-stu-linh", CreatedAt: now.Add(-6 * 24 * time.Hour), UpdatedAt: now.Add(-2 * 24 * time.Hour)},
		models.Task{ID: "demo-task-ai-2", ProjectID: projects[0].ID, Title: "Implement matching API integration", Description: "Connect Go backend to AI matching service with timeout and error handling.", Status: "done", Note: "Verified with sample resumes", AssignedTo: "demo-stu-khoa", FinishedBy: "demo-stu-khoa", CreatedAt: now.Add(-6 * 24 * time.Hour), UpdatedAt: now.Add(-24 * time.Hour)},
		models.Task{ID: "demo-task-ai-3", ProjectID: projects[0].ID, Title: "Build resume upload parser", Description: "Prototype file upload and extracted-skill preview flow.", Status: "in_progress", Note: "Parser edge cases pending", AssignedTo: "demo-stu-an", CreatedAt: now.Add(-4 * 24 * time.Hour), UpdatedAt: now.Add(-3 * time.Hour)},
		models.Task{ID: "demo-task-ai-4", ProjectID: projects[0].ID, Title: "Prepare investor demo dataset", Description: "Create sample candidates, roles, and evaluation screenshots.", Status: "review", Note: "Needs final copy polish", AssignedTo: "demo-biz-aurora", CreatedAt: now.Add(-2 * 24 * time.Hour), UpdatedAt: now.Add(-1 * time.Hour)},
		models.Task{ID: "demo-task-ai-5", ProjectID: projects[0].ID, Title: "Write integration test suite", Description: "Cover resume parsing, matching API, and candidate ranking endpoints.", Status: "todo", AssignedTo: "demo-stu-an", CreatedAt: now.Add(-6 * time.Hour)},
		models.Task{ID: "demo-task-crm-1", ProjectID: projects[1].ID, Title: "WebSocket notification spike", Description: "Validate realtime update architecture for customer health changes.", Status: "todo", AssignedTo: "demo-stu-khoa", CreatedAt: now.Add(-24 * time.Hour)},
		models.Task{ID: "demo-task-commerce-1", ProjectID: projects[2].ID, Title: "Define dashboard metric hierarchy", Description: "Prioritize GMV, conversion rate, repeat purchase, and product margin cards.", Status: "in_progress", AssignedTo: "demo-stu-mai", CreatedAt: now.Add(-48 * time.Hour)},
	}
	upsertMany(ctx, db.Collection("tasks"), docs)

	activities := []interface{}{
		activity("demo-act-1", "create", "demo-biz-aurora", projects[0].ID, "Design candidate ranking dashboard", "", "", now.Add(-6*24*time.Hour)),
		activity("demo-act-2", "Move", "demo-stu-linh", projects[0].ID, "Design candidate ranking dashboard", "review", "done", now.Add(-2*24*time.Hour)),
		activity("demo-act-3", "ChangeAssignee", "demo-biz-aurora", projects[0].ID, "Build resume upload parser", "", "Nguyen Minh An", now.Add(-3*time.Hour)),
		activity("demo-act-4", "Move", "demo-stu-an", projects[0].ID, "Build resume upload parser", "todo", "in_progress", now.Add(-90*time.Minute)),
	}
	upsertMany(ctx, db.Collection("activities"), activities)
}

func activity(id, typ, doneBy, projectID, title, from, to string, t time.Time) models.Activity {
	return models.Activity{ID: id, Type: typ, DoneBy: doneBy, ProjectID: projectID, Title: title, From: from, To: to, CreatedAt: t}
}

func seedChat(ctx context.Context, db *mongo.Database, projects []demoProject, now time.Time) {
	groups := []interface{}{}
	messages := []interface{}{}
	for _, p := range projects {
		groups = append(groups, models.Group{ID: "demo-group-" + p.ID, ProjectID: p.ID, Title: p.Title, CreatedAt: now.Add(-7 * 24 * time.Hour)})
	}
	messages = append(messages,
		models.Message{ID: "demo-msg-1", SenderID: "demo-biz-aurora", GroupID: projects[0].ID, Content: "Welcome team. The investor demo goal is simple: show how SkillForge turns student skills into verified project outcomes.", Type: "text", CreatedAt: now.Add(-28 * time.Hour)},
		models.Message{ID: "demo-msg-2", SenderID: "demo-stu-an", GroupID: projects[0].ID, Content: "I will connect the upload flow with the matching endpoint and keep the UX smooth for recruiters.", Type: "text", CreatedAt: now.Add(-27 * time.Hour)},
		models.Message{ID: "demo-msg-3", SenderID: "demo-stu-linh", GroupID: projects[0].ID, Content: "Dashboard cards are ready. I added explanation chips for skills, experience, and project fit.", Type: "text", CreatedAt: now.Add(-22 * time.Hour)},
		models.Message{ID: "demo-msg-4", SenderID: "demo-stu-khoa", GroupID: projects[0].ID, Content: "Backend timeout/error handling is done. AI service failures now degrade cleanly.", Type: "text", CreatedAt: now.Add(-18 * time.Hour)},
		models.Message{ID: "demo-msg-5", SenderID: "demo-biz-aurora", GroupID: projects[1].ID, Content: "For CRM, we need a convincing realtime notification moment in the demo.", Type: "text", CreatedAt: now.Add(-6 * time.Hour)},
	)
	upsertMany(ctx, db.Collection("groups"), groups)
	upsertMany(ctx, db.Collection("messages"), messages)
}

func seedBadges(ctx context.Context, db *mongo.Database, now time.Time) []models.Badge {
	badges := []models.Badge{
		{ID: "demo-badge-ai-builder", Code: "AI_BUILDER", Name: "AI Builder", Description: "Completed meaningful work on an AI-assisted product project.", Type: "skill", Prerequisites: map[string]interface{}{"skill": "AI", "projects": 1}, CreatedAt: now.Add(-30 * 24 * time.Hour)},
		{ID: "demo-badge-realtime", Code: "REALTIME_COLLAB", Name: "Realtime Collaborator", Description: "Contributed to chat, WebSocket, or realtime task workflows.", Type: "collaboration", Prerequisites: map[string]interface{}{"messages": 5}, CreatedAt: now.Add(-30 * 24 * time.Hour)},
		{ID: "demo-badge-finisher", Code: "PROJECT_FINISHER", Name: "Project Finisher", Description: "Finished key tasks and helped ship a project milestone.", Type: "completion", Prerequisites: map[string]interface{}{"completed_tasks": 2}, CreatedAt: now.Add(-30 * 24 * time.Hour)},
	}
	docs := make([]interface{}, len(badges))
	for i := range badges {
		docs[i] = badges[i]
	}
	upsertMany(ctx, db.Collection("badges"), docs)
	return badges
}

func seedGamification(ctx context.Context, db *mongo.Database, students []demoUser, badges []models.Badge, now time.Time) {
	levels := []interface{}{}
	skills := []interface{}{}
	userBadges := []interface{}{}
	for i, s := range students {
		levels = append(levels, models.UserLevel{ID: "demo-level-" + s.ID, UserID: s.ID, Level: 2 + i%4, XPNeed: 1000 + i*250, XPCurrent: 420 + i*180, CreatedAt: now.Add(-10 * 24 * time.Hour), UpdatedAt: now})
		for j, skill := range s.Skills {
			if j >= 5 {
				break
			}
			skills = append(skills, models.UserSkill{ID: fmt.Sprintf("demo-skill-%s-%d", s.ID, j), UserID: s.ID, Skill: skill, Level: 1 + (i+j)%3, PointNeeded: 5, PointCurrent: 2 + (i+j)%4, CreatedAt: now.Add(-9 * 24 * time.Hour), UpdatedAt: now})
		}
		badge := badges[i%len(badges)]
		userBadges = append(userBadges, models.UserBadge{ID: "demo-user-badge-" + s.ID, UserID: s.ID, BadgeID: badge.ID, Badge: &badge, AwardedAt: now.Add(time.Duration(-i-1) * 24 * time.Hour), RelatedEntityID: "demo-project-ai-resume"})
	}
	upsertMany(ctx, db.Collection("user_levels"), levels)
	upsertMany(ctx, db.Collection("user_skills"), skills)
	upsertMany(ctx, db.Collection("user_badges"), userBadges)
}

func seedTalentPool(ctx context.Context, db *mongo.Database, businesses []demoUser, students []demoUser, now time.Time) {
	docs := []interface{}{}
	for i, s := range students {
		business := businesses[i%len(businesses)]
		docs = append(docs, models.TalentPool{ID: fmt.Sprintf("demo-talent-%s-%s", business.ID, s.ID), BusinessID: business.ID, StudentID: s.ID, StudentName: s.Name, Skills: s.Skills, Notes: "Strong demo candidate: clear portfolio, active collaboration, and relevant skill overlap.", CreatedAt: now.Add(time.Duration(-i) * 24 * time.Hour), UpdatedAt: now})
	}
	upsertMany(ctx, db.Collection("talent_pool"), docs)
}

func seedFeedback(ctx context.Context, db *mongo.Database, projects []demoProject, now time.Time) {
	docs := []interface{}{
		models.Feedback{ID: "demo-feedback-1", ProjectID: projects[0].ID, FromID: "demo-biz-aurora", ToID: "demo-stu-an", Type: "student", Rating: 5, Content: "An connected AI, backend, and frontend quickly. Excellent ownership for investor-facing work.", CreatedAt: now.Add(-12 * time.Hour)},
		models.Feedback{ID: "demo-feedback-2", ProjectID: projects[0].ID, FromID: "demo-stu-linh", ToID: "demo-biz-aurora", Type: "business", Rating: 5, Content: "Aurora Labs gave clear requirements, fast feedback, and realistic product constraints.", CreatedAt: now.Add(-11 * time.Hour)},
		models.Feedback{ID: "demo-feedback-3", ProjectID: projects[3].ID, FromID: "demo-biz-nova", ToID: "demo-stu-quang", Type: "student", Rating: 4, Content: "Quang improved deployment reliability and documented the release flow clearly.", CreatedAt: now.Add(-2 * 24 * time.Hour)},
	}
	upsertMany(ctx, db.Collection("feedbacks"), docs)
}

func seedNotifications(ctx context.Context, db *mongo.Database, users []demoUser, now time.Time) {
	docs := []interface{}{}
	for i, u := range users {
		docs = append(docs, models.Notification{ID: "demo-notification-" + u.ID, ToUserID: u.ID, Type: "demo", Content: demoNotificationContent(u.Role), CreatedAt: now.Add(time.Duration(-i) * time.Hour)})
	}
	upsertMany(ctx, db.Collection("notifications"), docs)
}

func demoNotificationContent(role string) string {
	if role == "business" {
		return "You have new student applications and recommended talent to review."
	}
	return "A business invited you to join a project matching your skills."
}

func seedPortfolios(ctx context.Context, db *mongo.Database, projects []demoProject, students []demoUser, now time.Time) {
	portfolios := []interface{}{}
	for _, s := range students {
		projectIDs := []string{}
		for _, p := range projects {
			for _, m := range p.Members {
				if m.ID == s.ID {
					projectIDs = append(projectIDs, p.ID)
				}
			}
		}
		portfolios = append(portfolios, models.Portfolio{ID: "demo-portfolio-" + s.ID, UserID: s.ID, Projects: projectIDs, Skills: s.Skills, CreatedAt: now.Add(-7 * 24 * time.Hour), UpdatedAt: now})
		writePortfolioHTML(s, projectIDs)
	}
	upsertMany(ctx, db.Collection("portfolios"), portfolios)
}

func writePortfolioHTML(student demoUser, projectIDs []string) {
	if err := os.MkdirAll("./storage/portfolios", 0o755); err != nil {
		log.Printf("warning: create portfolio dir: %v", err)
		return
	}
	path := filepath.Join("./storage/portfolios", student.ID+".html")
	body := fmt.Sprintf(`<!doctype html><html><head><meta charset="utf-8"><title>%s - SkillForge Portfolio</title><style>body{font-family:Inter,Arial,sans-serif;max-width:880px;margin:48px auto;color:#111827} .badge{display:inline-block;background:#ede9fe;color:#5b21b6;padding:6px 10px;border-radius:999px;margin:4px}</style></head><body><h1>%s</h1><h2>%s</h2><p>Verified SkillForge demo portfolio with %d project(s), badges, feedback and skill evidence.</p><h3>Skills</h3><p>%s</p><h3>Highlight</h3><p>Contributed to real business projects through SkillForge collaboration, task tracking, and feedback loops.</p></body></html>`, student.Name, student.Name, student.Title, len(projectIDs), skillBadges(student.Skills))
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		log.Printf("warning: write portfolio %s: %v", path, err)
	}
}

func skillBadges(skills []string) string {
	out := ""
	for _, s := range skills {
		out += `<span class="badge">` + s + `</span>`
	}
	return out
}

func upsertMany(ctx context.Context, col *mongo.Collection, docs []interface{}) {
	for _, doc := range docs {
		id, err := docID(doc)
		if err != nil {
			log.Fatalf("seed %s: %v", col.Name(), err)
		}
		_, err = col.ReplaceOne(ctx, bson.M{"_id": id}, doc, options.Replace().SetUpsert(true))
		if err != nil {
			log.Fatalf("seed %s/%s: %v", col.Name(), id, err)
		}
	}
}

func docID(doc interface{}) (string, error) {
	switch v := doc.(type) {
	case models.User:
		return v.ID, nil
	case models.BusinessInfo:
		return v.UserID, nil
	case models.Project:
		return v.ID, nil
	case models.Project_student:
		return v.ID, nil
	case models.Application:
		return v.ID, nil
	case models.Invitation:
		return v.ID, nil
	case models.Group:
		return v.ID, nil
	case models.Message:
		return v.ID, nil
	case models.Task:
		return v.ID, nil
	case models.Activity:
		return v.ID, nil
	case models.Badge:
		return v.ID, nil
	case models.UserBadge:
		return v.ID, nil
	case models.UserLevel:
		return v.ID, nil
	case models.UserSkill:
		return v.ID, nil
	case models.TalentPool:
		return v.ID, nil
	case models.Feedback:
		return v.ID, nil
	case models.Notification:
		return v.ID, nil
	case models.Portfolio:
		return v.ID, nil
	default:
		return "", fmt.Errorf("unsupported doc type %T", doc)
	}
}
