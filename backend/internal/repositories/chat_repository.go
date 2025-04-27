package repositories

import (
	"context"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatRepository struct {
	MessageCollection *mongo.Collection
	GroupCollection   *mongo.Collection
	ProjectStudentCollection *mongo.Collection
	ProjectCollection *mongo.Collection
}

func NewChatRepository(db *mongo.Database) *ChatRepository {
	return &ChatRepository{
		MessageCollection: db.Collection("messages"),
		GroupCollection:   db.Collection("groups"),
		ProjectStudentCollection: db.Collection("project_student"),
		ProjectCollection: db.Collection("projects"),
	}
}

func (r *ChatRepository) GetGroups(ctx context.Context, userID string) ([]*models.Group, error) {
	// Get all project_student with studentID
	cursor, err := r.ProjectStudentCollection.Find(ctx, bson.M{"student_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var projectStudents []models.Project_student
	if err := cursor.All(ctx, &projectStudents); err != nil {
		return nil, err
	}

	// Get all group with projectID
	groups := []*models.Group{}
	for _, projectStudent := range projectStudents {
		// Find one group with projectID
		var group models.Group
		err := r.GroupCollection.FindOne(ctx, bson.M{"project_id": projectStudent.Project_id}).Decode(&group)
		if err == mongo.ErrNoDocuments {
			// Create new group
			// Get title
			var project models.Project
			err := r.ProjectCollection.FindOne(ctx, bson.M{"_id": projectStudent.Project_id}).Decode(&project)
			if err != nil {
				return nil, err
			}
			group = *&models.Group{
				ProjectID: projectStudent.Project_id,
				Title:     project.Title,
				CreatedAt: time.Now(),
			}
			groups = append(groups, &group)
		} else if err != nil {
			return nil, err
		}
		group = *&models.Group{
			ProjectID: projectStudent.Project_id,
			Title:     group.Title,
		}
		groups = append(groups, &group)
	}

	return groups, nil
}

func (r *ChatRepository) GetGroupMessages(ctx context.Context, groupID string) ([]*models.Message, error) {
	// Get all messages with groupID
	cursor, err := r.MessageCollection.Find(ctx, bson.M{"group_id": groupID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var messages []*models.Message
	if err := cursor.All(ctx, &messages); err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *ChatRepository) GetGroupMembers(ctx context.Context, groupID string) ([]string, error) {
	// Get all userID with groupID in project_student
	cursor, err := r.ProjectStudentCollection.Find(ctx, bson.M{"project_id": groupID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var projectStudents []models.Project_student
	if err := cursor.All(ctx, &projectStudents); err != nil {
		return nil, err
	}

	// Get all userID 
	users := []string{}
	for _, projectStudent := range projectStudents {
		users = append(users, projectStudent.Student_id)
	}
	// Get userID of the project creator
	var project models.Project
	err = r.ProjectCollection.FindOne(ctx, bson.M{"_id": groupID}).Decode(&project)
	if err != nil {
		return nil, err
	}
	users = append(users, project.CreatedByID)

	return users, nil
}

func (r *ChatRepository) InsertMessage(ctx context.Context, message *models.Message) (*models.Message, error) {
	// Insert message into database
	message.CreatedAt = time.Now()
	_, err := r.MessageCollection.InsertOne(ctx, message)
	if err != nil {
		return nil, err
	}
	return message, nil
}