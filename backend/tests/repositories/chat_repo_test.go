package repositories_test

import (
	"context"
	"testing"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"github.com/iknizzz1807/SkillForge/tests/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChatRepo_InsertMessage(t *testing.T) {
	groupID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "messages", "groups", "users", "projects", "project_student") })

	repo := repositories.NewChatRepository(testutil.GetTestDB())
	message := &models.Message{
		ID:        utils.GenerateUUID(),
		SenderID:  utils.GenerateUUID(),
		GroupID:   groupID,
		Content:   "Hello, world!",
		Type:      "text",
		CreatedAt: time.Now(),
	}

	saved, err := repo.InsertMessage(message)
	require.NoError(t, err)
	require.NotNil(t, saved)
	assert.Equal(t, "Hello, world!", saved.Content)
}

func TestChatRepo_GetGroupMessages(t *testing.T) {
	groupID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "messages", "groups", "users", "projects", "project_student") })

	repo := repositories.NewChatRepository(testutil.GetTestDB())
	for i := 0; i < 3; i++ {
		msg := &models.Message{
			ID:        utils.GenerateUUID(),
			SenderID:  utils.GenerateUUID(),
			GroupID:   groupID,
			Content:   "Message " + string(rune('0'+i)),
			Type:      "text",
			CreatedAt: time.Now(),
		}
		_, err := repo.InsertMessage(msg)
		require.NoError(t, err)
	}

	messages, err := repo.GetGroupMessages(groupID)
	require.NoError(t, err)
	assert.Len(t, messages, 3)
}

func TestChatRepo_GetGroupMessages_Empty(t *testing.T) {
	repo := repositories.NewChatRepository(testutil.GetTestDB())
	messages, err := repo.GetGroupMessages("nonexistent")
	require.NoError(t, err)
	assert.Empty(t, messages)
}

func TestChatRepo_CreateGroup(t *testing.T) {
	projectID := utils.GenerateUUID()
	userID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "groups", "projects", "users", "project_student") })

	testutil.CreateTestUser(t, userID, "Creator", "c@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Chat Project", userID, "Creator", 5, 0)

	_, err := testutil.GetTestDB().Collection("project_student").InsertOne(context.Background(), &models.Project_student{
		ID:         utils.GenerateUUID(),
		Project_id: projectID,
		Student_id: userID,
		CreatedAt:  time.Now(),
	})
	require.NoError(t, err)

	repo := repositories.NewChatRepository(testutil.GetTestDB())
	groups, err := repo.GetGroups(context.Background(), userID)
	require.NoError(t, err)
	require.Len(t, groups, 1)
	assert.Equal(t, projectID, groups[0].ProjectID)
}

func TestChatRepo_GetGroupMembers(t *testing.T) {
	projectID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	studentID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "groups", "messages", "project_student", "projects", "users") })

	testutil.CreateTestUser(t, businessID, "Biz", "biz@test.com", "business", nil)
	testutil.CreateTestUser(t, studentID, "Student", "s@test.com", "student", nil)
	testutil.CreateTestProject(t, projectID, "Team Project", businessID, "Biz", 5, 0)

	_, err := testutil.GetTestDB().Collection("project_student").InsertOne(context.Background(), &models.Project_student{
		ID:         utils.GenerateUUID(),
		Project_id: projectID,
		Student_id: studentID,
		CreatedAt:  time.Now(),
	})
	require.NoError(t, err)

	repo := repositories.NewChatRepository(testutil.GetTestDB())
	members, err := repo.GetGroupMembers(projectID)
	require.NoError(t, err)

	memberIDs := make([]string, 0, len(members))
	for _, m := range members {
		memberIDs = append(memberIDs, m.ID)
	}
	assert.Contains(t, memberIDs, businessID)
	assert.Contains(t, memberIDs, studentID)
}

func TestChatRepo_CheckProjectAccess(t *testing.T) {
	projectID := utils.GenerateUUID()
	studentID := utils.GenerateUUID()
	businessID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "project_student", "projects", "users") })

	testutil.CreateTestUser(t, studentID, "Student", "s@test.com", "student", nil)
	testutil.CreateTestUser(t, businessID, "Biz", "b@test.com", "business", nil)
	testutil.CreateTestProject(t, projectID, "Access Test", businessID, "Biz", 5, 0)

	_, err := testutil.GetTestDB().Collection("project_student").InsertOne(context.Background(), &models.Project_student{
		ID:         utils.GenerateUUID(),
		Project_id: projectID,
		Student_id: studentID,
		CreatedAt:  time.Now(),
	})
	require.NoError(t, err)

	repo := repositories.NewChatRepository(testutil.GetTestDB())

	hasAccess, err := repo.CheckProjectAccess(projectID, studentID)
	require.NoError(t, err)
	assert.True(t, hasAccess)

	hasAccess, err = repo.CheckProjectAccess(projectID, businessID)
	require.NoError(t, err)
	assert.True(t, hasAccess)

	hasAccess, err = repo.CheckProjectAccess(projectID, "outsider")
	require.NoError(t, err)
	assert.False(t, hasAccess)
}
