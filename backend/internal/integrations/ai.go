package integrations

// Nhiệm vụ: Tích hợp dịch vụ AI (FastAPI) để gọi API matching
// Liên kết:
// - Dùng net/http để gửi request HTTP tới FastAPI
// - Được gọi từ services/project_service.go để match sinh viên với dự án
// Vai trò trong flow:
// - Gọi API AI để tính toán sự phù hợp giữa skills của sinh viên và dự án

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type AIClient struct {
	// url là địa chỉ của dịch vụ AI (ví dụ: http://localhost:8000)
	url string
}

// NewAIClient khởi tạo AIClient với URL
// Input: url (string) - URL của FastAPI server
// Return: *AIClient - con trỏ đến AIClient
func NewAIClient(url string) *AIClient {
	return &AIClient{url}
}

// MatchSkills gọi API AI để match skills
// Input: userSkills ([]string), projectSkills ([]string)
// Return: float64 (độ phù hợp), error (nếu có lỗi)
func (c *AIClient) MatchSkills(userSkills, projectSkills []string) (float64, error) {
	// Tạo payload JSON
	payload := map[string][]string{
		"user_skills":    userSkills,
		"project_skills": projectSkills,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return 0, err
	}

	// Gửi POST request tới FastAPI
	resp, err := http.Post(c.url+"/match", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// Decode response
	var result struct {
		MatchScore float64 `json:"match_score"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, err
	}

	// Trả về độ phù hợp
	return result.MatchScore, nil
}
