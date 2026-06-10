package testutil

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iknizzz1807/SkillForge/internal/integrations"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MongoURI      = "mongodb://localhost:27017"
	DatabaseName  = "skillforge_test"
	JWTTestSecret = "test-secret-key-for-skillforge-tests-123456"
)

var (
	testClient *mongo.Client
	testDB      *mongo.Database
	aiServerURL string
	aiOnce      sync.Once
	dbMu        sync.Mutex
)

func init() {
	os.Setenv("JWT_SECRET", JWTTestSecret)
}

func GetAIServerURL() string {
	aiOnce.Do(func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/matching", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"match_score": 85.0}`))
		})
		mux.HandleFunc("/matching2", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"match_score": [85.0, 70.0, 90.0]}`))
		})
		server := &http.Server{Addr: ":0", Handler: mux}
		listener, err := net.Listen("tcp", ":0")
		if err != nil {
			log.Fatalf("Failed to start AI test server: %v", err)
		}
		aiServerURL = "http://" + listener.Addr().String()
		go server.Serve(listener)
	})
	return aiServerURL
}

func GetTestDB() *mongo.Database {
	dbMu.Lock()
	defer dbMu.Unlock()

	if testDB != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		pingErr := testClient.Ping(ctx, nil)
		if pingErr == nil {
			return testDB
		}
		log.Printf("MongoDB ping failed, reconnecting: %v", pingErr)
		testClient.Disconnect(context.Background())
		testClient = nil
		testDB = nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(MongoURI).
		SetMinPoolSize(1).
		SetMaxPoolSize(10).
		SetMaxConnIdleTime(1 * time.Second).
		SetServerSelectionTimeout(2 * time.Second).
		SetConnectTimeout(3 * time.Second).
		SetSocketTimeout(5 * time.Second).
		SetHeartbeatInterval(2 * time.Second)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	testClient = client
	testDB = client.Database(DatabaseName)
	return testDB
}

func CleanCollections(db *mongo.Database, collections ...string) {
	ctx := context.Background()
	for _, col := range collections {
		if err := db.Collection(col).Drop(ctx); err != nil {
			log.Printf("Warning: failed to drop collection %s: %v", col, err)
		}
	}
}

func CleanAll(db *mongo.Database) {
	CleanCollections(db,
		"users", "projects", "applications", "notifications",
		"user_levels", "user_skills", "badges", "user_badges",
		"business_info", "project_student", "invitations",
		"messages", "groups", "tasks", "feedbacks", "portfolios",
	)
}

func GenerateTestToken(userID, role, name string) string {
	token, err := utils.GenerateJWT(userID, userID+"@test.com", name, role)
	if err != nil {
		log.Fatalf("Failed to generate test token: %v", err)
	}
	return token
}

func GenerateExpiredToken(userID, role string) string {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(-24 * time.Hour).Unix(),
		"iat":     time.Now().Add(-48 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(JWTTestSecret))
	return tokenString
}

func NewDummyEmailClient() *integrations.EmailClient {
	return integrations.NewEmailClient("smtp.example.com", 587, "user", "pass")
}
