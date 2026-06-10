package services_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/iknizzz1807/SkillForge/internal/integrations"
	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/services"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"github.com/iknizzz1807/SkillForge/tests/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func connectNotifUser(t *testing.T, realtime *integrations.RealtimeClient, userID string) {
	t.Helper()
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}
		conn, err := upgrader.Upgrade(w, r, nil)
		require.NoError(t, err)
		defer conn.Close()
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				break
			}
		}
	}))
	t.Cleanup(s.Close)

	url := "ws" + strings.TrimPrefix(s.URL, "http")
	wsConn, _, err := websocket.DefaultDialer.Dial(url+"/ws", nil)
	require.NoError(t, err)
	t.Cleanup(func() { wsConn.Close() })

	realtime.AddConnection("notification", userID, wsConn)
}

func setupNotifSvc(t *testing.T) *services.NotificationService {
	t.Helper()
	db := testutil.GetTestDB()
	testutil.CleanAll(db)
	return services.NewNotificationService(
		integrations.NewRealtimeClient(),
		db,
		testutil.NewDummyEmailClient(),
	)
}

func TestNotifService_GetUserNotifications(t *testing.T) {
	db := testutil.GetTestDB()
	testutil.CleanAll(db)
	realtime := integrations.NewRealtimeClient()
	svc := services.NewNotificationService(realtime, db, testutil.NewDummyEmailClient())
	userID := utils.GenerateUUID()
	connectNotifUser(t, realtime, userID)

	require.NoError(t, svc.SendNotification(userID, "Notif 1", "info"))
	require.NoError(t, svc.SendNotification(userID, "Notif 2", "alert"))

	notifications, err := svc.GetUserNotifications(userID)
	require.NoError(t, err)
	assert.Len(t, notifications, 2)
}

func TestNotifService_SendNotification(t *testing.T) {
	db := testutil.GetTestDB()
	testutil.CleanAll(db)
	realtimeClient := integrations.NewRealtimeClient()
	svc := services.NewNotificationService(
		realtimeClient,
		db,
		testutil.NewDummyEmailClient(),
	)

	userID := utils.GenerateUUID()
	connectNotifUser(t, realtimeClient, userID)
	err := svc.SendNotification(userID, "Test content", "test_type")
	require.NoError(t, err)

	notificationRepo := repositories.NewNotificationRepository(db)
	notifications, err := notificationRepo.GetUserNotifications(context.Background(), userID)
	require.NoError(t, err)
	require.Len(t, notifications, 1)
	assert.Equal(t, "Test content", notifications[0].Content)
	assert.Equal(t, "test_type", notifications[0].Type)
	assert.Equal(t, userID, notifications[0].ToUserID)
}

func TestNotifService_InvalidData(t *testing.T) {
	svc := setupNotifSvc(t)

	err := svc.SendNotification("", "content", "type")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid realtime data")

	err = svc.SendNotification("user", "", "type")
	assert.Error(t, err)

	err = svc.SendNotification("user", "content", "")
	assert.Error(t, err)
}

func TestNotifService_SendEmail(t *testing.T) {
	svc := setupNotifSvc(t)
	err := svc.SendEmail("to@test.com", "Subject", "Body")
	assert.Error(t, err)
}

func TestNotifService_CreateMany(t *testing.T) {
	db := testutil.GetTestDB()
	testutil.CleanAll(db)
	notificationRepo := repositories.NewNotificationRepository(db)
	userID := utils.GenerateUUID()

	for i := 0; i < 5; i++ {
		notif := &models.Notification{
			ID:       utils.GenerateUUID(),
			ToUserID: userID,
			Content:  "Notif",
			Type:     "info",
		}
		require.NoError(t, notificationRepo.InsertNotification(context.Background(), notif))
	}

	notifications, err := notificationRepo.GetUserNotifications(context.Background(), userID)
	require.NoError(t, err)
	assert.Len(t, notifications, 5)
}
