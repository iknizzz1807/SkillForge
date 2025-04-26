package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TalentPoolRepository struct {
	collection *mongo.Collection
}

func NewTalentPoolRepository(db *mongo.Database) *TalentPoolRepository {
	return &TalentPoolRepository{
		collection: db.Collection("talent_pool"),
	}
}

// FindByBusinessID retrieves all talent pool entries for a specific business
func (r *TalentPoolRepository) FindByBusinessID(ctx context.Context, businessID string) ([]models.TalentPool, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	filter := bson.M{"business_id": businessID}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var talentPool []models.TalentPool
	if err := cursor.All(ctx, &talentPool); err != nil {
		return nil, err
	}

	return talentPool, nil
}

// FindByBusinessAndStudent retrieves a talent pool entry for a specific business and student
func (r *TalentPoolRepository) FindByBusinessAndStudent(ctx context.Context, businessID, studentID string) (*models.TalentPool, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	filter := bson.M{
		"business_id": businessID,
		"student_id":  studentID,
	}

	var talentPoolEntry models.TalentPool
	err := r.collection.FindOne(ctx, filter).Decode(&talentPoolEntry)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &talentPoolEntry, nil
}

// InsertTalentPool adds a new student to a business's talent pool
func (r *TalentPoolRepository) InsertTalentPool(ctx context.Context, talentPool *models.TalentPool) error {
	if ctx == nil {
		ctx = context.Background()
	}

	// Check if entry already exists
	existing, err := r.FindByBusinessAndStudent(ctx, talentPool.BusinessID, talentPool.StudentID)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("student already exists in talent pool")
	}

	// Set timestamps
	talentPool.CreatedAt = time.Now()

	_, err = r.collection.InsertOne(ctx, talentPool)
	return err
}

// UpdateTalentPool updates an existing talent pool entry
func (r *TalentPoolRepository) UpdateTalentPool(ctx context.Context, talentPool *models.TalentPool) error {
	if ctx == nil {
		ctx = context.Background()
	}

	filter := bson.M{
		"business_id": talentPool.BusinessID,
		"student_id":  talentPool.StudentID,
	}

	// Set update timestamp
	talentPool.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{
			"student_name": talentPool.StudentName,
			"skills":       talentPool.Skills,
			"notes":        talentPool.Notes,
			"updated_at":   talentPool.UpdatedAt,
		},
	}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

// RemoveFromTalentPool removes a student from a business's talent pool
func (r *TalentPoolRepository) RemoveFromTalentPool(ctx context.Context, businessID, studentID string) error {
	if ctx == nil {
		ctx = context.Background()
	}

	filter := bson.M{
		"business_id": businessID,
		"student_id":  studentID,
	}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("student not found in talent pool")
	}

	return nil
}

func (r *TalentPoolRepository) CheckStudentInTalentPool(ctx context.Context, businessID, studentID string) (bool, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	filter := bson.M{
		"business_id": businessID,
		"student_id":  studentID,
	}

	var talentPoolEntry models.TalentPool
	err := r.collection.FindOne(ctx, filter).Decode(&talentPoolEntry)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
