package services

import (
	"context"
	"fmt"
	"log"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/integrations"
	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type MatchingService struct {
	db          *mongo.Database
	aiClient    *integrations.AIClient
	userRepo    *repositories.UserRepository
	projectRepo *repositories.ProjectRepository
}

func NewMatchingService(db *mongo.Database, aiClient *integrations.AIClient) *MatchingService {
	return &MatchingService{
		db:          db,
		aiClient:    aiClient,
		userRepo:    repositories.NewUserRepository(db),
		projectRepo: repositories.NewProjectRepository(db),
	}
}

type ProjectMatchInfo struct {
	ProjectID     string    `json:"project_id"`
	ProjectTitle  string    `json:"project_title"`
	ProjectSkills []string  `json:"project_skills"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	CreatorName   string    `json:"creator_name"`
	MatchScore    float64   `json:"match_score"`
}

func joinSkills(skills []string) string {
	seen := make(map[string]struct{})
	var unique []string
	for _, s := range skills {
		trimmed := strings.TrimSpace(s)
		if trimmed == "" {
			continue
		}
		if _, ok := seen[trimmed]; !ok {
			seen[trimmed] = struct{}{}
			unique = append(unique, trimmed)
		}
	}
	return strings.Join(unique, " ")
}

func (s *MatchingService) GetScore(userID string, projectID string) (float64, error) {
	user, err := s.userRepo.FindUserByID(context.Background(), userID)
	if err != nil {
		return 0, err
	}
	if user == nil {
		return 0, fmt.Errorf("user not found")
	}

	project, err := s.projectRepo.FindProjectByID(context.Background(), projectID)
	if err != nil {
		return 0, err
	}

	if len(user.Skills) == 0 {
		return 0, nil
	}

	userInfo := joinSkills(user.Skills)
	if user.Title != "" {
		userInfo += " " + user.Title
	}

	projectInfo := joinSkills(project.Skills)
	if project.Title != "" {
		projectInfo += " " + project.Title
	}
	if project.Description != "" {
		projectInfo += " " + project.Description
	}

	matchScore, err := s.aiClient.MatchSkills(userInfo, projectInfo)
	if err != nil {
		return 0, err
	}

	matchScore = math.Round(matchScore)
	return matchScore, nil
}

func (s *MatchingService) GetScoreUserProjects(userID string, projects []*models.Project) ([]float64, error) {
	log.Printf("GetScoreUserProjects")

	user, err := s.userRepo.FindUserByID(context.Background(), userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	if len(user.Skills) == 0 {
		scores := make([]float64, len(projects))
		return scores, nil
	}

	userInfo := joinSkills(user.Skills)
	if user.Title != "" {
		userInfo += " " + user.Title
	}

	listProjectInfo := []string{}

	for _, project := range projects {
		if project == nil {
			listProjectInfo = append(listProjectInfo, "")
			continue
		}
		projectInfo := joinSkills(project.Skills)
		if project.Title != "" {
			projectInfo += " " + project.Title
		}
		if project.Description != "" {
			projectInfo += " " + project.Description
		}
		listProjectInfo = append(listProjectInfo, projectInfo)
	}

	matchScores, err := s.aiClient.MatchSkillsWithProject(userInfo, listProjectInfo)
	if err != nil {
		return nil, err
	}

	return matchScores, nil
}

func (s *MatchingService) GetTopProjectMatchesForStudent(userID string) ([]ProjectMatchInfo, error) {
	log.Printf("GetTopProjectMatchesForStudent")

	projects, err := s.projectRepo.FindAllProjects(context.Background(), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return []ProjectMatchInfo{}, nil
	}

	listScores, err := s.GetScoreUserProjects(userID, projects)
	if err != nil {
		return nil, err
	}

	if len(listScores) != len(projects) {
		return nil, fmt.Errorf("mismatch between projects count (%d) and scores count (%d)",
			len(projects), len(listScores))
	}

	listProjects := []ProjectMatchInfo{}

	for i, project := range projects {
		if i >= len(listScores) {
			break
		}

		matchScore := listScores[i]

		creator, err := s.userRepo.FindUserByID(context.Background(), project.CreatedByID)
		if err != nil {
			log.Printf("Error finding creator for project %s: %v", project.ID, err)
			continue
		}
		if creator == nil {
			log.Printf("Creator not found for project %s", project.ID)
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

	sort.Slice(listProjects, func(i, j int) bool {
		return listProjects[i].MatchScore > listProjects[j].MatchScore
	})

	if len(listProjects) > 10 {
		listProjects = listProjects[:10]
	}

	return listProjects, nil
}
