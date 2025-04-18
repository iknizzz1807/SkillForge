package services

import (
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

func (s *TalentPoolService) GetTalentPoolByBusinessID(businessID string) ([]StudentInfo, error) {
	// if businessID == "" {
	// 	return nil, errors.New("business id cannot be empty")
	// }

	// projectRepo := repositories.NewProjectRepository(s.db)
	// projects, err := projectRepo.GetProjectIDsByCreatorID(businessID)
	// if err != nil || projects == nil {
	// 	return nil, errors.New("projects not found")
	// }

	// projectStudentRepo := repositories.NewProjectStudentRepository(s.db)
	// studentIDs := []string{}
	// for _, project := range projects {
	// 	projectStudentIDs, err := projectStudentRepo.FindStudentsByProjectID(context.Background(), project)
	// 	if err != nil || projectStudentIDs == nil {
	// 		return nil, errors.New("student not found")
	// 	}
	// 	studentIDs = append(studentIDs, projectStudentIDs...)
	// }

	// studentIDs = utils.RemoveDuplicates(studentIDs)

	// StudentInfos := []StudentInfo{}
	// userRepo := repositories.NewUserRepository(s.db)
	// for _, studentID := range studentIDs {
	// 	userReturn, err := userRepo.FindUserByID(context.Background(), studentID)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	id := userReturn.ID
	// 	name := userReturn.Name
	// 	skills := userReturn.Skills

	// 	StudentInfos = append(StudentInfos, StudentInfo{
	// 		ID:        id,
	// 		Name:      name,
	// 		Skills:    skills,
	// 	})
	// }

	// if err != nil {
	// 	return nil, err
	// }
	// return StudentInfos, nil
}

func (s *TalentPoolService) AddStudentToTalentPool(studentID string, businessID string) (StudentInfo, error) {
	// Gọi đến repo để thêm vào database
}
