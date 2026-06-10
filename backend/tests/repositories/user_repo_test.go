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

func TestUserRepo_CreateUser(t *testing.T) {
	userID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "users") })

	repo := repositories.NewUserRepository(testutil.GetTestDB())
	user := &models.User{
		ID:        userID,
		Name:      "John Doe",
		Email:     "john@example.com",
		Password:  "hashed_password",
		Role:      "student",
		Skills:    []string{"Go", "Python"},
		Title:     "Developer",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.InsertUser(context.Background(), user)
	require.NoError(t, err)

	found, err := repo.FindUserByID(context.Background(), userID)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Equal(t, "John Doe", found.Name)
	assert.Equal(t, "john@example.com", found.Email)
	assert.Equal(t, "student", found.Role)
	assert.Contains(t, found.Skills, "Go")
	assert.Contains(t, found.Skills, "Python")
}

func TestUserRepo_FindUserByID_NotFound(t *testing.T) {
	repo := repositories.NewUserRepository(testutil.GetTestDB())
	user, err := repo.FindUserByID(context.Background(), "nonexistent")
	require.NoError(t, err)
	assert.Nil(t, user)
}

func TestUserRepo_FindByEmail(t *testing.T) {
	userID := utils.GenerateUUID()
	email := "unique@example.com"
	t.Cleanup(func() { testutil.CleanupDB(t, "users") })

	testutil.CreateTestUser(t, userID, "Jane", email, "business", nil)

	want := userID + "." + email
	repo := repositories.NewUserRepository(testutil.GetTestDB())
	found, err := repo.FindUserByEmail(context.Background(), want)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Equal(t, userID, found.ID)
	assert.Equal(t, want, found.Email)
}

func TestUserRepo_FindByEmail_NotFound(t *testing.T) {
	repo := repositories.NewUserRepository(testutil.GetTestDB())
	user, err := repo.FindUserByEmail(context.Background(), "doesnotexist@test.com")
	require.NoError(t, err)
	assert.Nil(t, user)
}

func TestUserRepo_UpdateUser(t *testing.T) {
	userID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "users") })

	testutil.CreateTestUser(t, userID, "Old Name", "old@test.com", "student", nil)

	repo := repositories.NewUserRepository(testutil.GetTestDB())
	user, err := repo.FindUserByID(context.Background(), userID)
	require.NoError(t, err)
	require.NotNil(t, user)
	user.Name = "New Name"
	user.Title = "Senior Developer"
	err = repo.UpdateUser(context.Background(), user)
	require.NoError(t, err)

	found, err := repo.FindUserByID(context.Background(), userID)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Equal(t, "New Name", found.Name)
}

func TestUserRepo_UpdateUserSkills(t *testing.T) {
	userID := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "users") })

	testutil.CreateTestUser(t, userID, "Skill User", "skills@test.com", "student", nil)

	repo := repositories.NewUserRepository(testutil.GetTestDB())
	newSkills := []string{"Go", "Rust", "Kubernetes"}
	err := repo.UpdateUserSkills(context.Background(), userID, newSkills)
	require.NoError(t, err)

	found, err := repo.FindUserByID(context.Background(), userID)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Equal(t, newSkills, found.Skills)
}

func TestUserRepo_FindUsersByIDs(t *testing.T) {
	id1 := utils.GenerateUUID()
	id2 := utils.GenerateUUID()
	t.Cleanup(func() { testutil.CleanupDB(t, "users") })

	testutil.CreateTestUser(t, id1, "User1", "u1@test.com", "student", nil)
	testutil.CreateTestUser(t, id2, "User2", "u2@test.com", "business", nil)

	repo := repositories.NewUserRepository(testutil.GetTestDB())
	users, err := repo.FindUsersByIDs(context.Background(), []string{id1, id2})
	require.NoError(t, err)
	assert.Len(t, users, 2)
}

func TestUserRepo_FindUsersByIDs_Empty(t *testing.T) {
	repo := repositories.NewUserRepository(testutil.GetTestDB())
	users, err := repo.FindUsersByIDs(context.Background(), []string{})
	require.NoError(t, err)
	assert.Empty(t, users)
}
