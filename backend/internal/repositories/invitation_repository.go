package repositories

import (
	"context"
	"errors"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type InvitationRepository struct {
	collection *mongo.Collection
}

func NewInvitationRepository(db *mongo.Database) *InvitationRepository {
	return &InvitationRepository{
		collection: db.Collection("invitations"),
	}
}

func (r *InvitationRepository) Create(ctx context.Context, invitation *models.Invitation) error {
	_, err := r.collection.InsertOne(ctx, invitation)
	return err
}

func (r *InvitationRepository) FindByStudentID(ctx context.Context, studentID string) ([]*models.Invitation, error) {
	filter := bson.M{"student_id": studentID, "status": "pending"}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var invitations []*models.Invitation
	if err := cursor.All(ctx, &invitations); err != nil {
		return nil, err
	}
	return invitations, nil
}

func (r *InvitationRepository) FindByID(ctx context.Context, invitationID string) (*models.Invitation, error) {
	var invitation models.Invitation
	err := r.collection.FindOne(ctx, bson.M{"_id": invitationID}).Decode(&invitation)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("invitation not found")
		}
		return nil, err
	}
	return &invitation, nil
}

func (r *InvitationRepository) UpdateStatus(ctx context.Context, invitationID string, status string) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": invitationID},
		bson.M{"$set": bson.M{"status": status}},
	)
	return err
}
