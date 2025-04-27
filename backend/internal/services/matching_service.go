package services

import (
	"context"
	"fmt"
	"log"
	"math"
	"sort"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/integrations"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type MatchingService struct {
	db       *mongo.Database
	aiClient *integrations.AIClient
}

func NewMatchingService(db *mongo.Database, aiClient *integrations.AIClient) *MatchingService {
	return &MatchingService{
		db:       db,
		aiClient: aiClient,
	}
}

// ProjectMatchInfo holds match score and basic project info
type ProjectMatchInfo struct {
	ProjectID     string    `json:"project_id"`
	ProjectTitle  string    `json:"project_title"`
	ProjectSkills []string  `json:"project_skills"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	CreatorName   string    `json:"creator_name"`
	MatchScore    float64   `json:"match_score"`
}

// GetTopMatchesForStudent retrieves the top 10 best project matches for a student
func (s *MatchingService) GetScore(userID string, projectID string) (float64, error) {
	userRepo := repositories.NewUserRepository(s.db)
	projectRepo := repositories.NewProjectRepository(s.db)

	user, err := userRepo.FindUserByID(context.Background(), userID)
	if err != nil {
		return 0, err
	}

	project, err := projectRepo.FindProjectByID(context.Background(), projectID)
	if err != nil {
		return 0, err
	}

	userInfo := ""
	for _, skill := range user.Skills {
		userInfo += skill + " "
	}
	userInfo += user.Title

	projectInfo := ""
	for _, skill := range project.Skills {
		projectInfo += skill + " "
	}
	projectInfo += project.Title + " " + project.Description

	matchScore, err := s.aiClient.MatchSkills(userInfo, projectInfo)
	if err != nil {
		return 0, err
	}

	matchScore = math.Round(matchScore)
	return matchScore, nil
}

func (s *MatchingService) GetScoreUserProjects(userID string) ([]float64, error) {
	fmt.Println("GetScoreUserProjects")
	userRepo := repositories.NewUserRepository(s.db)
	projectRepo := repositories.NewProjectRepository(s.db)

	user, err := userRepo.FindUserByID(context.Background(), userID)
	if err != nil {
		return nil, err
	}
	userInfo := ""
	for _, skill := range user.Skills {
		userInfo += skill + " "
	}
	userInfo += user.Title

	projects, err := projectRepo.FindAllProjects(context.Background())
	if err != nil {
		return nil, err
	}

	listProjectInfo := []string{}

	for _, project := range projects {
		projectInfo := ""
		for _, skill := range project.Skills {
			projectInfo += skill + " "
		}
		projectInfo += project.Title + " " + project.Description
		listProjectInfo = append(listProjectInfo, projectInfo)
	}

	matchScores, err := s.aiClient.MatchSkillsWithProject(userInfo, listProjectInfo)
	if err != nil {
		return nil, err
	}

	return matchScores, nil
}

func (s *MatchingService) GetTopProjectMatchesForStudent(userID string) ([]ProjectMatchInfo, error) {
	fmt.Println("GetTopProjectMatchesForStudent")
	userRepo := repositories.NewUserRepository(s.db)
	projectRepo := repositories.NewProjectRepository(s.db)

	// Lấy danh sách dự án
	projects, err := projectRepo.FindAllProjects(context.Background())
	if err != nil {
		return nil, err
	}

	// Kiểm tra số lượng project - thêm kiểm tra này
	if len(projects) == 0 {
		// Trả về slice rỗng nếu không có dự án
		return []ProjectMatchInfo{}, nil
	}

	// Lấy điểm match
	listScores, err := s.GetScoreUserProjects(userID)
	if err != nil {
		return nil, err
	}

	// Thêm kiểm tra độ dài của listScores và projects
	if len(listScores) != len(projects) {
		return nil, fmt.Errorf("mismatch between projects count (%d) and scores count (%d)",
			len(projects), len(listScores))
	}

	listProjects := []ProjectMatchInfo{}

	for i, project := range projects {
		// Kiểm tra index trước khi truy cập
		if i >= len(listScores) {
			break
		}

		matchScore := listScores[i]

		creator, err := userRepo.FindUserByID(context.Background(), project.CreatedByID)
		if err != nil {
			// Log lỗi nhưng không dừng toàn bộ xử lý
			log.Printf("Error finding creator for project %s: %v", project.ID, err)
			continue
		}

		listProjects = append(listProjects, ProjectMatchInfo{
			ProjectID:     project.ID,
			ProjectTitle:  project.Title,
			ProjectSkills: project.Skills,
			StartTime:     project.StartTime,
			EndTime:       project.EndTime,
			CreatorName:   creator.Name,
			MatchScore:    matchScore,
		})
	}

	// Sắp xếp dự án theo điểm match
	sort.Slice(listProjects, func(i, j int) bool {
		return listProjects[i].MatchScore > listProjects[j].MatchScore
	})

	if len(listProjects) > 10 {
		listProjects = listProjects[:10]
	}

	return listProjects, nil
}
