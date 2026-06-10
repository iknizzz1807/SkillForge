package testutil

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/integrations"
	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func withRetry(t *testing.T, fn func() error) {
	t.Helper()
	var err error
	for i := 0; i < 3; i++ {
		if i > 0 {
			time.Sleep(500 * time.Millisecond)
		}
		err = fn()
		if err == nil {
			return
		}
		if !isConnError(err) {
			t.Fatalf("failed: %v", err)
		}
	}
	t.Fatalf("failed after retries: %v", err)
}

func isConnError(err error) bool {
	msg := err.Error()
	return strings.Contains(msg, "connection(") &&
		(strings.Contains(msg, "socket") || strings.Contains(msg, "closed"))
}

func NewTestRealtimeClient() *integrations.RealtimeClient {
	return integrations.NewRealtimeClient()
}

func NewTestEmailClient() *integrations.EmailClient {
	return integrations.NewEmailClient("localhost", 1025, "", "")
}

func NewTestAIClient() *integrations.AIClient {
	return integrations.NewAIClient("http://localhost:8000")
}

func GetDB() *mongo.Database {
	return GetTestDB()
}

func CleanupDB(t *testing.T, collections ...string) {
	for _, col := range collections {
		if err := GetTestDB().Collection(col).Drop(context.Background()); err != nil {
			t.Logf("Warning: failed to drop collection %s: %v", col, err)
		}
	}
}

func CreateTestUser(t *testing.T, id, name, email, role string, skills []string) {
	user := &models.User{
		ID:        id,
		Name:      name,
		Email:     id + "." + email,
		Role:      role,
		Skills:    skills,
		Title:     "Test " + role,
		CreatedAt: time.Now(),
	}
	withRetry(t, func() error {
		_, err := GetTestDB().Collection("users").InsertOne(context.Background(), user)
		return err
	})
}

func CreateTestProject(t *testing.T, id, title, bizID, bizName string, maxMember, currentMember int) {
	project := &models.Project{
		ID:            id,
		Title:         title,
		Description:   "Test description",
		Skills:        []string{"Go"},
		MaxMember:     maxMember,
		CurrentMember: currentMember,
		Difficulty:    "beginner",
		CreatedByID:   bizID,
		CreatedByName: bizName,
		Status:        "open",
		CreatedAt:     time.Now(),
	}
	withRetry(t, func() error {
		_, err := GetTestDB().Collection("projects").InsertOne(context.Background(), project)
		return err
	})
}
