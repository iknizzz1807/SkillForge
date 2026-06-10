package helpers

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/integrations"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/services"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

var TestDB *mongo.Database
var TestClient *mongo.Client

func init() {
	gin.SetMode(gin.TestMode)
	os.Setenv("JWT_SECRET", "test-secret-key-for-skillforge-testing-1234567890")
	utils.GetJWTSecret()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := os.Getenv("TEST_MONGO_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to test MongoDB: %v", err)
	}
	TestClient = client
	TestDB = client.Database("skillforge_test")
}

func CleanupDB(t *testing.T) {
	t.Helper()
	ctx := context.Background()
	collections, err := TestDB.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		t.Fatalf("Failed to list collections: %v", err)
	}
	for _, c := range collections {
		if err := TestDB.Collection(c).Drop(ctx); err != nil {
			t.Logf("Warning: failed to drop collection %s: %v", c, err)
		}
	}
}

type TestUser struct {
	ID    string
	Email string
	Name  string
	Role  string
	Token string
}

func CreateTestUser(t *testing.T, role string) TestUser {
	t.Helper()
	id := utils.GenerateUUID()
	email := role + "_" + id[:8] + "@test.com"
	name := "Test " + role

	doc := bson.M{
		"_id":        id,
		"email":      email,
		"name":       name,
		"password":   utils.HashPassword("password123"),
		"role":       role,
		"title":      "Test Title",
		"skills":     []string{"Go", "Python"},
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}
	_, err := TestDB.Collection("users").InsertOne(context.Background(), doc)
	require.NoError(t, err)

	token, err := utils.GenerateJWT(id, email, name, role)
	require.NoError(t, err)

	return TestUser{ID: id, Email: email, Name: name, Role: role, Token: token}
}

func CreateTestProject(t *testing.T, businessID, businessName string) string {
	t.Helper()
	now := time.Now()
	id := utils.GenerateUUID()

	doc := bson.M{
		"_id":            id,
		"title":          "Test Project",
		"description":    "Test Description",
		"skills":         []string{"Go", "MongoDB"},
		"start_time":     now,
		"end_time":       now.Add(30 * 24 * time.Hour),
		"max_member":     5,
		"current_member": 0,
		"difficulty":     "intermediate",
		"created_by_id":  businessID,
		"created_by_name": businessName,
		"status":         "open",
		"created_at":     now,
	}
	_, err := TestDB.Collection("projects").InsertOne(context.Background(), doc)
	require.NoError(t, err)
	return id
}

func AddStudentToProject(t *testing.T, projectID, studentID string) {
	t.Helper()
	psID := utils.GenerateUUID()
	_, err := TestDB.Collection("project_student").InsertOne(context.Background(), bson.M{
		"_id":        psID,
		"project_id": projectID,
		"student_id": studentID,
		"created_at": time.Now(),
	})
	require.NoError(t, err)

	TestDB.Collection("projects").UpdateOne(
		context.Background(),
		bson.M{"_id": projectID},
		bson.M{"$inc": bson.M{"current_member": 1}},
	)
}

func GenerateToken(t *testing.T, userID, email, name, role string) string {
	t.Helper()
	token, err := utils.GenerateJWT(userID, email, name, role)
	require.NoError(t, err)
	return token
}

func SetupAuthContext(method, target string, body interface{}) (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	var req *http.Request
	if body != nil {
		jsonBytes, _ := json.Marshal(body)
		req = httptest.NewRequest(method, target, bytes.NewBuffer(jsonBytes))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req = httptest.NewRequest(method, target, nil)
	}
	c.Request = req
	return c, w
}

func SetupMultipartContext(method, target string, fields map[string]string) (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for k, v := range fields {
		writer.WriteField(k, v)
	}
	writer.Close()

	req := httptest.NewRequest(method, target, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	c.Request = req
	return c, w
}

func SetAuthContext(c *gin.Context, user TestUser) {
	c.Set("userID", user.ID)
	c.Set("role", user.Role)
	c.Set("email", user.Email)
	c.Set("name", user.Name)
	c.Request.Header.Set("Authorization", "Bearer "+user.Token)
}

func NewNotificationService() *services.NotificationService {
	realtimeClient := integrations.NewRealtimeClient()
	emailClient := integrations.NewEmailClient("localhost", 587, "test@test.com", "testpass")
	return services.NewNotificationService(realtimeClient, TestDB, emailClient)
}

func NewBadgeService() *services.BadgeService {
	badgeRepo := repositories.NewBadgeRepository(TestDB)
	return services.NewBadgeService(badgeRepo)
}

func NewGamificationService() *services.GamificationService {
	userService := services.NewUserService(TestDB)
	gamificationRepo := repositories.NewGamificationRepository(TestDB)
	return services.NewGamificationService(gamificationRepo, userService)
}
