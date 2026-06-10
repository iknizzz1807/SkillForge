package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/iknizzz1807/SkillForge/internal/handlers"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/services"
	"github.com/iknizzz1807/SkillForge/tests/helpers"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func newAvatarHandler() *handlers.AvatarHandler {
	userRepo := repositories.NewUserRepository(helpers.TestDB)
	fileService := services.NewFileService(userRepo)
	return handlers.NewAvatarHandler(fileService)
}

func TestUploadAvatar(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	student := helpers.CreateTestUser(t, "student")
	handler := newAvatarHandler()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("avatar", "test.png")
	// Minimal valid PNG bytes with correct magic signature
	pngBytes := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	part.Write(pngBytes)
	writer.Close()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest("POST", "/api/avatar/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	c.Request = req
	helpers.SetAuthContext(c, student)

	handler.UploadAvatar(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Contains(t, resp["message"], "uploaded")
}

func TestUploadAvatar_NoFile(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	student := helpers.CreateTestUser(t, "student")
	handler := newAvatarHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/avatar/upload", nil)
	helpers.SetAuthContext(c, student)
	c.Request.Header.Set("Content-Type", "multipart/form-data")

	handler.UploadAvatar(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUploadAvatar_Unauthorized(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	handler := newAvatarHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/avatar/upload", nil)

	handler.UploadAvatar(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestUploadAvatar_TooLarge(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	student := helpers.CreateTestUser(t, "student")
	handler := newAvatarHandler()

	c, w := helpers.SetupAuthContext("POST", "/api/avatar/upload", nil)
	helpers.SetAuthContext(c, student)

	handler.UploadAvatar(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestServeAvatar_NotFound(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	handler := newAvatarHandler()

	student := helpers.CreateTestUser(t, "student")

	c, w := helpers.SetupAuthContext("GET", "/api/avatars/"+student.ID, nil)
	c.Params = []gin.Param{{Key: "id", Value: student.ID}}

	handler.ServeAvatarByUserID(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestServeAvatar_MissingID(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	handler := newAvatarHandler()

	c, w := helpers.SetupAuthContext("GET", "/api/avatars/", nil)

	handler.ServeAvatarByUserID(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestServeAvatar_Success(t *testing.T) {
	helpers.CleanupDB(t)
	defer helpers.CleanupDB(t)

	student := helpers.CreateTestUser(t, "student")
	handler := newAvatarHandler()

	err := os.MkdirAll("./storage/avatars", 0755)
	assert.NoError(t, err)
	avatarPath := filepath.Join("./storage/avatars", student.ID+".png")
	err = os.WriteFile(avatarPath, []byte("fake-png"), 0644)
	assert.NoError(t, err)
	defer os.Remove(avatarPath)

	_, err = helpers.TestDB.Collection("users").UpdateOne(context.Background(),
		bson.M{"_id": student.ID},
		bson.M{"$set": bson.M{"avatar_name": student.ID + ".png"}},
	)
	assert.NoError(t, err)

	c, w := helpers.SetupAuthContext("GET", "/api/avatars/"+student.ID, nil)
	c.Params = []gin.Param{{Key: "id", Value: student.ID}}

	handler.ServeAvatarByUserID(c)

	assert.Equal(t, http.StatusOK, w.Code)
}


