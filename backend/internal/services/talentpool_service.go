package services

import (
	"context"
	"errors"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"github.com/iknizzz1807/SkillForge/internal/repositories"
	"github.com/iknizzz1807/SkillForge/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type TalentPoolService struct {
	// db để truy cập MongoDB database
	db *mongo.Database
}

func NewTalentPoolService(db *mongo.Database) *TalentPoolService {
	return &TalentPoolService{db}
}

type StudentInfo struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Skills []string `json:"skills"`
}

// GetTalentPoolByBusinessID retrieves all students in a business's talent pool
func (s *TalentPoolService) GetTalentPoolByBusinessID(businessID string) ([]StudentInfo, error) {
	if businessID == "" {
		return nil, errors.New("business ID cannot be empty")
	}

	ctx := context.Background()
	talentPoolRepo := repositories.NewTalentPoolRepository(s.db)

	// Get all talent pool entries for the business
	talentPoolEntries, err := talentPoolRepo.FindByBusinessID(ctx, businessID)
	if err != nil {
		return nil, errors.New("failed to retrieve talent pool: " + err.Error())
	}

	// If no entries found, return empty array rather than nil
	if len(talentPoolEntries) == 0 {
		return []StudentInfo{}, nil
	}

	// Convert to StudentInfo format
	var studentInfos []StudentInfo
	for _, entry := range talentPoolEntries {
		studentInfos = append(studentInfos, StudentInfo{
			ID:     entry.StudentID,
			Name:   entry.StudentName,
			Skills: entry.Skills,
		})
	}

	return studentInfos, nil
}

// AddStudentToTalentPool adds a student to a business's talent pool
func (s *TalentPoolService) AddStudentToTalentPool(studentID string, businessID string) (StudentInfo, error) {
	if studentID == "" || businessID == "" {
		return StudentInfo{}, errors.New("student ID and business ID cannot be empty")
	}

	ctx := context.Background()

	// Retrieve student details from user repository
	userRepo := repositories.NewUserRepository(s.db)
	student, err := userRepo.FindUserByID(ctx, studentID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return StudentInfo{}, errors.New("student not found")
		}
		return StudentInfo{}, errors.New("failed to fetch student information: " + err.Error())
	}

	// Check if student is actually a student
	if student.Role != "student" {
		return StudentInfo{}, errors.New("specified user is not a student")
	}

	// Use skills directly from the user model
	skills := student.Skills

	// Create talent pool entry
	talentPoolEntry := &models.TalentPool{
		ID:          utils.GenerateUUID(),
		BusinessID:  businessID,
		StudentID:   studentID,
		StudentName: student.Name,
		Skills:      skills,
	}

	// Save to database
	talentPoolRepo := repositories.NewTalentPoolRepository(s.db)
	err = talentPoolRepo.InsertTalentPool(ctx, talentPoolEntry)
	if err != nil {
		if err.Error() == "student already exists in talent pool" {
			return StudentInfo{}, errors.New("student already in talent pool")
		}
		return StudentInfo{}, errors.New("failed to add student to talent pool: " + err.Error())
	}

	// Return student info
	return StudentInfo{
		ID:     studentID,
		Name:   student.Name,
		Skills: skills,
	}, nil
}

// RemoveFromTalentPool removes a student from a business's talent pool
func (s *TalentPoolService) RemoveFromTalentPool(studentID string, businessID string) error {
	if studentID == "" || businessID == "" {
		return errors.New("student ID and business ID cannot be empty")
	}

	ctx := context.Background()
	talentPoolRepo := repositories.NewTalentPoolRepository(s.db)

	err := talentPoolRepo.RemoveFromTalentPool(ctx, businessID, studentID)
	if err != nil {
		return errors.New("failed to remove student from talent pool: " + err.Error())
	}

	return nil
}

func (s *TalentPoolService) CheckStudentInTalentPool(studentID string, businessID string) (bool, error) {
	if studentID == "" || businessID == "" {
		return false, errors.New("student ID and business ID cannot be empty")
	}

	ctx := context.Background()
	talentPoolRepo := repositories.NewTalentPoolRepository(s.db)

	exists, err := talentPoolRepo.CheckStudentInTalentPool(ctx, businessID, studentID)
	if err != nil {
		return false, errors.New("failed to check if student is in talent pool: " + err.Error())
	}

	return exists, nil
}
