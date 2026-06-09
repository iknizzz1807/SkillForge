package repositories

import (
	"context"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ChatRepository struct {
	MessageCollection        *mongo.Collection
	GroupCollection          *mongo.Collection
	ProjectStudentCollection *mongo.Collection
	ProjectCollection        *mongo.Collection
	UserCollection           *mongo.Collection
}

func NewChatRepository(db *mongo.Database) *ChatRepository {
	messageCollection := db.Collection("messages")
	groupCollection := db.Collection("groups")
	projectStudentCollection := db.Collection("project_student")

	_, _ = messageCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{{Key: "group_id", Value: 1}, {Key: "timestamp", Value: -1}},
	})
	_, _ = groupCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{{Key: "project_id", Value: 1}},
	})
	_, _ = projectStudentCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{{Key: "student_id", Value: 1}},
	})

	return &ChatRepository{
		MessageCollection:        messageCollection,
		GroupCollection:          groupCollection,
		ProjectStudentCollection: projectStudentCollection,
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

	var projectStudents []models.Project_student
	if err := cursor.All(ctx, &projectStudents); err != nil {
		cursor.Close(ctx)
		return nil, err
	}
	cursor.Close(ctx)

	// Collect project IDs from project_student
	projectIDSet := make(map[string]bool)
	for _, ps := range projectStudents {
		projectIDSet[ps.Project_id] = true
	}

	// Also query projects where user is the creator (business users)
	projectCursor, err := r.ProjectCollection.Find(ctx, bson.M{"created_by_id": userID})
	if err != nil {
		return nil, err
	}

	var projects []models.Project
	if err := projectCursor.All(ctx, &projects); err != nil {
		projectCursor.Close(ctx)
		return nil, err
	}
	projectCursor.Close(ctx)

	for _, p := range projects {
		projectIDSet[p.ID] = true
	}

	// Get or create group for each project ID
	groups := []*models.Group{}
	for projectID := range projectIDSet {
		var group models.Group
		err := r.GroupCollection.FindOne(ctx, bson.M{"project_id": projectID}).Decode(&group)
		if err == mongo.ErrNoDocuments {
			// Get project title
			var project models.Project
			err := r.ProjectCollection.FindOne(ctx, bson.M{"_id": projectID}).Decode(&project)
			if err != nil {
				return nil, err
			}
			group = models.Group{
				ProjectID: projectID,
				Title:     project.Title,
				CreatedAt: time.Now(),
			}
			_, err = r.GroupCollection.InsertOne(ctx, &group)
			if err != nil {
				return nil, err
			}
		} else if err != nil {
			return nil, err
		}
		groups = append(groups, &group)
	}

	return groups, nil
}

func (r *ChatRepository) GetGroupMessages(groupID string) ([]*models.Message, error) {
	// Get all messages with groupID
	ctx := context.Background()
	findOptions := options.Find().SetSort(bson.M{"created_at": 1})
	cursor, err := r.MessageCollection.Find(ctx, bson.M{"group_id": groupID}, findOptions)
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
	if len(usersID) > 0 {
		filter := bson.M{"_id": bson.M{"$in": usersID}}
		cursor, err := r.UserCollection.Find(ctx, filter)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err := cursor.All(ctx, &users); err != nil {
			return nil, err
		}
	}

	return users, nil
}

func (r *ChatRepository) CheckProjectAccess(projectID, userID string) (bool, error) {
	ctx := context.Background()
	count, err := r.ProjectStudentCollection.CountDocuments(ctx, bson.M{"project_id": projectID, "student_id": userID})
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	var project models.Project
	err = r.ProjectCollection.FindOne(ctx, bson.M{"_id": projectID, "created_by_id": userID}).Decode(&project)
	if err == mongo.ErrNoDocuments {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
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


func (r *ChatRepository) DeleteGroupMessages(ctx context.Context, groupID string) error {
	filter := bson.M{"group_id": groupID}
	_, err := r.MessageCollection.DeleteMany(ctx, filter)
	if err != nil {
		return err // avoid importing fmt if not present, will add later
	}
	return nil
}

func (r *ChatRepository) DeleteGroupByProjectID(ctx context.Context, projectID string) error {
	filter := bson.M{"project_id": projectID}
	// delete messages
	err := r.DeleteGroupMessages(ctx, projectID)
	if err != nil {
		return err
	}

	_, err = r.GroupCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
