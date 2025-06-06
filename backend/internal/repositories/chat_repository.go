package repositories

import (
	"context"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatRepository struct {
	MessageCollection        *mongo.Collection
	GroupCollection          *mongo.Collection
	ProjectStudentCollection *mongo.Collection
	ProjectCollection        *mongo.Collection
	UserCollection           *mongo.Collection
}

func NewChatRepository(db *mongo.Database) *ChatRepository {
	return &ChatRepository{
		MessageCollection:        db.Collection("messages"),
		GroupCollection:          db.Collection("groups"),
		ProjectStudentCollection: db.Collection("project_student"),
		ProjectCollection:        db.Collection("projects"),
		UserCollection:           db.Collection("users"),
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

func (r *ChatRepository) GetGroupMessages(groupID string) ([]*models.Message, error) {
	// Get all messages with groupID
	ctx := context.Background()
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

func (r *ChatRepository) GetGroupMembers(groupID string) ([]*models.User, error) {
	// Get all userID with groupID in project_student
	ctx := context.Background()
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
	usersID := []string{}
	for _, projectStudent := range projectStudents {
		usersID = append(usersID, projectStudent.Student_id)
	}
	// Get userID of the project creator
	var project models.Project
	err = r.ProjectCollection.FindOne(ctx, bson.M{"_id": groupID}).Decode(&project)
	if err != nil {
		return nil, err
	}
	usersID = append(usersID, project.CreatedByID)

	// Get all user with userID
	users := []*models.User{}
	for _, userID := range usersID {
		var user models.User
		err := r.UserCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *ChatRepository) InsertMessage(message *models.Message) (*models.Message, error) {
	// Insert message into database
	ctx := context.Background()
	message.CreatedAt = time.Now()
	_, err := r.MessageCollection.InsertOne(ctx, message)
	if err != nil {
		return nil, err
	}
	return message, nil
}
