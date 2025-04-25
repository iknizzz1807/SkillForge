package services

import (
	"context"
	"math"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/integrations"
	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type MatchingService struct {
	db        *mongo.Database
	aiClient  *integrations.AIClient
	matchRepo *repositories.MatchingRepository
}

func NewMatchingService(db *mongo.Database, aiClient *integrations.AIClient) *MatchingService {
	return &MatchingService{
		db:        db,
		aiClient:  aiClient,
		matchRepo: repositories.NewMatchingRepository(db),
	}
}

// ProjectMatchInfo holds match score and basic project info
type ProjectMatchInfo struct {
	ProjectID     string    `json:"project_id"`
	ProjectSkills []string  `json:"project_skills"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	CreatorName   string    `json:"creator_name"`
	MatchScore    float64   `json:"match_score"`
}

// GetTopMatchesForStudent retrieves the top 10 best project matches for a student
func (s *MatchingService) GetTopMatchesForStudent(studentID string) ([]models.Match, error) {
	ctx := context.Background()
	return s.matchRepo.GetTopMatchesForStudent(ctx, studentID, 10)
}

// GetTopProjectMatchesForStudent gets top 10 projects with highest match scores for a student
func (s *MatchingService) GetTopProjectMatchesForStudent(studentID string) ([]ProjectMatchInfo, error) {
	ctx := context.Background()
	userRepo := repositories.NewUserRepository(s.db)

	// Get top matches first
	matches, err := s.matchRepo.GetTopMatchesForStudent(ctx, studentID, 10)
	if err != nil {
		return nil, err
	}

	// Get project details
	projRepo := repositories.NewProjectRepository(s.db)
	var results []ProjectMatchInfo

	for _, match := range matches {
		proj, err := projRepo.FindProjectByID(ctx, match.ProjectID)
		if err != nil {
			continue
		}

		// Get creator info
		creator, err := userRepo.FindUserByID(ctx, proj.CreatedByID)
		if err != nil {
			continue
		}

		results = append(results, ProjectMatchInfo{
			ProjectID:     proj.ID,
			ProjectSkills: proj.Skills,
			StartTime:     proj.StartTime,
			EndTime:       proj.EndTime,
			CreatorName:   creator.Name,
			MatchScore:    match.Score,
		})
	}

	return results, nil
}

// GetMatchScore retrieves the match score between a student and project
// If no score exists, calculates it using AI service and saves it
func (s *MatchingService) GetMatchScore(studentID, projectID string, studentSkills, projectSkills []string) (float64, error) {
	ctx := context.Background()

	// Try to get existing score
	score, err := s.matchRepo.GetMatchScore(ctx, studentID, projectID)
	if err == nil && score != nil {
		return score.Score, nil
	}

	// Calculate new score using AI service
	matchScore, err := s.aiClient.MatchSkills(studentSkills, projectSkills)
	if err != nil {
		return 0, err
	}

	// Round score before returning
	matchScore = math.Round(matchScore)

	// Save new score
	match := &models.Match{
		StudentID: studentID,
		ProjectID: projectID,
		Score:     matchScore,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.matchRepo.SaveMatchScore(ctx, match); err != nil {
		return 0, err
	}

	return matchScore, nil
}

// GetMatchScoreForPair gets the match score for a specific student-project pair
func (s *MatchingService) GetMatchScoreForPair(studentID, projectID string) (float64, error) {
	ctx := context.Background()

	// Get match from repository
	match, err := s.matchRepo.GetMatchScore(ctx, studentID, projectID)
	if err != nil {
		return 0, err
	}

	if match == nil {
		return 0, nil
	}

	return match.Score, nil
}

// UpdateMatchScore updates the match score for a student-project pair
func (s *MatchingService) UpdateMatchScore(studentID, projectID string, studentSkills, projectSkills []string) (float64, error) {
	ctx := context.Background()

	// Calculate new score
	matchScore, err := s.aiClient.MatchSkills(studentSkills, projectSkills)
	if err != nil {
		return 0, err
	}

	// Round score before returning
	matchScore = math.Round(matchScore)

	// Update score in database
	match := &models.Match{
		StudentID: studentID,
		ProjectID: projectID,
		Score:     matchScore,
		UpdatedAt: time.Now(),
	}

	if err := s.matchRepo.UpdateMatchScore(ctx, match); err != nil {
		return 0, err
	}

	return matchScore, nil
}
