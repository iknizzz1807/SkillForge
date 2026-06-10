package integrations_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/integrations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAIClient_MatchSkills(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/matching", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		var reqBody map[string]string
		json.NewDecoder(r.Body).Decode(&reqBody)
		assert.Contains(t, reqBody["student_infos"], "Go")
		assert.Contains(t, reqBody["project_infos"], "MongoDB")

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"match_score": 85.5}`))
	}))
	defer server.Close()

	client := integrations.NewAIClient(server.URL)
	score, err := client.MatchSkills("Go Python", "MongoDB React")
	require.NoError(t, err)
	assert.Equal(t, 85.5, score)
}

func TestAIClient_MatchSkillsWithProject(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/matching2", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"match_score": [85.0, 70.5, 90.0]}`))
	}))
	defer server.Close()

	client := integrations.NewAIClient(server.URL)
	scores, err := client.MatchSkillsWithProject("Go Python", []string{"Go React", "Python Django", "Java Spring"})
	require.NoError(t, err)
	assert.Len(t, scores, 3)
	assert.Equal(t, 85.0, scores[0])
	assert.Equal(t, 70.5, scores[1])
	assert.Equal(t, 90.0, scores[2])
}

func TestAIClient_Non200Response(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"match_score": 0}`))
	}))
	defer server.Close()

	client := integrations.NewAIClient(server.URL)
	score, err := client.MatchSkills("Go", "Python")
	require.NoError(t, err)
	assert.Equal(t, float64(0), score)
}

func TestAIClient_ServiceDown(t *testing.T) {
	client := integrations.NewAIClient("http://localhost:19999")
	_, err := client.MatchSkills("Go", "Python")
	assert.Error(t, err)
}

func TestAIClient_MatchSkills_Timeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(50 * time.Millisecond)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"match_score": 85.0}`))
	}))
	defer server.Close()

	client := integrations.NewAIClient(server.URL)
	score, err := client.MatchSkills("Go", "Python")
	require.NoError(t, err)
	assert.Equal(t, 85.0, score)
}
