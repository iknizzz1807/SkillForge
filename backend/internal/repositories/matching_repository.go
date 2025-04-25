package repositories

import (
	"context"
	"math"

	"github.com/iknizzz1807/SkillForge/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MatchingRepository struct {
	collection *mongo.Collection
}

func NewMatchingRepository(db *mongo.Database) *MatchingRepository {
	return &MatchingRepository{
		collection: db.Collection("matches"),
	}
}

// GetMatchScore retrieves the match score for a student-project pair
func (r *MatchingRepository) GetMatchScore(ctx context.Context, studentID, projectID string) (*models.Match, error) {
	filter := bson.M{
		"student_id": studentID,
		"project_id": projectID,
	}

	var match models.Match
	err := r.collection.FindOne(ctx, filter).Decode(&match)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &match, nil
}

// GetTopMatchesForStudent gets the top N matches for a student
func (r *MatchingRepository) GetTopMatchesForStudent(ctx context.Context, studentID string, limit int) ([]models.Match, error) {
	filter := bson.M{"student_id": studentID}
	opts := options.Find().
		SetSort(bson.D{{Key: "score", Value: -1}}).
		SetLimit(int64(limit))

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var matches []models.Match
	if err := cursor.All(ctx, &matches); err != nil {
		return nil, err
	}

	// Round all scores to whole numbers
	for i := range matches {
		matches[i].Score = math.Round(matches[i].Score)
	}

	return matches, nil
}

// SaveMatchScore saves a new match score
func (r *MatchingRepository) SaveMatchScore(ctx context.Context, match *models.Match) error {
	// Round score before saving
	match.Score = math.Round(match.Score)
	_, err := r.collection.InsertOne(ctx, match)
	return err
}

// UpdateMatchScore updates an existing match score
func (r *MatchingRepository) UpdateMatchScore(ctx context.Context, match *models.Match) error {
	// Round score before updating
	match.Score = math.Round(match.Score)
	filter := bson.M{
		"student_id": match.StudentID,
		"project_id": match.ProjectID,
	}

	update := bson.M{
		"$set": bson.M{
			"score":      match.Score,
			"updated_at": match.UpdatedAt,
		},
	}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}
