package services

import (
	"context"
	"embed"
	"errors"
	"html/template"
	"os"
	"path/filepath"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:embed portfolio_template.html
var templateFS embed.FS

type PortfolioService struct {
	db *mongo.Database
}

func NewPortfolioService(db *mongo.Database) *PortfolioService {
	return &PortfolioService{db}
}

type PortfolioProject struct {
	Name        string
	Description string
}

type PortfolioBadge struct {
	Name        string
	Description string
	Icon        string
}

type PortfolioData struct {
	Name      string
	Email     string
	Title     string
	Skills    []string
	AvatarURL string
	Website   string
	UserID    string
	Projects  []PortfolioProject
	Badges    []PortfolioBadge
}

func badgeIcon(badgeType string) string {
	switch badgeType {
	case "skill":
		return "💡"
	case "completion", "project":
		return "🏆"
	case "collaboration":
		return "👥"
	case "achievement":
		return "⭐"
	default:
		return "🎯"
	}
}

func (s *PortfolioService) GeneratePortfolio(userID string) (string, error) {
	if userID == "" {
		return "", errors.New("user ID cannot be empty")
	}

	ctx := context.Background()

	userRepo := repositories.NewUserRepository(s.db)
	user, err := userRepo.FindUserByID(ctx, userID)
	if err != nil || user == nil {
		return "", errors.New("user not found")
	}

	psRepo := repositories.NewProjectStudentRepository(s.db)
	projectIDs, err := psRepo.FindProjectsByStudentID(ctx, userID)
	if err != nil {
		return "", errors.New("failed to fetch projects")
	}

	var projects []PortfolioProject
	projRepo := repositories.NewProjectRepository(s.db)
	foundProjects, err := projRepo.FindProjectsByIDs(ctx, projectIDs)
	if err == nil {
		for _, p := range foundProjects {
			if p != nil {
				projects = append(projects, PortfolioProject{
					Name:        p.Title,
					Description: p.Description,
				})
			}
		}
	}

	badgeRepo := repositories.NewBadgeRepository(s.db)
	userBadges, err := badgeRepo.FindUserBadges(ctx, userID)
	if err != nil {
		userBadges = nil
	}

	var badges []PortfolioBadge
	for _, ub := range userBadges {
		if ub.Badge != nil {
			badges = append(badges, PortfolioBadge{
				Name:        ub.Badge.Name,
				Description: ub.Badge.Description,
				Icon:        badgeIcon(ub.Badge.Type),
			})
		}
	}

	avatarURL := ""
	if user.AvatarName != "" {
		avatarURL = "/avatars/" + userID
	}

	data := PortfolioData{
		Name:      user.Name,
		Email:     user.Email,
		Title:     user.Title,
		Skills:    user.Skills,
		AvatarURL: avatarURL,
		Website:   user.Website,
		UserID:    userID,
		Projects:  projects,
		Badges:    badges,
	}

	tmpl, err := template.ParseFS(templateFS, "portfolio_template.html")
	if err != nil {
		return "", errors.New("failed to parse template")
	}

	dir := "./storage/portfolios"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", errors.New("failed to create storage directory")
	}

	filePath := filepath.Join(dir, userID+".html")
	f, err := os.Create(filePath)
	if err != nil {
		return "", errors.New("failed to create portfolio file")
	}
	defer f.Close()

	if err := tmpl.Execute(f, data); err != nil {
		return "", errors.New("failed to execute template")
	}

	return "/portfolios/" + userID + ".html", nil
}

func (s *PortfolioService) GetPortfolio(userID string) (*models.Portfolio, error) {
	if userID == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	portfolioRepo := repositories.NewPortfolioRepository(s.db)
	portfolio, err := portfolioRepo.FindPortfolioByUserID(context.Background(), userID)
	if err != nil || portfolio == nil {
		return nil, errors.New("portfolio not found")
	}

	return portfolio, nil
}
