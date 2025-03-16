package integrations

// Nhiệm vụ: Tích hợp GitHub API để tạo repository cho dự án
// Liên kết:
// - Dùng github.com/google/go-github/v50 để gọi GitHub API
// - Được gọi từ services/project_service.go để tạo repository
// Vai trò trong flow:
// - Tạo repository GitHub khi dự án được tạo

import (
	"context"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

type GitHubClient struct {
	// client là client GitHub API
	client *github.Client
}

// NewGitHubClient khởi tạo GitHubClient với token
// Input: token (string) - GitHub personal access token
// Return: *GitHubClient - con trỏ đến GitHubClient
func NewGitHubClient(token string) *GitHubClient {
	// Tạo OAuth2 token source
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	// Tạo HTTP client với token
	tc := oauth2.NewClient(context.Background(), ts)
	// Tạo GitHub client
	client := github.NewClient(tc)
	return &GitHubClient{client}
}

// CreateRepository tạo repository mới trên GitHub
// Input: name (string) - tên repository
// Return: string (URL repository), error (nếu có lỗi)
func (c *GitHubClient) CreateRepository(name string) (string, error) {
	// Tạo repository request
	repo := &github.Repository{
		Name:    github.String(name),
		Private: github.Bool(false), // Public repository
	}

	// Gửi request tạo repository
	createdRepo, _, err := c.client.Repositories.Create(context.Background(), "", repo)
	if err != nil {
		return "", err
	}

	// Trả về URL repository
	return *createdRepo.HTMLURL, nil
}
